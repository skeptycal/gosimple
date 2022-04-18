package main

// interesting pattern ...
// Reference: https://stackoverflow.com/a/26149656

import (
	"sync"
	"sync/atomic"
)

type Server struct {
	s     chan int
	r     chan int
	c     *sync.Cond
	state uint32
}

const (
	sNormal       = 0
	sQuitting     = 1
	sReconnecting = 2
)

func New() *Server {
	s := &Server{
		s: make(chan int),
		r: make(chan int),
		c: sync.NewCond(&sync.Mutex{}),
	}
	go s.sender()
	// go s.receiver()
	return s
}
func (s *Server) sender() {
	//
	for {
		select {
		case data := <-s.s:
			//do stuff with data
			_ = data
		default:
			s.c.L.Lock()
		L:
			for {
				switch atomic.LoadUint32(&s.state) {
				case sNormal:
					break L
				case sReconnecting:
				case sQuitting:
					s.c.L.Unlock()
					return
				}
				s.c.Wait()
			}
			s.c.L.Unlock()
		}
	}
}

//repeat for receiver

func (s *Server) Reconnect() {
	var cannotReconnect bool
	atomic.StoreUint32(&s.state, sReconnecting)
	//keep trying to reconnect
	if cannotReconnect {
		atomic.StoreUint32(&s.state, sQuitting)
	} else {
		atomic.StoreUint32(&s.state, sNormal)
	}
	s.c.Broadcast()
}
