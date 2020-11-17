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

package util

import (
	"errors"
	"fmt"
	"os"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/database"
)

// FindFile returns a curse.File for a MODID. It ensures the file matches the
// correct Minecraft version and Loader. Additionally it warns the user if the
// enter an unknown version or loader.
func FindFile(id int, minecraftVersion string, loader string) (curse.File, error) {
	files, err := curse.AddonFiles(id)
	// Validate the modloader and mc version
	mcVersions, err := curse.MinecraftVersionList()
	if minecraftVersion != "" {
		files = VersionFilter(files, minecraftVersion)
		if !ValidateMinecraftVersion(minecraftVersion, mcVersions) {
			fmt.Println("warning: Minecraft Version entered is not recognized")
		}
	}
	if loader != "" {
		files = LoaderFilter(files, loader)
		if !ValidateModLoader(loader) {
			fmt.Println("warning: Modloader entered is not recognized")
		}
	}
	files = TimeSort(files)
	if len(files) == 0 {
		err = errors.New("file not found for those search terms")
	}
	return files[0], err
}

// FindLocalMod returns a database.Mod for a MODID.
func FindLocalMod(id int, db database.Database) (database.Mod, error) {
	var mod database.Mod
	for _, mod := range db.Mods {
		if mod.ID == id {
			return mod, nil
		}
	}
	return mod, errors.New("file not found in database")
}

// ensureDir creates a directory if missing.
func ensureDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModeDir)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}
