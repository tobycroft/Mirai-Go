package BotAllowGroupModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "bot_allow_group"

func Api_find(qq, group_id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq":       qq,
		"group_id": group_id,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_insert(qq, group_id interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"qq":       qq,
		"group_id": group_id,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_delete(qq, group_id interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"qq":       qq,
		"group_id": group_id,
	}
	db.Data(data)
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
