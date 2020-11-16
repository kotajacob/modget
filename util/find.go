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
	"path/filepath"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/database"
)

// FindDatabase finds the .modget database at the path. Create the database if missing.
func FindDatabase(path string) (database.Database, error) {
	var db database.Database
	if path == "" {
		path = "."
	}
	err := EnsureDir(path)
	if err != nil {
		return db, err
	}
	path = filepath.Join(path, ".modget")
	db, err = database.Load(path)
	return db, err
}

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