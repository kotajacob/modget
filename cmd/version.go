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

	"github.com/spf13/cobra"
)

// Version represents the current version of modget. It is set in config.mk
var Version string

//SetVersion initializes the aerc version displayed in template functions
func SetVersion(v string) {
	Version = v
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version and exit.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("modget " + Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
