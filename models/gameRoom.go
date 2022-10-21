package models

type GameRoom struct {
	Id int `json:"roomId" form:"roomId" binding:"required"` //游戏房间号

}

func NewGameRoom() *GameRoom {
	room := &GameRoom{
		0,
	}
	return room
}
