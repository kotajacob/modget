package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/* https://twitchappapi.docs.apiary.io
 * official API documentation
 */

func get(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	check(err)

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	check(err)

	return respBody
}

func getAddonInfo(id int) []byte {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", id)
	return get(url)
}

func getAddonFiles(id int) []byte {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", id) + "/files"
	return get(url)
}
