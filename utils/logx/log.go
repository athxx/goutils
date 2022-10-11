package logx

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	AppName         string `yaml:"app_name" json:"app_name" toml:"app_name"`
	Level           string `yaml:"level" json:"level" toml:"level"`
	StacktraceLevel string `yaml:"stacktrace_level" json:"stacktrace_level" toml:"stacktrace_level"`
	IsStdOut        bool   `yaml:"is_stdout" json:"is_stdout" toml:"is_stdout"`
	IsStdErr        bool   `yaml:"is_stderr" json:"is_stderr" toml:"is_stderr"`
	TimeFormat      string `yaml:"time_format" json:"time_format" toml:"time_format"` // second, milli, nano, standard, iso,
	Encoding        string `yaml:"encoding" json:"encoding" toml:"encoding"`          // console, json
	Skip            int    `yaml:"skip" json:"skip" toml:"skip"`

	IsFileOut     bool   `yaml:"is_file_out" json:"is_file_out" toml:"is_file_out"`
	FileDir       string `yaml:"file_dir" json:"file_dir" toml:"file_dir"`
	FileName      string `yaml:"file_name" json:"file_name" toml:"file_name"`
	FileMaxSize   int    `yaml:"file_max_size" json:"file_max_size" toml:"file_max_size"` // MB
	FileMaxAge    int    `yaml:"file_max_age" json:"file_max_age" toml:"file_max_age"`    // Days
	FileMaxBackup int    `yaml:"file_max_backup" json:"file_max_backup" toml:"file_max_backup"`
}

var l = defaultLogger()

// defaultLogger default logger setting
func defaultLogger() *LogX {
	conf := &LogConfig{
		Level:           "debug",
		StacktraceLevel: "error",
		IsStdOut:        true,
		TimeFormat:      "standard",
		Encoding:        "console",
		Skip:            2,
	}
	writers := []zapcore.WriteSyncer{os.Stdout}
	lg, lv := newZapLogger(setLogLevel(conf.Level), setLogLevel(conf.StacktraceLevel), conf.Encoding, conf.TimeFormat, conf.Skip, zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(lg)
	return &LogX{logger: lg, atomLevel: lv}
}

// InitDefaultLogger initial standard log, if you don't init, it will use default logger setting
func InitDefaultLogger(cfg *LogConfig) {
	var writers []zapcore.WriteSyncer
	if cfg.IsStdOut || (!cfg.IsStdOut && !cfg.IsStdErr && !cfg.IsFileOut) {
		writers = []zapcore.WriteSyncer{os.Stdout}
	}

	if cfg.IsStdErr {
		writers = append(writers, os.Stderr)
	}

	if cfg.IsFileOut {
		writers = append(writers, NewRollingFile(cfg.FileDir, cfg.FileName, cfg.FileMaxSize, cfg.FileMaxAge, cfg.FileMaxBackup))
	}

	lg, lv := newZapLogger(setLogLevel(cfg.Level), setLogLevel(cfg.StacktraceLevel), cfg.Encoding, cfg.TimeFormat, cfg.Skip, zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(lg)
	if cfg.AppName != `` {
		lg = lg.With(zap.String("app", cfg.AppName)) // 加上应用名称
	}
	l = &LogX{logger: lg, atomLevel: lv}
}

// NewLogger create a new logger
func NewLogger(cfg *LogConfig) *LogX {
	var writers []zapcore.WriteSyncer
	if cfg.IsStdOut || (!cfg.IsStdOut && !cfg.IsStdErr && !cfg.IsFileOut) {
		writers = []zapcore.WriteSyncer{os.Stdout, os.Stderr}
	}

	if cfg.IsStdErr {
		writers = append(writers, os.Stderr)
	}

	if cfg.IsFileOut {
		writers = append(writers, NewRollingFile(cfg.FileDir, cfg.FileName, cfg.FileMaxSize, cfg.FileMaxAge, cfg.FileMaxBackup))
	}

	lg, lv := newZapLogger(setLogLevel(cfg.Level), setLogLevel(cfg.StacktraceLevel), cfg.Encoding, cfg.TimeFormat, cfg.Skip, zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(lg)
	if cfg.AppName != `` {
		lg = lg.With(zap.String("app", cfg.AppName)) // 加上应用名称
	}
	return &LogX{logger: lg, atomLevel: lv}
}

// create a new zaplog logger
func newZapLogger(level, stacktrace zapcore.Level, encoding, timeType string, skip int, output zapcore.WriteSyncer) (*zap.Logger, *zap.AtomicLevel) {
	encCfg := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "name",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stack",
		//LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
	}
	setTimeFormat(timeType, &encCfg) // set time type
	atmLvl := zap.NewAtomicLevel()   // set level
	atmLvl.SetLevel(level)
	encoder := zapcore.NewJSONEncoder(encCfg) // 确定encoder格式
	if encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encCfg)
	}
	if skip < 1 {
		skip = 2
	}
	return zap.New(zapcore.NewCore(encoder, output, atmLvl), zap.AddCaller(), zap.AddStacktrace(stacktrace), zap.AddCallerSkip(skip)), &atmLvl
}

