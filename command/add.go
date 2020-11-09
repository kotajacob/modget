package command

import (
	"errors"
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
func matchFiles(files []curse.File, version string, loader string) ([]curse.File, error) {
	var matchFiles []curse.File
	mcVersions, err := curse.MinecraftVersionList()
	if err != nil {
		return nil, err
	}
	if !validateMinecraftVersion(version, mcVersions) {
		fmt.Println("Warning: Minecraft Version entered is not recognized!")
	}
	if !validateModloader(loader) {
		fmt.Println("Warning: Modloader entered is not recognized!")
	}
	for _, file := range files {
		for _, fileVersion := range file.GameVersion {
			if fileVersion == version {
				matchFiles = append(matchFiles, file)
			}
		}
	}
	return matchFiles, nil
}

// Add searches and downloads one or more mods and records the result in the
// database. Additionally it can accept a manually specified mc version and
// loader, or fallback to the default one in the database.
func Add(mods []int, mc string, loader string) error {
	if len(mods) > 0 {
		for i := 0; i < len(mods); i++ {
			files, err := curse.AddonFiles(mods[i])
			if err != nil {
				return err
			}
			files, err = matchFiles(files, mc, loader)
			if err != nil {
				return err
			}
			for _, file := range files {
				fmt.Println(file.FileName)
				fmt.Println(file.FileDate)
				fmt.Println(file.Id)
				for _, fileVersion := range file.GameVersion {
					fmt.Println(fileVersion)
				}
			}
		}
	} else {
		return errors.New("modget add requires at least one MODID")
	}
	return nil
}
