package handlers

import (
	"LunaGO/server/interfaces"
	"bytes"
	"encoding/binary"
	"log"
)

func HandleClose(server interfaces.Server) func([]byte) []byte {
	return func(packet []byte) []byte {

		return close(packet)
	}
}

func close([]byte) []byte {
	log.Printf("close")
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, int32(-1))
	return buf.Bytes()
}
