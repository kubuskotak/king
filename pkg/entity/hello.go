// Package entity contains objects that represent the system data of the usecase or persist.
package entity

// Hello data transfer object.
type Hello struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Body        string `json:"body,omitempty"`
	Description string `json:"description,omitempty"`
	Slug        string `json:"slug,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
}

// HelloParams data transfer object.
type HelloParams struct {
	ID int `validate:"gt=0,required"`
}
