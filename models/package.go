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
	Privacy          string
	Repository       string
	ShortDescription *string `db:"short_description"`
	UpdatedAt        string  `db:"updated_at"`
	Version          string
	Volumes          *types.JSONText
}

// Port represents a port forward config
type Port struct {
	Local     int
	Container int
}

// Volume represents a volume forward config
type Volume struct {
	Local     string
	Container string
}

// Packages represents a list of packages
type Packages struct {
	Package []Package
}
