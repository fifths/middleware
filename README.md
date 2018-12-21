# middleware

## Use
```go
package main

import (
	"fmt"
	"github.com/fifths/middleware"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func testMiddleware(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("testMiddleware")
		handle(w, r, p)
	}
}

func Test(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("success"))
}

func Test1(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			w.Write([]byte("success"))
		}(w, r, p)
	}(w, r, p)
}

func main() {
	router := httprouter.New()
	m := middleware.New()
	m.Use(testMiddleware)
	router.GET("/test", m.Handle(Test))
	router.GET("/test1", Test1)
	log.Fatal(http.ListenAndServe(":8080", router))
}
```