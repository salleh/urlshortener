package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger
var AppLogger *zap.SugaredLogger
var ZapLevel zapcore.Level
var AppZapConfig zapcore.EncoderConfig
var AppZapCore zapcore.Core
var AppZapFileEncoder zapcore.Encoder
var AppZapConsoleEncoder zapcore.Encoder
var AppZapWriter zapcore.WriteSyncer

func init() {
	InitializeLogger()
}

func InitializeLogger() {
	AppZapConfig = zap.NewProductionEncoderConfig()
	AppZapConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	AppZapFileEncoder = zapcore.NewJSONEncoder(AppZapConfig)
	AppZapConsoleEncoder = zapcore.NewConsoleEncoder(AppZapConfig)
	//logFile, _ := os.OpenFile("cid-log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//writer := zapcore.AddSync(logFile)
	AppZapWriter = zapcore.AddSync(&lumberjack.Logger{
		Filename:   "otp.log.json",
		MaxSize:    128, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	})

	ZapLevel = zapcore.DebugLevel
	AppZapCore = zapcore.NewTee(
		zapcore.NewCore(AppZapFileEncoder, AppZapWriter, ZapLevel),
		zapcore.NewCore(AppZapConsoleEncoder, zapcore.AddSync(os.Stdout), ZapLevel),
	)
	Logger = zap.New(AppZapCore, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	AppLogger = Logger.Sugar()
}

func ChangeLogLevel(level string) {
	if ZapLevel.String() == level {
		return
	}

	switch level {
	case "debug":
		ZapLevel = zapcore.DebugLevel
	case "info":
		ZapLevel = zapcore.InfoLevel
	case "warn":
		ZapLevel = zapcore.WarnLevel
	case "error":
		ZapLevel = zapcore.ErrorLevel
	case "panic":
		ZapLevel = zapcore.PanicLevel
	case "fatal":
		ZapLevel = zapcore.FatalLevel
	default:
		ZapLevel = zapcore.InfoLevel
	}

	AppLogger.Infof("Changing log level to %v", ZapLevel.String())

	AppZapCore = zapcore.NewTee(
		zapcore.NewCore(AppZapFileEncoder, AppZapWriter, ZapLevel),
		zapcore.NewCore(AppZapConsoleEncoder, zapcore.AddSync(os.Stdout), ZapLevel),
	)
	Logger = zap.New(AppZapCore, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	AppLogger = Logger.Sugar()
}
