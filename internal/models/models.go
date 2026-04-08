package models

import "encoding/json"

// APIIndex is the JSON body of GET /api.
type APIIndex struct {
	Characters string `json:"characters" example:"https://thesimpsonsapi.com/api/characters"`
	Episodes   string `json:"episodes" example:"https://thesimpsonsapi.com/api/episodes"`
	Locations  string `json:"locations" example:"https://thesimpsonsapi.com/api/locations"`
}

// PaginatedCharacters is the list response for GET /api/characters.
type PaginatedCharacters struct {
	Count   int                 `json:"count"`
	Next    *string             `json:"next"`
	Prev    *string             `json:"prev"`
	Pages   int                 `json:"pages"`
	Results []CharacterListItem `json:"results"`
}

// CharacterListItem is one row in the characters list (fewer fields than detail).
type CharacterListItem struct {
	ID           int      `json:"id"`
	Age          *int     `json:"age"`
	Birthdate    *string  `json:"birthdate"`
	Gender       string   `json:"gender"`
	Name         string   `json:"name"`
	Occupation   string   `json:"occupation"`
	PortraitPath string   `json:"portrait_path"`
	Phrases      []string `json:"phrases"`
	Status       string   `json:"status"`
}

// CharacterDetail is GET /api/characters/:id (includes nested episode when present).
type CharacterDetail struct {
	ID                   int            `json:"id"`
	Age                  *int           `json:"age"`
	Birthdate            *string        `json:"birthdate"`
	Description          string         `json:"description"`
	FirstAppearanceEpID  *int           `json:"first_appearance_ep_id"`
	FirstAppearanceShID  *int           `json:"first_appearance_sh_id"`
	Gender               string         `json:"gender"`
	Name                 string         `json:"name"`
	Occupation           string         `json:"occupation"`
	Phrases              []string       `json:"phrases"`
	PortraitPath         string         `json:"portrait_path"`
	Status               string         `json:"status"`
	FirstAppearanceEp    *EpisodeDetail `json:"first_appearance_ep,omitempty"`
}

// PaginatedEpisodes is the list response for GET /api/episodes.
type PaginatedEpisodes struct {
	Count   int            `json:"count"`
	Next    *string        `json:"next"`
	Prev    *string        `json:"prev"`
	Pages   int            `json:"pages"`
	Results []EpisodeBrief `json:"results"`
}

// EpisodeBrief is one row in the episodes list.
type EpisodeBrief struct {
	ID            int    `json:"id"`
	Airdate       string `json:"airdate"`
	EpisodeNumber int    `json:"episode_number"`
	ImagePath     string `json:"image_path"`
	Name          string `json:"name"`
	Season        int    `json:"season"`
	Synopsis      string `json:"synopsis"`
}

// EpisodeDetail is GET /api/episodes/:id.
type EpisodeDetail struct {
	ID            int    `json:"id"`
	Airdate       string `json:"airdate"`
	Description   string `json:"description"`
	EpisodeNumber int  `json:"episode_number"`
	ImagePath     string `json:"image_path"`
	Name          string `json:"name"`
	Season        int    `json:"season"`
	Synopsis      string `json:"synopsis"`
}

// PaginatedLocations is the list response for GET /api/locations.
type PaginatedLocations struct {
	Count   int             `json:"count"`
	Next    *string         `json:"next"`
	Prev    *string         `json:"prev"`
	Pages   int             `json:"pages"`
	Results []LocationBrief `json:"results"`
}

// LocationBrief is one row in the locations list.
type LocationBrief struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
	Town      string `json:"town"`
	Use       string `json:"use"`
}

// LocationDetail is GET /api/locations/:id.
type LocationDetail struct {
	ID                  int             `json:"id"`
	Description         string          `json:"description"`
	FirstAppearanceEpID *int          `json:"first_appearance_ep_id"`
	FirstAppearanceShID *int          `json:"first_appearance_sh_id"`
	ImagePath           string          `json:"image_path"`
	Name                string          `json:"name"`
	Town                string          `json:"town"`
	Use                 string          `json:"use"`
	FirstAppearanceSh   json.RawMessage `json:"first_appearance_sh" swaggertype:"object"`
	FirstAppearanceEp   *EpisodeDetail  `json:"first_appearance_ep,omitempty"`
}
