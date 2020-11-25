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

var (
	one bool
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:     "show [mod]...",
	Aliases: []string{"sh"},
	Short:   "Query and print more information about a specific mod by MODID/Slug.",
	Run: func(cmd *cobra.Command, args []string) {
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
		fmt.Printf("Database: %s\nMinecraft: %s\nLoader: %s\n\n", db.Version, db.Minecraft, db.Loader)
		if !one {
			showNormal(IDs, db)
		} else {
			showOne(IDs, db)
		}
	},
}

// showNormal prints a list of mods and displays a reasonable amount of
// information for each one.
func showNormal(IDs []int, db *database.Database) {
	for _, ID := range IDs {
		v := db.Mods[ID].GameVersion[0]
		for i := 1; i < len(db.Mods[ID].GameVersion); i++ {
			v += ", "
			v += db.Mods[ID].GameVersion[i]
		}
		fmt.Printf("%d/%s - %d/%s\n\tDownloads: %d\n\tDate: %s\n\tVersions: %s\n\t%s\n\n",
			ID,
			db.Mods[ID].Slug,
			db.Mods[ID].FileID,
			db.Mods[ID].FileName,
			int(db.Mods[ID].DownloadCount),
			db.Mods[ID].FileDate,
			v,
			db.Mods[ID].Summary)
	}
}

// showOne prints a list of mods and displays each mod on a single line.
func showOne(IDs []int, db *database.Database) {
	for _, ID := range IDs {
		fmt.Printf("%d/%s - %s\n",
			ID,
			db.Mods[ID].Slug,
			db.Mods[ID].FileName)
	}
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVarP(&one, "oneline", "l", false, "Show mods one per line")
}
