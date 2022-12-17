package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	// ルーティングパス設定
	dir, _ := os.Getwd()
	dir2 := strings.ReplaceAll(dir, "/cmd", "")

	// ルーティング
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(dir2+"/assets"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/terms", terms)
	http.HandleFunc("/privacy", privacy)
	http.HandleFunc("/getSignup", getSignup)
	http.HandleFunc("/postSignup", postSignup)
	http.HandleFunc("/getLogin", getLogin)
	http.HandleFunc("/postLogin", postLogin)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/getReset", getReset)
	http.HandleFunc("/postReset", postReset)
	http.HandleFunc("/favicon.ico", faviconHandler)

	// ListenAndServe
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
