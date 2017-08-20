package logger

import (
	"fmt"
	"github.com/astaxie/beego"
)

var level int

func SetLevel(ll int) {
	level = ll
}

func Error(msg ...interface{}) {
	if level >= beego.LevelError {
		fmt.Println(msg)
	}
}

func Warn(msg ...interface{}) {
	if level >= beego.LevelWarning {
		fmt.Println(msg)
	}

}

func Info(msg ...interface{}) {
	if level >= beego.LevelInformational {
		fmt.Println(msg)
	}

}

func Debug(msg ...interface{}) {
	if level >= beego.LevelDebug {
		fmt.Println(msg)
	}

}
