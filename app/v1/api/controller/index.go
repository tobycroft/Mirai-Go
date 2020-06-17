package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main.go/extend/MirAi/api"
	"main.go/extend/MirAi/v1/model/BotModel"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("", index)
	route.Any("about", about)
	route.Any("auth", auth)
	route.Any("verify", verify)
}

func index(c *gin.Context) {
	c.JSON(200, RET.Ret_succ(0, index))
}

func about(c *gin.Context) {
	ret, err := api.About(2140300010)
	if err != nil {
		c.JSON(200, err.Error())
	} else {
		c.JSON(200, ret)
	}
}

func auth(c *gin.Context) {
	ret, err := api.Auth(2140300010)
	if err != nil {
		c.JSON(200, err.Error())
	} else {
		c.JSON(200, ret)
	}
}

func verify(c *gin.Context) {
	qq := 2140300010
	ret, err := api.Auth(qq)
	if err != nil {
		c.JSON(200, err.Error())
	} else {
		if ret["code"].(float64) == 0 {
			session := ret["session"].(string)
			//bind
			ret, err = api.Verify(qq, session)
			if err != nil {
				c.JSON(200, err.Error())
			} else {
				if BotModel.Api_update(qq, session) {
					c.JSON(200, err.Error())
				} else {
					c.JSON(200, errors.New("数据库插入失败"))
				}
			}
		} else {
			c.JSON(200, errors.New("获取session失败"))
		}
	}
}
