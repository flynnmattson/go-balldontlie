package balldontlie

import "fmt"

const (
	gamesPath string = "/api/v1/games"
	gamePath  string = "/api/v1/games/%d"
)

// Game - Structure
type Game struct {
	ID               int    `json:"id"`
	Date             string `json:"date"`
	Season           int    `json:"season"`
	HomeTeamScore    int    `json:"home_team_score"`
	VisitorTeamScore int    `json:"visitor_team_score"`
	Period           int    `json:"period"`
	Status           string `json:"status"`
	Time             string `json:"time"`
	Postseason       bool   `json:"postseason"`
	HomeTeam         Team   `json:"home_team"`
	VisitorTeam      Team   `json:"visitor_team"`
}

// GameService - service used to retrieve games
type GameService struct {
	client *Client
}

// GamesOpts - query param options for Find
type GamesOpts struct {
	Seasons    []string `url:"seasons,brackets,omitempty"`
	Dates      []string `url:"dates,brackets,omitempty"`
	TeamIds    []int    `url:"team_ids,brackets,omitempty"`
	Postseason bool     `url:"postseason,omitempty"`
	StartDate  string   `url:"start_date,omitempty"`
	EndDate    string   `url:"end_date,omitempty"`
	PageOpts
}

// Games - Client return struct
type Games struct {
	Data     []Game             `json:"data"`
	PageInfo *PaginatedResponse `json:"meta"`
}

// Find - Retrieve all games with given parameters
func (s GameService) Find(gamesOpts GamesOpts) (*Games, *Response, error) {
	pathStr, err := addOptions(gamesPath, gamesOpts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", pathStr)
	if err != nil {
		return nil, nil, err
	}

	var c struct {
		*Games
	}

	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c.Games, resp, nil
}

// Get - Retrieves single Game with given ID
func (s GameService) Get(id int) (*Game, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf(gamePath, id))
	if err != nil {
		return nil, nil, err
	}

	var c struct {
		*Game
	}

	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c.Game, resp, nil
}
