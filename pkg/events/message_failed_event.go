package events

type MessageFailedEvent struct {
	BaseSystemEvent `json:",inline"`
	MessageId       string `json:"messageId"`
	SentTo          string `json:"sentTo"`
	SentToUserId    string `json:"sentToUserId,omitempty"` // Business-scoped user ID (BSUID) of the recipient.
	FailReason      string `json:"failReason"`
	ErrorCode       int    `json:"errorCode"`
	ErrorMessage    string `json:"errorMessage"`
}

func NewMessageFailedEvent(baseSystemEvent BaseSystemEvent, messageId, sendTo, sendToUserId, failReason string, errCode int, errorMessage string) *MessageFailedEvent {
	return &MessageFailedEvent{
		BaseSystemEvent: baseSystemEvent,
		MessageId:       messageId,
		SentTo:          sendTo,
		SentToUserId:    sendToUserId,
		FailReason:      failReason,
		ErrorCode:       errCode,
		ErrorMessage:    errorMessage,
	}

}
