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
	"git.sr.ht/~kota/modget/filter"
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
		var mods []database.Mod
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
		fmt.Printf("Database: %s\nMinecraft: %s\nLoader: %s\n\n", db.Version, db.Minecraft, db.Loader)
		if !one {
			showNormal(mods)
		} else {
			showOne(mods)
		}
	},
}

// showNormal prints a list of mods and displays a reasonable amount of
// information for each one.
func showNormal(mods []database.Mod) {
	for _, mod := range mods {
		v := mod.GameVersion[0]
		for i := 1; i < len(mod.GameVersion); i++ {
			v += ", "
			v += mod.GameVersion[i]
		}
		fmt.Printf("%d/%s - %d/%s\n\tDownloads: %d\n\tDate: %s\n\tVersions: %s\n\t%s\n\n",
			mod.ID,
			mod.Slug,
			mod.FileID,
			mod.FileName,
			int(mod.DownloadCount),
			mod.FileDate,
			v,
			mod.Summary)
	}
}

// showOne prints a list of mods and displays each mod on a single line.
func showOne(mods []database.Mod) {
	for _, mod := range mods {
		v := mod.GameVersion[0]
		for i := 1; i < len(mod.GameVersion); i++ {
			v += ", "
			v += mod.GameVersion[i]
		}
		fmt.Printf("%d/%s - %s\n",
			mod.ID,
			mod.Slug,
			mod.FileName)
	}
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVarP(&one, "oneline", "l", false, "Show mods one per line")
}
