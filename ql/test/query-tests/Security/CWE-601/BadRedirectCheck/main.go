package main

import (
	"net/http"
	"strings"
)

func badRedirect(redirect string, rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, sanitizeUrl(redirect), 302)
}

func goodRedirect(redirect string, rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, sanitizeUrlGood(redirect), 302)
}

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
