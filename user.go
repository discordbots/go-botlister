package botlister

import "errors"

// UpvoteBot upvotes a bot, requires the UserToken to be set.
func (a *API) UpvoteBot(id string) error {
	if id == "" {
		return errors.New("bot id was not provided")
	}

	return a.doRequest("POST", true, EndpointBotUpvote(id), nil, nil) // TODO this needs a response
}

// DeleteBot deletes a bot, requires the UserToken to be set.
func (a *API) DeleteBot(id string) error {
	if id == "" {
		return errors.New("bot id was not provided")
	}

	return a.doRequest("DELETE", true, EndpointBotFetch(id), nil, nil) // TODO this needs a response
}

// FetchUser fetches a user.
func (a *API) FetchUser(id string) (User, error) {
	var ret User
	if id == "" {
		return ret, errors.New("user id was not provided")
	}

	err := a.doRequest("GET", false, EndpointUserFetch(id), nil, &ret) // TODO this needs a response
	return ret, err
}

// FetchUserBots fetches the bots belonging to a user.
func (a *API) FetchUserBots(id string) ([]Bot, error) {
	user, err := a.FetchUser(id)
	return user.Bots, err
}
