package user

import "fmt"

type Closer interface {
	Closer() error
}


type Address struct {
	port int	`json:port`
	ip string	`json:ip`
}

func (c Address) Ip() string {
	return c.ip
}

func (c Address) Port() int {
	return c.port
}


type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{port, ip}}
}



func NewConnection2(address Address, status int) *Connection {
	return &Connection{Address: address, status: status}
}

func (connection Connection)send(msg string) error  {
	fmt.Printf("发送给 [%s:%d]:%s", connection.ip, connection.port, msg)
	return nil
}

func (connection Connection)Close() error  {
	fmt.Printf("关闭连接 [%s:%d]", connection.ip, connection.port)
	return nil
}



