package main

import (
	"fmt"
	"strconv"
)

func add(mods []string, mc string, loader string) {
	// fmt.Println("  tail:", mods)
	// fmt.Println("  mc: ", mc)
	// fmt.Println("  loader: ", loader)

	/* Currently we only handle the first MODID in the list.
	 * Additionally there is no error checking about if it is a valid MODID.
	 */
	if mc == "" && loader == "" {
		addDefaultFile(mods)
	}
}

func addDefaultFile(mods []string) {
	modid, err := strconv.Atoi(mods[0])
	check(err)

	addon := parseAddonInfo(getAddonInfo(modid))
	fmt.Println("Addon Name: " + addon.Name)
	fmt.Printf("Default File ID: %v\n", addon.DefaultFileId)

	file := parseAddonFileInformation(getAddonFileInformation(modid, addon.DefaultFileId))
	fmt.Println("File Date: " + file.FileDate)
	fmt.Println("File Name: " + file.FileName)
	for i := 0; i < len(file.GameVersion); i++ {
		fmt.Println("Game Version: " + file.GameVersion[i])
	}
	fmt.Println("URL: " + file.DownloadUrl)

	store(file.DownloadUrl, file.FileName)
}
