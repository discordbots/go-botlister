package botlister

import "errors"

func (a *API) UpvoteBot(id string) error {
	if id == "" {
		return errors.New("bot id was not provided")
	}

	return a.doRequest("POST", true, EndpointBotUpvote(id), nil, nil) // TODO this needs a response
}

func (a *API) DeleteBot(id string) error {
	if id == "" {
		return errors.New("bot id was not provided")
	}

	return a.doRequest("DELETE", true, EndpointBotFetch(id), nil, nil) // TODO this needs a response
}

func (a *API) FetchUser(id string) (User, error) {
	var ret User
	if id == "" {
		return ret, errors.New("user id was not provided")
	}

	err := a.doRequest("GET", false, EndpointUserFetch(id), nil, &ret) // TODO this needs a response
	return ret, err
}

func (a *API) FetchUserBots(id string) ([]Bot, error) {
	user, err := a.FetchUser(id)
	return user.Bots, err
}
