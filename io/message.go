package kqio // Import path github.com/killer-queen-stats/libkq/io

import "time"

type MessageType string
type Message struct {
	Time time.Time
	Type MessageType
	Val  interface{}
}

type MessageString struct {
	Time    time.Time
	Message []byte
}
