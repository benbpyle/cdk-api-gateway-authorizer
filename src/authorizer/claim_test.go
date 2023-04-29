package main

import (
	"testing"

	"github.com/lestrrat-go/jwx/jwt"
)

func TestNewClaims_IsAdmin(t *testing.T) {
	tok := jwt.New()
	_ = tok.Set(jwt.IssuerKey, "sister/mary/margaret")
	_ = tok.Set(jwt.AudienceKey, "users")
	_ = tok.Set("isAdmin", "true")

	claims := NewClaims(tok)
	if !claims.IsAdmin {
		t.Fail()
	}
}

func TestNewClaims_IsNotAdmin(t *testing.T) {
	tok := jwt.New()
	_ = tok.Set(jwt.IssuerKey, "sister/mary/margaret")
	_ = tok.Set(jwt.AudienceKey, "users")
	_ = tok.Set("isAdmin", "false")

	claims := NewClaims(tok)
	if claims.IsAdmin {
		t.Fail()
	}
}

func TestNewClaims_NoLocations(t *testing.T) {
	tok := jwt.New()
	_ = tok.Set(jwt.IssuerKey, "sister/mary/margaret")
	_ = tok.Set(jwt.AudienceKey, "users")
	_ = tok.Set("isAdmin", "true")

	claims := NewClaims(tok)
	if len(claims.AllowedLocations) != 0 {
		t.Fail()
	}
}

func TestNewClaims_OneLocations(t *testing.T) {
	tok := jwt.New()
	_ = tok.Set(jwt.IssuerKey, "sister/mary/margaret")
	_ = tok.Set(jwt.AudienceKey, "users")
	_ = tok.Set("isAdmin", "true")
	_ = tok.Set("allowedLocations", "697190")

	claims := NewClaims(tok)

	if len(claims.AllowedLocations) != 1 {
		t.Fail()
	}
}

func TestNewClaims_ThreeLocations(t *testing.T) {
	tok := jwt.New()
	_ = tok.Set(jwt.IssuerKey, "sister/mary/margaret")
	_ = tok.Set(jwt.AudienceKey, "users")
	_ = tok.Set("isAdmin", "true")
	_ = tok.Set("allowedLocations", "697190,69434,194956")

	claims := NewClaims(tok)

	if len(claims.AllowedLocations) != 3 {
		t.Fail()
	}
}

func TestDumpClaims_IsAdmin(t *testing.T) {
	tok := jwt.New()
	_ = tok.Set(jwt.IssuerKey, "sister/mary/margaret")
	_ = tok.Set(jwt.AudienceKey, "users")
	_ = tok.Set("isAdmin", "true")

	claims := DumpClaims(tok)
	if _, ok := claims["isAdmin"]; !ok {
		t.Fail()
	}

	if v, ok := claims["isAdmin"]; ok {
		b := v.(string)
		if b != "true" {
			t.Fail()
		}
	}
}

func TestDumpClaims_IsNotAdmin(t *testing.T) {
	tok := jwt.New()
	_ = tok.Set(jwt.IssuerKey, "sister/mary/margaret")
	_ = tok.Set(jwt.AudienceKey, "users")
	_ = tok.Set("isAdmin", "false")

	claims := DumpClaims(tok)
	if _, ok := claims["isAdmin"]; !ok {
		t.Fail()
	}

	if v, ok := claims["isAdmin"]; ok {
		b := v.(string)
		if b != "false" {
			t.Fail()
		}
	}
}

func TestDumpClaims_NoLocations(t *testing.T) {
	tok := jwt.New()
	_ = tok.Set(jwt.IssuerKey, "sister/mary/margaret")
	_ = tok.Set(jwt.AudienceKey, "users")
	_ = tok.Set("isAdmin", "true")

	claims := DumpClaims(tok)
	if _, ok := claims["allowedLocations"]; ok {
		t.Fail()
	}
}

func TestDumpClaims_OneLocations(t *testing.T) {
	tok := jwt.New()
	_ = tok.Set(jwt.IssuerKey, "sister/mary/margaret")
	_ = tok.Set(jwt.AudienceKey, "users")
	_ = tok.Set("isAdmin", "true")
	_ = tok.Set("allowedLocations", "697190")

	claims := DumpClaims(tok)
	if _, ok := claims["allowedLocations"]; !ok {
		t.Fail()
	}

	if v, ok := claims["allowedLocations"]; ok {
		b := v.(string)
		if b != "697190" {
			t.Fail()
		}
	}
}

func TestDumpClaims_ThreeLocations(t *testing.T) {
	tok := jwt.New()
	_ = tok.Set(jwt.IssuerKey, "sister/mary/margaret")
	_ = tok.Set(jwt.AudienceKey, "users")
	_ = tok.Set("isAdmin", "true")
	_ = tok.Set("allowedLocations", "697190,69434,194956")

	claims := DumpClaims(tok)
	if _, ok := claims["allowedLocations"]; !ok {
		t.Fail()
	}

	if v, ok := claims["allowedLocations"]; ok {
		b := v.(string)
		if b != "697190,69434,194956" {
			t.Fail()
		}
	}
}
