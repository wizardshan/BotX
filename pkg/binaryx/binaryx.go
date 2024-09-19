package binaryx

import (
	"bytes"
	"encoding/binary"
)

func ToInt(b []byte) (int64, error) {
	buf := bytes.NewBuffer(b)
	var data int64
	return data, binary.Read(buf, binary.BigEndian, &data)
}
