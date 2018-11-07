package blc

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(num int64) []byte { // 将int64转换为自负切片
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
