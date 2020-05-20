package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/user/v1/action/UIAction"
	"main.go/app/user/v1/model/UserFollowModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func FollowController(route *gin.RouterGroup) {
	route.Use(initialize(), gin.Recovery())
	route.Any("/", index)
	route.Any("/set", set)
	route.Any("/unset", unset)
	route.Any("/list", list)
	route.Any("follow", follow)
	route.Any("/is_follow", is_follow)
}

func set(c *gin.Context) {
	project := c.PostForm("project")
	uid := c.PostForm("uid")
	roomid, ok := c.GetPostForm("roomid")
	if !ok {
		c.JSON(200, RET.Ret_fail(400, "roomid"))
		c.Abort()
		return
	}
	up_id, ok := c.GetPostForm("up_id")
	if !ok {
		c.JSON(200, RET.Ret_fail(400, "up_id"))
		c.Abort()
		return
	}
	user := UserFollowModel.Api_find(project, uid, roomid, up_id)
	if len(user) > 0 {
		c.JSON(200, RET.Ret_fail(403, "你已经关注这个主播了"))
	} else {
		if UserFollowModel.Api_insert(project, uid, roomid, up_id, "default", 0) {
			c.JSON(200, RET.Ret_succ(0, "关注成功"))
		} else {
			c.JSON(200, RET.Ret_fail(500, "关注失败"))
		}
	}
}

func unset(c *gin.Context) {
	project := c.PostForm("project")
	uid := c.PostForm("uid")
	roomid, ok := c.GetPostForm("roomid")
	if !ok {
		c.JSON(200, RET.Ret_fail(400, "roomid"))
		c.Abort()
		return
	}
	up_id, ok := c.GetPostForm("up_id")
	if !ok {
		c.JSON(200, RET.Ret_fail(400, "up_id"))
		c.Abort()
		return
	}
	user := UserFollowModel.Api_find(project, uid, roomid, up_id)
	if len(user) > 0 {
		if UserFollowModel.Api_delete(project, uid, roomid, up_id) {
			c.JSON(200, RET.Ret_succ(0, "取关成功"))
		} else {
			c.JSON(200, RET.Ret_fail(500, "取关失败"))
		}
	} else {
		c.JSON(200, RET.Ret_fail(403, "你没有关注这个主播"))
	}
}

func follow(c *gin.Context) {
	project := c.PostForm("project")
	uid := c.PostForm("uid")
	roomid, ok := c.GetPostForm("roomid")
	if !ok {
		c.JSON(200, RET.Ret_fail(400, "roomid"))
		c.Abort()
		return
	}
	up_id, ok := c.GetPostForm("up_id")
	if !ok {
		c.JSON(200, RET.Ret_fail(400, "up_id"))
		c.Abort()
		return
	}
	folow, ok := Input.Post("follow", c, false)
	if !ok {
		return
	}
	switch folow {
	case "yes":
		user := UserFollowModel.Api_find(project, uid, roomid, up_id)
		if len(user) > 0 {
			c.JSON(200, RET.Ret_fail(403, "你已经关注这个主播了"))
		} else {
			if UserFollowModel.Api_insert(project, uid, roomid, up_id, "default", 0) {
				c.JSON(200, RET.Ret_succ(0, "关注成功"))
			} else {
				c.JSON(200, RET.Ret_fail(500, "关注失败"))
			}
		}
		break

	case "no":
		user := UserFollowModel.Api_find(project, uid, roomid, up_id)
		if len(user) > 0 {
			if UserFollowModel.Api_delete(project, uid, roomid, up_id) {
				c.JSON(200, RET.Ret_succ(0, "取关成功"))
			} else {
				c.JSON(200, RET.Ret_fail(500, "取关失败"))
			}
		} else {
			c.JSON(200, RET.Ret_fail(403, "你没有关注这个主播"))
		}
		break

	default:
		break
	}

}

func list(c *gin.Context) {
	project := c.PostForm("project")
	uid := c.PostForm("uid")
	follow_list := UserFollowModel.Api_select(project, uid)
	for k, v := range follow_list {
		v["upinfo"] = UIAction.App_userinfo(project, v["up_id"])
		follow_list[k] = v
	}
	c.JSON(200, RET.Ret_succ(0, follow_list))
}

func is_follow(c *gin.Context) {
	project := c.PostForm("project")
	uid := c.PostForm("uid")
	roomid, ok := Input.Post("roomid", c, false)
	if !ok {
		return
	}
	up_id, ok := Input.Post("up_id", c, false)
	if !ok {
		return
	}
	if len(UserFollowModel.Api_find(project, uid, roomid, up_id)) > 0 {
		c.JSON(200, RET.Ret_succ(0, true))
	} else {
		c.JSON(200, RET.Ret_fail(0, false))
	}
}