// set log level
func setLogLevel(lvl string) zapcore.Level {
	lv := map[string]zapcore.Level{
		"panic": zapcore.PanicLevel,
		"fatal": zapcore.FatalLevel,
		"error": zapcore.ErrorLevel,
		"warn":  zapcore.WarnLevel,
		"info":  zapcore.InfoLevel,
	}
	if v, ok := lv[strings.ToLower(lvl)]; ok {
		return v
	}
	return zapcore.DebugLevel
}

// set time format
func setTimeFormat(timeType string, z *zapcore.EncoderConfig) {
	encodeTime := map[string]zapcore.TimeEncoder{
		"iso":    zapcore.ISO8601TimeEncoder,
		"sec":    func(t time.Time, enc zapcore.PrimitiveArrayEncoder) { enc.AppendInt64(t.Unix()) },
		"second": zapcore.EpochTimeEncoder,
		"milli":  zapcore.EpochMillisTimeEncoder,
		"nano":   zapcore.EpochNanosTimeEncoder,
	}
	z.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	if v, ok := encodeTime[strings.ToLower(timeType)]; ok {
		z.EncodeTime = v
	}
}

func GetLevel() string {
	lv := map[zapcore.Level]string{
		zapcore.FatalLevel: "fatal",
		zapcore.PanicLevel: "panic",
		zapcore.ErrorLevel: "error",
		zapcore.WarnLevel:  "warn",
		zapcore.InfoLevel:  "info",
	}
	if v, ok := lv[l.atomLevel.Level()]; ok {
		return v
	}
	return "debug"
}

func SetLevel(lvl string) {
	l.atomLevel.SetLevel(setLogLevel(lvl))
}

// AddCallerSkip temporary add call skip
func AddCallerSkip(skip int) *LogX {
	l.logger.WithOptions(zap.AddCallerSkip(skip))
	return l
}

// AddDepth permanent add call skip
func AddDepth(skip int) *LogX {
	l.logger = l.logger.WithOptions(zap.AddCallerSkip(skip))
	return l
}

// AddOptions permanent add options
func AddOptions(opts ...zap.Option) *LogX {
	l.logger = l.logger.WithOptions(opts...)
	return l
}

func AddField(k string, v interface{}) {
	l.logger.With(zap.Any(k, v))
}

func AddFields(fields map[string]interface{}) *LogX {
	for k, v := range fields {
		l.logger.With(zap.Any(k, v))
	}
	return l
}

func Bool(e interface{}) bool {
	if e == nil {
		return true
	}
	Warn(e)
	return false
}

// Normal log

func Debug(e interface{}, args ...interface{}) error {
	return l.Debug(e, args...)
}
func Info(e interface{}, args ...interface{}) error {
	return l.Info(e, args...)
}
func Warn(e interface{}, args ...interface{}) error {
	return l.Warn(e, args...)
}
func Error(e interface{}, args ...interface{}) error {
	return l.Error(e, args...)
}
func Panic(e interface{}, args ...interface{}) {
	l.Panic(e, args...)
}
func Fatal(e interface{}, args ...interface{}) {
	l.Fatal(e, args...)
}

// Format logs

func Debugf(format string, args ...interface{}) error {
	return l.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) error {
	return l.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) error {
	return l.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) error {
	return l.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

func formatFieldMap(m Fields) []zap.Field {
	var res []zap.Field
	for k, v := range m {
		res = append(res, zap.Any(k, v))
	}
	return res
}
