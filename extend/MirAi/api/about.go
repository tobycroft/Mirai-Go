package api

import (
	"main.go/config"
	"main.go/tuuz/Net"
)

func About() {
	ret, err := Net.Get(config.API_URL+"/about", nil, nil, nil)
}
