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
	"git.sr.ht/~kota/modget/util"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:     "del mod...",
	Aliases: []string{"d"},
	Short:   "Remove installed mod(s) based on MODID or Slug.",
	Run: func(cmd *cobra.Command, args []string) {
		var mods []database.Mod
		if len(args) == 0 {
			fmt.Println("modget del requires at least one MODID or Slug")
			os.Exit(1)
		}
		fmt.Printf("Reading database... ")
		db, err := util.FindDatabase(path)
		if err != nil {
			fmt.Printf("Failed to open database: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Done")
		ids, err := util.ToID(args, db)
		if err != nil {
			fmt.Printf("Failed read input: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Finding Mods... ")
		for _, id := range ids {
			mod, err := util.FindLocalMod(id, db)
			if err != nil {
				fmt.Printf("Failed to find mod: %v\n%v\n", id, err)
				os.Exit(1)
			}
			mods = append(mods, mod)
		}
		fmt.Println("Done")
		util.ShowRemove(mods)
		if !util.Ask() {
			os.Exit(0)
		}
		db, err = util.RemoveMods(mods, path, db)
		if err != nil {
			fmt.Printf("Failed to remove mod: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Updating database... ")
		err = db.Write(filepath.Join(path, ".modget"))
		if err != nil {
			fmt.Printf("Failed to write database: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
