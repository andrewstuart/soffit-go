package soffit

import (
	"encoding/base64"
	"net/http"
	"strings"

	jasypt "astuart.co/go-jasypt"

	"github.com/dgrijalva/jwt-go"
)

const headerPrefix = "X-Soffit-"

// Headers provides structured access to the several headers that may be
// sent in the soffit payload.
type Headers struct {
	Preferences map[string]interface{} `json:"preferences"`
	Definition  map[string]interface{} `json:"definition"`
	Request     map[string]interface{} `json:"request"`
}

// Receiver takes a password, provides utilities for handling incoming Soffit
// requests.
type Receiver struct {
	Password string
}

// GetHeaders takes url.Values and returns the decrypted headers.
func (d Receiver) GetHeaders(h http.Header) (*Headers, error) {
	var s Headers

	for k := range h {
		if strings.Index(k, "X-Soffit") != 0 {
			continue
		}

		bs, err := base64.StdEncoding.DecodeString(h.Get(k))
		if err != nil {
			return nil, err
		}
		err = jasypt.DecryptJasypt(bs, d.Password)
		if err != nil {
			return nil, err
		}

		token, err := jwt.Parse(string(bs), nil)

		if err != nil && !strings.Contains(err.Error(), "Keyfunc") {
			return nil, err
		}

		switch k {
		case "X-Soffit-Portalrequest":
			s.Request = token.Claims.(jwt.MapClaims)
		case "X-Soffit-Definition":
			s.Definition = token.Claims.(jwt.MapClaims)
		case "X-Soffit-Preferences":
			s.Preferences = token.Claims.(jwt.MapClaims)
		}
	}

	return &s, nil
}
