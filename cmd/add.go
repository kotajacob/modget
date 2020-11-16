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
	"path/filepath"
	"strconv"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/database"
	"git.sr.ht/~kota/modget/util"
	"github.com/spf13/cobra"
)

var (
	minecraftVersion string
	loader           string
	path             string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <MODID/Slug>",
	Short: "Download and install a mod based on its MODID or Slug.",
	Run: func(cmd *cobra.Command, args []string) {
		var files []curse.File
		if len(args) == 0 {
			fmt.Println("modget add requires at least one MODID or Slug")
			os.Exit(1)
		}
		fmt.Printf("Reading database... ")
		db, err := findDatabase()
		if err != nil {
			fmt.Printf("Failed to open database: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Done")
		ids := toID(args)
		fmt.Printf("Finding Mods... ")
		for _, id := range ids {
			file, err := findFile(id)
			if err != nil {
				fmt.Printf("Failed to find mod: %v\n%v\n", id, err)
				os.Exit(1)
			}
			files = append(files, file)
		}
		// for _, file := range files {
		// 	util.DebugFilePrint(file)
		// }
		fmt.Println("Done")
		showMods(files)
		if !util.Ask() {
			os.Exit(0)
		}
		err = getMods(files, db)
		if err != nil {
			fmt.Printf("Failed to download file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Updating database... ")
		err = db.Write(path)
		if err != nil {
			fmt.Printf("Failed to write database: %v\n", err)
			// TODO: remove failed downloaded files
			os.Exit(1)
		}
		fmt.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&minecraftVersion, "minecraft", "m", "", "Limit install for a specific minecraft version.")
	addCmd.Flags().StringVarP(&loader, "loader", "l", "", "Limit install for a specific minecraft mod loader.")
	addCmd.Flags().StringVarP(&path, "path", "p", "", "Mod install location.")
}

// Convert a list of strings to MODIDs
func toID(s []string) []int {
	// Convert string to int list of modids
	var mods []int
	for i := 0; i < len(s); i++ {
		id, err := strconv.Atoi(s[i])
		if err != nil {
			// Attempt to convert slug to modid
			id, err = util.GetModid(s[i])
			if err != nil {
				fmt.Printf("Failed to find: %v\n", s[i])
				os.Exit(1)
			}
		}
		mods = append(mods, id)
	}
	return mods
}

// Find the .modget database at the path. Create the database if missing.
func findDatabase() (database.Database, error) {
	var db database.Database
	if path == "" {
		path = "."
	}
	err := util.EnsureDir(path)
	if err != nil {
		return db, err
	}
	path = filepath.Join(path, ".modget")
	db, err = database.Load(path)
	return db, err
}

// findFile returns a curse.File for a MODID. It ensures the file matches the
// correct Minecraft version and Loader. Additionally it warns the user if the
// enter an unknown version or loader.
func findFile(id int) (curse.File, error) {
	files, err := curse.AddonFiles(id)
	// Validate the modloader and mc version
	mcVersions, err := curse.MinecraftVersionList()
	if minecraftVersion != "" {
		files = util.VersionFilter(files, minecraftVersion)
		if !util.ValidateMinecraftVersion(minecraftVersion, mcVersions) {
			fmt.Println("warning: Minecraft Version entered is not recognized")
		}
	}
	if loader != "" {
		files = util.LoaderFilter(files, loader)
		if !util.ValidateModLoader(loader) {
			fmt.Println("warning: Modloader entered is not recognized")
		}
	}
	files = util.TimeSort(files)
	if len(files) == 0 {
		err = errors.New("file not found for those search terms")
	}
	return files[0], err
}

func showMods(files []curse.File) {
	fmt.Println("The following mods will be installed:")
	var s string
	var d int
	for _, file := range files {
		s += " " + file.FileName
		d += file.FileLength
	}
	fmt.Printf("%v\n", s)
	fmt.Printf("After this operation, %d of additional disk space will be used.\n", d)
}

func getMods(files []curse.File, db database.Database) error {
	for i, file := range files {
		p := filepath.Join(filepath.Dir(path), file.FileName)
		fmt.Printf("Get:%d %v\n", i, file.DownloadURL)
		err := curse.Download(file.DownloadURL, p)
		db.Files = append(db.Files, file)
		if err != nil {
			return err
		}
	}
	return nil
}
