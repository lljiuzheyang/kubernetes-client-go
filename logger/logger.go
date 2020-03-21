package logger

import (
	"os"
	"../go-logging"
	"time"
	"../helper"
	"fmt"
)

var t = time.Now()
var log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

type logger struct {
}

func Debug(content string) {
	Instance("debug")
	log.Debug(content)
}

func Debugf(content string) {
	Instance("debugf")
	log.Debugf(content)
}

func Info(content string) {
	Instance("info")
	log.Info(content)
}

func Error(content string) {
	Instance("error")
	log.Error(content)
}

func Warning(content string) {
	Instance("warning")
	log.Warning(content)
}

func Critical(content string) {
	Instance("critical")
	log.Critical(content)
}

func Request(content string) {
	Instance("request")
	log.Info(content)
}

func Response(content string) {
	Instance("response")
	log.Info(content)
}

func Instance(logType string) {
	var timeStr = helper.GetTmStr(t, "Y-m-d")
	var files = "./tmp/" + logType + "-" + timeStr + ".txt"
	helper.CreateFile(files)
	logFile, err := os.OpenFile(files, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)
}
