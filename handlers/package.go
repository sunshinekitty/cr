package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"

	"github.com/sunshinekitty/cr/helpers"
	"github.com/sunshinekitty/cr/models"
)

// CreatePackage creates a new Package
func CreatePackage(c echo.Context) error {
	p := new(models.Package)
	if err := c.Bind(p); err != nil {
		return err
	}

	if err := helpers.ValidPackage(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// TODO: check for conflict here
	foundPackage, err := selectPackage(p.Name, p.Version)
	if foundPackage.Name != "" {
		return echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("Package %s:%s already exists", foundPackage.Name, foundPackage.Version))
	}

	query := `INSERT INTO packages(command_start, homepage, long_description, 
								   name, owner, pulls, ports, repository, 
								   short_description, version, volumes) 
			  VALUES(:command_start, :homepage, :long_description, :name, 
					 :owner, :pulls, :ports, :repository, :short_description, 
					 :version, :volumes)`

	_, err = DB.NamedExec(query, p)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, p)
}

// ReadPackage returns a Package by name
func ReadPackage(c echo.Context) error {
	// Params
	name := c.Param("name")
	version := c.QueryParam("version")

	if !helpers.ValidPackageName(name) {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// Query/unpack
	foundPackage, err := selectPackage(name, version)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, foundPackage)
}

// UpdatePackage updates a Package by name
func UpdatePackage(c echo.Context) error {
	// Params
	name := c.Param("name")
	//version := c.QueryParam("version")

	if !helpers.ValidPackageName(name) {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, "ok")
}

// DeletePackage deletes a Package by name
func DeletePackage(c echo.Context) error {
	// Params
	name := c.Param("name")
	//version := c.QueryParam("version")

	if !helpers.ValidPackageName(name) {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// Query
	delete := `DELETE FROM packages WHERE name=$1`
	res, err := DB.Exec(delete, name)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	deletedRows, _ := res.RowsAffected()
	if deletedRows == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}

func selectPackage(packageName string, version string) (models.Package, error) {
	var err error
	foundPackage := models.Package{}
	if version == "" {
		err = DB.Get(&foundPackage, "SELECT * FROM packages WHERE name=$1 ORDER BY created_at DESC LIMIT 1", packageName)
	} else {
		log.Error(version)
		err = DB.Get(&foundPackage, "SELECT * FROM packages WHERE name=$1 AND version=$2 ORDER BY created_at DESC LIMIT 1", packageName, version)
	}
	return foundPackage, err
}
