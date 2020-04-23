package main

import (
	"net/http"
	"strings"
)

// CVE-2018-15178
// Code from github.com/gogs/gogs
func isValidRedirect(url string) bool {
	return len(url) >= 2 && url[0] == '/' && url[1] != '/' // NOT OK
}

func isValidRedirect1(url string) bool {
	return len(url) >= 2 && url[0] == '/' && url[1] != '/' && url[1] != '\\' // OK
}

func dummyUse1(w http.ResponseWriter, r *http.Request) {
	url := ""
	if isValidRedirect(url) {
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func dummyUse2(w http.ResponseWriter, r *http.Request) {
	url := ""
	if isValidRedirect1(url) {
		http.Redirect(w, r, url, http.StatusFound)
	}
}

// CVE-2017-1000070 (both vulnerable!)
// Code from github.com/bitly/oauth2_proxy
func OAuthCallback(rw http.ResponseWriter, req *http.Request) {
	redirect := req.Form.Get("state")
	if !strings.HasPrefix(redirect, "/") { // NOT OK
		redirect = "/"
	}
	http.Redirect(rw, req, redirect, http.StatusFound)
}

func OAuthCallback1(rw http.ResponseWriter, req *http.Request) {
	redirect := req.Form.Get("state")
	if !strings.HasPrefix(redirect, "/") || strings.HasPrefix(redirect, "//") { // NOT OK
		redirect = "/"
	}
	http.Redirect(rw, req, redirect, http.StatusFound)
}
