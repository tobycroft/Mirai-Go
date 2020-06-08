package MirRoute

import (
	"fmt"
	"main.go/app/http/v1/model/LogModel"
)

func Do(qq, json string) {
	fmt.Println(qq)
	fmt.Println(json)
	LogModel.Api_insert(qq, json)
}

func Message(qq, json string) {

}

func Notice(qq, json string) {

}

func Request(qq, json string) {

}
