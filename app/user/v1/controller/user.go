package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/up/v1/model/UserMedalModel"
	"main.go/app/up/v1/model/UserRoomModel"
	"main.go/app/user/v1/model/UserInfoModel"
	"main.go/common/Action/AuthAction"
	"main.go/common/Model/ProjectModel"
	"main.go/tuuz"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func UserController(route *gin.RouterGroup) {
	route.Use(initialize(), gin.Recovery())
	route.Any("/", index)
	route.Any("/userinfo", userinfo)
	route.Any("/medal", medal)
	route.Any("/set_medal", set_medal)
}

func initialize() gin.HandlerFunc {
	return func(c *gin.Context) {
		project, is := c.GetPostForm("project")
		if is == false {
			project, is = c.GetQuery("project")
			if is == false {
				c.JSON(200, RET.Ret_succ(400, "project"))
				c.Abort()
				return
			}
		}

		token, is := c.GetPostForm("token")
		if is == false {
			project, is = c.GetQuery("token")
			if is == false {
				c.JSON(200, RET.Ret_succ(400, "token"))
				c.Abort()
				return
			}
		}
		if len(ProjectModel.Api_find_avail(project)) > 0 {

		} else {
			c.JSON(200, RET.Ret_fail(404, "project_not_found"))
			c.Abort()
			return
		}
		uid, ok := Input.PostInt("uid", c)
		if !ok {
			return
		}
		if AuthAction.TokenAction(project, uid, token, c.ClientIP()) {

		} else {
			//todo:在这里放上连接代码，刷新token
			c.JSON(200, RET.Ret_fail(-1, "token_expired"))
			c.Abort()
			return
		}
		c.Next()
	}
}

func index(c *gin.Context) {
	c.String(0, "index")
}

func userinfo(c *gin.Context) {
	project := c.PostForm("project")
	uid := c.PostForm("uid")
	userinfo := UserInfoModel.Api_find(project, uid)
	if len(userinfo) > 0 {
		c.JSON(200, RET.Ret_succ(0, userinfo))
	} else {
		c.JSON(200, RET.Ret_fail(404, "没有找到该用户"))
	}
}

func medal(c *gin.Context) {
	project := c.PostForm("project")
	uid := c.PostForm("uid")
	usermedal := UserMedalModel.Api_select(tuuz.Db(), project, uid)
	for k, v := range usermedal {
		v["room_info"] = UserRoomModel.Api_find(project, v["roomid"])
		usermedal[k] = v
	}
	c.JSON(200, RET.Ret_succ(0, usermedal))
}

func set_medal(c *gin.Context) {
	project := c.PostForm("project")
	uid := c.PostForm("uid")
	id, ok := Input.PostInt("id", c)
	if !ok {
		return
	}
	UserMedalModel.Api_update_deavtive(tuuz.Db(), project, uid)
	if UserMedalModel.Api_update_active(tuuz.Db(), project, uid, id) {
		c.JSON(200, RET.Ret_succ(0, "设定成功"))
	} else {
		c.JSON(200, RET.Ret_fail(500, "设定失败"))

	}
}
