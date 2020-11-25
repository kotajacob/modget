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
	"path/filepath"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/database"
	"github.com/spf13/cobra"
)

var (
	path string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "modget [command]",
	Short: "A package manager for minecraft curseforge mods.",
	Long:  `Modget provides commands for searching, installing, and querying information about mods.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Mod install location.")
}

// get downloads a map of mods and updates a Database
func get(mods map[int]database.Mod, path string, db *database.Database) error {
	for ID, mod := range mods {
		p := filepath.Join(filepath.Dir(path), mod.FileName)
		fmt.Printf("Get:%d %v\n", ID, mod.DownloadURL)
		err := curse.Download(mod.DownloadURL, p)
		if err != nil {
			return err
		}
		db.Add(ID, mod)
	}
	return nil
}

// remove deletes a list of local mods and updates a Database
func remove(IDs []int, path string, db *database.Database) error {
	for _, ID := range IDs {
		fmt.Printf("Deleted: %v\n", db.Mods[ID].FileName)
		err := os.Remove(filepath.Join(path, db.Mods[ID].FileName))
		if err != nil {
			return err
		}
		db.Del(ID)
	}
	return nil
}
