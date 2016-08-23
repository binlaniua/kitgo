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
func (rd *Reader) ReadByte() (byte, error) {
	byteList, err := rd.ReadBytes(1)
	if err != nil || len(byteList) == 0 {
		return -1, err
	} else {
		return byteList[0], nil
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader) ReadBytes(size int) ([]byte, error) {
	buffer := make([]byte, size)
	if size, err := rd.reader.Read(buffer); err != nil {
		return "", err
	} else {
		return buffer[:size], nil
	}

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
func (rd *Reader)ReadUInt16() (uint16, error) {
	buffer := make([]byte, 2)
	if _, err := rd.reader.Read(buffer); err != nil {
		return 0, err
	}
	return rd.order.Uint16(buffer), nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader)ReadUInt32() (uint32, error) {
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
func (rd *Reader)ReadInt16() (int16, error) {
	r, err := rd.ReadUInt16()
	return int16(r), err
}


//-------------------------------------
//
//
//
//-------------------------------------
func (rd *Reader)ReadInt32() (int32, error) {
	r, err := rd.ReadUInt32()
	return int32(r), err
}

//-------------------------------------
//
// 
//
//-------------------------------------
func (rd *Reader)ReadInt() (int, error) {
	return rd.ReadInt32()
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
	r, err := rd.ReadUInt32()
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

//-------------------------------------
//
// 忽略多少字节
//
//-------------------------------------
func (rd *Reader) Skip(size int) {
	buff := make([]byte, size)
	rd.reader.Read(buff)
	buff = nil
}