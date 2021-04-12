package model

import "encoding/json"

type Profile struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Age        string `json:"age"`
	Height     string `json:"height"`
	Weight     string `json:"weight"`
	Income     string `json:"income"`
	Marriage   string `json:"marriage"`
	Education  string `json:"education"`
	Occupation string `json:"occupation"`
	Hokou      string `json:"hokou"`
	Xinzuo     string `json:"xinzuo"`
	House      string `json:"house"`
	Car        string `json:"car"`
	Avatar     string `json:"avatar"`
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, nil
}
