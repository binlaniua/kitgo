package reader

import (
	"io"
	"encoding/binary"
	"math"
)

//-------------------------------------
//
//
//
//-------------------------------------
func ReadUInt(reader io.Reader, order binary.ByteOrder) (uint, error) {
	buffer := make([]byte, 4)
	if _, err := reader.Read(buffer); err != nil {
		return 0, err
	}
	return order.Uint32(buffer), nil
}

func ReadUInt64(reader io.Reader, order binary.ByteOrder) (int64, error) {
	buffer := make([]byte, 8)
	if _, err := reader.Read(buffer); err != nil {
		return 0, err
	}
	return order.Uint64(buffer), nil
}



//-------------------------------------
//
//
//
//-------------------------------------
func ReadInt(reader io.Reader, order binary.ByteOrder) (int, error) {
	r, err := ReadUInt(reader, order)
	return int(r), err
}

func ReadInt64(reader io.Reader, order binary.ByteOrder) (int64, error) {
	r, err := ReadUInt64(reader, order)
	return int64(r), err
}

//-------------------------------------
//
//
//
//-------------------------------------
func ReadFloat32(reader io.Reader, order binary.ByteOrder) (float32, error) {
	r, err := ReadUInt(reader, order)
	if err != nil {
		return 0, err
	} else {
		return math.Float32frombits(r), err
	}
}

func ReadFloat64(reader io.Reader, order binary.ByteOrder) (float64, error) {
	r, err := ReadUInt(reader, order)
	if err != nil {
		return 0, err
	} else {
		return math.Float64frombits(r), err
	}
}