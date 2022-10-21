package service

import (
	"fmt"
	"game/dao"
	"game/models"
)

type GameService interface {
	//玩家移动判断
	Move() (isVictory bool, info *dao.RoomInfo)
}
type Game struct {
	User *models.User
	Room *models.GameRoom
}
type Result struct {
	IsVictory bool          `json:"isVictory"` //返回结果，true表示游戏结束，有人胜利，false表示游戏继续
	Info      *dao.RoomInfo `json:"info"`      //房间信息

}

func (game *Game) Move() *Result {
	//根据游戏房间号查询该房间内的相关信息，拿到此房间内玩家的游戏状态
	roomId := game.Room.Id
	roomInfo := dao.NewRoomInfo()
	roomInfo.SelectRoomInfoByRoomId(roomId)
	fmt.Println(roomInfo)
	//判断房间是否存在或者房间是否关闭
	if roomInfo.Id == 0 {
		//房间不存在,或者已经关闭
		panic("房间不存在！")
	}
	//判断点数之和有没有超过总格子数，超过则后退
	user := game.User
	nowUser1Step := new(int)
	nowUser2Step := new(int)
	if user.Id == roomInfo.UserId1 {
		//如果是房间的玩家1
		if user.Step+roomInfo.User1Steps > roomInfo.Size {
			//超出了总格子数，原路返回
			*nowUser1Step = roomInfo.Size - (user.Step + roomInfo.User1Steps - roomInfo.Size)
		} else {
			*nowUser1Step = user.Step + roomInfo.User1Steps
		}
		//更新房间的最新状态，将新扔的点数写入
		roomInfo.User1Steps = *nowUser1Step
		roomInfo.UpdateUser1StepByRoomId(roomId)
		//判断是否胜利
		if *nowUser1Step == roomInfo.Size {
			//如果有人胜利则将房间关闭
			roomInfo.DeleteRoomByRoomId(roomId)
			return &Result{
				true, nil,
			}
		}
	}
	if user.Id == roomInfo.UserId2 {
		//如果是房间的玩家2
		if user.Step+roomInfo.User2Steps > roomInfo.Size {
			//超出了总格子数，原路返回
			*nowUser2Step = roomInfo.Size - (user.Step + roomInfo.User2Steps - roomInfo.Size)
		} else {
			*nowUser2Step = user.Step + roomInfo.User2Steps
		}
		//更新房间的最新状态，将新扔的点数写入
		roomInfo.User2Steps = *nowUser2Step
		roomInfo.UpdateUser2StepByRoomId(roomId)
		//判断是否胜利
		if *nowUser2Step == roomInfo.Size {
			//如果有人胜利则将房间关闭
			roomInfo.DeleteRoomByRoomId(roomId)
			return &Result{true, nil}
		}
	}
	//没有玩家胜利则返回最新的房间状态
	return &Result{
		false, roomInfo,
	}
}
