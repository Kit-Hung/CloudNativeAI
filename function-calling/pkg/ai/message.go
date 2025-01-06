package ai

var MessageStore ChatMessage

func init() {
	MessageStore = make(ChatMessage, 0)
	MessageStore.Clear()
}
