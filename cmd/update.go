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

	"git.sr.ht/~kota/modget/ask"
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
		var updates []int
		fmt.Printf("Reading database... ")
		db, err := database.Load(filepath.Join(path, ".modget"))
		if err != nil {
			fmt.Printf("failed to open database: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Done")
		IDs, err := slug.Slug(args, db)
		if err != nil {
			fmt.Printf("failed read input: %v\n", err)
			os.Exit(1)
		}
		if len(IDs) == 0 { // select all mods
			for ID := range db.Mods {
				IDs = append(IDs, ID)
			}
		}
		fmt.Printf("Checking for updates... ")
		for _, ID := range IDs {
			file, err := filter.FindFile(ID, minecraft, loader)
			if err != nil {
				fmt.Printf("failed to find mod: %v\n%v\n", ID, err)
				os.Exit(1)
			}
			mTime, err := time.Parse(time.RFC3339, db.Mods[ID].FileDate)
			if err != nil {
				fmt.Printf("failed to parse time: %v\n%v\n", ID, err)
				os.Exit(1)
			}
			fTime, err := time.Parse(time.RFC3339, file.FileDate)
			if err != nil {
				fmt.Printf("failed to parse time: %v\n%v\n", ID, err)
				os.Exit(1)
			}
			if fTime.After(mTime) {
				updates = append(updates, ID)
			}
		}
		fmt.Println("Done")
		ask.Show(updates, "updated", db.Mods)
		if !ask.Prompt() {
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
