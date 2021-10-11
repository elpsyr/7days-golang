package main

import (
	"fmt"
	"log"
	"net/http"
)
/**
标准库启动Web服务
 */

func main() {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/hello",helloHandler)
	log.Fatalln(http.ListenAndServe(":9999",nil))

}
// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w,"URL.Path =%q\n",req.URL.Path)
}
// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w,"Header[%q] =%q\n",k,v)
	}
}
