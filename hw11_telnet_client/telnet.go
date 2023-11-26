package main

import (
	"bufio"
	"context"
	"errors"
	"io"
	"net"
	"os"
	"os/signal"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
	HandleSignals()
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	ctx, cancel := context.WithCancel(context.Background())
	return &telnetClient{
		network: "tcp",
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
		ctx:     ctx,
		cancel:  cancel,
	}
}

type telnetClient struct {
	network string
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
	ctx     context.Context
	cancel  context.CancelFunc
}

func (c *telnetClient) Connect() error {
	conn, err := net.DialTimeout(c.network, c.address, c.timeout)
	c.conn = conn
	return err
}

func (c *telnetClient) Close() error {
	errC := c.conn.Close()
	errIn := c.in.Close()
	c.cancel()
	if errC != nil {
		return errC
	}
	if errIn != nil {
		return errIn
	}
	return nil
}

func (c *telnetClient) HandleSignals() {
	ctx, stop := signal.NotifyContext(c.ctx, os.Interrupt)
	defer stop()
	<-ctx.Done()
	c.Close()
	os.Stderr.Write([]byte("Connection was closed by peer\n"))
}

func (c *telnetClient) Send() error {
	r := bufio.NewReader(c.in)
	for {
		select {
		case <-c.ctx.Done():
			return nil
		default:
			b, errR := r.ReadBytes('\n')
			if errR == io.EOF {
				return nil
			}
			if errR != nil {
				c.cancel()
				return errR
			}
			_, errW := c.conn.Write(b)
			if errW != nil {
				c.cancel()
				return errW
			}
		}
	}
}

func (c *telnetClient) Receive() error {
	for {
		select {
		case <-c.ctx.Done():
			return nil
		default:
			b := make([]byte, 100)
			n, errR := c.conn.Read(b)
			if errors.Is(errR, io.EOF) {
				return nil
			}
			if errR != nil {
				c.cancel()
				return errR
			}
			if _, errW := c.out.Write(b[:n]); errW != nil {
				return errW
			}
		}
	}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
