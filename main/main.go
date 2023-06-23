package main

import (
	"io/ioutil"
	"net/http"
)

func main()  {
	resp,err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	data,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	 http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(data)
		 return
	 })
	 http.ListenAndServe(":80",nil)
}
