package controller

import (
	"fmt"
	"game/models"
	"game/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"runtime/debug"
)

// @Summary 步数更新
// @Description 更新房间内游戏玩家的步数
// @Tags 蛇梯游戏模块
// @Accept       json
// @Produce      json
// @Param “用户和房间信息” body string true "json"
// @Router /throwDice [post]
func ThrowDice(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("panic: %v\n", err)
			debug.PrintStack()
			c.JSON(http.StatusOK, gin.H{
				"code": "-1",
				"msg":  errorToString(err),
				"data": nil,
			})
			return
		}
	}()
	//绑定参数
	gameRoom := models.NewGameRoom()
	err := c.ShouldBindBodyWith(gameRoom, binding.JSON)
	if err != nil {
		fmt.Println(err)
	}
	newUser := models.NewUser()
	err = c.ShouldBindBodyWith(newUser, binding.JSON)
	if err != nil {
		fmt.Println(err)
	}
	game := &service.Game{
		User: newUser,
		Room: gameRoom,
	}
	result := game.Move()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": result,
	})

}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
