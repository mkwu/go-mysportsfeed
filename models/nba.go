package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type NBAPlayer struct {
	Id           int64  `msf:"player:ID"`
	LastName     string `msf:"player:LastName"`
	FirstName    string `msf:"player:FirstName"`
	JerseyNumber string `msf:"player:JerseyNumber"`
	Position     string `msf:"player:Position"`
	Height       string `msf:"player:Height"`
	Weight       string `msf:"player:Weight"`
	BirthDate    string `msf:"player:BirthDate"`
	Age          int32  `msf:"player:Age"`
	BirthCity    string `"msf:"player:BirthCity"`
	BirthCountry string `"msf:"player:BirthCountry"`
	IsRookie     bool   `"msf:"player:IsRookie"`
	IsActive     bool
}

type NBACumulativePlayerStats struct {
	PlayerId              int64   `msf:"player:ID"`
	TeamId                int64   `msf:"team:ID"`
	GamesPlayed           int32   `msf:"stats:GamesPlayed"`
	Fg2PtAtt              int32   `msf:"stats:Fg2PtAtt"`
	Fg2PtAttPerGame       float32 `msf:"stats:Fg2PtAttPerGame"`
	Fg2PtMade             int32   `msf:"stats:Fg2PtMade"`
	Fg2PtMadePerGame      float32 `msf:"stats:Fg2PtMadePerGame"`
	Fg2PtPct              float32 `msf:"stats:Fg2PtPct"`
	Fg3PtAtt              int32   `msf:"stats:Fg3PtAtt"`
	Fg3PtAttPerGame       float32 `msf:"stats:Fg3PtAttPerGame"`
	Fg3PtMade             int32   `msf:"stats:Fg3PtMade"`
	Fg3PtMadePerGame      float32 `msf:"stats:Fg3PtMadePerGame"`
	Fg3PtPct              float32 `msf:"stats:Fg3PtPct"`
	FgAtt                 int32   `msf:"stats:FgAtt"`
	FgAttPerGame          float32 `msf:"stats:FgAttPerGame"`
	FgMade                int32   `msf:"stats:FgMade"`
	FgMadePerGame         float32 `msf:"stats:FgMadePerGame"`
	FgPct                 float32 `msf:"stats:FgPct"`
	FtAtt                 int32   `msf:"stats:FtAtt"`
	FtAttPerGame          float32 `msf:"stats:FtAttPerGame"`
	FtMade                int32   `msf:"stats:FtMade"`
	FtMadePerGame         float32 `msf:"stats:FtMadePerGame"`
	FtPct                 float32 `msf:"stats:FtPct"`
	OffReb                int32   `msf:"stats:OffReb"`
	OffRebPerGame         int32   `msf:"stats:OffRebPerGame"`
	DefReb                int32   `msf:"stats:DefReb"`
	DefRebPerGame         float32 `msf:"stats:DefRebPerGame"`
	Reb                   int32   `msf:"stats:Reb"`
	RebPerGame            float32 `msf:"stats:RebPerGame"`
	Ast                   int32   `msf:"stats:Ast"`
	AstPerGame            float32 `msf:"stats:AstPerGame"`
	Pts                   int32   `msf:"stats:Pts"`
	PtsPerGame            float32 `msf:"stats:PtsPerGame"`
	Tov                   int32   `msf:"stats:Tov"`
	TovPerGame            float32 `msf:"stats:TovPerGame"`
	Stl                   int32   `msf:"stats:Stl"`
	StlPerGame            float32 `msf:"stats:StlPerGame"`
	Blk                   int32   `msf:"stats:Blk"`
	BlkPerGame            float32 `msf:"stats:BlkPerGame"`
	BlkAgainst            int32   `msf:"stats:BlkAgainst"`
	BlkAgainstPerGame     float32 `msf:"stats:BlkAgainstPerGame"`
	Fouls                 int32   `msf:"stats:Fouls"`
	FoulsPerGame          float32 `msf:"stats:FoulsPerGame"`
	FoulsDrawn            int32   `msf:"stats:FoulsDrawn"`
	FoulsDrawnPerGame     float32 `msf:"stats:FoulsDrawnPerGame"`
	FoulPers              int32   `msf:"stats:FoulPers"`
	FoulPersPerGame       float32 `msf:"stats:FoulPersPerGame"`
	FoulPersDrawn         int32   `msf:"stats:FoulPersDrawn"`
	FoulPersDrawnPerGame  float32 `msf:"stats:FoulPersDrawnPerGame"`
	FoulTech              int32   `msf:"stats:FoulTech"`
	FoulTechPerGame       float32 `msf:"stats:FoulTechPerGame"`
	FoulTechDrawn         int32   `msf:"stats:FoulTechDrawn"`
	FoulTechDrawnPerGame  float32 `msf:"stats:FoulTechDrawnPerGame"`
	FoulFlag1             int32   `msf:"stats:FoulFlag1"`
	FoulFlag1PerGame      float32 `msf:"stats:FoulFlag1PerGame"`
	FoulFlag1Drawn        int32   `msf:"stats:FoulFlag1Drawn"`
	FoulFlag1DrawnPerGame float32 `msf:"stats:FoulFlag1DrawnPerGame"`
	FoulFlag2             int32   `msf:"stats:FoulFlag2"`
	FoulFlag2PerGame      float32 `msf:"stats:FoulFlag2PerGame"`
	FoulFlag2Drawn        int32   `msf:"stats:FoulFlag2Drawn"`
	FoulFlag2DrawnPerGame float32 `msf:"stats:FoulFlag2DrawnPerGame"`
	Ejections             int32   `msf:"stats:Ejections"`
	PlusMinus             float32 `msf:"stats:PlusMinus"`
	PlusMinusPerGame      float32 `msf:"stats:PlusMinusPerGame"`
	MinSeconds            int32   `msf:"stats:MinSeconds"`
	MinSecondsPerGame     float32 `msf:"stats:MinSecondsPerGame"`
}

func (s *NBACumulativePlayerStats) MarshalJSON() ([]byte, error) {
	d := map[string]interface{}{
		"player": map[string]interface{}{"ID": 0},
		"team":   map[string]interface{}{"ID": 0},
		"stats":  map[string]interface{}{},
	}
	r := reflect.ValueOf(s).Elem()
	t := reflect.TypeOf(s).Elem()
	for i := 0; i < r.NumField(); i++ {
		var tmp interface{}
		field := r.Field(i)
		tag := t.Field(i).Tag.Get("msf")
		tags := strings.Split(tag, ":")
		section := ""
		name := ""
		if len(tags) == 2 {
			section, name = tags[0], tags[1]
		}
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			tmp = strconv.FormatInt(field.Int(), 10)
		case reflect.Float32:
			tmp = strconv.FormatFloat(field.Float(), 'E', -1, 32)
		case reflect.Float64:
			tmp = strconv.FormatFloat(field.Float(), 'E', -1, 64)
		case reflect.String:
			tmp = field.String()
		}
		fmt.Printf("%s %s\n", section, name)
		if section == "stats" {
			d["stats"].(map[string]interface{})[name] = map[string]interface{}{
				"#text": 0,
			}
			d["stats"].(map[string]interface{})[name].(map[string]interface{})["#text"] = tmp
		} else {
			d[section].(map[string]interface{})[name] = tmp
		}
	}
	fmt.Printf("%+v\n", d)
	return json.Marshal(d)
}
