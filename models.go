package main

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Favorite    bool   `json:"favorite"`
}
