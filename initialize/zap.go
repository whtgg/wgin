package initialize

import (
	"bytes"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
	"wgin/global"
)

func init() {
	global.Logger = NewZap()
}

// NewZap 初始化日志和日志切割配置
func NewZap() *zap.Logger {
	logger := zap.New(zapcore.NewTee(new(wZap).GetZapCors()...))
	return logger
}

type wZap struct{}

func (z *wZap) GetZapCors() []zapcore.Core {
	cores := make([]zapcore.Core, 0)
	for level := GetConfigLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelEnablerFunc(level)))
	}
	return cores
}

// EncoderConfig 定义日志格式
func (z *wZap) EncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     z.GetTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// GetTimeEncoder 设置日志日期显示格式
func (z *wZap) GetTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	buf := bytes.NewBufferString("")
	buf.WriteString(t.Format("2006/01/02 - 15:04:05.000"))
	enc.AppendString(buf.String())
}

// GetEncoderCore 配置日志切割
func (z *wZap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	wLogs, err := rotateLogs.New("WG_log.%Y%m%d")
	if err != nil {
		panic(err)
	}
	ws := zapcore.NewMultiWriteSyncer(zapcore.AddSync(wLogs), zapcore.AddSync(os.Stdout))
	return zapcore.NewCore(z.GetEncoder(), ws, level)
}

// GetEncoder 设置日志打印格式
func (z *wZap) GetEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(z.EncoderConfig())
}

func (z *wZap) GetLevelEnablerFunc(lv zapcore.Level) zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level == lv
	}
}

func GetConfigLevel() (level zapcore.Level) {
	Level := strings.ToLower("info")
	level, _ = zapcore.ParseLevel(Level)
	return
}
