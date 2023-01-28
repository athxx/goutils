package logx

import (
	"errors"
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

type LogX struct {
	logger    *zap.Logger
	atomLevel *zap.AtomicLevel
}

type Fields map[string]interface{}

// 判断其他类型--start
func getFields(msg string, format bool, args ...interface{}) (string, []zap.Field) {
	var str []interface{}
	var fields []zap.Field
	if len(args) > 0 {
		for _, v := range args {
			if f, ok := v.(zap.Field); ok {
				fields = append(fields, f)
			} else if f, ok := v.(Fields); ok {
				fields = append(fields, formatFieldMap(f)...)
			} else {
				str = append(str, ` `+AnyToString(v))
			}
		}
		if format {
			return fmt.Sprintf(msg, str...), fields
		}
		str = append([]interface{}{msg}, str...)
		return fmt.Sprint(str...), fields
	}
	return msg, []zap.Field{}
}

func (l *LogX) Debug(s interface{}, args ...interface{}) error {
	e := checkErr(s)
	if e != nil {
		msg, field := getFields(e.Error(), false, args...)
		l.logger.Debug(msg, field...)
		return errors.New(msg)
	}
	return e
}
func (l *LogX) Info(s interface{}, args ...interface{}) error {
	e := checkErr(s)
	if e != nil {
		msg, field := getFields(e.Error(), false, args...)
		l.logger.Info(msg, field...)
	}
	return e
}
func (l *LogX) Warn(s interface{}, args ...interface{}) error {
	e := checkErr(s)
	if e != nil {
		msg, field := getFields(e.Error(), false, args...)
		l.logger.Warn(msg, field...)
	}
	return e
}
func (l *LogX) Error(s interface{}, args ...interface{}) error {
	e := checkErr(s)
	if e != nil {
		msg, field := getFields(e.Error(), false, args...)
		l.logger.Error(msg, field...)
	}
	return e
}
func (l *LogX) DPanic(s interface{}, args ...interface{}) {
	e := checkErr(s)
	if e != nil {
		msg, field := getFields(e.Error(), false, args...)
		l.logger.DPanic(msg, field...)
	}
}
func (l *LogX) Panic(s interface{}, args ...interface{}) {
	e := checkErr(s)
	if e != nil {
		msg, field := getFields(e.Error(), false, args...)
		l.logger.Panic(msg, field...)
	}
}
func (l *LogX) Fatal(s interface{}, args ...interface{}) {
	e := checkErr(s)
	if e != nil {
		msg, field := getFields(e.Error(), false, args...)
		l.logger.Fatal(msg, field...)
	}
}

func checkErr(s interface{}) error {
	var msg string
	switch e := s.(type) {
	case error:
		return e
	case string:
		return errors.New(e)
	case []byte:
		return errors.New(string(e))
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		msg = fmt.Sprintf("%v", e)
		return errors.New(msg)
	case nil:
		return nil
	default:
		return errors.New(fmt.Sprintf("%+v", e))
	}
}

func (l *LogX) LogError(err error) error {
	return l.Error(err.Error())
}

func (l *LogX) Debugf(msg string, args ...interface{}) error {
	s, f := getFields(msg, true, args...)
	l.logger.Debug(s, f...)
	return errors.New(s)
}

func (l *LogX) Infof(msg string, args ...interface{}) error {
	s, f := getFields(msg, true, args...)
	l.logger.Info(s, f...)
	return errors.New(s)
}

func (l *LogX) Warnf(msg string, args ...interface{}) error {
	s, f := getFields(msg, true, args...)
	l.logger.Warn(s, f...)
	return errors.New(s)
}

func (l *LogX) Errorf(msg string, args ...interface{}) error {
	s, f := getFields(msg, true, args...)
	l.logger.Error(s, f...)
	return errors.New(s)
}

func (l *LogX) DPanicf(msg string, args ...interface{}) error {
	s, f := getFields(msg, true, args...)
	l.logger.DPanic(s, f...)
	return errors.New(s)
}

func (l *LogX) Panicf(msg string, args ...interface{}) {
	s, f := getFields(msg, true, args...)
	l.logger.Panic(s, f...)
}

func (l *LogX) Fatalf(msg string, args ...interface{}) {
	s, f := getFields(msg, true, args...)
	l.logger.Fatal(s, f...)
}

func AnyToString(raw interface{}) string {
	switch i := raw.(type) {
	case error:
		return i.Error()
	case string:
		return i
	case []byte:
		return string(i)
	case int:
		return strconv.FormatInt(int64(i), 10)
	case int64:
		return strconv.FormatInt(i, 10)
	case uint:
		return strconv.FormatInt(int64(i), 10)
	case uint8:
		return strconv.FormatInt(int64(i), 10)
	case uint16:
		return strconv.FormatInt(int64(i), 10)
	case uint32:
		return strconv.FormatInt(int64(i), 10)
	case uint64:
		return strconv.FormatInt(int64(i), 10)
	case int8:
		return strconv.FormatInt(int64(i), 10)
	case int16:
		return strconv.FormatInt(int64(i), 10)
	case int32:
		return strconv.FormatInt(int64(i), 10)
	case float32:
		return strconv.FormatFloat(float64(i), 'f', 2, 64)
	case float64:
		return strconv.FormatFloat(i, 'f', 2, 64)
	}
	return fmt.Sprintf("%#v", raw)
}
