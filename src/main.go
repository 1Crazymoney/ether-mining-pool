package main

import (
    "net/http"
)

func    main () {
    http.Handle("/", http.FileServer(http.Dir("../www")))
    http.Handle("/miner", http.FileServer(http.Dir("../www/miner")))
    http.ListenAndServeTSL(":8080", "cert.pem", "key.pem", nil)
}