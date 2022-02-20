package logger

import (
	"os"
	"time"

	"example.com/portto/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func InitLogger(cfg *config.LogConfig) (err error) {

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = SyslogTimeEncoder
	encoderCfg.EncodeLevel = CustomLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	// use json formate
	// encoder := zapcore.NewJSONEncoder(encoderCfg)

	var l = new(zapcore.Level)
	if err = l.UnmarshalText([]byte(cfg.Level)); err != nil {
		return
	}

	consoleOut := zapcore.Lock(os.Stdout)

	core := zapcore.NewCore(encoder, consoleOut, l)
	Logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Logger)

	return
}
