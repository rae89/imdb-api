package themoviedb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// https://developers.themoviedb.org/3/find/find-by-id
// FindID search objects by exteral id

func (c *Client) FindID(externalID string, externalSource string, queryParam string) (Find, error) {
	var f Find
	data, err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s", FindEndpoint, queryParam), nil)
	if err != nill {
		return f, err
	}
	err = json.Unmarshal(data, &f)
	if err != nil {
		return s, errors.Wrap(err, "failed to unmarshal find response")
	}

	return f, nil
}
