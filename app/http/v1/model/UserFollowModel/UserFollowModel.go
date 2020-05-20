package UserFollowModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "live_user_follow"

func Api_insert(project, uid, roomid, up_id, fenzu, exp interface{}) bool {
	db := tuuz.DB().Table(table)
	data := map[string]interface{}{
		"project": project,
		"uid":     uid,
		"roomid":  roomid,
		"up_id":   up_id,
		"fenzu":   fenzu,
		"exp":     exp,
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

func Api_find(project, uid, roomid, up_id interface{}) gorose.Data {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{
		"project": project,
		"uid":     uid,
		"roomid":  roomid,
		"up_id":   up_id,
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

func Api_find_byUpId(project, uid, up_id interface{}) gorose.Data {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{
		"project": project,
		"uid":     uid,
		"up_id":   up_id,
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

func Api_delete(project, uid, roomid, up_id interface{}) bool {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{
		"project": project,
		"uid":     uid,
		"roomid":  roomid,
		"up_id":   up_id,
	}
	db.Where(where)
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select(project, uid interface{}) []gorose.Data {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{
		"project": project,
		"uid":     uid,
	}
	db.Where(where)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_upId(project, up_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"project": project,
		"up_id":   up_id,
	}
	db.Where(where)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_count_upId(project, up_id interface{}) int64 {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"project": project,
		"up_id":   up_id,
	}
	db.Where(where)
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_count(project, roomid interface{}) int64 {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{
		"project": project,
		"roomid":  roomid,
	}
	db.Where(where)
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}
