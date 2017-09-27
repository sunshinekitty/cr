package helpers

import "regexp"

var (
	match = regexp.MustCompile

	packageName = match(`([a-z\d]){1}([a-z0-9-*_*]){0,48}([a-z\d]){1}`)
)

// ValidPackageName validates a package's name
func ValidPackageName(n string) bool {
	return len(packageName.FindString(n)) == len(n)
}
