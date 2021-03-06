package Input

import (
	"github.com/gin-gonic/gin"
	"main.go/tuuz/Calc"
	"main.go/tuuz/RET"
)

func Post(key string, c *gin.Context, xss bool) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(200, RET.Ret_fail(400, key))
		c.Abort()
		return "", false
	} else {
		return in, true
	}
}

func PostInt(key string, c *gin.Context) (int, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(200, RET.Ret_fail(400, key))
		c.Abort()
		return 0, false
	} else {
		i, e := Calc.String2Int(in)
		if e != nil {
			c.JSON(200, RET.Ret_fail(407, key+" should be int"))
			c.Abort()
			return 0, false
		}
		return i, true
	}
}

func PostInt64(key string, c *gin.Context) (int64, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(200, RET.Ret_fail(400, key))
		c.Abort()
		return 0, false
	} else {
		i, e := Calc.String2Int64(in)
		if e != nil {
			c.JSON(200, RET.Ret_fail(407, key+" should be int64"))
			c.Abort()
			return 0, false
		}
		return i, true
	}
}

func PostFloat64(key string, c *gin.Context) (float64, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(200, RET.Ret_fail(400, key))
		c.Abort()
		return 0, false
	} else {
		i, e := Calc.String2Float64(in)
		if e != nil {
			c.JSON(200, RET.Ret_fail(407, key+" should be float64"))
			c.Abort()
			return 0, false
		}
		return i, true
	}
}
