package guild

const (
	EnvCommandTopic                    = "COMMAND_TOPIC_GUILD"
	CommandTypeRequestCreate           = "REQUEST_CREATE"
	CommandTypeRequestInvite           = "REQUEST_INVITE"
	CommandTypeRequestDisband          = "REQUEST_DISBAND"
	CommandTypeRequestCapacityIncrease = "REQUEST_CAPACITY_INCREASE"
	CommandTypeCreationAgreement       = "CREATION_AGREEMENT"
	CommandTypeChangeEmblem            = "CHANGE_EMBLEM"
	CommandTypeChangeNotice            = "CHANGE_NOTICE"
	CommandTypeChangeTitles            = "CHANGE_TITLES"
	CommandTypeChangeMemberTitle       = "CHANGE_MEMBER_TITLE"
	CommandTypeLeave                   = "LEAVE"
)

type command[E any] struct {
	CharacterId uint32 `json:"characterId"`
	Type        string `json:"type"`
	Body        E      `json:"body"`
}

type requestCreateBody struct {
	WorldId   byte   `json:"worldId"`
	ChannelId byte   `json:"channelId"`
	MapId     uint32 `json:"mapId"`
	Name      string `json:"name"`
}

type creationAgreementBody struct {
	Agreed bool `json:"agreed"`
}

type changeEmblemBody struct {
	GuildId             uint32 `json:"guildId"`
	Logo                uint16 `json:"logo"`
	LogoColor           byte   `json:"logoColor"`
	LogoBackground      uint16 `json:"logoBackground"`
	LogoBackgroundColor byte   `json:"logoBackgroundColor"`
}

type changeNoticeBody struct {
	GuildId uint32 `json:"guildId"`
	Notice  string `json:"notice"`
}

type leaveBody struct {
	GuildId uint32 `json:"guildId"`
	Force   bool   `json:"force"`
}

type requestInviteBody struct {
	GuildId  uint32 `json:"guildId"`
	TargetId uint32 `json:"targetId"`
}

type changeTitlesBody struct {
	GuildId uint32   `json:"guildId"`
	Titles  []string `json:"titles"`
}

type changeMemberTitleBody struct {
	GuildId  uint32 `json:"guildId"`
	TargetId uint32 `json:"targetId"`
	Title    byte   `json:"title"`
}

type requestDisbandBody struct {
	WorldId   byte `json:"worldId"`
	ChannelId byte `json:"channelId"`
}

type requestCapacityIncreaseBody struct {
	WorldId   byte `json:"worldId"`
	ChannelId byte `json:"channelId"`
}
