package middleware

import (
	"net/http"

	"github.com/lukegrn/days/pkg/hash"
)

func PasswordProtect(compare string, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	f := func(w http.ResponseWriter, r *http.Request) {
		pw := r.PostFormValue("password")

		if !hash.EqToHash(pw, compare) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		handler(w, r)
	}

	return f
}
