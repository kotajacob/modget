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
	"git.sr.ht/~kota/modget/input"
	"git.sr.ht/~kota/modget/printer"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete mod...",
	Aliases: []string{"d"},
	Short:   "Remove installed mod(s) based on MODID or Slug.",
	Run:     del,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func del(cmd *cobra.Command, args []string) {
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
	ids, err := input.Slug(args, db)
	if err != nil {
		fmt.Printf("failed read input: %v\n", err)
		os.Exit(1)
	}
	printer.Show(ids, "deleted", db.Mods)
	if !printer.Continue() {
		os.Exit(0)
	}
	for _, id := range ids {
		fmt.Printf("Delete: %v\n", db.Mods[id].FileName)
		err := os.Remove(filepath.Join(path, db.Mods[id].FileName))
		if err != nil {
			if err != nil {
				fmt.Printf("failed to remove mod: %v\n", err)
			}
		}
		db.Del(id)
	}
	fmt.Printf("Updating database... ")
	err = db.Write(filepath.Join(path, ".modget"))
	if err != nil {
		fmt.Printf("failed to write database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Done")
}
