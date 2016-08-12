package kitgo

import (
	"io"
	"encoding/binary"
	"math"
	"bytes"
	"io/ioutil"
)

//-------------------------------------
//
//
//
//-------------------------------------
type Reader struct {
	reader io.Reader
	order  binary.ByteOrder
}

//-------------------------------------
//
//
//
//-------------------------------------
func NewReader(reader io.Reader, order binary.ByteOrder) *Reader {
	return &Reader{
		reader,
		order,
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func NewReaderByByte(byteList []byte, order binary.ByteOrder) *Reader {
	buf := bytes.NewBuffer(byteList)
	return NewReader(buf, order)
}

//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader) ReadString(len int, isTrim bool) (string, error) {
	buffer := make([]byte, len)
	if rl, err := rd.reader.Read(buffer); err != nil {
		return "", err
	} else {
		result := string(buffer[:rl])
		if isTrim {
			result = StringReplace(result, "\u0000", "")
		}
		return result, err
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader)ReadUInt() (uint32, error) {
	buffer := make([]byte, 4)
	if _, err := rd.reader.Read(buffer); err != nil {
		return 0, err
	}
	return rd.order.Uint32(buffer), nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader)ReadUInt64() (uint64, error) {
	buffer := make([]byte, 8)
	if _, err := rd.reader.Read(buffer); err != nil {
		return 0, err
	}
	return rd.order.Uint64(buffer), nil
}



//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader)ReadInt() (int, error) {
	r, err := rd.ReadUInt()
	return int(r), err
}

//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader)ReadInt64() (int64, error) {
	r, err := rd.ReadUInt64()
	return int64(r), err
}

//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader)ReadFloat32() (float32, error) {
	r, err := rd.ReadUInt()
	if err != nil {
		return 0, err
	} else {
		return math.Float32frombits(r), err
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader)ReadFloat64() (float64, error) {
	r, err := rd.ReadUInt64()
	if err != nil {
		return 0, err
	} else {
		return math.Float64frombits(r), err
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader) ReaderToEnd() ([]byte, error) {
	return ioutil.ReadAll(rd.reader)
}