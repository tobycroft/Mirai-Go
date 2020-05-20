package UserBalanceModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "live_user_balance"

func Api_insert_all(Db gorose.IOrm, project, uid, balance, freeze_balance interface{}) bool {
	db := Db.Table(table)
	data := map[string]interface{}{
		"project":        project,
		"uid":            uid,
		"balance":        balance,
		"freeze_balance": freeze_balance,
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

func Api_insert(Db gorose.IOrm, project, uid interface{}) bool {
	db := Db.Table(table)
	data := map[string]interface{}{
		"project": project,
		"uid":     uid,
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

func Api_inc_balance(Db gorose.IOrm, project, uid, inc_balance interface{}) bool {
	db := Db.Table(table)
	where := map[string]interface{}{
		"project": project,
		"uid":     uid,
	}
	db.Where(where)
	_, err := db.Increment("balance", inc_balance)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find(Db gorose.IOrm, project, uid interface{}) gorose.Data {
	db := Db.Table(table)
	where := map[string]interface{}{
		"project": project,
		"uid":     uid,
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
