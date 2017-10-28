package helpers

import (
	"fmt"
	"os"
	"os/user"
)

// EnsureConfigDirs ensures necessary Crackle config dirs are setup
func EnsureConfigDirs() error {
	home := ConfigDir()
	if _, err := os.Stat(home); os.IsNotExist(err) {
		err = os.Mkdir(home, 0755)
		if err != nil {
			return err
		}
	}
	homePackages := fmt.Sprintf("%s/packages", home)
	if _, err := os.Stat(homePackages); os.IsNotExist(err) {
		err = os.Mkdir(homePackages, 0755)
		if err != nil {
			return err
		}
	}
	homeBin := fmt.Sprintf("%s/bin", home)
	if _, err := os.Stat(homeBin); os.IsNotExist(err) {
		err = os.Mkdir(homeBin, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// ConfigDir returns the location Crackle Home should be set to
func ConfigDir() string {
	xdg := os.Getenv("XDG_CONFIG_HOME")
	usr, _ := user.Current()
	home := fmt.Sprintf("%s/.cr", usr.HomeDir)
	if xdg != "" {
		home = fmt.Sprintf("%s/.cr", xdg)
	}
	return home
}

// CreatePackageFiles returns a file handler for config file and creates
// a file in bin based on package name that is executable
func CreatePackageFiles(packageName string) (*os.File, error) {
	bPath := fmt.Sprintf("%s/bin/%s", ConfigDir(), packageName)
	b, err := os.OpenFile(bPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil, err
	}
	_, err = b.Write([]byte(fmt.Sprintf("cr exec %s", packageName)))
	if err != nil {
		return nil, err
	}
	b.Close()
	err = os.Chmod(bPath, 0755)
	if err != nil {
		return nil, err
	}
	pPath := fmt.Sprintf("%s/packages/%s.toml", ConfigDir(), packageName)
	return os.Create(pPath)
}
