package logger

import (
	"github.com/astaxie/beego"
	"testing"
)

func init() {
	SetLevel(beego.LevelInformational)
}

func TestLogger(t *testing.T) {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
}
