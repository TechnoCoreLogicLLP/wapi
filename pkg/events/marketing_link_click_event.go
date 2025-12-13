package events

// MarketingLinkClickComponent represents the part of the message clicked
type MarketingLinkClickComponent string

const (
	// MarketingLinkClickComponentCTA indicates a CTA button was clicked
	MarketingLinkClickComponentCTA MarketingLinkClickComponent = "cta"
	// MarketingLinkClickComponentBody indicates a link in the body was clicked
	MarketingLinkClickComponentBody MarketingLinkClickComponent = "body"
)

// MarketingMessagesLinkClickData contains the click tracking details
type MarketingMessagesLinkClickData struct {
	ClickComponent MarketingLinkClickComponent `json:"click_component"`
	ProductId      string                      `json:"product_id,omitempty"`
	ClickId        string                      `json:"click_id,omitempty"`
	TrackingToken  string                      `json:"tracking_token,omitempty"`
}

// MarketingMessagesLinkClickEvent represents a marketing message link click event.
// This event is triggered when a user clicks on a CTA button or body link
// in a marketing message sent via WhatsApp.
type MarketingMessagesLinkClickEvent struct {
	BaseBusinessAccountEvent `json:",inline"`
	PhoneNumber              BusinessPhoneNumber            `json:"phone_number"`
	ClickData                MarketingMessagesLinkClickData `json:"click_data"`
}

// NewMarketingMessagesLinkClickEvent creates a new instance of MarketingMessagesLinkClickEvent.
func NewMarketingMessagesLinkClickEvent(
	baseEvent BaseBusinessAccountEvent,
	phoneNumber BusinessPhoneNumber,
	clickData MarketingMessagesLinkClickData,
) *MarketingMessagesLinkClickEvent {
	return &MarketingMessagesLinkClickEvent{
		BaseBusinessAccountEvent: baseEvent,
		PhoneNumber:              phoneNumber,
		ClickData:                clickData,
	}
}
