package main

import ("fmt"
"math/rand"
"net/http"
"time")

func main() {
http.HandleFunc("/webfortune", webfortunehandler)
seed := time.Now().UnixNano()
rand.Seed(seed)

http.ListenAndServe(":8080", nil)
}

func webfortunehandler(w http.ResponseWriter, r *http.Request) {
n := rand.Intn(100)
var result string

if n > 90 {
result = "大吉"
} else if n > 60 {
result = "中吉"
} else if n > 30 {
result = "吉"
} else {
result = "凶"
}

fmt.Fprintf(w, "今の運勢は %s です！\n", result)
}
