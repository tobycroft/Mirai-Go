package UserInfoModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "live_user_info"

func Api_insert(project, uid, username, uname, face, quhao, phone interface{}) bool {
	db := tuuz.DB().Table(table)
	data := map[string]interface{}{}
	data["project"] = project
	data["uid"] = uid
	data["username"] = username
	data["uname"] = uname
	data["quhao"] = quhao
	data["phone"] = phone
	data["face"] = face
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update(project, uid, username, uname, face, quhao, phone interface{}) bool {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{}
	where["project"] = project
	where["uid"] = uid
	db.Where(where)
	data := map[string]interface{}{}
	data["username"] = username
	data["uname"] = uname
	data["face"] = face
	data["quhao"] = quhao
	data["phone"] = phone
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update_uname(project, uid, uname interface{}) bool {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{}
	where["project"] = project
	where["uid"] = uid
	db.Where(where)
	data := map[string]interface{}{}
	data["uname"] = uname
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update_face(project, uid, face interface{}) bool {
	db := tuuz.DB().Table(table)
	where := map[string]interface{}{}
	where["project"] = project
	where["uid"] = uid
	db.Where(where)
	data := map[string]interface{}{}
	data["face"] = face
	db.Data(data)
	_, err := db.Update()
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

func Api_select(project interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"project": project,
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

func Api_search_byUsername(project, username interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"project":  project,
		"username": username,
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
