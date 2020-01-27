package balldontlie

import "fmt"

const (
	playersPath string = "/api/v1/players"
	playerPath  string = "/api/v1/players/%d"
)

// PlayerService - service used to retrieve Players
type PlayerService struct {
	client *Client
}

// Player - Structure
type Player struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Position     string `json:"position"`
	HeightFeet   int    `json:"height_feet"`
	HeightInches int    `json:"height_inches"`
	Weight       int    `json:"weight_pounds"`
	Team         Team   `json:"team"`
}

// PlayerOpts - query param options for Find
type PlayerOpts struct {
	Search string `url:"search,omitempty"`
	PageOpts
}

// Players - Client return struct
type Players struct {
	Data     []Player           `json:"data"`
	PageInfo *PaginatedResponse `json:"meta"`
}

// Find - Retrieve any Players with given parameters
func (s PlayerService) Find(playerOpts PlayerOpts) (*Players, *Response, error) {
	pathStr, err := addOptions(playersPath, playerOpts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", pathStr)
	if err != nil {
		return nil, nil, err
	}

	var c struct {
		*Players
	}

	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c.Players, resp, nil
}

// Get - Retrieves single Player with given ID
func (s PlayerService) Get(id int) (*Player, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf(playerPath, id))
	if err != nil {
		return nil, nil, err
	}

	var c struct {
		*Player
	}

	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c.Player, resp, nil
}
