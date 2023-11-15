package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	debugSLogger    *zap.SugaredLogger
	infoSLogger     *zap.SugaredLogger
	warnSLogger     *zap.SugaredLogger
	errorSLogger    *zap.SugaredLogger
	analysisSLogger *zap.SugaredLogger

	DebugLogger    *zap.Logger
	InfoLogger     *zap.Logger
	WarnLogger     *zap.Logger
	ErrorLogger    *zap.Logger
	AnalysisLogger *zap.Logger
)

func init() {
	debugPath := "./logs/debug.log"
	infoPath := "./logs/info.log"
	warnPath := "./logs/warn.log"
	errorPath := "./logs/rerror.log"
	analysisPath := "./logs/analysis.log"

	DebugLogger = zap.New(zapcore.NewTee(
		zapcore.NewCore(ConsoleEncoder(), ConsoleWriter(), zapcore.InfoLevel),
		zapcore.NewCore(FileEncoder(), FileWriter(debugPath), zapcore.DebugLevel),
	), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(1))

	InfoLogger = zap.New(zapcore.NewTee(
		zapcore.NewCore(ConsoleEncoder(), ConsoleWriter(), zapcore.InfoLevel),
		zapcore.NewCore(FileEncoder(), FileWriter(infoPath), zapcore.InfoLevel),
	), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(1))

	WarnLogger = zap.New(zapcore.NewTee(
		zapcore.NewCore(ConsoleEncoder(), ConsoleWriter(), zapcore.WarnLevel),
		zapcore.NewCore(FileEncoder(), FileWriter(warnPath), zapcore.WarnLevel),
	), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(1))

	ErrorLogger = zap.New(zapcore.NewTee(
		zapcore.NewCore(ConsoleEncoder(), ConsoleWriter(), zapcore.ErrorLevel),
		zapcore.NewCore(FileEncoder(), FileWriter(errorPath), zapcore.ErrorLevel),
	), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(1))

	AnalysisLogger = zap.New(zapcore.NewTee(
		zapcore.NewCore(ConsoleEncoder(), ConsoleWriter(), zapcore.DebugLevel),
		zapcore.NewCore(FileEncoder(), FileWriter(analysisPath), zapcore.DebugLevel),
	), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(1))

	debugSLogger = DebugLogger.Sugar()
	infoSLogger = InfoLogger.Sugar()
	warnSLogger = WarnLogger.Sugar()
	errorSLogger = ErrorLogger.Sugar()
	analysisSLogger = AnalysisLogger.Sugar()
}

func encodeConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // format: 2006-01-02T15:04:05.000Z0700
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return encoderConfig
}

func ConsoleEncoder() zapcore.Encoder {
	c := encodeConfig()
	c.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(c)
}

func FileEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(encodeConfig())
}

func ConsoleWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func FileWriter(path string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	})
}

// for compatibility

func InfoF(template string, args ...interface{}) {
	infoSLogger.Infof(template, args...)
}

func DebugF(template string, args ...interface{}) {
	debugSLogger.Debugf(template, args...)
}

func WarnF(template string, args ...interface{}) {
	warnSLogger.Warnf(template, args...)
}

func ErrorF(template string, args ...interface{}) {
	errorSLogger.Errorf(template, args...)
}

func AnalysisF(template string, args ...interface{}) {
	analysisSLogger.Debugf(template, args...)
}

func Debug(s string) {
	debugSLogger.Debug(s)
}

func Info(s string) {
	infoSLogger.Info(s)
}

func Warn(s string) {
	warnSLogger.Warn(s)
}

func Error(s string) {
	errorSLogger.Error(s)
}

func Analysis(s string) {
	analysisSLogger.Debug(s)
}
