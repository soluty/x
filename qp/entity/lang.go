package entity

import (
	"os"
	"strings"
)

type Lang int

const (
	En_us Lang = iota
	Zh_cn
)

func GetLang() Lang {
	lang := os.Getenv("LANG")
	if strings.HasPrefix(lang, "zh_") {
		return Zh_cn
	} else {
		return En_us
	}
}
