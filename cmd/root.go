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

// showRemove prints a list of mods that will be installed.
func showRemove(mods []database.Mod) {
	fmt.Println("The following mods will be removed:")
	var s string
	var d int
	for _, mod := range mods {
		s += " " + mod.Slug
		d += mod.FileLength
	}
	fmt.Printf("%v\n", s)
	fmt.Printf("After this operation, %s of additional disk space will be freed.\n", humanize.Bytes(uint64(d)))
}

// showNew prints a list of mods that will be installed.
func showNew(addons []curse.Addon, files []curse.File) {
	fmt.Println("The following mods will be installed:")
	var s string
	var d int
	for i, addon := range addons {
		s += " " + addon.Slug
		d += files[i].FileLength
	}
	fmt.Printf("%v\n", s)
	fmt.Printf("After this operation, %s of additional disk space will be used.\n", humanize.Bytes(uint64(d)))
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

// getMods downloads a list of files and updates a Database
func getMods(addons []curse.Addon, files []curse.File, path string, db *database.Database) error {
	for i, file := range files {
		p := filepath.Join(filepath.Dir(path), file.FileName)
		fmt.Printf("Get:%d %v\n", i, file.DownloadURL)
		err := curse.Download(file.DownloadURL, p)
		db.Add(addons[i], file)
		if err != nil {
			return fmt.Errorf("add mod to database: %s: %v\n", file.FileName, err)
		}
	}
	return nil
}

// deleteMods deleted a list of local mods and updates a Database
func deleteMods(mods []database.Mod, path string, db *database.Database) error {
	for _, mod := range mods {
		fmt.Printf("Delete: %v\n", mod.FileName)
		err := os.Remove(filepath.Join(path, mod.FileName))
		if err != nil {
			return fmt.Errorf("delete mod: %s: %v\n", mod.FileName, err)
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
