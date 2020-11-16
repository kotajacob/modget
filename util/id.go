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
	"strconv"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/database"
)

// searchModid takes a string, which is meant to be an addon's slug and attempts
// to convert it to a MODID by using a search.
func searchModid(s string) (int, error) {
	var search curse.Search
	search.GameID = 432     // Set game to minecraft
	search.SectionID = 6    // Set section to mods
	search.SearchFilter = s // Search string
	addons, err := curse.AddonSearch(search)
	if err != nil {
		return 0, err
	}
	for _, addon := range addons {
		if addon.Slug == s {
			return addon.ID, nil
		}
	}
	err = errors.New("Could not find: " + s)
	return 0, err
}

// readModid takes a string, which is meant to be an addon's slug and attempts to convert it to a MODID by using the local database.
func readModid(s string, db database.Database) (int, error) {
	for _, mod := range db.Mods {
		if mod.Slug == s {
			return mod.ID, nil
		}
	}
	err := errors.New("Could not find: " + s)
	return 0, err
}

// ToID converts a list of strings to MODIDs
func ToID(s []string, db database.Database) ([]int, error) {
	var mods []int
	for i := 0; i < len(s); i++ {
		id, err := strconv.Atoi(s[i])
		if err != nil {
			// Attempt to convert slug to modid
			id, err = readModid(s[i], db)
			if err != nil {
				id, err = searchModid(s[i])
				if err != nil {
					return nil, err
				}
			}
		}
		mods = append(mods, id)
	}
	return mods, nil
}
