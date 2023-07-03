package entity

// Pagination is the process of dividing a document into discrete pages.
type Pagination struct {
	Page  int `validate:"gte=0,default=1"`
	Limit int `validate:"gte=0,default=10"`
}

// Filter searching custom of keyword.
type Filter struct {
	Query string `json:"q"`
	Sort  string `json:"sort"`
}

// Keys unique identity key.
type Keys struct {
	ID int `validate:"gt=0,required"`
}
