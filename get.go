package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get(id int) []byte {
	client := &http.Client{}

	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", id) + "/files"
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	check(err)

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp_body
}
