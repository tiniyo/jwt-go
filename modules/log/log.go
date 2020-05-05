package log

import (
	lg "github.com/labstack/gommon/log"
)

type Logger struct {
	*lg.Logger
}

var (
	global        = lg.New("log")
	defaultHeader = `{"time":"${time_rfc3339}","level":"${level}","prefix":"${prefix}",` +
		`"file":"${long_file}","line":"${line}"}`
)

func init() {
	lg.SetHeader(defaultHeader)
	lg.SetLevel(lg.DEBUG)
	global.SetHeader(defaultHeader)
	global.SetLevel(lg.DEBUG)
}

func SetLevel(v lg.Lvl) {
	lg.SetLevel(v)
	global.SetLevel(v)
}

func Debug(i ...interface{}) {
	global.Debug(i)
}

func Debugf(format string, values ...interface{}) {
	global.Debugf(format, values...)
}

func Infoj(j lg.JSON) {
	global.Infoj(j)
}

func Errorj(j lg.JSON) {
	global.Errorj(j)
}

func Info(i ...interface{}) {
	global.Info(i)
}

func Infof(format string, values ...interface{}) {
	global.Infof(format, values...)
}

func Warn(i ...interface{}) {
	global.Warn(i)
}

func Warnf(format string, values ...interface{}) {
	global.Warnf(format, values...)
}

func Error(i ...interface{}) {
	global.Error(i)
}

func Errorf(format string, values ...interface{}) {
	global.Errorf(format, values...)
}

func Fatal(i ...interface{}) {
	global.Fatal(i)
}

func Fatalf(format string, values ...interface{}) {
	global.Fatalf(format, values...)
}

func Panic(i ...interface{}) {
	global.Panic(i)
}

func Panicf(format string, args ...interface{}) {
	global.Panicf(format, args)
}

