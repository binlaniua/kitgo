package tcp

import (
	"net"
	"bufio"
	"io"
)

//-------------------------------------
//
// 按行写入, 按行读取
//
//-------------------------------------
type LineConnection struct {
	conn net.Conn
	buf  *bufio.Reader
}

//-------------------------------------
//
//
//
//-------------------------------------
func NewLineConnection(conn net.Conn) *LineConnection {
	c := &LineConnection{
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
func (c *LineConnection) Read() (string, error) {
	line, err := c.buf.ReadString('\n')
	if err == io.EOF || err == nil {
		return line, nil
	} else {
		return line, err
	}
}

//-------------------------------------
//
// 写入, 自动加\n
//
//-------------------------------------
func (c *LineConnection) Write(src string) (error) {
	buff := []byte(src + "\n")
	_, err := c.conn.Write(buff)
	return err
}

//-------------------------------------
//
// 关闭
//
//-------------------------------------
func (c *LineConnection) Close() {
	c.conn.Close()
}




