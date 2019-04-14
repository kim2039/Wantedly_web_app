package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "log"
    "net/http"
)

func main() {
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
        w.WriteJson(map[string]string{"message": "Hello World!"})
    }))
    log.Fatal(http.ListenAndServe(":80", api.MakeHandler())) 
}
