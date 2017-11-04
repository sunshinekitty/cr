package crackle

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sunshinekitty/cr/models"
)

// VersionService retrieves version information
type VersionService service

// Server fetchs the Package object for a given Package name
func (s *VersionService) Server(ctx context.Context) (*models.Version, *http.Response, error) {
	u := fmt.Sprintf("version")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", accept)

	version := new(models.Version)

	resp, err := s.client.Do(ctx, req, version)
	if err != nil {
		return nil, resp, err
	}

	return version, resp, nil
}
