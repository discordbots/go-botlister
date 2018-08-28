package botlister

import "errors"

func (a *API) FetchAllBots() ([]Bot, error) {
	var ret []Bot
	err := a.doRequest("GET", false, EndpointBots, nil, &ret)
	return ret, err
}

func (a *API) FetchBot(id string) (Bot, error) {
	var ret Bot
	if id == "" {
		return ret, errors.New("bot id was not provided")
	}

	err := a.doRequest("GET", false, EndpointBotFetch(id), nil, &ret)
	return ret, err
}

func (a *API) FetchBotStatistics(id string) (Stats, error) {
	bot, err := a.FetchBot(id)
	return bot.Stats, err
}

func (a *API) UpdateBotStatistics(id string, stats *Stats) error {
	if id == "" {
		return errors.New("bot id was not provided")
	}

	return a.doRequest("POST", false, EndpointBotStats(id), stats, nil)
}

func (a *API) ResetBotStatstics(id string) error {
	if id == "" {
		return errors.New("bot id was not provided")
	}

	return a.doRequest("DELETE", false, EndpointBotStats(id), nil, nil)
}
