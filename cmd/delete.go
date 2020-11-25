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

	"git.sr.ht/~kota/modget/ask"
	"git.sr.ht/~kota/modget/database"
	"git.sr.ht/~kota/modget/slug"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete mod...",
	Aliases: []string{"d"},
	Short:   "Remove installed mod(s) based on MODID or Slug.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("modget delete requires at least one MODID or Slug")
			os.Exit(1)
		}
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
		ask.Show(IDs, "deleted", db.Mods)
		if !ask.Prompt() {
			os.Exit(0)
		}
		err = remove(IDs, path, db)
		if err != nil {
			fmt.Printf("failed to remove mod: %v\n", err)
		}
		fmt.Printf("Updating database... ")
		err = db.Write(filepath.Join(path, ".modget"))
		if err != nil {
			fmt.Printf("failed to write database: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
