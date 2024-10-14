package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	core := getCore()
	logger := zap.New(core)
	logger.Info("Logger initialized")
	Logger = logger
}

func getCore() zapcore.Core {
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// 创建文件写入器
	file, err := os.Create("./logs/blog-service.log")
	if err != nil {
		log.Fatalf("Error creating log file: %s", err)
	}

	// 创建多重写入器
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(file),
	)

	// 使用多重写入器创建核心
	core := zapcore.NewCore(encoder, multiWriteSyncer, zap.InfoLevel)
	return core
}
