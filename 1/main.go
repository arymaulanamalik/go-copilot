package main

import "net/http"

// html form
// name
// email

func serveForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
		<form method="POST" action="/process">
			<input type="text" name="name">
			<input type="text" name="email">
			<input type="submit">
		</form>
	`))
}

func processForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Hello " + r.FormValue("name")))
}

func main() {
	http.HandleFunc("/", serveForm)
	http.HandleFunc("/process", processForm)
	http.ListenAndServe(":8080", nil)
}
