package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://localhost/phpmyadmin/db_structure.php?server=1&db=latihan_api"

	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("Url\t: %s\n", urlString)

	fmt.Printf("Protocol : %s\n", u.Scheme)
	fmt.Printf("host\t: %s\n", u.Host)
	fmt.Printf("Path\t: %s\n", u.Path)

	var server = u.Query()["server"][0]
	var db = u.Query()["db"][0]
	fmt.Printf("Server\t: %s\nDB\t: %s\n", server, db)
}
