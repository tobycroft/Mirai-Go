package AuthAction

import (
	"fmt"
	"main.go/app/http/v1/model/UserInfoModel"
	"main.go/common/Action/ProjectAction"
	"main.go/common/Model/UserTokenModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
	"time"
)

func TokenAction(project, uid, token, ip interface{}) bool {
	//return true
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("----painc_recover----", tuuz.FUNCTION_ALL(), err)
		}
	}()
	if len(UserTokenModel.Api_find_avail(project, uid, token)) > 0 {
		return true
	} else {
		ok, proj := ProjectAction.App_project(project)
		if ok {
			auth_url := proj["auth_url"].(string)
			ts := Calc.Any2String(time.Now().Unix())
			query := map[string]interface{}{
				"uid":   uid,
				"token": token,
				"ip":    ip,
				"ts":    ts,
				"sign":  Calc.Md5(ts + Calc.Any2String(token)),
			}
			ret, err := Net.Post(auth_url, nil, query, nil, nil)
			if err != nil {
				Log.Crrs(err, tuuz.FUNCTION_ALL())
				return false
			} else {
				rtt, err := Jsong.JObject(ret)
				if err != nil {
					Log.Crrs(err, tuuz.FUNCTION_ALL())
					return false
				} else {
					if rtt["code"].(float64) == 0 {
						data, err := Jsong.ParseObject(rtt["data"])
						if err != nil {

						} else {
							if len(UserInfoModel.Api_find(project, uid)) > 0 {
								UserInfoModel.Api_update(project, uid, data["username"], data["uname"], data["face"], nil, nil)
							} else {
								UserInfoModel.Api_insert(project, uid, data["username"], data["uname"], data["face"], nil, nil)
							}
						}
						return true
					} else {
						return false
					}
				}
			}
		} else {
			return false
		}
	}
}
