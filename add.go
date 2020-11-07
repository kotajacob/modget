package main

import (
	"fmt"
	"strconv"
)

func add(mods []string, mc string, loader string) {
	// DEBUG //
	// fmt.Println("  tail:", mods)
	// fmt.Println("  mc: ", mc)
	// fmt.Println("  loader: ", loader)

	/* Currently we only handle the first MODID in the list.
	 * Additionally there is no error checking about if it is a valid MODID.
	 */
	modid, err := strconv.Atoi(mods[0])
	check(err)

	addon := parseAddonInfo(getAddonInfo(modid))
	fmt.Println("Addon Name: " + addon.Name)
	fmt.Println("Addon Summary: " + addon.Summary)
	fmt.Printf("Popularity: %f\n", addon.PopularityScore)
	fmt.Printf("Downloads: %f\n", addon.DownloadCount)
	fmt.Printf("Default File ID: %v\n", addon.DefaultFileId)
	fmt.Println("")

	file := parseAddonFileInformation(getAddonFileInformation(modid, addon.DefaultFileId))
	fmt.Println("File ID: " + strconv.Itoa(file.Id))
	fmt.Println("File Date: " + file.FileDate)
	fmt.Println("File Name: " + file.FileName)
	for ii := 0; ii < len(file.GameVersion); ii++ {
		fmt.Println("Game Version: " + file.GameVersion[ii])
	}
	fmt.Println("URL: " + file.DownloadUrl)
	fmt.Println("")
	fmt.Println("Downloading: " + strconv.Itoa(file.Id))

	store(file.DownloadUrl, file.FileName)
}
