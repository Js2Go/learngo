package main

import (
	"fmt"
	"learngo/lang/retriever/mock"
	real2 "learngo/lang/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func post(poster Poster) {
	poster.Post("https://www.imooc.com", map[string]string{
		"name": "mazi",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout: time.Minute,
	}
	inspect(r)

	// Type assertion
	//realRetriever := r.(*real2.Retriever)
	//fmt.Println(realRetriever.Timeout)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever)  {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println()
}
