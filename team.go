package balldontlie

import "fmt"

const (
	teamsPath string = "/api/v1/teams"
	teamPath  string = "/api/v1/teams/%d"
)

// TeamService - service used to retrieve teams
type TeamService struct {
	client *Client
}

// Team - Structure
type Team struct {
	ID         int    `json:"id"`
	Abbr       string `json:"abbreviation"`
	City       string `json:"city"`
	Conference string `json:"conference"`
	Division   string `json:"division"`
	FullName   string `json:"full_name"`
	Name       string `json:"name"`
}

// Teams - Client return struct
type Teams struct {
	Data     []Team             `json:"data"`
	PageInfo *PaginatedResponse `json:"meta"`
}

// All - Retrieves all teams with pagination options
func (s TeamService) All(pageOpts PageOpts) (*Teams, *Response, error) {
	pathStr, err := addOptions(teamsPath, pageOpts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", pathStr)
	if err != nil {
		return nil, nil, err
	}

	var c struct {
		*Teams
	}

	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c.Teams, resp, nil
}

// Get - Retrieves single Team with given ID
func (s TeamService) Get(id int) (*Team, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf(teamPath, id))
	if err != nil {
		return nil, nil, err
	}

	var c struct {
		*Team
	}

	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c.Team, resp, nil
}
