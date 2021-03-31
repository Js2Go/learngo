package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)

	request.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Printf("Redirect: %v\n", req.URL)
			return nil
		},
	}
	resp, err := client.Do(request)
	//resp, err := http.DefaultClient.Do(request)
	//resp, err := http.Get("https://mazi233.gitee.io/blogs")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("\n%s\n", s)
}
