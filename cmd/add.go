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
	"git.sr.ht/~kota/modget/filter"
	"git.sr.ht/~kota/modget/printer"
	"git.sr.ht/~kota/modget/slug"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add mod...",
	Aliases: []string{"a"},
	Short:   "Download and install mod(s) based on MODID or Slug.",
	Run:     add,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&minecraft, "minecraft", "m", "", "Limit install for a specific minecraft version.")
	addCmd.Flags().StringVarP(&loader, "loader", "l", "", "Limit install for a specific minecraft mod loader.")
}

func add(cmd *cobra.Command, args []string) {
	mods := make(map[int]*database.Mod)
	if len(args) == 0 {
		fmt.Println("modget add requires at least one MODID or Slug")
		os.Exit(1)
	}
	fmt.Printf("Reading database... ")
	db, err := database.Load(filepath.Join(path, ".modget"))
	if err != nil {
		fmt.Println("not found!")
		if minecraft == "" {
			minecraft = printer.Ask("minecraft version")
		}
		if loader == "" {
			loader = printer.Ask("modloader")
		}
		db = &database.Database{
			Version:   Version,
			Minecraft: minecraft,
			Loader:    loader,
			Mods:      make(map[int]*database.Mod),
		}
	}
	minecraft = db.Minecraft
	loader = db.Loader
	fmt.Println("Done")
	ids, err := slug.Slug(args, db)
	if err != nil {
		fmt.Printf("failed read input: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Finding Mods... ")
	for _, id := range ids {
		addon, err := curse.AddonInfo(id)
		file, err := filter.FindFile(id, minecraft, loader)
		if err != nil {
			fmt.Printf("\n%v\n", err)
			os.Exit(1)
		}
		mods[id] = database.NewMod(addon, file)
	}
	fmt.Println("Done")
	if len(mods) == 0 {
		fmt.Println("no mods found with those search terms")
		os.Exit(0)
	}
	printer.Show(ids, "added", mods)
	if !printer.Prompt() {
		os.Exit(0)
	}
	for id, mod := range mods {
		p := filepath.Join(filepath.Dir(path), mod.FileName)
		fmt.Printf("Get:%d %v\n", id, mod.DownloadURL)
		err := curse.Download(mod.DownloadURL, p)
		if err != nil {
			fmt.Printf("failed to download file: %v\n", err)
			os.Exit(1)
		}
		db.Add(id, mod)
	}
	fmt.Printf("Updating database... ")
	err = db.Write(filepath.Join(path, ".modget"))
	if err != nil {
		fmt.Printf("failed to write database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done")
}
