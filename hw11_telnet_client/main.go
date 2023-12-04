package main

import (
	"flag"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var timeout time.Duration

func main() {
	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout")
	flag.Parse()
	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)
	c := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		if err := c.Send(); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := c.Receive(); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()
		c.HandleSignals()
	}()
	wg.Wait()
}
