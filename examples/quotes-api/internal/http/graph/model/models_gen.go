// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Quote struct {
	ID     string  `json:"id"`
	Text   string  `json:"text"`
	Author *Author `json:"author"`
}
