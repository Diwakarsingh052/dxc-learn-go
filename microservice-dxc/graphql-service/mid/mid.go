package mid

import (
	"context"
	"graphql-service/auth"
	"net/http"
)

// FetchToken put the token value in the context
// This helps to check necessary details about the use in the request lifetime
func FetchToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if token is passed when doing request then we fetch it from the auth header
		authHeader := r.Header.Get("Authorization")

		//token format :- bearer <token> // assignment:- separate both strings with space

		//putting the token inside the auth header in the context, so it could be fetched in the request lifetime
		ctx := context.WithValue(r.Context(), auth.Key, authHeader)

		// setting our request to use the context
		r = r.WithContext(ctx)

		//doing next thing in the chain
		next.ServeHTTP(w, r)

	})
}
