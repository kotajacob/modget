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
	"errors"
	"fmt"
	"os"
	"strconv"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/util"
	"github.com/spf13/cobra"
)

var (
	MinecraftVersion string
	Loader           string
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
			if err != nil {
				// Attempt to convert slug to modid
				id, err = util.GetModid(args[i])
				if err != nil {
					fmt.Printf("Failed to find: %v\n", args[i])
					os.Exit(1)
				}
			}
			mods = append(mods, id)
		}
		// Exit if no mods listed
		if len(mods) == 0 {
			fmt.Println("modget add requires at least one MODID")
			os.Exit(1)
		}
		for _, mod := range mods {
			err := add(mod, MinecraftVersion, Loader)
			if err != nil {
				fmt.Printf("Failed to add mod: %v\n", mod)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&MinecraftVersion, "minecraft", "m", "", "Limit install for a specific minecraft version.")
	addCmd.Flags().StringVarP(&Loader, "loader", "l", "", "Limit install for a specific minecraft mod loader.")
}

// Add searches and downloads a mod and records the result in the database.
// Additionally it can accept a manually specified mc version and loader, or
// fallback to the default one in the database.
func add(mod int, version string, loader string) error {
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
		files = util.VersionFilter(files, version)
		if !util.ValidateMinecraftVersion(version, mcVersions) {
			fmt.Println("Warning: Minecraft Version entered is not recognized!")
		}
	}
	if loader != "" {
		files = util.LoaderFilter(files, loader)
		if !util.ValidateModLoader(loader) {
			fmt.Println("Warning: Modloader entered is not recognized!")
		}
	}
	files = util.TimeSort(files)
	if len(files) == 0 {
		errors.New("File not found for those search terms.")
	}
	selected := files[0]
	// TODO: Make this toggle-able with a verbose flag
	util.DebugFilePrint(selected)
	curse.Download(selected.DownloadUrl, selected.FileName)
	return nil
}
