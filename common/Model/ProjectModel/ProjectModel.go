package ProjectModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "live_project"

func Api_insert() {

}

func Api_find(token interface{}) gorose.Data {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{}
	where["`token`"] = token
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_avail(token interface{}) gorose.Data {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{}
	where["`token`"] = token
	where["`status`"] = true
	where["`lock`"] = false
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
