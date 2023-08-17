package cores

type Soldier struct {
	Rank       string `json:"rank"`
	Wife       string `json:"wife"`
	Salary     int    `json:"salary"`
	Home       bool   `json:"home"`
	Car        bool   `json:"car"`
	Corruption bool   `json:"corruption"`
}
