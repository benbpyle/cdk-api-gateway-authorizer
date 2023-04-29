package main

import (
	"github.com/lestrrat-go/jwx/jwt"
)

type Claims struct {
	IsAdmin          bool
	AllowedLocations []string
}

func DumpClaims(token jwt.Token) map[string]interface{} {
	m := make(map[string]interface{})

	m["customKey"] = "SomeValueHere"

	return m
}
