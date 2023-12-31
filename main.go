package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "Serve on port")
}

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir("./"))

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")

		slog.Info("Request", "host", r.Host, "method", r.Method, "path", r.URL.Path, "origin", r.Header.Get("Origin"))

		fs.ServeHTTP(w, r)
	})

	slog.Info("Server starting...", "port", port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		slog.Error(err.Error())
		return
	}

	slog.Info("Server started", "URL", fmt.Sprintf("http://localhost:%v", port))

	if err := http.Serve(listener, nil); err != nil {
		slog.Error(err.Error())
	}
}
