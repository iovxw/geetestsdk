package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Bluek404/geetestsdk"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("validate.gtpl")
		if err != nil {
			log.Println(err)
		}
		t.Execute(w, nil)
	})

	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			challenge := r.FormValue("geetest_challenge")
			validate := r.FormValue("geetest_validate")
			seccode := r.FormValue("geetest_seccode")

			sdk := geetestsdk.New("3c3f7cfcdcf21e216bf4ed7930c191e2")

			ok, err := sdk.Validate(challenge, validate, seccode)
			if err != nil {
				log.Println(err)
			}

			if ok {
				fmt.Fprintf(w, "验证成功")
			} else {
				fmt.Fprintf(w, "验证失败")
			}
		} else {
			// Method错误，重定向到首页
			http.Redirect(w, r, "http://127.0.0.1:8080", http.StatusMovedPermanently)
		}
	})

	http.ListenAndServe(":8080", nil)
}
