package main

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type yamlpath struct {
	Data map[string]string `json:"data"` //data will not work because it cant be imported when unmarshaling
}

func getyaml() string {
	var path string
	path = `
data:
   /home: https://github.com/Kudistalbekov
   /email: https://mail.google.com/mail/u/0/#inbox
   /telegram: https://web.telegram.org/#/im?p=u608708913_16198332196946819856
   /whatsapp: https://web.whatsapp.com/
   /udemy: https://www.udemy.com
   /heroku: https://dashboard.heroku.com/login`
	return path

}
func getpath() map[string]string {
	path := make(map[string]string)
	path["/home"] = "https://github.com/Kudistalbekov"
	path["/email"] = "https://mail.google.com/mail/u/0/#inbox"
	path["/telegram"] = "https://web.telegram.org/#/im?p=u608708913_16198332196946819856"
	path["/whatsapp"] = "https://web.whatsapp.com/"
	path["/udemu"] = "https://www.udemy.com"
	path["/heroku"] = "https://dashboard.heroku.com/login"
	return path
}

func myhandler(path map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userpath := r.URL.Path
		if v, ok := path[userpath]; ok {
			http.Redirect(w, r, v, http.StatusFound)
			return
		}
		http.Error(w, "Not found page", http.StatusNotFound)
	}
}

func myyamlhandler(y string) http.HandlerFunc {
	var data yamlpath
	err := yaml.Unmarshal([]byte(y), &data)
	if err != nil {
		log.Fatal(err)
	}
	return myhandler(data.Data)
}
