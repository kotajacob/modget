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

	"git.sr.ht/~kota/modget/database"
	"git.sr.ht/~kota/modget/slug"
	"github.com/spf13/cobra"
)

// markCmd represents the mark command
var markCmd = &cobra.Command{
	Use:     "mark [mod]...",
	Aliases: []string{"m"},
	Short:   "Change the update status of a mod.",
	Long: `Mark allows you to "hold" or "unhold" your mods.

When a mod has a hold status it will be ignored by the update command and will
remain at its exact version until the hold is removed or a different version is
manually installed. By default the hold status of your selected mod(s) is
toggled and printed.`,
	Run: mark,
}

func init() {
	rootCmd.AddCommand(markCmd)
}

func mark(cmd *cobra.Command, args []string) {
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
	if len(ids) == 0 { // select all mods
		for id := range db.Mods {
			ids = append(ids, id)
		}
	}
	for _, id := range ids {
		db.Mods[id].Hold = !db.Mods[id].Hold
		fmt.Printf("%s - %t\n", db.Mods[id].Slug, db.Mods[id].Hold)
	}
	fmt.Printf("Updating database... ")
	err = db.Write(filepath.Join(path, ".modget"))
	if err != nil {
		fmt.Printf("failed to write database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done")
}
