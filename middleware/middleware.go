package middleware

import (
	"compress/gzip"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/utils"
	"net/http"
	"strings"
)

type wrapper struct {
	http.ResponseWriter
	request    *http.Request
	statusCode int
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *wrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
	return
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		wrappedWriter := &wrapper{
			ResponseWriter: writer,
			request:        request,
			statusCode:     http.StatusOK,
		}

		if strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
			wrappedWriter.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(writer)
			defer gz.Close()

			gzWritter := gzipResponseWriter{
				ResponseWriter: writer,
				Writer:         gz,
			}
			next.ServeHTTP(gzWritter, request)
			app.Server.Logger.Printf("%s %s %s %v \n", utils.ClientIP(request), request.Method, request.RequestURI, wrappedWriter.statusCode)
			return
		}

		next.ServeHTTP(wrappedWriter, request)
		app.Server.Logger.Printf("%s %s %s %v \n", utils.ClientIP(request), request.Method, request.RequestURI, wrappedWriter.statusCode)
	})
}
