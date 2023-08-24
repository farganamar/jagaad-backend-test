package entities

type User struct {
	ID       string   `json:"_id"`
	Index    int64    `json:"index"`
	GUID     string   `json:"guid"`
	IsActive bool     `json:"isActive"`
	Balance  string   `json:"balance"`
	Tags     []string `json:"tags"`
	Friends  []Friend `json:"friends"`
}

type Friend struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
