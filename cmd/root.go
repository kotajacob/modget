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
	"strings"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/database"
	"github.com/dustin/go-humanize"
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

// show prints a list of mods that will be added, deleted, or updated.
// mode should be "added", "deleted", or "updated"
func show(mods []database.Mod, mode string) {
	fmt.Printf("The following mods will be %s:\n", mode)
	var s string
	var d int
	for _, mod := range mods {
		s += " " + mod.Slug
		d += mod.FileLength
	}
	fmt.Printf("%v\n", s)
	if mode == "deleted" {
		fmt.Printf("After this operation, %s of disk space will be freed.\n", humanize.Bytes(uint64(d)))
	} else {
		fmt.Printf("After this operation, %s of additional disk space will be used.\n", humanize.Bytes(uint64(d)))
	}
}

// ask prompts the user with a Yes/No question about continuing
func ask() bool {
	fmt.Printf("Do you want to continue? [Y/n] ")
	var answer string
	fmt.Scanln(&answer)
	answer = strings.ToLower(answer)
	if answer == "y" || answer == "" {
		return true
	}
	return false
}

// get downloads a list of mods and updates a Database
func get(mods []database.Mod, path string, db *database.Database) error {
	for i, mod := range mods {
		p := filepath.Join(filepath.Dir(path), mod.FileName)
		fmt.Printf("Get:%d %v\n", i, mod.DownloadURL)
		err := curse.Download(mod.DownloadURL, p)
		if err != nil {
			return fmt.Errorf("failed to download mod: %s: %v\n", mod.FileName, err)
		}
		db.Add(mod)
	}
	return nil
}

// remove deletes a list of local mods and updates a Database
func remove(mods []database.Mod, path string, db *database.Database) error {
	for _, mod := range mods {
		fmt.Printf("Deleted: %v\n", mod.FileName)
		err := os.Remove(filepath.Join(path, mod.FileName))
		if err != nil {
			return fmt.Errorf("remove mod: %s: %v\n", mod.FileName, err)
		}
		db.Del(mod.ID)
	}
	return nil
}

// debugFilePrint shows debug info about a curse.File
func debugFilePrint(file curse.File) {
	fmt.Println(file.FileName)
	fmt.Println(file.FileDate)
	fmt.Println(file.ID)
	for _, fileVersion := range file.GameVersion {
		fmt.Println(fileVersion)
	}
}

// debugAddonPrint shows debug info about a curse.Addon
func debugAddonPrint(addon curse.Addon) {
	fmt.Println(addon.Name)
	fmt.Println(addon.Slug)
	fmt.Println(addon.ID)
	fmt.Printf("%d\n", int(addon.DownloadCount))
}

// debugModPrint shows debug info about a database.Mod
func debugModPrint(mod database.Mod) {
	fmt.Println(mod.Name)
	fmt.Println(mod.Slug)
	fmt.Println(mod.ID)
	fmt.Println(mod.FileName)
	fmt.Println(mod.FileLength)
}
