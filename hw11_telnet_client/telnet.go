package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Telnet struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func (tl *Telnet) Connect() error {
	var err error

	tl.conn, err = net.DialTimeout("tcp", tl.address, tl.timeout)

	return err
}

func (tl *Telnet) Close() error {
	if tl.conn == nil {
		return fmt.Errorf("connection is not open")
	}

	return tl.conn.Close()
}

func (tl *Telnet) Send() error {
	_, err := io.Copy(tl.conn, tl.in)

	return err
}

func (tl *Telnet) Receive() error {
	_, err := io.Copy(tl.out, tl.conn)

	return err
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Telnet{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
