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
	"fmt"
	"os"
	"path/filepath"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/database"
)

// GetMods downloads a list of files and updates a Database
func GetMods(addons []curse.Addon, files []curse.File, path string, db database.Database) (database.Database, error) {
	for i, file := range files {
		p := filepath.Join(filepath.Dir(path), file.FileName)
		fmt.Printf("Get:%d %v\n", i, file.DownloadURL)
		err := curse.Download(file.DownloadURL, p)
		db = db.AddMod(addons[i], file)
		if err != nil {
			return db, err
		}
	}
	return db, nil
}

// RemoveMods removes a list of local mods and updates a Database
func RemoveMods(mods []database.Mod, path string, db database.Database) (database.Database, error) {
	for _, mod := range mods {
		fmt.Printf("Remove: %v\n", mod.FileName)
		err := os.Remove(filepath.Join(path, mod.FileName))
		if err != nil {
			return db, err
		}
		db = db.DelMod(mod.ID)
	}
	return db, nil
}
