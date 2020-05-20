package txlive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//HTTPGet Get
func HTTPGet(baseurl string, param map[string]string) []byte {
	paramstr := "?"
	for k, v := range param {
		paramstr += k + "=" + v + "&"
	}
	fmt.Println(baseurl + paramstr)
	resp, err := http.Get(baseurl + paramstr)
	if err != nil {
		// handle error
		fmt.Println("err1", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("err2", err)
	}

	// fmt.Println(string(body))
	return []byte(body)

}

//HTTPPost Post
func HTTPPost(baseURL string, rbody interface{}) []byte {
	fmt.Println("post data:", rbody)
	b, err := json.Marshal(rbody)
	body := bytes.NewBuffer([]byte(b))
	fmt.Println("HTTPPost body", body)
	resp, err := http.Post(baseURL, "application/json", body)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	// fmt.Println(string(body))
	return []byte(resBody)
}
