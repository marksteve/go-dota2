package dota2

import (
	"encoding/json"
)

type League struct {
	Name          string
	LeagueId      uint
	Description   string
	TournamentUrl string
	ItemDef       uint
}

func UnmarshalLeagues(b []byte) []*League {
	var j interface{}
	err := json.Unmarshal(b, &j)
	if err != nil {
		panic(err)
	}
	leagues := make([]*League, 0)
	for _, v := range j.(map[string]interface{})["result"].(map[string]interface{})["leagues"].([]interface{}) {
		l := v.(map[string]interface{})
		leagues = append(leagues, &League{
			Name:          l["name"].(string),
			LeagueId:      uint(l["leagueid"].(float64)),
			Description:   l["description"].(string),
			TournamentUrl: l["tournament_url"].(string),
			ItemDef:       uint(l["itemdef"].(float64)),
		})
	}
	return leagues
}
