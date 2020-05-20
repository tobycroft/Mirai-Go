package ProjectAction

import (
	"github.com/gohouse/gorose/v2"
	"main.go/common/Model/ProjectModel"
)

func App_project(project interface{}) (bool, gorose.Data) {
	proj := ProjectModel.Api_find_avail(project)
	if len(proj) > 0 {
		return true, proj
	} else {
		return false, nil
	}
}

func App_project_avail(project interface{}) bool {
	b, _ := App_project(project)
	return b
}
