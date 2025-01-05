package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(tl.in)

	for scanner.Scan() {
		text := scanner.Text() + "\n"
		_, err := tl.conn.Write([]byte(text))
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

func (tl *Telnet) Receive() error {
	scanner := bufio.NewScanner(tl.conn)
	for scanner.Scan() {
		_, err := fmt.Fprintln(tl.out, scanner.Text())
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Telnet{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
