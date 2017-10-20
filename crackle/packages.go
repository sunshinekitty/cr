package crackle

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sunshinekitty/cr/models"
)

// PackageService handles communication with Crackle API relating to Package endpoint
type PackageService service

// GetPackage fetchs the Package object for a given Package name.
func (s *PackageService) GetPackage(ctx context.Context, p string) (*models.Package, *http.Response, error) {
	u := fmt.Sprintf("package/%s", p)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", accept)

	c := new(models.Package)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}
