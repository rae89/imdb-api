package themoviedb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// https://developers.themoviedb.org/3/find/find-by-id
// FindID search objects by exteral id

func (c *Client) FindID(externalID string, queryparams map[string]string) (Find, error) {
	var f Find
	data, err := c.doRequest(http.MethodGet, fmt.Sprintf("/%s/%s", FindEndpoint, externalID), nil, queryparams)
	if err != nil {
		return f, err
	}
	err = json.Unmarshal(data, &f)
	if err != nil {
		return f, errors.Wrap(err, "failed to unmarshal find response")
	}

	return f, nil
}
