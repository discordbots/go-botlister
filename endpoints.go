package botlister

var (
	EndpointAPI = "https://discordbotlist.com/api"

	EndpointBots = EndpointAPI + "/bots"
	EndpointBotFetch = func(botID string) string {return EndpointBots + "/" + botID}
	EndpointBotStats = func(botID string) string {return EndpointBotFetch(botID) + "/stats"}
	EndpointBotUpvote = func(botID string) string {return EndpointBotFetch(botID) + "/upvotes"}

	EndpointUsers = EndpointAPI + "/users"
	EndpointUserFetch = func(userID string) string {return EndpointUsers + "/" + userID}
)