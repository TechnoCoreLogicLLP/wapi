package events

// MessageSentEvent represents an event indicating that a message has been sent.
type MessageSentEvent struct {
	BaseSystemEvent `json:",inline"`
	MessageId       string `json:"messageId"`
	SentTo          string `json:"sentTo"`
	SentToUserId    string `json:"sentToUserId,omitempty"` // Business-scoped user ID (BSUID) of the recipient.
}

// NewMessageSentEvent creates a new instance of MessageSentEvent.
func NewMessageSentEvent(baseSystemEvent BaseSystemEvent, messageId, sendTo, sendToUserId string) *MessageSentEvent {
	return &MessageSentEvent{
		BaseSystemEvent: baseSystemEvent,
		MessageId:       messageId,
		SentTo:          sendTo,
		SentToUserId:    sendToUserId,
	}
}
