package models

// Package represents a package in the package table
type Package struct {
	Name             string
	Owner            string
	Version          string
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
	Config           string
	Pulls            int
	Privacy          string
	ShortDescription *string `db:"short_description"`
	LongDescription  *string `db:"long_description"`
	DockerRepo       string  `db:"docker_repo"`
	Homepage         *string
}

// Packages represents a list of packages
type Packages struct {
	Package []Package
}
