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
	MinecraftVersion string
	Loader           string
	Path             string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <MODID/Slug>",
	Short: "Download and install a mod based on its MODID or Slug.",
	Run: func(cmd *cobra.Command, args []string) {
		var files []curse.File
		db, err := findDatabase()
		if err != nil {
			fmt.Printf("Failed to open database: %v\n", err)
			os.Exit(1)
		}
		ids := toId(args)
		if len(ids) == 0 {
			fmt.Println("modget add requires at least one MODID or Slug")
			os.Exit(1)
		}
		for _, id := range ids {
			file, err := findId(id)
			if err != nil {
				fmt.Printf("Failed to find mod: %v\n%v\n", id, err)
				os.Exit(1)
			}
			files = append(files, file)
		}
		for _, file := range files {
			err := get(file)
			if err != nil {
				fmt.Printf("Failed to download file: %v\n%v\n", file.FileName, err)
				os.Exit(1)
			}
			db.Files = append(db.Files, file)
		}
		err = db.Write(Path)
		if err != nil {
			fmt.Printf("Failed to write database: %v\n", err)
			// TODO: remove failed downloaded files
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&MinecraftVersion, "minecraft", "m", "", "Limit install for a specific minecraft version.")
	addCmd.Flags().StringVarP(&Loader, "loader", "l", "", "Limit install for a specific minecraft mod loader.")
	addCmd.Flags().StringVarP(&Path, "path", "p", "", "Mod install location.")
}

// Convert a list of strings to MODIDs
func toId(s []string) []int {
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
	if Path == "" {
		Path = "."
	}
	err := util.EnsureDir(Path)
	if err != nil {
		return db, err
	}
	Path = filepath.Join(Path, ".modget")
	db, err = database.Load(Path)
	return db, err
}

// Find returns a curse.File for a MODID. It ensures the file matches the
// correct Minecraft version and Loader. Additionally it warns the user if the
// enter an unknown version or loader.
func findId(id int) (curse.File, error) {
	files, err := curse.AddonFiles(id)
	// Validate the modloader and mc version
	mcVersions, err := curse.MinecraftVersionList()
	if MinecraftVersion != "" {
		files = util.VersionFilter(files, MinecraftVersion)
		if !util.ValidateMinecraftVersion(MinecraftVersion, mcVersions) {
			fmt.Println("Warning: Minecraft Version entered is not recognized!")
		}
	}
	if Loader != "" {
		files = util.LoaderFilter(files, Loader)
		if !util.ValidateModLoader(Loader) {
			fmt.Println("Warning: Modloader entered is not recognized!")
		}
	}
	files = util.TimeSort(files)
	if len(files) == 0 {
		err = errors.New("File not found for those search terms.")
	}
	return files[0], err
}

func get(f curse.File) error {
	// TODO: Make this toggle-able with a verbose flag
	p := filepath.Join(filepath.Dir(Path), f.FileName)
	util.DebugFilePrint(f)
	err := curse.Download(f.DownloadUrl, p)
	return err
}
