package whatsappx

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type FacebookMessageRequest struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		Body string `json:"body"`
	} `json:"text"`
}

type FacebookMessageResponse struct {
	MessagingProduct string `json:"messaging_product"`
	Contacts         []struct {
		Input string `json:"input"`
		WaID  string `json:"wa_id"`
	} `json:"contacts"`
	Messages []struct {
		ID string `json:"id"`
	} `json:"messages"`
}

func SendWhatsappMessage(recipientID, messageBody, WhatsappAccessToken, WhatsappPhoneNumberID string) (*FacebookMessageResponse, error) {
	client := resty.New()
	accessToken := WhatsappAccessToken
	phoneNumberID := WhatsappPhoneNumberID
	url := fmt.Sprintf("https://graph.facebook.com/v12.0/%s/messages", phoneNumberID)
	request := FacebookMessageRequest{
		MessagingProduct: "whatsapp",
		To:               recipientID,
		Type:             "text",
	}
	request.Text.Body = messageBody

	response := &FacebookMessageResponse{}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(accessToken).
		SetBody(request).
		SetResult(response).
		Post(url)

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error sending message: %s", resp.Status())
	}

	return response, nil
}
