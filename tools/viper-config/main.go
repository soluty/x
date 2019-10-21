package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Config struct {
	Package string
	Config  string
	Env     string
	File    string
	Struct  string
	Output  string

	Remote       string
	RemoteUrl    string
	RemoteType   string
	RemoteKey    string
	RemoteSecret string

	FlagValues    []string
	EnvValues     []string
	DefaultValues []string
	HasNet        bool
	HasTime       bool
}

var SupportRemote = []string{"etcd", "consul"}
var SupportType = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"}

func containString(arr []string, s string) bool {
	for _, value := range arr {
		if value == s {
			return true
		}
	}
	return false
}

var C = &Config{}

func main() {
	flag.StringVar(&C.File, "f", "config.go", "结构体所在的go文件位置")
	flag.StringVar(&C.Package, "p", "config", "需要生成的包名")
	flag.StringVar(&C.Config, "c", "config", "生成的viper读取配置的文件名，无后缀")
	flag.StringVar(&C.Struct, "s", "Config", "配置结构体的名字")
	flag.StringVar(&C.Env, "e", "app", "环境变量前缀")
	flag.StringVar(&C.Output, "o", "config_gen.go", "生成的文件路径")
	flag.StringVar(&C.Remote, "r", "", "远程配置中心类型, 可选值etcd，consul")
	flag.StringVar(&C.RemoteUrl, "rurl", "", "远程配置中心url")
	flag.StringVar(&C.RemoteType, "rtype", "json", "远程配置中心文件类型")
	flag.StringVar(&C.RemoteKey, "rkey", "app_config_key", "远程配置中心k/v store的key")
	flag.StringVar(&C.RemoteSecret, "rsecret", "", "远程配置中心的secret keyring.gpg")

	flag.Parse()

	if C.RemoteUrl == "" {
		if C.Remote == "etcd" {
			C.RemoteUrl = "http://127.0.0.1:4001"
		} else if C.Remote == "consul" {
			C.RemoteUrl = "127.0.0.1:8500"
		}
	}

	if C.Remote != "" {
		if !containString(SupportRemote, C.Remote) {
			panic(fmt.Sprintf("不支持的远程配置中心，目前只支持%v", SupportRemote))
		}
		if !containString(SupportType, C.RemoteType) {
			panic(fmt.Sprintf("不支持的远程文件类型，目前只支持%v", SupportType))
		}
	}

	node, err := parser.ParseFile(token.NewFileSet(), C.File, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var configPos token.Pos

	ast.Inspect(node, func(nn ast.Node) bool {
		switch n := nn.(type) {
		case *ast.TypeSpec:
			if n.Name.Name == C.Struct {
				configPos = n.Type.Pos()
				return false
			}
		}
		return true
	})

	if !configPos.IsValid() {
		panic("need Config as Your Type")
	}

	var isStruct bool

	ast.Inspect(node, func(nn ast.Node) bool {
		switch n := nn.(type) {
		case *ast.StructType:
			if n.Pos() == configPos {
				process(n, "")
				isStruct = true
				return false
			}
		}
		return true
	})

	if !isStruct {
		panic("config should be a struct")
	}
	t, err := template.New("test").Parse(temple)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	err = t.Execute(buf, C)
	if err != nil {
		panic(err)
	}
	_ = os.Remove(C.Output)
	err = ioutil.WriteFile(C.Output, buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
	_ = os.Chmod(C.Output, 0555)
}

func process(n *ast.StructType, parent string) {
	for _, value := range n.Fields.List {
		var str string
		switch nt := value.Type.(type) {
		case *ast.ArrayType:
			switch i := nt.Elt.(type) {
			case *ast.Ident:
				processBasicArrayType(i, value, parent)
			case *ast.SelectorExpr:
				str = fmt.Sprintf("%v", i)
				processStructArrayType(value, str, parent)
			}
			str = ""
		case *ast.StructType:
			process(nt, getNameByParent(value, parent))
		case *ast.Ident:
			processBasicType(nt, value, parent)
		case *ast.SelectorExpr:
			str = fmt.Sprintf("%v", nt)
		//case *ast.StarExpr: 考虑了一下，还是不支持指针的好
		//	str = fmt.Sprintf("%v", nt.X)
		default:
		}
		if str != "" {
			processStructType(value, str, parent)
		}
	}
}

func processStructArrayType(value *ast.Field, str string, parent string) {
	tag := createTag(value)
	switch str {
	case "&{net IP}":
		C.HasNet = true
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.IPSlice("%v", %v, "%v")`, getNameByParent(value, parent), getIpSliceView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getIpSliceView(tag.Get("default"))))
	case "&{time Duration}":
		C.HasTime = true
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.DurationSlice("%v", %v, "%v")`, getNameByParent(value, parent), getDurationSliceView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getDurationSliceView(tag.Get("default"))))
	default:
		// panic("not support type " + getNameByParent(value, parent))
	}
}

