package themoviedb

type Find struct {
	MovieResults     []map[string]interface{} `json:"movie_results,omitempty"`
	PersonResults    []map[string]interface{} `json:"person_results,omitempty"`
	TVResults        []map[string]interface{} `json:"tv_results,omitempty"`
	TVEpisodeResults []map[string]interface{} `json:"tv_episode_results,omitempty"`
	TVSeasonResults  []map[string]interface{} `json:"tv_season_results,omitempty"`
}
