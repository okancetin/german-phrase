package entity

type Phrase struct {
	Id          string      `json:"id"`
	Content     string      `json:"content"`
	Title       string      `json:"title"`
	Link        string      `json:"link"`
	Translation Translation `json:"translation"`
}

type Translation struct {
	En string `json:"en"`
	De string `json:"de"`
}
