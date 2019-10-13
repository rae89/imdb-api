package themoviedb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind_FindID(t *testing.T) {
	setup()
	defer teardown()

	testExternalID := "tt0234215"

	testQueryParams := map[string]string{
		"api_key":         "09876545",
		"external_source": "imdb_id",
		"language":        "en-US",
	}
	endpoint := fmt.Sprintf("/%s/%s/%s/", apiVersion, FindEndpoint, testExternalID)
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"movie_results": [
				{
					"adult": false,
					"backdrop_path": "/Fp3piEuHXxKnPBO5R0Wj4wjZHg.jpg",
					"genre_ids": [
						12,
						28,
						53,
						878
					],
					"id": 604,
					"original_language": "en",
					"original_title": "The Matrix Reloaded",
					"overview": "Six months after the events depicted in The Matrix, Neo has proved to be a good omen for the free humans, as more and more humans are being freed from the matrix and brought to Zion, the one and only stronghold of the Resistance.  Neo himself has discovered his superpowers including super speed, ability to see the codes of the things inside the matrix and a certain degree of pre-cognition. But a nasty piece of news hits the human resistance: 250,000 machine sentinels are digging to Zion and would reach them in 72 hours. As Zion prepares for the ultimate war, Neo, Morpheus and Trinity are advised by the Oracle to find the Keymaker who would help them reach the Source.  Meanwhile Neo's recurrent dreams depicting Trinity's death have got him worried and as if it was not enough, Agent Smith has somehow escaped deletion, has become more powerful than before and has fixed Neo as his next target.",
					"poster_path": "/ezIurBz2fdUc68d98Fp9dRf5ihv.jpg",
					"release_date": "2003-05-15",
					"title": "The Matrix Reloaded",
					"video": false,
					"vote_average": 6.9,
					"vote_count": 6052,
					"popularity": 28.87
				}
			],
			"person_results": [],
			"tv_results": [],
			"tv_episode_results": [],
			"tv_season_results": []
		}`)
	})
	actual, err := client.FindID(testExternalID, testQueryParams)
	expected := Find{
		MovieResults: []map[string]interface{}{
			{
				"adult":         false,
				"backdrop_path": "/Fp3piEuHXxKnPBO5R0Wj4wjZHg.jpg",
				"genre_ids": []interface{}{
					12.0,
					28.0,
					53.0,
					878.0,
				},
				"id":                604.0,
				"original_language": "en",
				"original_title":    "The Matrix Reloaded",
				"overview":          "Six months after the events depicted in The Matrix, Neo has proved to be a good omen for the free humans, as more and more humans are being freed from the matrix and brought to Zion, the one and only stronghold of the Resistance.  Neo himself has discovered his superpowers including super speed, ability to see the codes of the things inside the matrix and a certain degree of pre-cognition. But a nasty piece of news hits the human resistance: 250,000 machine sentinels are digging to Zion and would reach them in 72 hours. As Zion prepares for the ultimate war, Neo, Morpheus and Trinity are advised by the Oracle to find the Keymaker who would help them reach the Source.  Meanwhile Neo's recurrent dreams depicting Trinity's death have got him worried and as if it was not enough, Agent Smith has somehow escaped deletion, has become more powerful than before and has fixed Neo as his next target.",
				"poster_path":       "/ezIurBz2fdUc68d98Fp9dRf5ihv.jpg",
				"release_date":      "2003-05-15",
				"title":             "The Matrix Reloaded",
				"video":             false,
				"vote_average":      6.9,
				"vote_count":        6052.0,
				"popularity":        28.87,
			},
		},
		PersonResults:    []map[string]interface{}{},
		TVResults:        []map[string]interface{}{},
		TVEpisodeResults: []map[string]interface{}{},
		TVSeasonResults:  []map[string]interface{}{},
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
