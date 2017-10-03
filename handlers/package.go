package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"

	"github.com/sunshinekitty/cr/models"
)

// CreatePackage creates a new Package
func CreatePackage(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}

// ReadPackage returns a Package by name
func ReadPackage(c echo.Context) error {
	// Params
	name := c.Param("name")

	// Query/unpack
	foundPackage := models.Package{}
	err := DB.Get(&foundPackage, "SELECT * FROM packages WHERE name=$1", name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, foundPackage)
}

// UpdatePackage updates a Package by name
func UpdatePackage(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}

// DeletePackage deletes a Package by name
func DeletePackage(c echo.Context) error {
	// Params
	name := c.Param("name")

	// Query
	delete := `DELETE FROM packages WHERE name=$1`
	_, err := DB.Exec(delete, name)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
