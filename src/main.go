package main

import (
     "net/http"
)

func    main () {
    http.Handle("/", http.FileServer(http.Dir("../www")))
    http.Handle("/miner", http.FileServer(http.Dir("../www/miner")))
    http.Handle("/miner", http.FileServer(http.Dir("../www/login/login.html")))
    http.ListenAndServe(":8080", nil)
}