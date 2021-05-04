package network

import (
	"encoding/gob"
	"fmt"
	"time"
)

type Message struct {
	ID   int
	Data []byte
}

const (
	OverAndOut = iota
	MapRequest
	NewPlayerRequest
	PlayerListRequest
	PlayerPosListRequest
	MovePlayerLeftRequest
	MovePlayerRightRequest
	MovePlayerUpRequest
	MovePlayerDownRequest
	StopPlayerRequest
)

func logerr(err error) bool {
	if err != nil {
		fmt.Println(time.Now(), err)
	}
	return false
}

func ReceiveMessage(dec *gob.Decoder) (*Message, error) {
	msg := new(Message)
	err := dec.Decode(msg)
	return msg, err
}
