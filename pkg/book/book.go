package book

// Book is a struct which represents a record of book
type Book struct {
	// ID is an unique value in the library
	ID       int `json:"id,omitempty"`
	Revision int `json:"revision,omitempty"`
	// ISBN number validation has to be performed
	// can contain hyphens. you can use an existing library
	// for the validation
	ISBN string `json:"isbn,omitempty"`
	// Can contain commas
	Title string `json:"title,omitempty"`
	// Can contain commas
	Author string `json:"author,omitempty"`
}
