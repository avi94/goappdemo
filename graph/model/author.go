package model

type Author struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Posts []*Post `json:"posts"`
}
