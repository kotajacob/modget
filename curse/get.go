/*
Copyright © 2020 Dakota Walsh

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

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

// Search represents a new curseforge search. Note not all fields are
// implemented yet, see AddonSearch for details.
type Search struct {
	CategoryId   int    // categoryId: 0
	GameId       int    // gameId: 432 = minecraft, 1 = wow
	GameVersion  string // gameVersion: 1.12.2
	Index        int    // index: 0
	PageSize     int    // pageSize: 25
	SearchFilter string // searchFilter: appleskin (the actual search string)
	SectionId    int    // sectionId: 6 = mods, 4561 = resource packs, 4471 = modpacks, 4560 = worlds
	Sort         int    // sort: 0
}

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

// AddonSearch fetches a list of Addons based on a Search.
func AddonSearch(s Search) ([]Addon, error) {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/search?"
	url += "gameId="
	url += fmt.Sprintf("%d", s.GameId)
	url += "&sectionId="
	url += fmt.Sprintf("%d", s.SectionId)
	// url += "categoryId="
	// url += fmt.Sprintf("%d", s.CategoryId)
	// url += "&gameVersion="
	// url += fmt.Sprintf("%v", s.GameVersion)
	// url += "&index="
	// url += fmt.Sprintf("%d", s.Index)
	// url += "&pageSize="
	// url += fmt.Sprintf("%d", s.PageSize)
	// url += "&sort="
	// url += fmt.Sprintf("%d", s.Sort)
	url += "&searchFilter="
	url += fmt.Sprintf("%v", s.SearchFilter)
	response, err := get(url)
	addons, err := parseAddonSearch(response)
	return addons, err
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
