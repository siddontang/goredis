package goredis

import (
	"github.com/alicebob/miniredis"
	"testing"
)

func Test(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	addr := s.Addr()

	c := NewClient(addr)
	defer c.Close()

	conn, err := c.Get()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	if pong, err := String(conn.Do("PING")); err != nil {
		t.Fatal(err)
	} else if pong != "PONG" {
		t.Fatal(pong)
	} else if conn.GetTotalReadSize() != 7 {
		t.Fatal(conn.GetTotalReadSize())
	}

	if pong, err := String(conn.Do("PING")); err != nil {
		t.Fatal(err)
	} else if pong != "PONG" {
		t.Fatal(pong)
	} else if conn.GetTotalReadSize() != 14 {
		t.Fatal(conn.GetTotalReadSize())
	}
}
