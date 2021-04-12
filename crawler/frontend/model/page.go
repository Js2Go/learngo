package model

import "learngo/crawler/engine"

type SearchResult struct {
	Hits     int64
	Start    int
	PrevFrom int
	NextFrom int
	Query    string
	Items    []engine.Item
}

type SearchRes struct {
	Count int64         `json:"count"`
	Data  []engine.Item `json:"data"`
}

type Req struct {
	Query  string `form:"q"`
	Offset int    `form:"offset"`
	Size   int    `form:"size"`
}
