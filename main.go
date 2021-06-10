package main

import (
  "fmt"
  "os"
  "log"
  "net/http"
  "strconv"
  "encoding/json"
)

func mul(a, b int) int {
  return a * b
}

func add(a, b int) int {
  return a + b
}

func minus(a, b int) int {
  return a - b
}

type Nums struct {
  N1 int `json:"n1"`
  N2 int `json:"n2"`
}

var cal_method = "add"

func main() {
  tmpStr, ex := os.LookupEnv("CAL_METHOD")
  if !ex {
    log.Printf("The env variable %s is not set.\n", "CAL_METHOD")
    cal_method = "add"
  } else {
    switch tmpStr {
    case "add":
      cal_method = "add"
    case "delete":
      cal_method = "minus"
    case "multiply":
      cal_method = "mul"
    default:
      cal_method = "add"
    }
  }

  http.HandleFunc("/", handler)
  fmt.Println("Server is listening port 7777, calculation method id is ", cal_method)
  log.Fatal(http.ListenAndServe("0.0.0.0:7777", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  var ns Nums
  if err := json.NewDecoder(r.Body).Decode(&ns); err != nil {
    r.Body.Close()
    log.Fatal(err)
  }

  var result int
  switch cal_method {
  case "add":
    result = add(ns.N1, ns.N2)
  case "minus":
    result = minus(ns.N1, ns.N2)
  case "mul":
    result = mul(ns.N1, ns.N2)
  }
  fmt.Fprintln(w, "the result is :" + strconv.Itoa(result))
}
