package main

// Tag define a tag name
type Tag struct {
	Description string `json:"description"`
}

// Todo define a to-do task
type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
	Favorite    bool   `json:"favorite"`
	Tags        []Tag  `json:"tags"`
}
