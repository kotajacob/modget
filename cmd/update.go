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

	"git.sr.ht/~kota/modget/database"
	"git.sr.ht/~kota/modget/filter"
	"git.sr.ht/~kota/modget/slug"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update [mod]...",
	Aliases: []string{"u"},
	Short:   "Check installed mod(s) and prompt to install any new mods.",
	Run: func(cmd *cobra.Command, args []string) {
		var mods []database.Mod
		var updates []database.Mod
		fmt.Printf("Reading database... ")
		db, err := database.Load(filepath.Join(path, ".modget"))
		if err != nil {
			fmt.Printf("failed to open database: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Done")
		ids, err := slug.Slug(args, db)
		if err != nil {
			fmt.Printf("failed read input: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Finding Mods... ")
		if len(ids) == 0 {
			mods = db.Mods
		} else {
			for _, id := range ids {
				mod, err := filter.FindLocalMod(id, db)
				if err != nil {
					fmt.Printf("failed to find mod: %v\n%v\n", id, err)
					os.Exit(1)
				}
				mods = append(mods, mod)
			}
		}
		fmt.Println("Done")
		fmt.Printf("Checking for updates... ")
		for _, mod := range mods {
			file, err := filter.FindFile(mod.ID, minecraft, loader)
			if err != nil {
				fmt.Printf("failed to find mod: %v\n%v\n", mod.ID, err)
				os.Exit(1)
			}
			mTime, err := time.Parse(time.RFC3339, mod.FileDate)
			if err != nil {
				fmt.Printf("failed to parse time: %v\n%v\n", mod.ID, err)
				os.Exit(1)
			}
			fTime, err := time.Parse(time.RFC3339, file.FileDate)
			if err != nil {
				fmt.Printf("failed to parse time: %v\n%v\n", mod.ID, err)
				os.Exit(1)
			}
			if fTime.After(mTime) {
				updates = append(updates, mod)
			}
		}
		fmt.Println("Done")
		show(updates, "updated")
		if !prompt() {
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
