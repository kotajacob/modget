package curse

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/* https://twitchappapi.docs.apiary.io
 * official API documentation
 */

func get(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
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

// Download gets a file at a specific url and saves it to a specific filename.
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

// AddonInfo fetches a detailed json response for a specific addon.
func AddonInfo(modId int) (Addon, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId)
	response, err := get(url)
	addon, err := parseAddonInfo(response)
	return addon, err
}

// AddonFileInformation fetches a detailed json response for a specific file of
// a specific addon.
func AddonFileInformation(modId, fileId int) (File, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/file/" + fmt.Sprintf("%d", fileId)
	response, err := get(url)
	file, err := parseAddonFileInformation(response)
	return file, err
}

// AddonFileDownloadURL fetches only the download url as a string for a
// specific file of a specific addon. Note you might already have this URL if
// you did GetAddonFiles or GetAddonFileInformation.
func AddonFileDownloadURL(modId, fileId int) (string, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/file/" + fmt.Sprintf("%d", fileId) + "/download-url"
	response, err := get(url)
	downloadURL := parseAddonFileDownloadURL(response)
	return downloadURL, err
}

// AddonFiles fetches a detailed json response for all the files of a specific
// addon.
func AddonFiles(modId int) ([]File, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + fmt.Sprintf("%d", modId) + "/files"
	response, err := get(url)
	files, err := parseAddonFiles(response)
	return files, err
}

// MinecraftVersionList returns a list of curseforge's current supported
// released of minecraft.
func MinecraftVersionList() ([]MinecraftVersion, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/minecraft/version"
	response, err := get(url)
	minecraftVersions, err := parseMinecraftVersionList(response)
	return minecraftVersions, err
}

// ModloaderList returns a list of curseforge's current supported modloaders.
func ModloaderList() ([]Modloader, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/minecraft/modloader"
	response, err := get(url)
	modloaders, err := parseModloaderList(response)
	return modloaders, err
}
