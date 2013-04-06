package main

import (
	"./otp"
	"fmt"
	//"html"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func SendFile(w io.Writer, filename string) {
	fmt.Printf("filename: %s\n", filename)
	file, err := os.Open(filename)
	if err == nil {
		io.Copy(w, file)
		file.Close()
	}
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	SendFile(w, "views/index.html")
}

func HandleRequires(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		str := r.URL.Query()["otp"][0]
		j, _ := json.Marshal(gotp.Totp(str))

		fmt.Fprintln(w, string(j))
	}
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mail := r.URL.Query()["mail"][0]
		key := r.URL.Query()["key"][0]
		fmt.Printf("The mail is %s and the key %s\n", mail, key)
	}
}

func main() {
	//password := "ymybvnckruprgkgr"
	password := "coucou"
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/requireOTP", HandleRequires)
	http.HandleFunc("/register", HandleRegister)
	fmt.Printf("%d\n", gotp.Totp(password))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
