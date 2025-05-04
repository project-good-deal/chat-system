package types

type ChatRoomType string

const (
	GROUP_CHAT ChatRoomType = "GROUP_CHAT"
	DIRECT_CHAT ChatRoomType = "DIRECT_CHAT"
)

func GetByString(chatRoomType string) ChatRoomType {
	switch chatRoomType {
	case "GROUP_CHAT":
		return GROUP_CHAT
	case "DIRECT_CHAT":
		return DIRECT_CHAT
	default:
		return GROUP_CHAT
	}
}
