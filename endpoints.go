package botlister

var (
	// EndpointAPI is the base URL of the API.
	EndpointAPI = "https://discordbotlist.com/api"

	// EndpointBots is the bots endpoint.
	EndpointBots = EndpointAPI + "/bots"
	// EndpointBotFetch returns the endpoint to of a bot.
	EndpointBotFetch = func(botID string) string { return EndpointBots + "/" + botID }
	// EndpointBotStats returns the endpoint of a bot's stats.
	EndpointBotStats = func(botID string) string { return EndpointBotFetch(botID) + "/stats" }
	// EndpointBotUpvote returns the endpoint of a bot's upvotes.
	EndpointBotUpvote = func(botID string) string { return EndpointBotFetch(botID) + "/upvotes" }

	// EndpointUsers is the users endpoint.
	EndpointUsers = EndpointAPI + "/users"
	// EndpointUserFetch returns the endpoint of a user.
	EndpointUserFetch = func(userID string) string { return EndpointUsers + "/" + userID }
)
