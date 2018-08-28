package botlister

import "errors"

// FetchAllBots fetches all of the bots.
func (a *API) FetchAllBots() ([]Bot, error) {
	var ret []Bot
	err := a.doRequest("GET", false, EndpointBots, nil, &ret)
	return ret, err
}

// FetchBot fetches a bot.
func (a *API) FetchBot(id string) (Bot, error) {
	var ret Bot
	if id == "" {
		return ret, errors.New("bot id was not provided")
	}

	err := a.doRequest("GET", false, EndpointBotFetch(id), nil, &ret)
	return ret, err
}

// FetchBotStatistics fetches a bot's statistics.
func (a *API) FetchBotStatistics(id string) (Stats, error) {
	bot, err := a.FetchBot(id)
	return bot.Stats, err
}

// UpdateBotStatistics updates a bot's statistics.
func (a *API) UpdateBotStatistics(id string, stats *Stats) error {
	if id == "" {
		return errors.New("bot id was not provided")
	}

	return a.doRequest("POST", false, EndpointBotStats(id), stats, nil)
}

// ResetBotStatistics resets a bot's statistics.
func (a *API) ResetBotStatistics(id string) error {
	if id == "" {
		return errors.New("bot id was not provided")
	}

	return a.doRequest("DELETE", false, EndpointBotStats(id), nil, nil)
}
