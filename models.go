package main

type Tag struct {
	Description string `json:"description"`
}

// TODO archived
type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Favorite    bool   `json:"favorite"`
	Tags        []Tag  `json:"tags"`
}
