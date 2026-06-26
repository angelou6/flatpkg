package flathub

type Hits struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
	Id      string `json:"app_id"`
}

type Results struct {
	Hits []Hits `json:"hits"`
}
