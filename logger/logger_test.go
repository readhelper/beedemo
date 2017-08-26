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

func TestCritical(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			//这里的err其实就是panic传入的内容
			Warn("critical:", err)    
		}
	}()
	Critical("critical")
}