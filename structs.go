package botlister

import "time"

type User struct {
	Username      string    `json:"username"`
	Discriminator string    `json:"discriminator"`
	Avatar        string    `json:"avatar"`
	Admin         bool      `json:"admin"`
	Banned        bool      `json:"banned"`
	CreatedAt     time.Time `json:"created_at"`
	Bots          []Bot     `json:"bots"`
}

type Stats struct {
	Online           bool `json:"online"`
	ShardID          uint `json:"shard_id"`
	Guilds           uint `json:"guilds"`
	Users            uint `json:"users"`
	VoiceConnections uint `json:"voice_connections"`
}

type Bot struct {
	ID               string    `json:"id"`
	Username         string    `json:"username"`
	Discriminator    string    `json:"discriminator"`
	Avatar           string    `json:"avatar"`
	ShortDescription string    `json:"short_description"`
	LongDescription  string    `json:"Long_description"`
	Prefix           string    `json:"prefix"`
	Website          string    `json:"website"`
	BotInvite        string    `json:"bot_invite"`
	ServerInvite     string    `json:"server_invite"`
	Verified         bool      `json:"verified"`
	Upvotes          uint      `json:"upvotes"`
	IsUpvoted        bool      `json:"is_upvoted"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Owner            User      `json:"owner"`
	Stats            Stats     `json:"stats"`
}
