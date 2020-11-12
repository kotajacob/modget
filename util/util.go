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

type StringSet map[string]bool

// Check if two StringSets conflict
func (a StringSet) Conflicts(b StringSet) bool {
	for k, _ := range a {
		if b[k] {
			return true
		}
	}
	return false
}

// Copy a StringSet
func (a StringSet) Copy() StringSet {
	copy := make(StringSet)
	for k, v := range a {
		copy[k] = v
	}
	return copy
}

// Remove a string from a StringSet
func (a StringSet) RemoveString(s string) StringSet {
	if a[s] {
		delete(a, s)
	}
	return a
}

// List of valid modloaders because curseforge doesn't provide one...
// dont set to false instead use delete()
var ModLoaders = StringSet{
	"forge":      true,
	"fabric":     true,
	"liteloader": true,
	"rift":       true,
}
