package main

import (
  "html/template"
  "fmt"
  "net/http"
  "strings"
  "flag"
  "strconv"
  "github.com/fatih/color"
  "os"
)

type Page struct {
  Url string
}

func handler(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path[1:]
  templateName := getFileName(path)
  t, err := template.ParseFiles(templateName)

  if err != nil {
    fmt.Fprintf(w, "File %s not found", templateName)
  } else {
    t.Execute(w, &Page{Url: templateName})
  }
}

func getFileName(value string) string {
  var result string

  if strings.Contains(value, ".") {
    result = value;
  } else if value == "" || value == "/" {
    result = "index.html"
  } else {
    result = value + ".html"
  }

  path, _ := os.Getwd()

  return  path + "/" + result
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