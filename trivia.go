package trivia

type Trivia struct {
	Id       int64  `json:"id" db:"trivia_id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
