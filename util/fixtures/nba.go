package fixtures

import (
	"github.com/bluele/factory-go/factory"
	"github.com/mkwu/go-mysportsfeed/models"
)

var NBACumulativePlayerStatsFactory = factory.NewFactory(
	&models.NBACumulativePlayerStats{},
)

func MakeNBACumulativePlayerStats(data map[string]interface{}) *models.NBACumulativePlayerStats {
	return NBACumulativePlayerStatsFactory.MustCreateWithOption(data).(*models.NBACumulativePlayerStats)
}
