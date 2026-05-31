package events

// MessageDeliveredEvent represents an event related to an undelivered message.
type MessageDeliveredEvent struct {
	BaseSystemEvent `json:",inline"`
	MessageId       string `json:"messageId"`
	SentTo          string `json:"sentTo"`
	SentToUserId    string `json:"sentToUserId,omitempty"` // Business-scoped user ID (BSUID) of the recipient.
}

// MessageDeliveredEvent creates a new instance of MessageUndeliveredEvent.
func NewMessageDeliveredEvent(baseSystemEvent BaseSystemEvent, messageId, sendTo, sendToUserId string) *MessageDeliveredEvent {
	return &MessageDeliveredEvent{
		BaseSystemEvent: baseSystemEvent,
		MessageId:       messageId,
		SentTo:          sendTo,
		SentToUserId:    sendToUserId,
	}
}
