package middleware

import (
	"fmt"
	"net/http"
)

type wrapper struct {
	http.ResponseWriter
	request    *http.Request
	statusCode int
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		wrappedWriter := &wrapper{
			ResponseWriter: writer,
			request:        request,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrappedWriter, request)
		fmt.Printf("%s %s %v", request.Method, request.RequestURI, wrappedWriter.statusCode)
	})
}
