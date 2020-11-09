package command

import (
	"fmt"
	"strings"

	"git.sr.ht/~kota/modget/curse"
)

// Check if the passed loader string is a valid modloader
func validateModloader(loader string) bool {
	loader = strings.ToLower(loader)
	modLoaders := []string{"fabric", "forge", "rift"}
	for _, modLoader := range modLoaders {
		if modLoader == loader {
			return true
		}
	}
	return false
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

// Filters a list of Files returning only the ones that match the Minecraft Version
func matchFiles(files []curse.File, version string, loader string) []curse.File {
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
	if !validateMinecraftVersion(version, mcVersions) {
		fmt.Println("Warning: Minecraft Version entered is not recognized!")
	}
	if !validateModloader(loader) {
		fmt.Println("Warning: Modloader entered is not recognized!")
	}
	files = matchFiles(files, version, loader)
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
