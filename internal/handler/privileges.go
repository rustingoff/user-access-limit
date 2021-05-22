package handler

import (
	"net/http"
)

func CheckPrivileges(url string, fn func(w http.ResponseWriter, r *http.Request, privileg map[uint]bool)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, user.Privileges)
	}
}

func TestingRoute(url string, fn func(w http.ResponseWriter, r *http.Request, privileg map[uint]bool)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, user.Privileges)
	}
}
