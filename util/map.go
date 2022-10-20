package util

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"github.com/codecat/go-libs/log"
	"os"
)

func ReadMapFromFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return data, err
	}

	log.Debug("UNCOMPRESSED MAP SIZE: ", len(data))

	return data, nil
}

func CompressMap(data []byte) []byte {
	var compressed bytes.Buffer
	w, err := zlib.NewWriterLevel(&compressed, zlib.DefaultCompression)
	if err != nil {
		return []byte{}
	}
	w.Write(data)
	w.Close()

	compressedBytes := compressed.Bytes()

	log.Debug("COMPRESSED MAP SIZE: ", len(compressedBytes))

	return compressedBytes
}

func GetMapAndSize(path string) ([]byte, []byte, error) {
	mapBytes, err := ReadMapFromFile(path)
	if err != nil {
		return mapBytes, []byte{}, err
	}

	compressed := CompressMap(mapBytes)
	bigEndianMapSize := len(compressed)

	mapSizeBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(mapSizeBuf, uint32(bigEndianMapSize))

	return compressed, mapSizeBuf, nil
}
