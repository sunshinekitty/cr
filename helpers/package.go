package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"

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
	// ErrLongShortDescription is thrown when short description is too long (>200)
	ErrLongShortDescription = errors.New("short description is too long (>200 chars)")
	// ErrLongLongDescription is thrown when long description is too long (>25000)
	ErrLongLongDescription = errors.New("long description is too long (>25000 chars)")
	// ErrLongHomepage is thrown when home page is too long (>100)
	ErrLongHomepage = errors.New("homepage is too long (>100 chars)")
	// ErrLongCommandStart is thrown when command start is too long (>100)
	ErrLongCommandStart = errors.New("command start is too long (>100 chars)")
	// ErrMissingUsername is thrown when a username isn't set in client config
	ErrMissingUsername = errors.New("username is not set in client config")
)

// ConfigFileToPackageToml takes a path to toml config and translates to PackageToml struct
func ConfigFileToPackageToml(path string) (*models.PackageToml, error) {
	var returnPackageToml models.PackageToml
	_, err := toml.DecodeFile(path, &returnPackageToml)
	return &returnPackageToml, err
}

// PackageTomlToPackage takes a PackageToml struct and converts it to a Package struct
func PackageTomlToPackage(pt *models.PackageToml) (*models.Package, error) {
	var p models.Package
	p.CommandStart = pt.CommandStart
	p.Homepage = pt.Homepage
	p.LongDescription = pt.LongDescription
	p.Name = pt.Package
	p.Pulls = 0
	p.ShortDescription = pt.ShortDescription

	username := viper.GetString("crackle.auth.username")
	if len(username) == 0 {
		return nil, ErrMissingUsername
	}
	p.Owner = username

	ptPorts, err := json.Marshal(pt.Ports)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(ptPorts, &p.Ports)
	if err != nil {
		return nil, err
	}

	splitRepository := strings.Split(pt.Repository, ":")
	p.Version = splitRepository[1]
	p.Repository = splitRepository[0]

	ptVolumes, err := json.Marshal(pt.Volumes)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(ptVolumes, &p.Volumes)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// ValidPackageToml validates a PackageToml object
func ValidPackageToml(pt *models.PackageToml) error {
	if !ValidPackageName(pt.Package) {
		return ErrInvalidPackageName
	}
	if !ValidRepositoryName(pt.Repository) {
		return ErrInvalidRepositoryName
	}
	for _, port := range pt.Ports {
		if !ValidPort(port.Container) {
			ErrInvalidPort = fmt.Errorf("Container port \"%v\" is invalid", port.Container)
			return ErrInvalidPort
		}
		if !ValidPort(port.Local) {
			ErrInvalidPort = fmt.Errorf("Local port \"%v\" is invalid", port.Local)
			return ErrInvalidPort
		}
	}
	if pt.ShortDescription != nil {
		if len(fmt.Sprintf("%s", *pt.ShortDescription)) > 200 {
			return ErrLongShortDescription
		}
	}
	if pt.LongDescription != nil {
		if len(fmt.Sprintf("%s", *pt.LongDescription)) > 25000 {
			return ErrLongLongDescription
		}
	}
	if pt.Homepage != nil {
		if len(fmt.Sprintf("%s", *pt.Homepage)) > 100 {
			return ErrLongHomepage
		}
	}
	if pt.CommandStart != nil {
		if len(fmt.Sprintf("%s", *pt.CommandStart)) > 100 {
			return ErrLongCommandStart
		}
	}
	return nil
}

// ValidPackage validates a Package object
func ValidPackage(p *models.Package) error {
	if !ValidPackageName(p.Name) {
		return ErrInvalidPackageName
	}
	if !ValidRepositoryName(fmt.Sprintf("%s:%s", p.Repository, p.Version)) {
		return ErrInvalidRepositoryName
	}

	portsBytes, err := json.Marshal(p.Ports)
	ports := new(models.Ports)
	if err != nil {
		return err
	}
	err = json.Unmarshal(portsBytes, &ports)
	for _, port := range *ports {
		if !ValidPort(port.Container) {
			ErrInvalidPort = fmt.Errorf("Container port \"%v\" is invalid", port.Container)
			return ErrInvalidPort
		}
		if !ValidPort(port.Local) {
			ErrInvalidPort = fmt.Errorf("Local port \"%v\" is invalid", port.Local)
			return ErrInvalidPort
		}
	}

	if p.ShortDescription != nil {
		if len(fmt.Sprintf("%s", *p.ShortDescription)) > 200 {
			return ErrLongShortDescription
		}
	}
	if p.LongDescription != nil {
		if len(fmt.Sprintf("%s", *p.LongDescription)) > 25000 {
			return ErrLongLongDescription
		}
	}
	if p.Homepage != nil {
		if len(fmt.Sprintf("%s", *p.Homepage)) > 100 {
			return ErrLongHomepage
		}
	}
	if p.CommandStart != nil {
		if len(fmt.Sprintf("%s", *p.CommandStart)) > 100 {
			return ErrLongCommandStart
		}
	}

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
