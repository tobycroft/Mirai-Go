package TypeModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "live_balance_type"

func Api_find(Db gorose.IOrm, Type interface{}) gorose.Data {
	db := Db.Table(table)
	where := map[string]interface{}{
		"type": Type,
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
