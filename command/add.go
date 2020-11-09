package command

import (
	"errors"
	"fmt"

	"git.sr.ht/~kota/modget/curse"
)

// Add searches and downloads one or more mods and records the result in the
// database. Additionally it can accept a manually specified mc version and
// loader, or fallback to the default one in the database.
func Add(mods []int, mc string, loader string) error {
	if len(mods) > 0 {
		for i := 0; i < len(mods); i++ {
			files, err := curse.AddonFiles(mods[i])
			if err != nil {
				return err
			}
			files, err = matchFiles(files, mc, loader)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("modget add requires at least one MODID")
	}
	return nil
}

// Filters a list of Files returning only the ones that match the loader and mc version.
func matchFiles(files []curse.File, mc, loader string) ([]curse.File, error) {
	for _, file := range files {
		fmt.Println(file.FileName)
		fmt.Println(file.FileDate)
		fmt.Println(file.Id)
	}
	return files, nil
}