func processStructType(value *ast.Field, str string, parent string) {
	tag := createTag(value)
	switch str {
	case "&{net IP}":
		C.HasNet = true
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.IP("%v", %v, "%v")`, getNameByParent(value, parent), getIpView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getIpView(tag.Get("default"))))
	case "&{time Duration}":
		C.HasTime = true
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.Duration("%v", %v, "%v")`, getNameByParent(value, parent), getDurationView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getDurationView(tag.Get("default"))))
	case "&{time Time}":
		// todo
		//C.HasTime = true
		//C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.Int("%v", %v, "%v")`, getNameByParent(value, parent), getIntView(tag.Get("default")), tag.Get("desc")))
		//C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		//C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getIntView(tag.Get("default"))))
	default:
		//panic("not support type " + getNameByParent(value, parent))
	}
}

func getNameByParent(value *ast.Field, parent string) string {
	if parent == "" {
		return strings.ToLower(value.Names[0].Name)
	} else {
		return parent + "." + strings.ToLower(value.Names[0].Name)
	}
}

func processBasicArrayType(nt *ast.Ident, value *ast.Field, parent string) {
	tag := createTag(value)
	switch nt.Name {
	case "int":
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.IntSlice("%v", %v, "%v")`, getNameByParent(value, parent), getIntSliceView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getIntSliceView(tag.Get("default"))))
	case "uint":
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.UintSlice("%v", %v, "%v")`, getNameByParent(value, parent), getUIntSliceView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getUIntSliceView(tag.Get("default"))))
	case "bool":
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.BoolSlice("%v", %v, "%v")`, getNameByParent(value, parent), getBoolSliceView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getBoolSliceView(tag.Get("default"))))
	case "string":
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.StringSlice("%v", %v, "%v")`, getNameByParent(value, parent), getStringSliceView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getStringSliceView(tag.Get("default"))))
	default:
		panic("not support type " + nt.Name)
	}
}

func processBasicType(nt *ast.Ident, value *ast.Field, parent string) {
	tag := createTag(value)
	switch nt.Name {
	case "int":
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.Int("%v", %v, "%v")`, getNameByParent(value, parent), getIntView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getIntView(tag.Get("default"))))
	case "uint":
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.UInt("%v", %v, "%v")`, getNameByParent(value, parent), getIntView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getIntView(tag.Get("default"))))
	case "string":
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.String("%v", %v, "%v")`, getNameByParent(value, parent), getStringView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getStringView(tag.Get("default"))))
	case "bool":
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.Bool("%v", %v, "%v")`, getNameByParent(value, parent), getBoolView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getBoolView(tag.Get("default"))))
	case "float32":
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.Float32("%v", %v, "%v")`, getNameByParent(value, parent), getIntView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getIntView(tag.Get("default"))))
	case "float64":
		C.FlagValues = append(C.FlagValues, fmt.Sprintf(`pflag.Float64("%v", %v, "%v")`, getNameByParent(value, parent), getIntView(tag.Get("default")), tag.Get("desc")))
		C.EnvValues = append(C.EnvValues, fmt.Sprintf(`_ = viper.BindEnv("%v")`, getNameByParent(value, parent)))
		C.DefaultValues = append(C.DefaultValues, fmt.Sprintf(`viper.SetDefault("%v", %v)`, getNameByParent(value, parent), getIntView(tag.Get("default"))))
	default:
		panic("------")
	}
}

func createTag(value *ast.Field) reflect.StructTag {
	var tag reflect.StructTag
	if value.Tag != nil {
		ss := value.Tag.Value
		ss = strings.TrimPrefix(ss, "`")
		ss = strings.TrimSuffix(ss, "`")
		tag = reflect.StructTag(ss)
	}
	return tag
}

func getStringSliceView(str string) string {
	if str == "" {
		return "nil"
	}
	ret := "[]string{"
	arr := strings.Split(str, ",")
	for idx, value := range arr {
		if idx < len(arr)-1 {
			ret += `"` + value + `", `
		} else {
			ret += `"` + value + `"`
		}
	}
	ret += "}"
	return ret
}

func getStringView(str string) string {
	if str == "" {
		return `""`
	}
	return `"` + str + `"`
}

func getIntView(str string) string {
	if str == "" {
		return "0"
	}
	return str
}

func getIntSliceView(str string) string {
	if str == "" {
		return "nil"
	}
	arr := strings.Split(str, ",")
	ret := "[]int{"
	for idx, value := range arr {
		d, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		if idx < len(arr)-1 {
			ret += fmt.Sprintf("%v", d) + `, `
		} else {
			ret += fmt.Sprintf("%v", d)
		}
	}
	ret += "}"
	return ret
}

func getUIntSliceView(str string) string {
	if str == "" {
		return "nil"
	}
	arr := strings.Split(str, ",")
	ret := "[]uint{"
	for idx, value := range arr {
		d, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		if idx < len(arr)-1 {
			ret += fmt.Sprintf("%v", d) + `, `
		} else {
			ret += fmt.Sprintf("%v", d)
		}
	}
	ret += "}"
	return ret
}

