package models

import "github.com/jmoiron/sqlx/types"

// Package represents a package in the package table
type Package struct {
	CommandStart     *string `db:"command_start"`
	CreatedAt        string  `db:"created_at"`
	Homepage         *string
	LongDescription  *string `db:"long_description"`
	Name             string
	Owner            string
	Pulls            int
	Ports            *types.JSONText
	Repository       string
	ShortDescription *string `db:"short_description"`
	UpdatedAt        string  `db:"updated_at"`
	Version          string
	Volumes          *types.JSONText
}

// Packages represents a list of Package structs
type Packages struct {
	Package []Package
}

// PackageToml represents a raw toml config object
type PackageToml struct {
	CommandStart     *string `toml:"command_start"`
	Homepage         *string
	LongDescription  *string
	Package          string
	Ports            []Port `toml:"port"`
	Repository       string
	ShortDescription *string
	Volumes          []Volume `toml:"volume"`
}

// Port represents a port forward config
type Port struct {
	Local     string
	Container string
}

// Volume represents a volume forward config
type Volume struct {
	Local     string
	Container string
}
