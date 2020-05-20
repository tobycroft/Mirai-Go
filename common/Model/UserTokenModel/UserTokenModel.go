package UserTokenModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"time"
)

const table = "live_user_token"

func Api_insert(project, uid, token, exp interface{}) bool {
	db := tuuz.DB().Table(table)
	data := map[string]interface{}{}
	data["project"] = project
	data["uid"] = uid
	data["token"] = token
	data["expire"] = exp
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find(project, uid interface{}) gorose.Data {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{}
	where["project"] = project
	where["uid"] = uid
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_avail(project, uid, token interface{}) gorose.Data {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{}
	where["project"] = project
	where["uid"] = uid
	where["token"] = token
	db.Where(where)
	db.Where("expire", ">", time.Now().Unix())
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
