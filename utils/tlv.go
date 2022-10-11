package utils

//--------------------------------- TLV: tag length value / type length value ---------------------------------

type TagValue struct {
	Tag   uint8
	Value []byte
}

func TlvEncode(tag uint8, data []byte) []byte {
	var res []byte
	res = append(res, tag)
	v := len(data)
	b := make([]byte, 4)
	b[3] = byte(v)
	b[2] = byte(v >> 8)
	b[1] = byte(v >> 16)
	b[0] = byte(v >> 24)
	res = append(res, b...)
	return append(res, data...)
}

func TlvDecode(data []byte) (tag uint8, b []byte) {
	tag = data[0]
	v := uint32(data[4]) | uint32(data[3])<<8 | uint32(data[2])<<16 | uint32(data[1])<<24
	return tag, data[5 : 5+v]
}

func TlvDecodeAll(data []byte) []*TagValue {
	var res []*TagValue
	for {
		v := uint32(data[4]) | uint32(data[3])<<8 | uint32(data[2])<<16 | uint32(data[1])<<24
		if len(data) < 5+int(v) {
			break
		}
		res = append(res, &TagValue{Tag: data[0], Value: data[5 : 5+v]})
		data = data[5+v:]
		if len(data) < 5 {
			break
		}
	}
	return res
}

// ---------------------------------LV : LENGTH & VALUE ---------------------------------

func LvEncode(data []byte) []byte {
	v := len(data)
	if v == 0 {
		return nil
	}
	b := make([]byte, 4)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	return append(b, data...)
}

func LvDecode(data []byte) []byte {
	if len(data) < 5 {
		return nil
	}
	v := uint32(data[3]) | uint32(data[2])<<8 | uint32(data[1])<<16 | uint32(data[0])<<24
	if len(data) < 4+int(v) {
		return nil
	}
	return data[4 : 4+v]
}

func LvEncodeAll(data [][]byte) (res []byte) {
	for _, v := range data {
		res = append(res, LvEncode(v)...)
	}
	return res
}

func LvDecodeAll(data []byte) [][]byte {
	var res [][]byte
	if len(data) < 5 {
		return nil
	}
	for {
		v := uint32(data[3]) | uint32(data[2])<<8 | uint32(data[1])<<16 | uint32(data[0])<<24
		if len(data) < 4+int(v) {
			break
		}
		res = append(res, data[4:4+v])
		data = data[4+v:]
		if len(data) < 5 {
			break
		}
	}
	return res
}
