package servicehelpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Telktia-LTD/longswipe-reuse/interfacesx"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type ServiceHelper interface {
	FetchUser(email string) (*interfacesx.UserServiceResponse, error)
	VerifyToken(token string) (*interfacesx.UserServiceResponse, error)
	VerifyTransactionPin(pin string, token string) (*bool, error)
	FetcBusinessAccountBySearch(search string) (*interfacesx.FetchBusinessByResponse, error)
	FetchCollector(code string, collectorID uuid.UUID) (*interfacesx.FetchCollectorResponse, error)
	FetcBusinessAccountByUserID(userID uuid.UUID) (*interfacesx.FetchBusinessByResponse, error)
}

type serviceHelperClient struct {
	baseURL   string
	secretKey string
	client    *http.Client
}

func NewServiceHelperClient(baseURL, secretKey string) ServiceHelper {

	return &serviceHelperClient{
		baseURL:   baseURL,
		secretKey: secretKey,
		client:    &http.Client{},
	}
}

func (p *serviceHelperClient) FetchUser(email string) (*interfacesx.UserServiceResponse, error) {
	url := fmt.Sprintf("open/fetch-user-by-email/" + email)

	res, err := p.makePlanRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response interfacesx.UserServiceResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching user: %v", res.Status)
	}
	return &response, nil
}

func (p *serviceHelperClient) FetcBusinessAccountBySearch(search string) (*interfacesx.FetchBusinessByResponse, error) {
	url := fmt.Sprintf("open/fetch-business/%s", search)
	res, err := p.makePlanRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response interfacesx.FetchBusinessByResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching user: %v", res.Status)
	}
	return &response, nil
}

func (p *serviceHelperClient) FetcBusinessAccountByUserID(userID uuid.UUID) (*interfacesx.FetchBusinessByResponse, error) {
	url := fmt.Sprintf("open/fetch-business-by-id/%s", userID)
	res, err := p.makePlanRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response interfacesx.FetchBusinessByResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching user: %v", res.Status)
	}
	return &response, nil
}

func (p *serviceHelperClient) FetchCollector(code string, collectorID uuid.UUID) (*interfacesx.FetchCollectorResponse, error) {
	url := fmt.Sprintf("open/fetch-collector/%s/%s", code, collectorID)
	res, err := p.makePlanRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response interfacesx.FetchCollectorResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching user: %v", res.Status)
	}
	return &response, nil
}

func (p *serviceHelperClient) VerifyToken(token string) (*interfacesx.UserServiceResponse, error) {
	res, err := p.makeRequest("GET", "microservices/verify-token", nil, token)
	if err != nil {
		return nil, err
	}

	var response interfacesx.UserServiceResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error verifying token: %v", res.Status)
	}

	logrus.Infof("Response: %v", response.Email)
	return &response, nil
}

func (p *serviceHelperClient) VerifyTransactionPin(pin string, token string) (*bool, error) {
	request := map[string]string{
		"transactionPin": pin,
	}

	res, err := p.makeRequest("POST", "microservices/verify-transaction-pin", request, token)
	if err != nil {
		return nil, err
	}

	var response bool
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *serviceHelperClient) makeRequest(method, endpoint string, body interface{}, token string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", p.baseURL, endpoint)
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
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("secret", p.secretKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *serviceHelperClient) makePlanRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", p.baseURL, endpoint)
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
	req.Header.Set("secret", p.secretKey)
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
