package curse

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func Download(url, filename string) {
	response, err := http.Get(url)
	check(err)
	defer response.Body.Close()

	file, err := os.Create(filename)
	check(err)
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	check(err)
}

func GetAddonInfo(modId int) []byte {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId)
	return get(url)
}

func GetAddonFileInformation(modId, fileId int) []byte {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/file/" + fmt.Sprintf("%d", fileId)
	return get(url)
}

func GetAddonFileDownloadURL(modId, fileId int) []byte {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/file/" + fmt.Sprintf("%d", fileId) + "/download-url"
	return get(url)
}

func GetAddonFiles(modId int) []byte {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/files"
	return get(url)
}
