package config

//go:generate vc
type Config struct {
	Base struct {
		Env string `default:"dev"`
	}
	Etcd struct {
		Url string
	}
	ServiceGate struct {
		Port int
	}
	ServiceGame struct {
		Port int
	}
	ServiceLogin struct {
		Port int
	}
	ServiceReplay struct {
		Port int
	}
}

type goEnv int

const (
	Dev goEnv = iota
	Prod
)

func GoEnv() goEnv {
	if C.Base.Env == "dev" {
		return Dev
	} else {
		return Prod
	}
}

const (
	ServerMsgStart = 0
	GateServerMsgEnd = 10
	GameMsgStart = 100
	GameServerMsgEnd = 30000
	OtherServerMsgEnd = 32768
)