package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logi *zap.Logger

func LogSetting(app *fiber.App) {
	setHttpLogSetting(app)
	setCustomLogSetting()
}

func setHttpLogSetting(app *fiber.App) {
	// HTTP Setting
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format:   "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
		TimeZone: "Asia/Seoul",
	}))
}

func setCustomLogSetting() {
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	// Cusomt Logging Setting
	_log, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logi = _log
}

func Info(message string, fields ...zap.Field) {
	logi.Info(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logi.Error(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	logi.Debug(message, fields...)
}
