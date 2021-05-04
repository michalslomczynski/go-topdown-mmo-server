package main

import (
	"encoding/gob"
	"fmt"
	"github.com/michalslomczynski/go-topdown-mmo-server/network"
	"github.com/michalslomczynski/go-topdown-mmo-server/players"
	"log"
	"net"
	"time"
)

const (
	port = ":8081"
)

func read(conn net.Conn, connID *int16) {
	dec := gob.NewDecoder(conn)
	enc := gob.NewEncoder(conn)

	for {
		msg, err := network.ReceiveMessage(dec)
		if err != nil {
			return
		}

		switch msg.ID {
		case 0:
			continue

		case network.MapRequest:
			network.SendMap(enc)

		case network.NewPlayerRequest:
			network.CreatePlayer(enc, connID)

		case network.PlayerListRequest:
			network.SendPlayerList(enc)

		case network.PlayerPosListRequest:
			network.SendPlayerPosList(enc)

		case network.MovePlayerLeftRequest:
			network.MovePlayerLeft(msg.Data)

		case network.MovePlayerRightRequest:
			network.MovePlayerRight(msg.Data)

		case network.MovePlayerUpRequest:
			network.MovePlayerUp(msg.Data)

		case network.MovePlayerDownRequest:
			network.MovePlayerDown(msg.Data)

		case network.StopPlayerRequest:
			network.StopPlayer(msg.Data)
		}
	}
}

func handle(conn net.Conn) {
	var ID int16
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(time.Now(), "Client connected from "+remoteAddr)
	read(conn, &ID)
	players.DeletePlayer(ID)
	fmt.Println("Client connected from " + remoteAddr + " has ended the session.")
}

func main() {
	server, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			return
		}
		go handle(conn)
	}
}