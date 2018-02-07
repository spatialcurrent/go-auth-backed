package authbackend

import (
	"net/http"
)

type User interface {
	GetId() string
	GetUsername() string
}

type Team interface {
	GetId() string
	GetName() string
}

type Backend interface {
	Type() string
	Connect(map[string]string) error
	AuthenticateRequest(r *http.Request) (bool, User, Team, error)
}
