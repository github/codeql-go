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

func alsoABadRedirect1(url string, rw http.ResponseWriter, req *http.Request) {
	if isValidRedir(url) {
		http.Redirect(rw, req, url, 302)
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

func goodRedirect2(url string, rw http.ResponseWriter, req *http.Request) {
	if isValidRedirectGood(url) {
		http.Redirect(rw, req, url, 302)
	}
}
