package worker

import (
	"errors"
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests,
			SerializeRequest(req))
	}

	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) (engine.ParseResult, error) {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserialzing "+
				"request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}

	return result, nil
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if profile, ok := p.Args.(map[string]interface{}); ok {
			return parser.NewProfileParser(
				profile["UserName"].(string), profile["Gender"].(string),
				profile["Hokou"].(string), profile["Age"].(string),
				profile["Edu"].(string), profile["Income"].(string),
				profile["Mar"].(string), profile["Height"].(string),
				profile["Avatar"].(string)), nil
		} else {
			return nil, fmt.Errorf("invalid "+
				"args: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}
