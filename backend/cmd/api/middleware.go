package main

import "net/http"

// TODO: PanicRecovery
// Goal: Prevent the entire server from crashing if a handler encounters a panic.
// Logic: Recover from panic, log the error using jsonlog, and send a 500 Internal Server Error.
// Priority: High (Add before Production)
func (app *Application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// defer func() { ... if err := recover(); err != nil { ... } }()
		next.ServeHTTP(w, r)
	})
}

// TODO: BodyLimit
// Goal: Protect the server from memory exhaustion attacks (e.g., uploading 1GB video).
// Logic: Wrap the request body in a MaxBytesReader (e.g., limit to 5MB).
// Priority: Medium (Add before Image Upload feature)
func (app *Application) enableBodyLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.Body = http.MaxBytesReader(w, r.Body, 5 * 1024 * 1024)
		next.ServeHTTP(w, r)
	})
}
