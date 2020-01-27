package balldontlie

const (
	statsPath     string = "/api/v1/stats"
	seasonAvgPath string = "/api/v1/season_averages"
)

// Stat - Structure
type Stat struct {
	ID         int     `json:"id"`
	Assists    int     `json:"ast"`
	Blocks     int     `json:"blk"`
	DefRebs    int     `json:"dreb"`
	FGPctThree float64 `json:"fg3_pct"`
	FGAThree   int     `json:"fg3a"`
	FGMThree   int     `json:"fg3m"`
	FGPct      float64 `json:"fg_pct"`
	FGA        int     `json:"fga"`
	FGM        int     `json:"fgm"`
	FTPct      float64 `json:"ft_pct"`
	FTA        int     `json:"fta"`
	FTM        int     `json:"ftm"`
	Game       struct {
		ID               int    `json:"id"`
		Date             string `json:"date"`
		HomeTeamScore    int    `json:"home_team_score"`
		VisitorTeamScore int    `json:"visitor_team_score"`
		Season           int    `json:"season"`
		HomeTeamID       int    `json:"home_team_id"`
		VisitorTeamID    int    `json:"visitor_team_id"`
	} `json:"game"`
	Minutes string `json:"min"`
	OffRebs int    `json:"oreb"`
	Fouls   int    `json:"pf"`
	Player  struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Position  string `json:"position"`
		TeamID    int    `json:"team_id"`
	} `json:"player"`
	Points    int  `json:"pts"`
	Rebounds  int  `json:"reb"`
	Steals    int  `json:"stl"`
	Team      Team `json:"team"`
	Turnovers int  `json:"turnover"`
}

// SeasonAvg - Structure
type SeasonAvg struct {
	GamesPlayed int     `json:"games_played"`
	PlayerID    int     `json:"player_id"`
	Season      int     `json:"season"`
	Minutes     string  `json:"min"`
	Assists     float64 `json:"ast"`
	Blocks      float64 `json:"blk"`
	DefRebs     float64 `json:"dreb"`
	OffRebs     float64 `json:"oreb"`
	Rebounds    float64 `json:"reb"`
	FGPctThree  float64 `json:"fg3_pct"`
	FGAThree    float64 `json:"fg3a"`
	FGMThree    float64 `json:"fg3m"`
	FGPct       float64 `json:"fg_pct"`
	FGA         float64 `json:"fga"`
	FGM         float64 `json:"fgm"`
	FTPct       float64 `json:"ft_pct"`
	FTA         float64 `json:"fta"`
	FTM         float64 `json:"ftm"`
	Turnovers   float64 `json:"turnover"`
	Steals      float64 `json:"stl"`
	Fouls       float64 `json:"pf"`
	Points      float64 `json:"pts"`
}

// StatService - service used to retrieve Stats
type StatService struct {
	client *Client
}

// StatsOpts - query param options for Find
type StatsOpts struct {
	Seasons    []string `url:"seasons,brackets,omitempty"`
	Dates      []string `url:"dates,brackets,omitempty"`
	TeamIds    []string `url:"team_ids,brackets,omitempty"`
	PlayerIds  []string `url:"player_ids,brackets,omitempty"`
	GameIds    []string `url:"game_ids,brackets,omitempty"`
	Postseason bool     `url:"postseason,omitempty"`
	StartDate  string   `url:"start_date,omitempty"`
	EndDate    string   `url:"end_date,omitempty"`
	PageOpts
}

// SeasonAvgOpts - query param options for FindSeasonAvg
type SeasonAvgOpts struct {
	Season    string `url:"season,omitempty"`
	PlayerIds []int  `url:"player_ids,brackets,omitempty"`
}

// Stats - Client return struct
type Stats struct {
	Data     []Stat             `json:"data"`
	PageInfo *PaginatedResponse `json:"meta"`
}

// Find - Retrieve all Individual Player Stats for Games with given parameters
func (s StatService) Find(statsOpts StatsOpts) (*Stats, *Response, error) {
	pathStr, err := addOptions(statsPath, statsOpts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", pathStr)
	if err != nil {
		return nil, nil, err
	}

	var c struct {
		*Stats
	}

	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c.Stats, resp, nil
}

// FindSeasonAvg - Retrieve Player Season Averages with given parameters
func (s StatService) FindSeasonAvg(seasonAvgOpts SeasonAvgOpts) ([]SeasonAvg, *Response, error) {
	pathStr, err := addOptions(seasonAvgPath, seasonAvgOpts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", pathStr)
	if err != nil {
		return nil, nil, err
	}

	var c struct {
		SeasonAvgs []SeasonAvg `json:"data"`
	}

	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c.SeasonAvgs, resp, nil
}
