package helpers

import (
	"regexp"

	"github.com/sunshinekitty/cr/models"
)

var (
	match = regexp.MustCompile

	packageName = match(`([a-z\d]){1}([a-z0-9-*_*]){0,48}([a-z\d]){1}`)
	repoName    = match(`([A-Za-z\d\./:-]*){3,141}`)
)

// ConfigToPackage takes a blob of toml config and translates to Package struct
func ConfigToPackage(s string) (*models.Package, error) {
	return nil, nil
}

// PackageToConfig takes a Package struct and spits out a blob of toml config
func PackageToConfig() (*string, error) {
	return nil, nil
}

// ValidPackage validates a package object
func ValidPackage(p *models.Package) error {
	return nil
}

// ValidPackageName validates a package's name
func ValidPackageName(n string) bool {
	if len(n) == 0 {
		return false
	}
	return len(packageName.FindString(n)) == len(n)
}

// ValidRepositoryName validates a repository name
func ValidRepositoryName(n string) bool {
	// We could pull in Docker and use their regexp matching, but I don't think it really matters
	// We should just verify it meets database constraints and is alphanumeric and/or ":" and/or "/"'s
	if len(n) > 141 || len(n) < 3 {
		return false
	}
	return len(repoName.FindString(n)) == len(n)
}

// ValidPort validate's a port number
func ValidPort(i int) bool {
	return i >= 1 && i <= 65535
}
