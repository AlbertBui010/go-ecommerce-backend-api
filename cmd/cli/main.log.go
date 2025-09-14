package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 1.
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d, abc: %s", "albert", 15, "hello world")

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "albert"), zap.Int("age", 40))

	// 2.
	// logger := zap.NewExample()
	// logger.Info("Hello NewExamle") // {"level":"info","msg":"Hello NewExamle"}

	// // Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment") // 2025-09-14T11:21:19.123+0700    INFO    cli/main.log.go:17    Hello NewDevelopment

	// // Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction") // {"level":"info","ts":1757823679.1244318,"caller":"cli/main.log.go:21","msg":"Hello NewProduction"}

	// 3. Format logs
	encoder := getEncoderLog()
	writeSyncer := getWriterSync()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Test", zap.String("Name", "Line 1"), zap.Int("Age", 16))
	logger.Error("Test", zap.String("Name", "Line 2"), zap.Int("Age", 16))

	// {"level":"INFO","time":"2025-09-14T11:55:28.549+0700","caller":"cli/main.log.go:36","msg":"Test","Name":"Line 1","Age":16}
	// {"level":"ERROR","time":"2025-09-14T11:55:28.549+0700","caller":"cli/main.log.go:37","msg":"Test","Name":"Line 2","Age":16}
}

// Format logs
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1757823679.1244318 --> 2025-09-14T11:21:19.123+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts --> time
	encodeConfig.TimeKey = "time"

	// info --> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.log.go:21"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

// Config read/write files
func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
