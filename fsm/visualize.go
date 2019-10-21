package fsm

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"text/template"
)

const temple = `digraph "fsm" {
	{{ range $v := .States }}
		"{{$v}}";{{ end }}
	{{ range $t := .Transitions }}
		"{{$t.From}}" -> "{{$t.To}}" [ label=" {{$t.Name}} " ];{{ end }}
}
`

func (fsm *StateMachine) Dot() string {
	type Vis struct {
		States      []State
		Transitions []*Transition
	}
	vis := &Vis{
		States: fsm.states,
	}
	for key, value := range fsm.transitions {
		vis.Transitions = append(vis.Transitions, &Transition{
			From: key.src,
			To:   value,
			Name: key.event,
		})
	}
	t, err := template.New("test").Parse(temple)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	err = t.Execute(buf, vis)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (fsm *StateMachine) Export(path string) error {
	dot := fsm.Dot()
	cmd := fmt.Sprintf("dot -o%s -T%s -K%s -s%s %s", path, "png", "dot", "72", "-Gsize=10,5 -Gdpi=200")
	return system(cmd, dot)
}

func system(cmdString string, stdin string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command(`cmd`, `/C`, cmdString)
	} else {
		cmd = exec.Command(`/bin/sh`, `-c`, cmdString)
	}
	cmd.Stdin = strings.NewReader(stdin)
	return cmd.Run()
}
