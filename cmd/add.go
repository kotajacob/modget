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
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/util"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <MODID>",
	Short: "Download and install a mod based on its MODID.",
	Run: func(cmd *cobra.Command, args []string) {
		// Convert args to int list of modids
		var mods []int
		for i := 0; i < len(args); i++ {
			id, err := strconv.Atoi(args[i])
			mods = append(mods, id)
			if err != nil {
				fmt.Printf("Failed to read MODID: %v\n", args[i])
				os.Exit(1)
			}
		}
		// Exit if no mods listed
		if len(mods) == 0 {
			fmt.Println("modget add requires at least one MODID")
			os.Exit(1)
		}
		for _, mod := range mods {
			err := Add(mod, "", "")
			if err != nil {
				fmt.Printf("Failed to add mod: %v\n", mod)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	// addCmd.Flags().BoolP("version", "v", false, "Limit install for a specific minecraft version.")
}

// Check if the passed loader string is a valid modloader
func validateModLoader(loader string) bool {
	loader = strings.ToLower(loader)
	if util.ModLoaders[loader] {
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
	conflicts := util.ModLoaders.Copy()
	conflicts = conflicts.RemoveString(loader)
	var matches []curse.File
	for _, file := range files {
		// create a string set of the GameVersions for the file
		var fileVersions = make(util.StringSet)
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

// Print some debug info about a list of curse.File
func debugFilePrint(files []curse.File) {
	for _, file := range files {
		fmt.Println(file.FileName)
		fmt.Println(file.FileDate)
		fmt.Println(file.Id)
		for _, fileVersion := range file.GameVersion {
			fmt.Println(fileVersion)
		}
	}
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
	if version != "" {
		files = versionFilter(files, version)
		if !validateMinecraftVersion(version, mcVersions) {
			fmt.Println("Warning: Minecraft Version entered is not recognized!")
		}
	}
	if loader != "" {
		files = loaderFilter(files, loader)
		if !validateModLoader(loader) {
			fmt.Println("Warning: Modloader entered is not recognized!")
		}
	}
	debugFilePrint(files)
	return nil
}
