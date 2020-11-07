package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

func store(url, filename string) {
	response, err := http.Get(url)
	check(err)
	defer response.Body.Close()

	file, err := os.Create(filename)
	check(err)
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	check(err)
}

func getAddonInfo(modId int) []byte {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId)
	return get(url)
}

func getAddonFileInformation(modId, fileId int) []byte {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/file/" + fmt.Sprintf("%d", fileId)
	return get(url)
}

func getAddonFileDownloadURL(modId, fileId int) []byte {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/file/" + fmt.Sprintf("%d", fileId) + "/download-url"
	return get(url)
}

func getAddonFiles(modId int) []byte {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/files"
	return get(url)
}
