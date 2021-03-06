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

	"git.sr.ht/~kota/modget/curse"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search string...",
	Aliases: []string{"s"},
	Short:   "Search for mods and print the results.",
	Run:     search,
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func search(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("modget search requires at least one search term")
		os.Exit(1)
	}
	for _, arg := range args {
		var search curse.Search
		search.GameID = 432           // Set game to minecraft
		search.SectionID = 6          // Set section to mods
		search.SearchFilter = args[0] // Search string
		addons, err := curse.AddonSearch(search)
		if err != nil {
			fmt.Printf("Search failed: %v\n", arg)
			os.Exit(1)
		}
		for _, addon := range addons {
			fmt.Println(addon)
		}
	}
}
