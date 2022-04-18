package main

import (
	"fmt"
	"time"
)

// Server Addr,Port为必须参数，而其他为可选参数
type Server struct {
	Addr     string
	Port     int
	Protocol string
	MaxConns int
	Timeout  time.Duration
}

// Option 函数类型
type Option func(*Server)

func Protocol(p string) Option {
	return func(server *Server) {
		server.Protocol = p
	}
}

func Maxconns(maxconns int) Option {
	return func(server *Server) {
		server.MaxConns = maxconns
	}
}

func Timeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.Timeout = timeout
	}
}

// NewServer server初始化
func NewServer(addr string, port int, option ...Option) (*Server, error) {
	svr := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		MaxConns: 100,
		Timeout:  10 * time.Second,
	}
	for _, opt := range option {
		opt(&svr)
	}
	return &svr, nil
}

func main() {
	s1, _ := NewServer("127.0.0.1", 9443)
	s2, _ := NewServer("127.0.0.1", 9443, Protocol("udp"))
	s3, _ := NewServer("127.0.0.1", 944, Protocol("tcp"), Maxconns(1000), Timeout(1*time.Second))
	fmt.Println(s1, s2, s3)
}
