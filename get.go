package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/* https://twitchappapi.docs.apiary.io
 * official API documentation
 */

func GetAddonFiles(id int) []byte {
	client := &http.Client{}

	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", id) + "/files"
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	check(err)

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	check(err)

	return respBody
}
