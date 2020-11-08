package commands

import (
	"errors"
	"fmt"
	"strconv"

	"git.sr.ht/~kota/modget/curse"
)

// Add searches and downloads one or more mods and records the result in the
// database. Additionally it can accept a manually specified mc version and
// loader, or fallback to the default one in the database.
func Add(mods []string, mc string, loader string) error {
	if len(mods) > 0 {
		if mc == "" && loader == "" {
			for i := 0; i < len(mods); i++ {
				err := addDefaultFile(mods[i])
				if err != nil {
					return err
				}
			}
		}
	} else {
		return errors.New("modget add requires at least one MODID")
	}
	return nil
}

func addDefaultFile(mod string) error {
	modid, err := strconv.Atoi(mod)
	if err != nil {
		return err
	}

	addon, err := curse.AddonInfo(modid)
	if err != nil {
		return err
	}
	fmt.Println("Addon Name: " + addon.Name)
	fmt.Printf("Default File ID: %v\n", addon.DefaultFileId)

	file, err := curse.AddonFileInformation(modid, addon.DefaultFileId)
	if err != nil {
		return err
	}
	fmt.Println("File Date: " + file.FileDate)
	fmt.Println("File Name: " + file.FileName)
	for i := 0; i < len(file.GameVersion); i++ {
		fmt.Println("Game Version: " + file.GameVersion[i])
	}

	fmt.Println("URL: " + file.DownloadUrl)
	err = curse.Download(file.DownloadUrl, file.FileName)
	return err
}
