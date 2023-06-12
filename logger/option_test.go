package logger

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"testing"

	"github.com/onqlavelabs/onqlave.core/logger/filter"
	"gopkg.in/go-playground/assert.v1"
)

func TestWithFilter(t *testing.T) {
	loggerWithFilter := NewLog(
		"TestWithFilter",
		WithFilters(filter.Field("test"), filter.Value("with"), filter.ValueRegex("filter")),
	)

	assert.Equal(t, len(loggerWithFilter.filters), 3)
	assert.Equal(t, fmt.Sprintf("%T", loggerWithFilter.filters[0]), "*filter.FieldFilter")
	assert.Equal(t, fmt.Sprintf("%T", loggerWithFilter.filters[1]), "*filter.ValueFilter")
	assert.Equal(t, fmt.Sprintf("%T", loggerWithFilter.filters[2]), "*filter.ValueRegexFilter")
}

func TestWithCid(t *testing.T) {
	loggerWithoutCid := NewLog("TestWithFilter", WithCid(""))
	loggerWithCid := NewLog("TestWithFilter", WithCid("custom-correlation-id"))

	assert.NotEqual(t, loggerWithoutCid.correlationID, "")
	assert.Equal(t, loggerWithCid.correlationID, "custom-correlation-id")
}

func TestWithTimeEncoder(t *testing.T) {
	loggerTimeEncoder := NewLog("TestWithTimeEncoder", WithTimeEncoder(zapcore.RFC3339TimeEncoder))

	assert.Equal(t, fmt.Sprintf("%T", loggerTimeEncoder.encodeTime), "zapcore.TimeEncoder")
}

func TestWithDurationEncoder(t *testing.T) {
	loggerDurationEncoder := NewLog("TestWithDurationEncoder", WithDurationEncoder(zapcore.NanosDurationEncoder))

	assert.Equal(t, fmt.Sprintf("%T", loggerDurationEncoder.encodeDuration), "zapcore.DurationEncoder")
}

func TestWithDevelopment(t *testing.T) {
	loggerDevelopment := NewLog("loggerDevelopment", WithDevelopment(true))

	assert.Equal(t, loggerDevelopment.isDevelopment, true)
}

func TestWithLevel(t *testing.T) {
	loggerDebug := NewLog("loggerDebug", WithLevel(LevelDebug))
	loggerInfo := NewLog("loggerInfo")
	loggerWarn := NewLog("loggerWarn", WithLevel(LevelWarn))

	assert.Equal(t, loggerDebug.level, zapcore.DebugLevel)
	assert.Equal(t, loggerInfo.level, zapcore.InfoLevel)
	assert.Equal(t, loggerWarn.level, zapcore.WarnLevel)
}
