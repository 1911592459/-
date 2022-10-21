package dao

//游戏房间，一个房间俩个玩家
type RoomInfo struct {
	Id         int `json:"id" gorm:"column:id"`                 //房间号
	UserId1    int `json:"userId1" gorm:"column:userId1"`       //玩家1id
	UserId2    int `json:"userId2" gorm:"column:userId2"`       //玩家2id
	Size       int `json:"size" gorm:"column:size"`             //总格子数
	User1Steps int `json:"user1Steps" gorm:"column:user1Steps"` //玩家1目前的总步数
	User2Steps int `json:"user2Steps" gorm:"column:user2Steps"` //玩家2目前的总步数

}

func NewRoomInfo() *RoomInfo {
	return &RoomInfo{}
}

// 根据房间id查询房间信息
func (roomInfo *RoomInfo) SelectRoomInfoByRoomId(roomId int) {

	scan := DB.Raw("select rm.id id,rm.userId1 userId1,rm.userId2 userId2,rm.size size,rm.user1Steps user1Steps,rm.user2Steps user2Steps from room rm where rm.delete_flag =0 and rm.id=?", roomId).Scan(roomInfo)
	if scan.Error != nil {
		panic(scan.Error)
	}
}

//更新玩家1最新步数
func (r *RoomInfo) UpdateUser1StepByRoomId(roomId int) {
	exec := DB.Exec("update room rm set rm.user1Steps =? where rm.delete_flag=0 and rm.id=?", r.User1Steps, roomId)
	if exec.Error != nil {
		panic(exec.Error)
	}
}

//更新玩家2最新步数
func (r *RoomInfo) UpdateUser2StepByRoomId(roomId int) {
	exec := DB.Exec("update room rm set rm.user2Steps =? where rm.delete_flag=0 and rm.id=?", r.User2Steps, roomId)
	if exec.Error != nil {
		panic(exec.Error)
	}
}

//关闭房间，逻辑删除
func (r *RoomInfo) DeleteRoomByRoomId(roomId int) {
	exec := DB.Exec("update room rm set rm.delete_flag=1 and rm.delete_flag=0 and rm.id=?", roomId)
	if exec.Error != nil {
		panic(exec.Error)
	}
}
