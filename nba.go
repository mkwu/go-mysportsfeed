package mysportsfeed

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/dghubble/sling"
	"github.com/mkwu/go-mysportsfeed/models"
)

const (
	cumulativePlayerStatsPath = "cumulative_player_stats"
	activePlayersPath         = "active_players"
)

type CumulativePlayerStatsOptions struct {
	Limit    int    `url:"limit"`
	Team     string `url:"team"`
	position string `url:"position"`
}

type ActivePlayersOptions struct {
	Limit    int    `url:"limit"`
	Team     string `url:"team"`
	position string `url:"position"`
}

type NBAService struct {
	sling *sling.Sling
}

// newFollowerService returns a new FollowerService.
func newNBAService(sling *sling.Sling) *NBAService {
	return &NBAService{
		sling: sling.Path("nba/"),
	}
}

func (s *NBAService) CumulativePlayerStats(season string, opts CumulativePlayerStatsOptions) (players []*models.NBACumulativePlayerStats, err error) {
	var results map[string]interface{}
	players = make([]*models.NBACumulativePlayerStats, 0)
	_, err = s.sling.New().Get(buildPath(season, cumulativePlayerStatsPath)).QueryStruct(opts).Receive(&results, err)
	if err != nil {
		return
	}
	if data, ok := results["cumulativeplayerstats"]; ok {
		data := data.(map[string]interface{})
		if _, ok := data["playerstatsentry"]; ok {
			switch reflect.TypeOf(data["playerstatsentry"]).Kind() {
			case reflect.Slice:
				list := reflect.ValueOf(data["playerstatsentry"])
				for i := 0; i < list.Len(); i++ {
					player := new(models.NBACumulativePlayerStats)
					p := list.Index(i).Interface().(map[string]interface{})
					err = Unmarshal(p, player)
					players = append(players, player)
				}
			default:
				err = fmt.Errorf("Expecting array, got %s", reflect.TypeOf(data["playerstatsentry"]).Kind())
			}
		} else {
			err = errors.New("Expecting playerstatsentry in response")
		}
	} else {
		err = errors.New("Expecting cumulativeplayerstats in response")
	}
	return
}

func (s *NBAService) ActivePlayers(season string, opts ActivePlayersOptions) (players []*models.NBAPlayer, err error) {
	var results map[string]interface{}
	players = make([]*models.NBAPlayer, 0)
	_, err = s.sling.New().Get(buildPath(season, activePlayersPath)).QueryStruct(opts).Receive(&results, err)
	if err != nil {
		return
	}
	if data, ok := results["activeplayers"]; ok {
		data := data.(map[string]interface{})
		if _, ok := data["playerentry"]; ok {
			switch reflect.TypeOf(data["playerentry"]).Kind() {
			case reflect.Slice:
				list := reflect.ValueOf(data["playerentry"])
				for i := 0; i < list.Len(); i++ {
					player := new(models.NBAPlayer)
					player.IsActive = true
					p := list.Index(i).Interface().(map[string]interface{})
					err = Unmarshal(p, player)
					players = append(players, player)
				}
			default:
				err = fmt.Errorf("Expecting array, got %s", reflect.TypeOf(data["playerentry"]).Kind())
			}
		} else {
			err = errors.New("Expecting playerentry in response")
		}
	} else {
		err = errors.New("Expecting activeplayers in response")
	}
	return
}
