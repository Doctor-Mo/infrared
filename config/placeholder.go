package config

type placeholder struct {
	Version     version     `json:"version"`
	Players     players     `json:"players"`
	Description description `json:"description"`
	Favicon     string      `json:"favicon"`
}

type version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type players struct {
	Max    int      `json:"max"`
	Online int      `json:"online"`
	Sample []sample `json:"sample"`
}

type sample struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type description struct {
	Text string `json:"text"`
}
