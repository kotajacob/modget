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
