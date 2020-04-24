package main

import "net/http"

func sanitizeUrl(redir string) string {
	if len(redir) > 0 && redir[0] == '/' {
		return redir
	}
	return "/"
}

func badRedirect(redirect string, rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, sanitizeUrl(redirect), 302)
}