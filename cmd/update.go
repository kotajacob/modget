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
	"time"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/database"
	"git.sr.ht/~kota/modget/filter"
	"git.sr.ht/~kota/modget/input"
	"git.sr.ht/~kota/modget/printer"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update [mod]...",
	Aliases: []string{"u"},
	Short:   "Check installed mod(s) and prompt to install any new mods.",
	Run:     update,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&minecraft, "minecraft", "m", "", "Limit install for a specific minecraft version.")
	updateCmd.Flags().StringVarP(&loader, "loader", "l", "", "Limit install for a specific minecraft mod loader.")
}

// update does the following
// 1. Load the database
// 2. Read input and make a list of selected mods (or all mods)
// 3. If version or loader changed, build list of incompatible mods
// 4. Build list of the latest files for minecraft version and loader
// 5. Remove old versions
// 6. Add new versions
func update(cmd *cobra.Command, args []string) {
	updateMods := make(map[int]*database.Mod)
	var updateIDs []int
	fmt.Printf("Reading database... ")
	db, err := database.Load(filepath.Join(path, ".modget"))
	if err != nil {
		fmt.Printf("failed to open database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done")
	ids, err := input.Slug(args, db)
	if err != nil {
		fmt.Printf("failed read input: %v\n", err)
		os.Exit(1)
	}
	if len(ids) == 0 { // select all mods
		for id := range db.Mods {
			ids = append(ids, id)
		}
	}
	fmt.Printf("Checking for updates... ")
	for _, id := range ids {
		// Skip if the mod has a hold
		if db.Mods[id].Hold == true {
			continue
		}
		addon, err := curse.AddonInfo(id)
		if err != nil {
			fmt.Printf("%v\n%v\n", id, err)
			continue
		}
		file, err := filter.FindFile(id, minecraft, loader)
		if err != nil {
			fmt.Printf("%v\n%v\n", id, err)
			continue
		}
		mTime, err := time.Parse(time.RFC3339, db.Mods[id].FileDate)
		if err != nil {
			fmt.Printf("failed to parse time: %v\n%v\n", id, err)
			os.Exit(1)
		}
		fTime, err := time.Parse(time.RFC3339, file.FileDate)
		if err != nil {
			fmt.Printf("failed to parse time: %v\n%v\n", id, err)
			os.Exit(1)
		}
		// Is the file newer than the installed mod?
		if fTime.After(mTime) {
			updateMods[id] = database.NewMod(addon, file)
			updateIDs = append(updateIDs, id)
		}
	}
	fmt.Println("Done")
	if len(updateIDs) == 0 {
		fmt.Println("your mods are up to date")
		os.Exit(0)
	}
	printer.Changes(updateIDs, "updated", updateMods)
	if !printer.Continue() {
		os.Exit(0)
	}
	for id, mod := range updateMods {
		fmt.Printf("Delete: %v\n", mod.FileName)
		err := os.Remove(filepath.Join(path, mod.FileName))
		if err != nil {
			if err != nil {
				fmt.Printf("failed to remove mod: %v\n", err)
			}
		}
		db.Del(id)
	}
	for id, mod := range updateMods {
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
