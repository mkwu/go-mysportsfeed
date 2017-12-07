package mysportsfeed

import (
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/dghubble/sling"
)

/*
Cumulative Player Stats https://api.mysportsfeeds.com/v1.1/pull/nba/{season-name}/cumulative_player_stats.{format}
Full Game Schedule
Daily Game Schedule
Daily Player Stats
Scoreboard
Game BoxScore
Game Play-by-Play
Game Starting Lineup
Player Game Logs
Team Game Logs
Roster Players
Active Players
Overall Team Standings
Conference Team Standings
Division Team Standings
Playoff Team Standings
Player Injuries
Current Season
Latest Updates
Daily DFS
*/

const (
	msfUrl         = "https://api.mysportsfeeds.com/v1.1/pull/"
	tagName        = "msf"
	responseFormat = "json"
)

type Client struct {
	sling *sling.Sling
	NBA   *NBAService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, username string, password string) *Client {
	base := sling.New().Client(httpClient).Base(msfUrl)
	base.SetBasicAuth(username, password)
	return &Client{
		sling: base,
		NBA:   newNBAService(base.New()),
	}
}

func buildPath(season string, path string) string {
	return season + "/" + path + "." + responseFormat
}

func Unmarshal(data map[string]interface{}, target interface{}) (err error) {
	// Get the field so we can set the Value
	s := reflect.ValueOf(target).Elem()
	// Get the type so we can extract the map's key.
	t := reflect.TypeOf(target).Elem()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		tag := t.Field(i).Tag.Get(tagName)
		tags := strings.Split(tag, ":")
		section := ""
		name := ""
		if len(tags) == 2 {
			section, name = tags[0], tags[1]
		}
		val := ""
		if section == "stats" {
			// handle stats fields
			if _, ok := data["stats"]; ok {
				stat := data["stats"].(map[string]interface{})
				// make sure the stat we are unmarshaling exits
				if _, ok := stat["name"]; ok {
					// the stat is in a name/value pair named #text
					vals := stat[name].(map[string]interface{})
					if v, ok := vals["#text"]; ok {
						// all returned values are always strings unless it's null
						if v != nil {
							val = v.(string)
						}
					}
				}
			}
		} else if section != "" {
			// handle all non stat fields
			if v, ok := data[section]; ok {
				// all returned values are always strings unless it's null
				if v.(map[string]interface{})[name] != nil {
					val = v.(map[string]interface{})[name].(string)
				}

			}
		}
		// map the value to its proper type
		switch field.Kind() {
		case reflect.Int32:
			x, err := strconv.ParseInt(val, 10, 32)
			if !field.OverflowInt(x) && err == nil {
				field.SetInt(x)
			}
		case reflect.Int64:
			x, err := strconv.ParseInt(val, 10, 64)
			if !field.OverflowInt(x) && err == nil {
				field.SetInt(x)
			}
		case reflect.Float32:
			x, err := strconv.ParseFloat(val, 32)
			if !field.OverflowFloat(x) && err == nil {
				field.SetFloat(x)
			}
		case reflect.Float64:
			x, err := strconv.ParseFloat(val, 64)
			if !field.OverflowFloat(x) && err == nil {
				field.SetFloat(x)
			}
		case reflect.Bool:
			x, err := strconv.ParseBool(val)
			if err == nil {
				field.SetBool(x)
			}
		case reflect.String:
			field.SetString(val)
		default:
			log.Printf("unrecognized type %v, %+v\n", field.Kind(), val)
		}
	}
	return
}
