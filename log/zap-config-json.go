package main

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var logger *zap.SugaredLogger
var Glogger = GetLogger("ahui")

// type LoggerX struct {
// 	logger *zap.SugaredLogger
// 	name   string
// }

// func (lx *LoggerX) Debug(data interface{}) {
// 	lx.logger.Debugw(lx.name, "log", data)
// }
// func (lx *LoggerX) Info(data interface{}) {
// 	lx.logger.Infow(lx.name, "log", data)
// }

func GetLogger(name string) *zap.SugaredLogger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		// EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	config := zap.Config{
		Level:            atom,
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		InitialFields:    map[string]interface{}{},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	loggerRaw, _ := config.Build()
	logger := loggerRaw.Named(name).Sugar()
	// logger.core.callerSkip += 2
	return logger
}

func JsonEncode(data interface{}) string {
	jsonBytes, _ := json.Marshal(data)
	return (string(jsonBytes))
}

func main() {
	lx := GetLogger("root")
	lx.Info("字符")
	data := (map[string]interface{}{"key": "value"})
	lx.Info(JsonEncode(data))
	Glogger.Info(JsonEncode(data))
}
