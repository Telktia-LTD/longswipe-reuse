package sms

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioClient struct {
	twilioClient *twilio.RestClient
}

func NewTwilioClient(TwilioAccountSID, TwilioAuthToken string) *TwilioClient {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TwilioAccountSID,
		Password: TwilioAuthToken,
	})
	return &TwilioClient{
		twilioClient: client,
	}
}

func (tc *TwilioClient) SendTextMessage(body, to, TwilioPhoneNumber string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(TwilioPhoneNumber)
	params.SetBody(body)

	_, err := tc.twilioClient.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	return nil
}

func (tc *TwilioClient) SendVoiceMessage(body string, to string) error {
	params := &openapi.CreateCallParams{}
	params.SetTo(to)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetTwiml(fmt.Sprintf("<Response><Say>%s</Say><Say>%s</Say><Say>%s</Say></Response>", body, body, body))

	_, err := tc.twilioClient.Api.CreateCall(params)
	if err != nil {
		return err
	}
	return nil
}

func (tc *TwilioClient) SendWhatsAppMessage(body string, to string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("whatsapp:%s", to))
	params.SetFrom(fmt.Sprintf("whatsapp:%s", os.Getenv("TWILIO_PHONE_NUMBER")))
	params.SetBody(body)

	_, err := tc.twilioClient.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	return nil
}
