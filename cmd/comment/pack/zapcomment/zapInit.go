package zapcomment

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Init() {
	var coreArr []zapcore.Core

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConsole := zapcore.NewConsoleEncoder(encoderConfig)

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	encoderFile := zapcore.NewConsoleEncoder(encoderConfig)

	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})
	allPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.FatalLevel && lev >= zap.DebugLevel
	})

	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/info.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	})
	infoFileCore := zapcore.NewCore(encoderFile, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer), lowPriority)

	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/error.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	})
	errorFileCore := zapcore.NewCore(encoderFile, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer), highPriority)

	consoleCore := zapcore.NewCore(encoderConsole, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), allPriority)

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	coreArr = append(coreArr, consoleCore)

	Logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())

	Logger.Info("the log system started successfully")
}
