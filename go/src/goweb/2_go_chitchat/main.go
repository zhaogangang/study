package main

import "net/http"

func main() {
	p("Chitchat", version(), "started at", config.Address)

	// handle static assetst
	mux := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// all route patterns matched here
	// route handler functions defined in other files

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFun("/signup", signup)
	mux.HandleFunc("/signup_accout", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFun("/thread/new", newThread)
	mux.HandleFun("/thread/create", createThread)
	mux.HandleFun("/thread/post", postThread)
	mux.HandleFun("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr: config.Address,
		Handler: mux,
		ReadTimeout: time.Duration(config.ReadTimeout * int64(time.Second))
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second))
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
