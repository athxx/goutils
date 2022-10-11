package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ToString(raw interface{}, e error) (res string) {
	if e != nil {
		return ""
	}
	return AnyToString(raw)
}

func ToInt64(raw interface{}, e error) int64 {
	if e != nil {
		return 0
	}
	return AnyToInt64(raw)
}

func AnyToBool(raw interface{}) bool {
	switch i := raw.(type) {
	case float32, float64, int, int64, uint, uint8, uint16, uint32, uint64, int8, int16, int32:
		return i != 0
	case []byte:
		return i != nil
	case string:
		if i == "false" {
			return false
		}
		return i != ``
	case error:
		return false
	case nil:
		return true
	}
	val := fmt.Sprint(raw)
	val = strings.TrimLeft(val, "&")
	if strings.TrimLeft(val, "{}") == `` {
		return false
	}
	if strings.TrimLeft(val, "[]") == `` {
		return false
	}
	// ptr type
	b, err := json.Marshal(raw)
	if err != nil {
		return false
	}
	if strings.TrimLeft(string(b), "\"\"") == `` {
		return false
	}
	if strings.TrimLeft(string(b), "{}") == `` {
		return false
	}
	return true
}

func AnyToInt64(raw interface{}) int64 {
	switch i := raw.(type) {
	case string:
		res, _ := strconv.ParseInt(i, 10, 64)
		return res
	case []byte:
		return BytesToInt64(i)
	case int:
		return int64(i)
	case int64:
		return i
	case uint:
		return int64(i)
	case uint8:
		return int64(i)
	case uint16:
		return int64(i)
	case uint32:
		return int64(i)
	case uint64:
		return int64(i)
	case int8:
		return int64(i)
	case int16:
		return int64(i)
	case int32:
		return int64(i)
	case float32:
		return int64(i)
	case float64:
		return int64(i)
	case error:
		return 0
	case bool:
		if i {
			return 1
		}
		return 0
	}
	return 0
}

func AnyToString(raw interface{}) string {
	switch i := raw.(type) {
	case []byte:
		return string(i)
	case int:
		return strconv.FormatInt(int64(i), 10)
	case int64:
		return strconv.FormatInt(i, 10)
	case float32:
		return Float64ToStr(float64(i))
	case float64:
		return Float64ToStr(i)
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
	case string:
		return i
	case error:
		return i.Error()
	case bool:
		return strconv.FormatBool(i)
	}
	return fmt.Sprintf("%#v", raw)
}

func AnyToFloat64(raw interface{}) float64 {
	switch i := raw.(type) {
	case []byte:
		f, _ := strconv.ParseFloat(string(i), 64)
		return f
	case int:
		return float64(i)
	case int64:
		return float64(i)
	case float32:
		return float64(i)
	case float64:
		return i
	case uint:
		return float64(i)
	case uint8:
		return float64(i)
	case uint16:
		return float64(i)
	case uint32:
		return float64(i)
	case uint64:
		return float64(i)
	case int8:
		return float64(i)
	case int16:
		return float64(i)
	case int32:
		return float64(i)
	case string:
		f, _ := strconv.ParseFloat(i, 64)
		return f
	case bool:
		if i {
			return 1
		}
	}
	return 0
}

func ToByte(raw interface{}, e error) []byte {
	if e != nil {
		return []byte{}
	}
	switch i := raw.(type) {
	case string:
		return []byte(i)
	case int:
		return Int64ToBytes(int64(i))
	case int64:
		return Int64ToBytes(i)
	case float32:
		return Float32ToByte(i)
	case float64:
		return Float64ToByte(i)
	case uint:
		return Int64ToBytes(int64(i))
	case uint8:
		return Int64ToBytes(int64(i))
	case uint16:
		return Int64ToBytes(int64(i))
	case uint32:
		return Int64ToBytes(int64(i))
	case uint64:
		return Int64ToBytes(int64(i))
	case int8:
		return Int64ToBytes(int64(i))
	case int16:
		return Int64ToBytes(int64(i))
	case int32:
		return Int64ToBytes(int64(i))
	case []byte:
		return i
	case error:
		return []byte(i.Error())
	case bool:
		if i {
			return []byte("true")
		}
		return []byte("false")
	}
	return []byte(fmt.Sprintf("%#v", raw))
}

func Int32ToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func StrToInt(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func StrToInt64(s string) int64 {
	res, _ := strconv.ParseInt(s, 10, 64)
	return res
}

func StrToUint64(s string) uint64 {
	res, _ := strconv.ParseUint(s, 10, 64)
	return res
}

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}
func Float64ToStrPrec1(f float64) string {
	return strconv.FormatFloat(f, 'f', 1, 64)
}

func Float32ToStr(f float32) string {
	return Float64ToStr(float64(f))
}

func StrToFloat64(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return res
}

func StrToFloat32(s string) float32 {
	res, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}
	return float32(res)
}

func StrToBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

func BoolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func FloatToInt64(f float64) int64 {
	return int64(f)
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Uint64ToStr(i uint64) string {
	return strconv.FormatUint(i, 10)
}
