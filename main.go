package main

import "fjpc/go-oauth2/http"

func main() {
	sc, err := http.NewServer("config.yml")
	if err != nil {
		panic(err)
	}
	if err = sc.Start(); err != nil {
		panic(err)
	}
}
