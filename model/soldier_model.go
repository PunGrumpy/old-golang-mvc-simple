package model

type Soldier struct {
	ID         int    `json:"id"`
	Rank       string `json:"rank"`
	Name       string `json:"name"`
	Wife       string `json:"wife"`
	Salary     int    `json:"salary"`
	Home       bool   `json:"home"`
	Car        bool   `json:"car"`
	Corruption bool   `json:"corruption"`
}
