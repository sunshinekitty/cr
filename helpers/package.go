package helpers

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/BurntSushi/toml"

	"github.com/sunshinekitty/cr/models"
)

var (
	match = regexp.MustCompile

	packageName = match(`([a-z\d]){1}([a-z0-9-*_*]){0,48}([a-z\d]){1}`)
	repoName    = match(`([A-Za-z\d\./:-]*){3,141}`)

	// ErrInvalidPackageName is thrown when an invalid package name is given
	ErrInvalidPackageName = errors.New("package name is invalid")
	// ErrInvalidRepositoryName is thrown when an invalid repository name is given
	ErrInvalidRepositoryName = errors.New("repository name is invalid")
	// ErrInvalidPort is thrown when an invalid port is given
	ErrInvalidPort = errors.New("port number is invalid")
)

// ConfigToPackageToml takes a blob of toml config and translates to PackageToml struct
func ConfigToPackageToml(path string) (*models.PackageToml, error) {
	var returnPackageToml models.PackageToml
	_, err := toml.DecodeFile(path, &returnPackageToml)
	return &returnPackageToml, err
}

// PackageTomlToPackage takes a PackageToml struct and converts it to a Package struct
func PackageTomlToPackage(srcPackageToml *models.PackageToml) (*models.Package, error) {
	var returnPackage models.Package
	if !ValidPackageName(srcPackageToml.Package) {
		return nil, ErrInvalidPackageName
	}
	if !ValidRepositoryName(srcPackageToml.Repository) {
		return nil, ErrInvalidRepositoryName
	}
	for _, port := range srcPackageToml.Ports {
		if !ValidPort(port.Container) {
			ErrInvalidPort = fmt.Errorf("Container port \"%v\" is invalid", port.Container)
			return nil, ErrInvalidPort
		}
		if !ValidPort(port.Local) {
			ErrInvalidPort = fmt.Errorf("Local port \"%v\" is invalid", port.Local)
			return nil, ErrInvalidPort
		}
	}
	// TODO: finish validating rest of PackageToml then actually convert, probably pull out validation to separate func
	return &returnPackage, nil
}

// PackageToTomlConfig takes a Package struct and spits out a blob of toml config
func PackageToTomlConfig() (*string, error) {
	return nil, nil
}

// ValidPackage validates a Package object
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
func ValidPort(s string) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return i >= 1 && i <= 65535
}
