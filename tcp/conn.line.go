package tcp

import (
	"net"
	"bufio"
	"encoding/binary"
)

//-------------------------------------
//
// 按行写入, 按行读取
//
//-------------------------------------
type TcpConnection struct {
	conn net.Conn
	buf  *bufio.Reader
}

//-------------------------------------
//
//
//
//-------------------------------------
func NewLineConnection(conn net.Conn) *TcpConnection {
	c := &TcpConnection{
		conn,
		bufio.NewReader(conn),
	}
	return c
}

//-------------------------------------
//
// 读取, 自动移除\n
//
//-------------------------------------
func (c *TcpConnection) ReadLine() (string, error) {
	line, err := c.buf.ReadString('\n')
	return line, err
}

//-------------------------------------
//
// 读取对象
//
//-------------------------------------
func (c *TcpConnection) ReadObject(obj interface{}) (error) {
	err := binary.Read(c.conn, binary.LittleEndian, obj)
	return err;
}

//-------------------------------------
//
// 读取固定行数
//
//-------------------------------------
func (c *TcpConnection) ReadLength(size int) ([]byte, error) {
	buff := make([]byte, size)
	size, err := c.conn.Read(buff)
	if err != nil {
		return nil, err
	} else {
		return buff[0: size], nil
	}
}

//-------------------------------------
//
// 写入
//
//-------------------------------------
func (c *TcpConnection) Write(src string) (error) {
	buff := []byte(src)
	_, err := c.conn.Write(buff)
	return err
}

//-------------------------------------
//
// 写入一行
//
//-------------------------------------
func (c *TcpConnection) WriteLine(src string) (error) {
	return c.Write(src + "\n")
}

//-------------------------------------
//
// 关闭
//
//-------------------------------------
func (c *TcpConnection) Close() {
	c.conn.Close()
}




