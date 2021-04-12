package controller

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"learngo/crawler/engine"
	"learngo/crawler/frontend/model"
	"learngo/crawler/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// TODO
// support paging
// add start page
type SearchResult interface {
	GetSearchResult(q string, from, size int) (model.SearchResult, error)
}

type SearchResultNoViewHandler struct {
	client *elastic.Client
}

func (h SearchResultNoViewHandler) GetSearchResult(
	q string, from, size int) (model.SearchRes, error) {
	var result model.SearchRes

	search := h.client.
		Search("dating_profile")
	
	if q != "" {
		search = search.
			Query(elastic.NewQueryStringQuery(
				rewriteQueryString(q)))
	}
	
	resp, err := search.From(from).
		Size(size).
		Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Count = resp.TotalHits()

	for _, v := range resp.Each(reflect.TypeOf(engine.Item{})) {
		item := v.(engine.Item)
		result.Data = append(result.Data, item)
	}

	return result, nil
}

func CreateSearchResultNoViewHandler() SearchResultNoViewHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultNoViewHandler{
		client: client,
	}
}

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(
		req.FormValue("from"))

	if err != nil {
		from = 0
	}
	page, err := h.GetSearchResult(q, from, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) GetSearchResult(
	q string, from, _ int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(
			rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from

	for _, v := range resp.Each(reflect.TypeOf(engine.Item{})) {
		item := v.(engine.Item)
		result.Items = append(result.Items, item)
	}

	result.PrevFrom =
		result.Start - len(result.Items)
	result.NextFrom =
		result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
