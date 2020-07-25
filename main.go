package main

import (
	"Project/auth-cookie/auth"
	"log"
	"net/http"
)

func main() {

	// handler to login and create auth cookie
	http.HandleFunc("/login/", Login)
	// page authorized only with valid cookie
	http.HandleFunc("/welcome/", Welcome)
	// endpoint to refresh auth cookie
	http.HandleFunc("/refresh", Refresh)
	// remove auth cookie
	http.HandleFunc("/logout/", Logout)

	log.Println("linsten port 8000...") // DEBUG
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")

		form := `
			<form action="/login/" method="post">
				<label for="email">Email</label>
				<input type="text" id="email" name="email"><br><br>
				<label for="password">Password:</label>
				<input type="text" id="password" name="password"><br><br>
				<input type="submit" value="Submit">
			</form>
		`

		w.Write([]byte(form))
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("err to parse:", err) // DEBUG
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		creds := auth.Credentials{Email: r.FormValue("email"), Password: r.FormValue("password")}

		user, err := auth.ValidCredentials(creds)
		if err != nil {
			log.Println("err to valid credential:", err) // DEBUG
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}

		cookie, err := auth.CreateCookie(user)
		if err != nil {
			log.Println("err to create:", err) // DEBUG
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("sorry, try again"))
			return
		}

		http.SetCookie(w, cookie)
		w.Write([]byte(user.Name + " is logged!"))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func Welcome(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		cookie, claims, err := auth.ValidCookie(r)
		if err != nil {
			log.Println("err to valid cookie:", err) // DEBUG
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}

		err = auth.UpdateCookie(cookie, claims)
		if err != nil {
			log.Println("err to update:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("sorry, try again"))
			return
		}

		http.SetCookie(w, cookie)
		w.Write([]byte("welcome " + claims.Name))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func Refresh(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		cookie, claims, err := auth.ValidCookie(r)
		if err != nil {
			log.Println("err to valid:", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}

		err = auth.UpdateCookie(cookie, claims)
		if err != nil {
			log.Println("err to update:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("sorry, try again"))
			return
		}

		http.SetCookie(w, cookie)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		cookie := &http.Cookie{}
		auth.SetBasicCookie(cookie)
		http.SetCookie(w, cookie)
		w.Write([]byte("logout"))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
