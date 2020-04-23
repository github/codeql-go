package main

import (
	"net/http"
	"strings"
)

func isValidRedir(redirect string) bool {
	switch {
	// Not OK: does not check for '/\'
	case strings.HasPrefix(redirect, "/") && !strings.HasPrefix(redirect, "//"):
		return true
	default:
		return false
	}
}

func isValidRedir1(redirect string) bool {
	switch {
	// OK
	case strings.HasPrefix(redirect, "/") && !strings.HasPrefix(redirect, "//") && !strings.HasPrefix(redirect, "/\\"):
		return true
	default:
		return false
	}
}

func dummyUse(w http.ResponseWriter, r *http.Request) {
	url := ""
	if isValidRedir(url) && isValidRedir1(url) {
		http.Redirect(w, r, url, http.StatusFound)
	}

	http.Redirect(w, r, sanitizeUrl(url), http.StatusFound)
}
