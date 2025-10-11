package models

import "net/http"

type UserFields struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (uf UserFields) Validate() (msgs []string, ok bool) {
	if uf.FirstName == "" {
		msgs = append(msgs, "first_name is required")
	}
	if uf.LastName == "" {
		msgs = append(msgs, "last_name is required")
	}
	if uf.Email == "" {
		msgs = append(msgs, "email is required")
	}
	return msgs, len(msgs) == 0
}

type User struct {
	// Embed the fields from UserFields.
	UserFields
	ID int `json:"id"`
}

func (u User) Validate() (msgs []string, ok bool) {
	return u.UserFields.Validate()
}

// GET /users
type UsersGetResponse struct {
	Users []User `json:"users"`
}

// POST /users
type UsersPostRequest struct {
	UserFields
}
type UsersPostResponse struct {
	OK bool `json:"ok"`
}

// Error responses.
type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func NewInvalidRequestResponse(msgs []string) InvalidRequestResponse {
	return InvalidRequestResponse{
		ErrorResponse: ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  "Invalid request",
		},
		Messages: msgs,
	}
}

type InvalidRequestResponse struct {
	ErrorResponse
	Messages []string `json:"messages"`
}
