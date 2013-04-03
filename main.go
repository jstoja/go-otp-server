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

func HandleStaticFiles(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		SendFile(w, "views/index.html")
	case "/chart.js":
		SendFile(w, "assets/chart.js")
	}
}

func HandleRequires(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		str := r.URL.Query()["otp"][0]
		j, _ := json.Marshal(otp.Totp(str))
		fmt.Println(str)
		fmt.Printf("OTP required for pass: %s\n", string(j))
		fmt.Fprintln(w, string(j))
	}
}

func main() {
	//password := "ymybvnckruprgkgr"
	http.HandleFunc("/", HandleStaticFiles)
	http.HandleFunc("/require", HandleRequires)
	//fmt.Printf("%d\n", otp.Totp(password))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
