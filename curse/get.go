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

func get(url string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func Download(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func GetAddon(modId int) ([]byte, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId)
	return get(url)
}

func GetAddonFileInformation(modId, fileId int) ([]byte, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/file/" + fmt.Sprintf("%d", fileId)
	return get(url)
}

func GetAddonFileDownloadURL(modId, fileId int) ([]byte, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/file/" + fmt.Sprintf("%d", fileId) + "/download-url"
	return get(url)
}

func GetAddonFiles(modId int) ([]byte, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/files"
	return get(url)
}
