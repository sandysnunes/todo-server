package main

type Tag struct {
	Description string `json:"description"`
}

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
	Favorite    bool   `json:"favorite"`
	Tags        []Tag  `json:"tags"`
}
