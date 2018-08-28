package botlister

import (
	"testing"
)

func TestAPI_Bot(t *testing.T) {

	api := NewAPI("", "")

	t.Run("TestFetchBot", func(t *testing.T) {
		bot, err := api.FetchBot("159985870458322944")
		if err != nil {
			t.Error(err)
			return
		}

		if bot.Username != "MEE6" {
			t.Error("The bot loaded isn't MEE6")
		}
	})

	t.Run("TestFetchBotStatistics", func(t *testing.T) {
		stats, err := api.FetchBotStatistics("159985870458322944")
		if err != nil {
			t.Error(err)
			return
		}

		if !stats.Online {
			t.Error("MEE6 is down?")
		}
	})

	t.Run("TestFetchAllBots", func(t *testing.T) {
		bots, err := api.FetchAllBots()
		if err != nil {
			t.Error(err)
			return
		}

		if len(bots) < 5 {
			t.Error("There are less than 5 bots?")
		}
	})

	t.Run("TestUpdateBotStatistics", func(t *testing.T) {
		err := api.UpdateBotStatistics("416717866411360258", &Stats{
			Online:           true,
			Guilds:           420,
			ShardID:          5,
			Users:            4,
			VoiceConnections: 8,
		})

		if err != nil {
			t.Error(err)
			return
		}
	})

	t.Run("TestResetBotStatistics", func(t *testing.T) {
		err := api.ResetBotStatstics("416717866411360258")

		if err != nil {
			t.Error(err)
			return
		}
	})
}
