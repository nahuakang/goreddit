package web

import "encoding/gob"

func init() {
	gob.Register(CreatePostForm{})
	gob.Register(FormErrors{})
}

// FormErrors store errors for validating forms
type FormErrors map[string]string

// CreatePostForm stores form values for new posts
type CreatePostForm struct {
	Title   string
	Content string
	Errors  FormErrors
}

// Validate valites the post forms
func (f *CreatePostForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Title == "" {
		f.Errors["Title"] = "Please enter a title."
	}
	if f.Content == "" {
		f.Errors["Content"] = "Please enter a text."
	}

	return len(f.Errors) == 0
}
