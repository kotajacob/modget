package printer

import (
	"fmt"
	"strings"

	"git.sr.ht/~kota/modget/database"
	"github.com/dustin/go-humanize"
)

// Ask the user to provide a string
func Ask(q string) string {
	fmt.Printf("Enter a %s for the new database: ", q)
	var a string
	fmt.Scanln(&a)
	return strings.ToLower(a)
}

// Prompt the user with a Yes/No question about continuing
func Prompt() bool {
	fmt.Printf("Do you want to continue? [Y/n] ")
	var a string
	fmt.Scanln(&a)
	a = strings.ToLower(a)
	if a == "y" || a == "" {
		return true
	}
	return false
}

// Show prints a list of mods that will be added, deleted, or updated.
// mode should be "added", "deleted", or "updated"
func Show(ids []int, mode string, mods map[int]*database.Mod) {
	fmt.Printf("The following mods will be %s:\n", mode)
	var s string
	var d int
	for _, id := range ids {
		s += " " + mods[id].Slug
		d += mods[id].FileLength
	}
	fmt.Printf("%v\n", s)
	if mode == "deleted" {
		fmt.Printf("After this operation, %s of disk space will be freed.\n", humanize.Bytes(uint64(d)))
	} else {
		fmt.Printf("After this operation, %s of additional disk space will be used.\n", humanize.Bytes(uint64(d)))
	}
}
