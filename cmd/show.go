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
	"text/tabwriter"

	"git.sr.ht/~kota/modget/database"
	"git.sr.ht/~kota/modget/slug"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:     "show [mod]...",
	Aliases: []string{"sh"},
	Short:   "Query and print more information about a specific mod by MODID/Slug.",
	Run:     show,
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVarP(&one, "oneline", "l", false, "Show mods one per line")
}

func show(cmd *cobra.Command, args []string) {
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
	fmt.Printf("Database: %s\nMinecraft: %s\nLoader: %s\n\n",
		db.Version,
		db.Minecraft,
		db.Loader)
	if !one {
		showNormal(ids, db)
	} else {
		showOneLine(ids, db)
	}
}

// showNormal prints a list of mods and displays a reasonable amount of
// information for each one.
func showNormal(ids []int, db *database.Database) {
	for _, id := range ids {
		v := db.Mods[id].GameVersion[0]
		for i := 1; i < len(db.Mods[id].GameVersion); i++ {
			v += ", "
			v += db.Mods[id].GameVersion[i]
		}
		fmt.Printf("%s/%d\n\t%s\n\t%d/%s\n\tHold: %t\n\tDownloads: %d\n\tDate: %s\n\tVersions: %s\n\t\n",
			db.Mods[id].Slug,
			id,
			db.Mods[id].Summary,
			db.Mods[id].FileID,
			db.Mods[id].FileName,
			db.Mods[id].Hold,
			int(db.Mods[id].DownloadCount),
			db.Mods[id].FileDate, // TODO: Format this a little better
			v)
	}
}

// showOneLine prints a list of mods and displays each mod on a single line.
func showOneLine(ids []int, db *database.Database) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, '-', 0)
	for _, id := range ids {
		fmt.Fprintf(w, "%s/%d \t %s\n",
			db.Mods[id].Slug,
			id,
			db.Mods[id].FileName)
	}
	w.Flush()
}
