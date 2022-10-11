package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func StrToTime(s string) (int64, error) {
	// delete all not int characters
	if s == `` {
		return time.Now().Unix(), nil
	}
	r := make([]rune, 14)
	l := 0
	// 过滤除数字以外的字符
	for _, v := range s {
		if '0' <= v && v <= '9' {
			r[l] = v
			l++
			if l == 14 {
				break
			}
		}
	}
	for l < 14 {
		r[l] = '0' // 补0
		l++
	}
	t, err := time.Parse(`20060102150405`, string(r))
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func TimeStdStr() string {
	return time.Now().Format(`2006-01-02 15:04:05`)
}

func TimeShortStr() string {
	return time.Now().Format(`20060102150405`)
}

func TimeToStr(unixSecTime interface{}, layout ...string) string {
	i := AnyToInt64(unixSecTime)
	if i == 0 {
		return ""
	}
	f := "2006-01-02 15:04:05"
	if len(layout) > 0 {
		f = layout[0]
	}
	return time.Unix(i, 0).Format(f)
}

func TimeFmtNanoUnix() string {
	return strings.Replace(time.Now().Format(`20060102150405.0000000`), ".", "", 1)
}

func TimeParse(format, src string) (time.Time, error) {
	return time.ParseInLocation(format, src, time.Local)
}

func TimeParseStd(src string) time.Time {
	t, _ := TimeParse("2006-01-02 15:04:05", src)
	return t
}

func TimeStdParseUnix(src string) int64 {
	t, err := TimeParse("2006-01-02 15:04:05", src)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 获取一个当前时间 时间间隔 时间戳
func TimeInterval(unit string, amount int) (startTime, endTime int64) {
	t := time.Now()
	nowTime := t.Unix()
	tmpTime := int64(0)
	switch unit {
	case "years":
		tmpTime = time.Date(t.Year()+amount, t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location()).Unix()
	case "months":
		tmpTime = time.Date(t.Year(), t.Month()+time.Month(amount), t.Day(), t.Hour(), 0, 0, 0, t.Location()).Unix()
	case "days":
		tmpTime = time.Date(t.Year(), t.Month(), t.Day()+amount, t.Hour(), 0, 0, 0, t.Location()).Unix()
	case "hours":
		tmpTime = time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+amount, 0, 0, 0, t.Location()).Unix()
	}
	if amount > 0 {
		startTime = nowTime
		endTime = tmpTime
	} else {
		startTime = tmpTime
		endTime = nowTime
	}
	return
}

// 时分秒字符串转时间戳，传入示例：8:40 or 8:40:10
func TimeHmsToUnix(str string) (int64, error) {
	t := time.Now()
	arr := strings.Split(str, ":")
	if len(arr) < 2 {
		return 0, errors.New("Time format error")
	}
	h, _ := strconv.Atoi(arr[0])
	m, _ := strconv.Atoi(arr[1])
	s := 0
	if len(arr) == 3 {
		s, _ = strconv.Atoi(arr[3])
	}
	formatted1 := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), h, m, s)
	res, err := time.ParseInLocation("20060102150405", formatted1, time.Local)
	if err != nil {
		return 0, err
	} else {
		return res.Unix(), nil
	}
}