func getBoolSliceView(str string) string {
	if str == "" {
		return "nil"
	}
	arr := strings.Split(str, ",")
	ret := "[]bool{"
	for idx, value := range arr {
		s := strings.TrimSpace(value)
		b := "false"
		if s != "" {
			b = "true"
		}
		if idx < len(arr)-1 {
			ret += fmt.Sprintf("%v", b) + `, `
		} else {
			ret += fmt.Sprintf("%v", b)
		}
	}
	ret += "}"
	return ret
}

func getIpView(str string) string {
	if str == "" {
		return "net.IPv4zero"
	}
	ips := strings.Split(str, ".")
	if len(ips) != 4 {
		panic("todo")
	}
	return fmt.Sprintf("net.IPv4(%v,%v,%v,%v)", ips[0], ips[1], ips[2], ips[3])
}

func getDurationView(str string) string {
	if str == "" {
		return "0 * time.Nanosecond"
	}
	d, err := time.ParseDuration(str)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v * time.Nanosecond", d.Nanoseconds())
}

func getDurationSliceView(str string) string {
	if str == "" {
		return "nil"
	}
	arr := strings.Split(str, ",")
	ret := "[]time.Duration{"
	for idx, value := range arr {
		d, err := time.ParseDuration(value)
		if err != nil {
			panic(err)
		}
		if idx < len(arr)-1 {
			ret += fmt.Sprintf("%v * time.Nanosecond", d.Nanoseconds()) + `, `
		} else {
			ret += fmt.Sprintf("%v * time.Nanosecond", d.Nanoseconds())
		}
	}
	ret += "}"
	return ret
}

func getIpSliceView(str string) string {
	if str == "" {
		return "nil"
	}
	arr := strings.Split(str, ",")
	ret := "[]net.IP{"
	for idx, value := range arr {
		ips := strings.Split(value, ".")
		if len(ips) != 4 {
			panic("todo")
		}
		if idx < len(arr)-1 {
			ret += fmt.Sprintf("net.IPv4(%v,%v,%v,%v)", ips[0], ips[1], ips[2], ips[3]) + `, `
		} else {
			ret += fmt.Sprintf("net.IPv4(%v,%v,%v,%v)", ips[0], ips[1], ips[2], ips[3])
		}
	}
	ret += "}"
	return ret
}

func getBoolView(str string) string {
	if str == "" {
		return "false"
	}
	return "true"
}

const temple = `// Code generated by "viper-config"; DO NOT EDIT.

package {{ .Package }}

import (
	"flag"
	"fmt"
{{if .HasNet}}	"net"{{end}}
{{if .HasTime}}	"time"{{end}}
{{if ne .Remote ""}}	_ "github.com/spf13/viper/remote"{{end}}
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var C = {{ .Struct }}{}
var watchFunc func()

func Init(watch func(), filePath ...string) {
	watchFunc = watch
	setFlags()
	setEnv()
	readConfigFile(filePath)
	setDefaults()
	readFromRemote()
	err := viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}
}

func setFlags() {
{{ range $v := .FlagValues }}
	{{ $v }}{{ end }}
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

func setEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("{{ .Env }}")
{{ range $v := .EnvValues }}
	{{ $v }}{{ end }}
}

func setDefaults() {
{{ range $v := .DefaultValues }}
	{{ $v }}{{ end }}
}

func readFromRemote() {
	{{if ne .Remote ""}}
	{{if eq .RemoteSecret ""}}
	viper.AddRemoteProvider("{{.Remote}}", "{{.RemoteUrl}}", "{{.RemoteKey}}")
	{{else}}
	viper.AddSecureRemoteProvider("{{.Remote}}", "{{.RemoteUrl}}", "{{.RemoteKey}}", "{{.RemoteSecret}}")
	{{end}}
	viper.SetConfigType("{{.RemoteType}}")
	err := viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
	if watchFunc != nil {
		go func() {
			for {
				time.Sleep(time.Second * 5) // delay after each request
				err := viper.WatchRemoteConfig()
				if err != nil {
					log.Errorf("unable to read remote config: %v", err)
					continue
				}
				err := viper.Unmarshal(&C)
				if err != nil {
					fmt.Println(err)
					return
				}
				watchFunc()
			}
		} ()
	}
	{{ end }}
}

func readConfigFile(path []string) {
	env := os.Getenv("GO_ENV")
	if env != "" {
		viper.SetConfigName(env)
	} else {
		viper.SetConfigName("{{ .Config }}")
	}
	if len(path) == 0 {
		viper.AddConfigPath(".")
	}else {
		for _, value := range path {
			viper.AddConfigPath(value)
		}
	}
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	if watchFunc != nil {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			err := viper.Unmarshal(&C)
			if err != nil {
				fmt.Println(err)
				return
			}
			watchFunc()
		})
	}
}

`
