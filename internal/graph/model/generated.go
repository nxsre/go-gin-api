// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Query struct {
}

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Mobile string `json:"mobile"`
}

type UpdateUserMobileInput struct {
	ID     string `json:"id"`
	Mobile string `json:"mobile"`
}
