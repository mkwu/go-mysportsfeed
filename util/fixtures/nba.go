package fixtures

import (
	"github.com/bluele/factory-go/factory"
	"github.com/mkwu/go-mysportsfeed/models"
)

var NBAPlayerFactory = factory.NewFactory(
	&models.NBAPlayer{},
)

func MakeNBAPlayer(data map[string]interface{}) *models.NBAPlayer {
	data["IsActive"] = true
	return NBAPlayerFactory.MustCreateWithOption(data).(*models.NBAPlayer)
}

var NBACumulativePlayerStatsFactory = factory.NewFactory(
	&models.NBACumulativePlayerStats{},
)

func MakeNBACumulativePlayerStats(data map[string]interface{}) *models.NBACumulativePlayerStats {
	return NBACumulativePlayerStatsFactory.MustCreateWithOption(data).(*models.NBACumulativePlayerStats)
}
