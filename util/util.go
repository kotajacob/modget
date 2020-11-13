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
package util

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"git.sr.ht/~kota/modget/curse"
)

type StringSet map[string]bool

// Check if two StringSets conflict
func (a StringSet) Conflicts(b StringSet) bool {
	for k, _ := range a {
		if b[k] {
			return true
		}
	}
	return false
}

// Copy a StringSet
func (a StringSet) Copy() StringSet {
	copy := make(StringSet)
	for k, v := range a {
		copy[k] = v
	}
	return copy
}

// Remove a string from a StringSet
func (a StringSet) RemoveString(s string) StringSet {
	if a[s] {
		delete(a, s)
	}
	return a
}

// List of valid modloaders because curseforge doesn't provide one...
// dont set to false instead use delete()
var ModLoaders = StringSet{
	"forge":      true,
	"fabric":     true,
	"liteloader": true,
	"rift":       true,
}

// Check if the passed loader string is a valid modloader
func ValidateModLoader(loader string) bool {
	loader = strings.ToLower(loader)
	if ModLoaders[loader] {
		return true
	} else {
		return false
	}
}

// Check if the passed mc version string is a valid Minecraft Version
func ValidateMinecraftVersion(version string, mcVersions []curse.MinecraftVersion) bool {
	for _, mcVersion := range mcVersions {
		if mcVersion.VersionString == version {
			return true
		}
	}
	return false
}

// Attempts to filer a list of Files and return only those that are
// "compatible" with a specified loader. Unfortunately curseforge doesn't allow
// mod authors to select a loader from a dropdown when they upload a file.
// Instead mod authors have taken to adding a loader as a "GameVersion"
// variable alongside minecraft version numbers. This of course extends to the
// data we can pull from the API so it kinda just sucks and this function will
// never be perfect until curseforge fixes this issue.
func LoaderFilter(files []curse.File, loader string) []curse.File {
	loader = strings.ToLower(loader)
	conflicts := ModLoaders.Copy()
	conflicts = conflicts.RemoveString(loader)
	var matches []curse.File
	for _, file := range files {
		// create a string set of the GameVersions for the file
		var fileVersions = make(StringSet)
		for _, fileVersion := range file.GameVersion {
			fileVersion = strings.ToLower(fileVersion)
			fileVersions[fileVersion] = true
		}
		// add non-conflicting files
		if !fileVersions.Conflicts(conflicts) {
			matches = append(matches, file)
		}
	}
	return matches
}

// Filters a list of Files returning only the ones that match the Minecraft Version.
func VersionFilter(files []curse.File, version string) []curse.File {
	var matchFiles []curse.File
	for _, file := range files {
		for _, fileVersion := range file.GameVersion {
			if fileVersion == version {
				matchFiles = append(matchFiles, file)
			}
		}
	}
	return matchFiles
}

// Sorts a list of files returning them in order from newest to old.
func TimeSort(files []curse.File) []curse.File {
	sort.SliceStable(files, func(i, j int) bool {
		iTime, err := time.Parse(time.RFC3339, files[i].FileDate)
		if err != nil {
			return false
		}
		jTime, err := time.Parse(time.RFC3339, files[j].FileDate)
		if err != nil {
			return false
		}
		return iTime.After(jTime)
	})
	return files
}

// GetModid takes a string, which is meant to be an addon's slug and attempts
// to convert it to a MODID. It returns an error on failure.
func GetModid(s string) (int, error) {
	var search curse.Search
	search.GameId = 432     // Set game to minecraft
	search.SectionId = 6    // Set section to mods
	search.SearchFilter = s // Search string
	addons, err := curse.AddonSearch(search)
	if err != nil {
		return 0, err
	}
	for _, addon := range addons {
		if addon.Slug == s {
			return addon.Id, nil
		}
	}
	err = errors.New("Could not find: " + s)
	return 0, err
}

// Print some debug info about a curse.File
func DebugFilePrint(file curse.File) {
	fmt.Println(file.FileName)
	fmt.Println(file.FileDate)
	fmt.Println(file.Id)
	for _, fileVersion := range file.GameVersion {
		fmt.Println(fileVersion)
	}
}

// Print some debug info about a curse.Addon
func DebugAddonPrint(addon curse.Addon) {
	fmt.Println(addon.Name)
	fmt.Println(addon.Slug)
	fmt.Println(addon.Id)
	fmt.Printf("%d\n", int(addon.DownloadCount))
}
