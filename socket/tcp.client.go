package socket

import (
	"bufio"
	"encoding/binary"
	"io"
	"net"
)

//-------------------------------------
//
// 按行写入, 按行读取
//
//-------------------------------------
type TcpClient struct {
	conn  net.Conn
	order binary.ByteOrder
	buf   *bufio.Reader
}

//-------------------------------------
//
//
//
//-------------------------------------
func NewTcpConnection(conn net.Conn, order binary.ByteOrder) *TcpClient {
	c := &TcpClient{
		conn,
		order,
		bufio.NewReader(conn),
	}
	return c
}

//-------------------------------------
//
// 读取, 自动移除\n
//
//-------------------------------------
func (c *TcpClient) ReadLine() (string, error) {
	line, err := c.buf.ReadString('\n')
	return line, err
}

//-------------------------------------
//
// 读取对象
//
//-------------------------------------
func (c *TcpClient) ReadObject(obj interface{}) error {
	err := binary.Read(c.conn, c.order, obj)
	return err
}

//-------------------------------------
//
// 读取固定行数
//
//-------------------------------------
func (c *TcpClient) ReadLength(size int) ([]byte, error) {
	buff := make([]byte, size)
	size, err := io.ReadFull(c.conn, buff)
	if err != nil {
		return nil, err
	} else {
		return buff[:size], nil
	}
}

//-------------------------------------
//
// 写入
//
//-------------------------------------
func (c *TcpClient) Write(src string) error {
	buff := []byte(src)
	_, err := c.conn.Write(buff)
	return err
}

//-------------------------------------
//
// 写入一行
//
//-------------------------------------
func (c *TcpClient) WriteLine(src string) error {
	return c.Write(src + "\n")
}

//-------------------------------------
//
// 关闭
//
//-------------------------------------
func (c *TcpClient) Close() {
	c.conn.Close()
}
