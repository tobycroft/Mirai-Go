package LogRecvModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "log_recv"

func Api_insert(qq, log interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"qq":  qq,
		"log": log,
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
