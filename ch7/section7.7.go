// package main

// import (
// 	"log"
// 	"fmt"
// 	// "strings"
// 	"net/http"
// )

// type dollars float32
// func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// type database map[string]dollars
// func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	for k, v := range d {
// 		fmt.Fprintf(w, "%s: %s\n", k, v)
// 	}
// }

// type database2 map[string]dollars
// func (db database2) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	switch r.URL.Path {
// 		case "/list":
// 			for k, v := range db {
// 				fmt.Fprintf(w, "%s: %s\n", k, v)
// 			}
// 		case "/price":
// 			item := r.URL.Query().Get("item")

// 			price, ok := db[item]
// 			if !ok {
// 			w.WriteHeader(http.StatusNotFound) // 404
// 			fmt.Fprintf(w, "no such item: %q\n", item)
// 			return
// 			}
// 			fmt.Fprintf(w, "%s\n", price)
// 		default:
// 			w.WriteHeader(http.StatusNotFound) // 404
// 			fmt.Fprintf(w, "no such page: %s\n", r.URL)
// 			// msg := fmt.Sprintf("no such page: %s\n", r.URL)
// 			// http.Error(w, msg, http.StatusNotFound) // 404

// //***********My Edition***********
// 			// if strings.HasPrefix(r.URL.Path, "/price") {
// 			// 	fmt.Println(r.URL.RawQuery)
// 			// 	if "" == r.URL.RawQuery {
// 			// 		fmt.Fprintf(w, "Please specify query parameter\n")
// 			// 		return
// 			// 	}
		
// 			// 	rawQuery := strings.SplitN(r.URL.RawQuery, "=", 2)
// 			// 	val := rawQuery[len(rawQuery)-1]
		
// 			// 	price, ok := db[val]
// 			// 	if !ok {
// 			// 		fmt.Fprintf(w, "%s's price is not exist\n", val)
// 			// 	}else {
// 			// 		fmt.Fprintf(w, "%s's price = %v\n", val, price)
// 			// 	}
// 			// }
// //***********My Edition***********
// 	}

// 	fmt.Fprintf(w, "ok !!!\n")
// }

// func main() {
// 	// db := database{"shoes": 50, "socks": 5}
// 	// log.Fatal(http.ListenAndServe("172.26.88.88:8000", db))

// 	db2 := database2{"shoes": 50, "socks": 5}
// 	log.Fatal(http.ListenAndServe("172.26.88.88:8001", db2))
// }
//*******************************************

package main

import (
	"log"
	"fmt"
	"net/http"
)

type dollars float32
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	// db := database{"shoes": 50, "socks": 5}
	// mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	// log.Fatal(http.ListenAndServe("172.26.88.88:8001", mux))


	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("172.26.88.88:8001", nil))
}