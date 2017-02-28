package socket

import (
	"github.com/binlaniua/kitgo"
	"net"
)

//-------------------------------------
//
//
//
//-------------------------------------
type UdpServer struct {
	server *net.UDPConn
}

//-------------------------------------
//
//
//
//-------------------------------------
func NewUdpServer(addressString string) *UdpServer {
	server := &UdpServer{}
	address, err := net.ResolveUDPAddr("udp4", addressString)
	if err != nil {
		kitgo.ErrorLog.Fatalln(address, "启动失败 => ", err)
	}
	server.server, err = net.ListenUDP("udp4", address)
	if err != nil {
		kitgo.ErrorLog.Fatalln(address, "启动失败 => ", err)
	}
	return server
}

//-------------------------------------
//
//
//
//-------------------------------------
func (us *UdpServer) ReadLength(size int64) ([]byte, *net.UDPAddr, error) {
	buff := make([]byte, size)
	readSize, addr, err := us.server.ReadFromUDP(buff)
	return buff[:readSize], addr, err
}

//-------------------------------------
//
//
//
//-------------------------------------
func (us *UdpServer) Close() {
	if us.server != nil {
		us.server.Close()
	}
}
