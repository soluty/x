module github.com/soluty/x

go 1.13

require (
	github.com/alexeyco/simpletable v0.0.0-20191023080658-fe3ac9971811
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gogf/gf v1.11.4
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/hajimehoshi/ebiten v1.10.3
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/mattn/go-runewidth v0.0.8 // indirect
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/petergtz/pegomock v2.7.0+incompatible
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/prometheus/procfs v0.0.7 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.1
	github.com/tidwall/gjson v1.9.3
	github.com/tidwall/sjson v1.0.4
	github.com/twitchtv/twirp v5.10.1+incompatible
	go.etcd.io/bbolt v1.3.3 // indirect
	go.etcd.io/etcd v3.3.17+incompatible
	go.uber.org/zap v1.12.0 // indirect
	golang.org/x/crypto v0.0.0-20191122220453-ac88ee75c92c // indirect
	golang.org/x/exp v0.0.0-20191030013958-a1ab85dbe136 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8
	golang.org/x/net v0.0.0-20191116160921-f9c825593386 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys v0.0.0-20200212091648-12a6c2dcc1e4 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20191112195655-aa38f8e97acc // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20191115221424-83cc0476cb11 // indirect
	google.golang.org/grpc v1.25.1
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637
	gopkg.in/yaml.v2 v2.2.5 // indirect
	logur.dev/logur v0.15.1
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace github.com/soluty/x/xrand v1.0.0 => ./xrand

replace github.com/soluty/x/ecs v1.0.0 => ./ecs

replace github.com/soluty/x/cc v1.0.0 => ./cc

replace github.com/soluty/x/emitter v1.0.0 => ./emitter

replace github.com/soluty/x/xnet v1.0.0 => ./xnet
