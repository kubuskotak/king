package entity

// Article data transfer object.
type Article struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty" validate:"required,min=10"`
	Body        string `json:"body,omitempty" validate:"required,min=10"`
	Description string `json:"description,omitempty" validate:"required,min=10"`
	Slug        string `json:"slug,omitempty" validate:"required,min=10"`
	UserID      int    `json:"user_id,omitempty"`
}
