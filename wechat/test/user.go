package main

import "net"

type User struct {
	Addr string
	Name string
	conn net.Conn
	msg  chan string
}

func Lister(c net.Conn) *User {
	addr := c.RemoteAddr().String()
	user := User{
		Addr: addr,
		Name: addr,
		msg:  make(chan string),
		conn: c,
	}
	go user.Write()
	return &user
}
func (u User) Write() {

	c1 := <-u.msg
	u.conn.Write([]byte(c1))
}
