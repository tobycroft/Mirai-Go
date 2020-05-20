package SystemParamModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "system_param"

func Api_find(project, key interface{}) gorose.Data {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{
		"project": project,
		"key":     key,
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

func Api_find_value(project, key interface{}) string {
	ret := Api_find(project, key)
	if len(ret) > 0 {
		return ret["value"].(string)
	} else {
		return ""
	}
}
