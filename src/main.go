package main

import (
  "html/template"
  "fmt"
  "net/http"
  "strings"
  "flag"
  "strconv"
  "github.com/fatih/color"
)

type Page struct {
  Title string
  Url string
}

func handler(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path[1:]
  templateName := getFileName(path)
  t, err := template.ParseFiles(templateName)

  if err != nil {
    fmt.Fprintf(w, "File %s not found", templateName)
  } else {
    t.Execute(w, &Page{Title: "Accueil", Url: templateName})
  }
}

func getFileName(value string) string {
  if strings.Contains(value, ".") {
    return value;
  } else if value == "" || value == "/" {
    return "index.html"
  } else {
    return value + ".html"
  }
}

func setupFlags() int {
  portPtr := flag.Int("port", 1337, "Port listen")

  flag.Parse()

  return *portPtr
}

func main() {
  port := strconv.Itoa(setupFlags())
  color.Yellow("Server running on localhost:"+ port)

  http.HandleFunc("/", handler)
  http.ListenAndServe(":" + port, nil)
}