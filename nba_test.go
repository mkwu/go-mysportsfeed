package mysportsfeed

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/mkwu/go-mysportsfeed/models"
	"github.com/mkwu/go-mysportsfeed/util/fixtures"
	"github.com/stretchr/testify/assert"
)

var data = map[string]interface{}{
	"PlayerId": int64(9158),
	"TeamId":   int64(86),
}

func TestActivePlayers(t *testing.T) {
	players := make([]*models.NBAPlayer, 0)
	players = append(players, fixtures.MakeNBAPlayer(data))
	rawData := map[string]interface{}{
		"activeplayers": map[string]interface{}{
			"playerentry": players,
		},
	}
	httpClient, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc(buildPath("/v1.1/pull/nba/2017-2018-regular", activePlayersPath), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		jsonStr, err := json.Marshal(rawData)
		if err != nil {
			fmt.Fprintf(w, "Error")
		} else {
			fmt.Fprintf(w, string(jsonStr))
		}
	})
	client := NewClient(httpClient, "test", "test")
	opts := ActivePlayersOptions{
		Limit: 1,
	}
	data, err := client.NBA.ActivePlayers("2017-2018-regular", opts)
	assert.Equal(t, players, data)
	assert.Nil(t, err)
}

func TestCumulativePlayerStats(t *testing.T) {
	players := make([]*models.NBACumulativePlayerStats, 0)
	players = append(players, fixtures.MakeNBACumulativePlayerStats(data))
	cumulativeRawData := map[string]interface{}{
		"cumulativeplayerstats": map[string]interface{}{
			"playerstatsentry": players,
		},
	}
	httpClient, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc(buildPath("/v1.1/pull/nba/2017-2018-regular", cumulativePlayerStatsPath), func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		jsonStr, err := json.Marshal(cumulativeRawData)
		if err != nil {
			fmt.Fprintf(w, "Error")
		} else {
			fmt.Fprintf(w, string(jsonStr))
		}
	})
	client := NewClient(httpClient, "test", "test")

	opts := CumulativePlayerStatsOptions{
		Limit: 1,
	}
	data, err := client.NBA.CumulativePlayerStats("2017-2018-regular", opts)
	assert.Equal(t, players, data)
	assert.Nil(t, err)
}
