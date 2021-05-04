package network

import (
	"bytes"
	"encoding/gob"
	"github.com/michalslomczynski/go-topdown-mmo-server/mappkg"
	"github.com/michalslomczynski/go-topdown-mmo-server/players"
	"log"
)

func SendMap(enc *gob.Encoder) error {
	err := enc.Encode(mappkg.Worldmap)
	logerr(err)

	return nil
}

func CreatePlayer(enc *gob.Encoder, connID *int16) error {
	newPlayer := *players.CreatePlayer()
	*connID = newPlayer.ID

	err := enc.Encode(newPlayer)
	logerr(err)

	return nil
}

func SendPlayerList(enc *gob.Encoder) error {
	err := enc.Encode(players.PlayerList)
	logerr(err)

	return nil
}

func SendPlayerPosList(enc *gob.Encoder) error {
	err := enc.Encode(players.PlayerPosList)
	logerr(err)

	return nil
}

func MovePlayerLeft(data []byte) {
	buff := bytes.NewBuffer(data)
	ID := new(int16)

	objdec := gob.NewDecoder(buff)
	err := objdec.Decode(ID)

	if err != nil {
		log.Fatal(err)
	}

	players.MoveLeft(*ID)
}

func MovePlayerRight(data []byte) {
	buff := bytes.NewBuffer(data)
	ID := new(int16)

	objdec := gob.NewDecoder(buff)
	err := objdec.Decode(ID)

	if err != nil {
		log.Fatal(err)
	}

	players.MoveRight(*ID)
}

func MovePlayerUp(data []byte) {
	buff := bytes.NewBuffer(data)
	ID := new(int16)

	objdec := gob.NewDecoder(buff)
	err := objdec.Decode(ID)

	if err != nil {
		log.Fatal(err)
	}

	players.MoveUp(*ID)
}

func MovePlayerDown(data []byte) {
	buff := bytes.NewBuffer(data)
	ID := new(int16)

	objdec := gob.NewDecoder(buff)
	err := objdec.Decode(ID)

	if err != nil {
		log.Fatal(err)
	}

	players.MoveDown(*ID)
}

func StopPlayer(data []byte) {
	buff := bytes.NewBuffer(data)
	ID := new(int16)

	objdec := gob.NewDecoder(buff)
	err := objdec.Decode(ID)

	if err != nil {
		log.Fatal(err)
	}

	players.MoveNone(*ID)
}
