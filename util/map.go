package util

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"os"
)

func ReadMapFromFile(path string) ([]byte, int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return data, 0, err
	}

	return data, len(data), nil
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

	return compressedBytes
}

func GetMapAndSize(path string) ([]byte, []byte, int, error) {
	mapBytes, mapSize, err := ReadMapFromFile(path)
	if err != nil {
		return mapBytes, []byte{}, 0, err
	}

	compressed := CompressMap(mapBytes)
	bigEndianMapSize := len(compressed)

	mapSizeBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(mapSizeBuf, uint32(bigEndianMapSize))

	return compressed, mapSizeBuf, mapSize, nil
}
