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

package input

import (
	"errors"
	"fmt"
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
		return 0, fmt.Errorf("failed to search for slug: %s: %v", s, err)
	}
	for _, addon := range addons {
		if addon.Slug == s {
			return addon.ID, nil
		}
	}
	return 0, fmt.Errorf("slug not found in search results: %s: %v", s, err)
}

// readModid takes a string, which is meant to be an addon's slug and attempts
// to convert it to a MODID by using the local database.
func readModid(s string, db *database.Database) (int, error) {
	for id, mod := range db.Mods {
		if mod.Slug == s {
			return id, nil
		}
	}
	err := errors.New("Could not find: " + s)
	return 0, err
}

// Slug converts a list of slugs to MODIDs
func Slug(s []string, db *database.Database) ([]int, error) {
	var ids []int
	for _, v := range s {
		id, err := strconv.Atoi(v)
		if err != nil {
			// Attempt to convert slug to modid
			id, err = readModid(v, db)
			if err != nil {
				id, err = searchModid(v)
				if err != nil {
					return nil, err
				}
			}
		}
		ids = append(ids, id)
	}
	return ids, nil
}
