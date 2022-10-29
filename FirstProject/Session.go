package main

import (
	"html/template"
	"net/http"
)

// aktuell angemeldeter user und dementsprechende funktionen wie :
// nutzer sich anmelden kann
// einen account erstellen kann
// Anzeigen gespeicherten Kombos über Linkliste zu diesen, löschen und speichern von kombos (nutze funktionen aus database.go)
// füge einen neuen handler im mux ein und ein neues template, das anmeldungen und account erstellung verwaltet
// bennen evtl Dateinamen um, scheint mir nicht passend zu sein

var currentUser *User
var savedCombosPath = "/saves/"

type UserAndCookies struct {
	User    User
	Cookies []*http.Cookie
}

func signIn(usersame string, password string) bool {
	user := getUser(usersame)

	if user == nil {
		return false
	}
	if user.Password == password {
		currentUser = user
		return true
	}
	return false
}

func signOut() {
	currentUser = nil
}

func createAccount(userName string, password string, comboLinks []string) bool {
	if !addUser(userName, password) {
		return false
	}
	user := getUser(userName)
	for _, c := range comboLinks {
		addCombo(*user, c)
	}
	return true
}

func deleteAccount(password string) bool {
	if currentUser.Password == password {
		return deleteUser(*currentUser)
	}
	return false
}

func getComboLinks() []string {
	return currentUser.ComboLinks.ComboLinks
}

// hier aufgehört: Über func map sollen Go-Funktionen im HTML Template aufgerufen werden können (hier um Kombos in die db zu laden)
// testuser soll noch durch 'current user' (der den aktuell angemeldeten User darstellt) ersetzt werden
func savedCombosHandler(w http.ResponseWriter, r *http.Request) {
	user := testBentutzer()
	cookies := r.Cookies()
	userAndCookies := UserAndCookies{User: user, Cookies: cookies}
	t, _ := template.New("comboLinks.html").Funcs(template.FuncMap{
		"saveCombo": func(user User, comboLink string) {
			err := addCombo(user, comboLink)
			if err != nil {

			}
		},
	}).ParseFiles("comboLinks.html")
	t.Execute(w, userAndCookies)
}

func testBentutzer() User {
	var comboLink ComboLinks
	comboLink.ComboLinks = []string{"test", "test2", "test3"}
	user := User{ComboLinks: comboLink}
	return user
}
