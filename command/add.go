package command

import (
	"fmt"
	"strings"

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

// Remove a string from a StringSet
func (a StringSet) RemoveString(s string) StringSet {
	if a[s] {
		delete(a, s)
	}
	return a
}

// Get a StringSet of conflicting ModLoaders. Expects loader to be lowercase.
func getConflicts(loader string) StringSet {
	conflicts := ModLoaders
	if ModLoaders[loader] {
		delete(conflicts, loader)
	}
	return conflicts
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
func validateModLoader(loader string) bool {
	loader = strings.ToLower(loader)
	if ModLoaders[loader] {
		return true
	} else {
		return false
	}
}

// Check if the passed mc version string is a valid Minecraft Version
func validateMinecraftVersion(version string, mcVersions []curse.MinecraftVersion) bool {
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
func loaderFilter(files []curse.File, loader string) []curse.File {
	loader = strings.ToLower(loader)
	conflicts := ModLoaders.RemoveString(loader)
	var matchFiles []curse.File
	for _, file := range files {
		// create a string set of the GameVersions for the file
		var fileVersions = make(StringSet)
		for _, fileVersion := range file.GameVersion {
			fileVersion = strings.ToLower(fileVersion)
			fileVersions[fileVersion] = true
		}
		if loader == "" {
			matchFiles = append(matchFiles, file)
			break
		}
		if fileVersions[loader] {
			matchFiles = append(matchFiles, file)
			break
		}
		if !fileVersions.Conflicts(conflicts) {
			matchFiles = append(matchFiles, file)
			break
		}
	}
	return matchFiles
}

// Filters a list of Files returning only the ones that match the Minecraft Version
func versionFilter(files []curse.File, version string) []curse.File {
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

// Add searches and downloads a mod and records the result in the database.
// Additionally it can accept a manually specified mc version and loader, or
// fallback to the default one in the database.
func Add(mod int, version string, loader string) error {
	files, err := curse.AddonFiles(mod)
	if err != nil {
		return err
	}
	// Validate the modloader and mc version
	mcVersions, err := curse.MinecraftVersionList()
	if err != nil {
		return err
	}
	if version != "" && !validateMinecraftVersion(version, mcVersions) {
		fmt.Println("Warning: Minecraft Version entered is not recognized!")
	}
	if loader != "" && !validateModLoader(loader) {
		fmt.Println("Warning: Modloader entered is not recognized!")
	}
	files = versionFilter(files, version)
	files = loaderFilter(files, loader)
	for _, file := range files {
		fmt.Println(file.FileName)
		fmt.Println(file.FileDate)
		fmt.Println(file.Id)
		for _, fileVersion := range file.GameVersion {
			fmt.Println(fileVersion)
		}
	}
	return nil
}
