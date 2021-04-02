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
