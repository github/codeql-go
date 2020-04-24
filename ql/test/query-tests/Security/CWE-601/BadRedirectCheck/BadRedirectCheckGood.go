package main

import "net/http"

func sanitizeUrlGood(redir string) string {
	if len(redir) > 1 && redir[0] == '/' && redir[1] != '/' && redir[1] != '\\' {
		return redir
	}
	return "/"
}

func goodRedirect(redirect string, rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, sanitizeUrlGood(redirect), 302)
}