package numberx

import (
	"bytes"
	"encoding/binary"
)

func ToBytes(data int64) ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, &data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
