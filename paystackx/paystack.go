package paystackx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Telktia-LTD/longswipe-reuse/interfacesx"

	"github.com/sirupsen/logrus"
)

type PaystackService interface {
	CreateVirtualAccount(customerID *interfacesx.CreatePaystackVirtualAccountRequest) (*VirtualAccountResponse, error)
	CreateUser(data PaystackCreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(data PaystackUpdateUserRequest, pastackCode string) error
	CreateTransferRecipient(data *PaystackCreateTransferRecipientRequest) (*CreateTransferRecipientResponse, error)
	InitiateTransfer(data *TransferFundsRequest) (*TransferOTPResponse, error)
	FetchBalance() (*BalanceResponse, error)
	FetchBanks() (*BanksResponse, error)
	GetBankByPrefix(prefix string) (*BanksResponse, error)
	GetBankNameByCode(bankCode string) (string, error)
	ResolveAccountNumber(account *interfacesx.ResolveBankAccountRequest) (*AccountResponse, error)
}

type paystackClient struct {
	baseURL   string
	secretKey string
	client    *http.Client
}

func NewPaystackClient(baseUrl, secretKey string) PaystackService {

	return &paystackClient{
		baseURL:   baseUrl,
		secretKey: secretKey,
		client:    &http.Client{},
	}
}

func (p *paystackClient) CreateVirtualAccount(customer *interfacesx.CreatePaystackVirtualAccountRequest) (*VirtualAccountResponse, error) {
	response, err := p.makeRequest("POST", "dedicated_account", customer)
	if err != nil {
		return nil, err
	}

	var virtualAccountResponse VirtualAccountResponse

	if err := json.NewDecoder(response.Body).Decode(&virtualAccountResponse); err != nil {
		return nil, err
	}

	return &virtualAccountResponse, nil
}

func (p *paystackClient) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", p.baseURL, endpoint)
	var requestBody []byte
	var err error

	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.secretKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *paystackClient) CreateUser(data PaystackCreateUserRequest) (*CreateUserResponse, error) {
	res, err := p.makeRequest("POST", "customer", data)
	if err != nil {
		return nil, err
	}

	var response CreateUserResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *paystackClient) UpdateUser(data PaystackUpdateUserRequest, pastackCode string) error {
	url := fmt.Sprintf("customer/%s", pastackCode)

	res, err := p.makeRequest("PUT", url, data)
	if err != nil {
		return err
	}

	var response CreateUserResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}

func (p *paystackClient) CreateTransferRecipient(data *PaystackCreateTransferRecipientRequest) (*CreateTransferRecipientResponse, error) {
	res, err := p.makeRequest("POST", "transferrecipient", data)
	if err != nil {
		return nil, err
	}

	var response CreateTransferRecipientResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *paystackClient) InitiateTransfer(data *TransferFundsRequest) (*TransferOTPResponse, error) {
	res, err := p.makeRequest("POST", "transfer", data)
	if err != nil {
		return nil, err
	}

	var response TransferOTPResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(response.Message)
	}

	return &response, nil
}

func (p *paystackClient) FetchBalance() (*BalanceResponse, error) {
	res, err := p.makeRequest("GET", "balance", nil)
	if err != nil {
		return nil, err
	}

	var response BalanceResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *paystackClient) FetchBanks() (*BanksResponse, error) {
	res, err := p.makeRequest("GET", "bank?country=nigeria", nil)
	if err != nil {
		return nil, err
	}

	var response BanksResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *paystackClient) GetBankByPrefix(prefix string) (*BanksResponse, error) {
	banksResponse, err := p.FetchBanks()
	if err != nil {
		return nil, err
	}

	lowerPrefix := strings.ToLower(prefix)
	var bank []Banks
	for _, banks := range banksResponse.Data {
		if strings.HasPrefix(strings.ToLower(banks.Name), lowerPrefix) {
			bank = append(bank, banks)
		}
	}

	logrus.Info(bank)
	return &BanksResponse{
		Message: "Banks fetched successfully",
		Status:  true,
		Data:    bank,
	}, nil
}

func (p *paystackClient) GetBankNameByCode(bankCode string) (string, error) {
	banksResponse, err := p.FetchBanks()
	if err != nil {
		return "", err
	}

	for _, bank := range banksResponse.Data {
		if bank.Code == bankCode {
			return bank.Name, nil
		}
	}

	return "", fmt.Errorf("bank with code %s not found", bankCode)
}

func (p *paystackClient) ResolveAccountNumber(account *interfacesx.ResolveBankAccountRequest) (*AccountResponse, error) {
	res, err := p.makeRequest("GET", fmt.Sprintf("bank/resolve?account_number=%s&bank_code=%s", account.AccountNumber, account.BankCode), nil)
	if err != nil {
		return nil, err
	}

	var response AccountResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
