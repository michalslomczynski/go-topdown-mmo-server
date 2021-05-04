package players

import (
	"fmt"
	"sync"
)

const (
	None = iota
	Left
	Right
	Up
	Down
)

type Position struct {
	X int16
	Y int16
	MovingDirection int8
}

type Player struct {
	ID              int16
	Pos *Position
	Level           int16
	Outfit          int16
}

var mutex = &sync.Mutex{}

var PlayerList = make(map[int16]*Player)
var PlayerPosList = make(map[int16]*Position)

func CreatePlayer() *Player {
	p := new(Player)
	p.ID = getNextID()
	fmt.Println(p.ID)
	p.Pos = &Position{15, 15, None}
	p.Level = 10
	mutex.Lock()
	PlayerList[p.ID] = p
	PlayerPosList[p.ID] = p.Pos
	mutex.Unlock()
	return p
}

func getNextID() int16 {
	for i := 0; i < len(PlayerList); i++ {
		if _, ok := PlayerList[int16(i)]; !ok {
			return int16(i)
		}
	}
	return int16(len(PlayerList))
}

func DeletePlayer(ID int16) {
	mutex.Lock()
	delete(PlayerList, ID)
	mutex.Unlock()
}

func MoveLeft(ID int16) {
	mutex.Lock()
	PlayerList[ID].Pos.X--
	PlayerList[ID].Pos.MovingDirection = Left
	mutex.Unlock()
}

func MoveRight(ID int16) {
	mutex.Lock()
	PlayerList[ID].Pos.X++
	PlayerList[ID].Pos.MovingDirection = Right
	mutex.Unlock()
}

func MoveUp(ID int16) {
	mutex.Lock()
	PlayerList[ID].Pos.Y--
	PlayerList[ID].Pos.MovingDirection = Up
	mutex.Unlock()
}

func MoveDown(ID int16) {
	mutex.Lock()
	PlayerList[ID].Pos.Y++
	PlayerList[ID].Pos.MovingDirection = Down
	mutex.Unlock()
}

func MoveNone(ID int16) {
	mutex.Lock()
	PlayerList[ID].Pos.MovingDirection = None
	mutex.Unlock()
}
