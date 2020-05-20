package UIAction

import "main.go/app/http/v1/model/UserInfoModel"

func App_userinfo(project, uid interface{}) map[string]interface{} {
	ui := UserInfoModel.Api_find(project, uid)
	if len(ui) > 0 {
		userinfo := map[string]interface{}{
			"uname":    ui["uname"],
			"username": ui["username"],
			"face":     ui["face"],
			"uid":      ui["uid"],
			"exp":      ui["exp"],
			"level":    ui["exp"].(float32) / 1000,
		}
		return userinfo
	} else {
		userinfo := map[string]interface{}{
			"uname":    "火星网友",
			"username": nil,
			"face":     "",
			"level":    0,
			"uid":      0,
			"exp":      0,
		}
		return userinfo
	}
}
