package paystackx

import "time"

type UserVirtualAccountResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Code    int                    `json:"code"`
	Data    UserVirtualAccountData `json:"data"`
}

type UserVirtualAccountData struct {
	BankName      string `json:"bank"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Currency      string `json:"currency"`
}

type VirtualAccountResponse struct {
	Status  bool                `json:"status"`
	Message string              `json:"message"`
	Data    VirtaualAccountData `json:"data"`
}

type VirtaualAccountData struct {
	Bank          Bank       `json:"bank"`
	AccountName   string     `json:"account_name"`
	AccountNumber string     `json:"account_number"`
	Assigned      bool       `json:"assigned"`
	Currency      string     `json:"currency"`
	Metadata      *string    `json:"metadata"` // Using a pointer to handle potential null values
	Active        bool       `json:"active"`
	ID            int        `json:"id"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	Assignment    Assignment `json:"assignment"`
	Customer      Customer   `json:"customer"`
}

type Bank struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Slug string `json:"slug"`
}

type Assignment struct {
	Integration  int    `json:"integration"`
	AssigneeID   int    `json:"assignee_id"`
	AssigneeType string `json:"assignee_type"`
	Expired      bool   `json:"expired"`
	AccountType  string `json:"account_type"`
	AssignedAt   string `json:"assigned_at"`
}

type Customer struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	CustomerCode string `json:"customer_code"`
	Phone        string `json:"phone"`
	RiskAction   string `json:"risk_action"`
}

type PaystackCreateUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type PaystackUpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type CreateUserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Email           string    `json:"email"`
	Integration     int       `json:"integration"`
	Domain          string    `json:"domain"`
	CustomerCode    string    `json:"customer_code"`
	ID              int       `json:"id"`
	Identified      bool      `json:"identified"`
	Identifications *string   `json:"identifications"` // null values in JSON can be represented as pointers
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type UpdateUserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type UpdateUserData struct {
	Integration     int           `json:"integration"`
	FirstName       string        `json:"first_name"`
	LastName        string        `json:"last_name"`
	Email           string        `json:"email"`
	Phone           *string       `json:"phone"`
	Metadata        Metadata      `json:"metadata"`
	Identified      bool          `json:"identified"`
	Identifications *string       `json:"identifications"`
	Domain          string        `json:"domain"`
	CustomerCode    string        `json:"customer_code"`
	ID              int           `json:"id"`
	Transactions    []interface{} `json:"transactions"`
	Subscriptions   []interface{} `json:"subscriptions"`
	Authorizations  []interface{} `json:"authorizations"`
	CreatedAt       time.Time     `json:"createdAt"`
	UpdatedAt       time.Time     `json:"updatedAt"`
}

type Metadata struct {
	Photos []Photo `json:"photos"`
	Email  string  `json:"email"`
	UserId string  `json:"userId"`
}

type Photo struct {
	Type      string `json:"type"`
	TypeID    string `json:"typeId"`
	TypeName  string `json:"typeName"`
	URL       string `json:"url"`
	IsPrimary bool   `json:"isPrimary"`
}

type LogEntry struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

type PaystackEventLog struct {
	TimeSpent      int        `json:"time_spent"`
	Attempts       int        `json:"attempts"`
	Authentication string     `json:"authentication"`
	Errors         int        `json:"errors"`
	Success        bool       `json:"success"`
	Mobile         bool       `json:"mobile"`
	Input          []string   `json:"input"`
	Channel        string     `json:"channel"`
	History        []LogEntry `json:"history"`
}

type PaystackEventCustomer struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	CustomerCode string `json:"customer_code"`
	Phone        string `json:"phone"`
	Metadata     string `json:"metadata"`
	RiskAction   string `json:"risk_action"`
}

type PaystackEventAuthorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	AccountName       string `json:"account_name"`
}

type PaystackEventData struct {
	ID              int                        `json:"id"`
	Domain          string                     `json:"domain"`
	Status          string                     `json:"status"`
	Reference       string                     `json:"reference"`
	Amount          float64                    `json:"amount"`
	Message         string                     `json:"message"`
	GatewayResponse string                     `json:"gateway_response"`
	PaidAt          string                     `json:"paid_at"`
	CreatedAt       string                     `json:"created_at"`
	Channel         string                     `json:"channel"`
	Currency        string                     `json:"currency"`
	IPAddress       string                     `json:"ip_address"`
	Metadata        interface{}                `json:"metadata"`
	Log             PaystackEventLog           `json:"log"`
	Fees            interface{}                `json:"fees"`
	Customer        Customer                   `json:"customer"`
	Authorization   PaystackEventAuthorization `json:"authorization"`
	Plan            interface{}                `json:"plan"`
}

type PaystackEventPayload struct {
	Event string            `json:"event"`
	Data  PaystackEventData `json:"data"`
}

type PaystackCreateTransferRecipientRequest struct {
	Type          string   `json:"type"`
	Name          string   `json:"name"`
	AccountNumber string   `json:"account_number"`
	BankCode      string   `json:"bank_code"`
	Currency      string   `json:"currency"`
	Metadata      Metadata `json:"metadata"`
}

type CreateTransferRecipientResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Active        bool   `json:"active"`
		CreatedAt     string `json:"createdAt"`
		Currency      string `json:"currency"`
		Domain        string `json:"domain"`
		ID            int    `json:"id"`
		Integration   int    `json:"integration"`
		Name          string `json:"name"`
		RecipientCode string `json:"recipient_code"`
		Type          string `json:"type"`
		UpdatedAt     string `json:"updatedAt"`
		IsDeleted     bool   `json:"is_deleted"`
		Details       struct {
			AuthorizationCode interface{} `json:"authorization_code"`
			AccountNumber     string      `json:"account_number"`
			AccountName       string      `json:"account_name"`
			BankCode          string      `json:"bank_code"`
			BankName          string      `json:"bank_name"`
		} `json:"details"`
	} `json:"data"`
}

type TransferFundsRequest struct {
	Source    string  `json:"source"`
	Reason    string  `json:"reason"`
	Amount    float64 `json:"amount"`
	Recipient string  `json:"recipient"`
	Reference string  `json:"reference"`
	Currency  string  `json:"currency"`
}

type TransferOTPResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Integration  int    `json:"integration"`
		Domain       string `json:"domain"`
		Amount       int    `json:"amount"`
		Currency     string `json:"currency"`
		Source       string `json:"source"`
		Reason       string `json:"reason"`
		Recipient    int    `json:"recipient"`
		Status       string `json:"status"`
		TransferCode string `json:"transfer_code"`
		ID           int    `json:"id"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
	} `json:"data"`
}

type BalanceResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		Currency string  `json:"currency"`
		Balance  float64 `json:"balance"`
	} `json:"data"`
}

type Banks struct {
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Code        string    `json:"code"`
	Longcode    string    `json:"longcode"`
	Gateway     *string   `json:"gateway"`
	PayWithBank bool      `json:"pay_with_bank"`
	Active      bool      `json:"active"`
	IsDeleted   bool      `json:"is_deleted"`
	Country     string    `json:"country"`
	Currency    string    `json:"currency"`
	Type        string    `json:"type"`
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Meta struct {
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	PerPage  int     `json:"perPage"`
}

type BanksResponse struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
	Data    []Banks `json:"data"`
}

type AccountData struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankName      string `json:"bank_name"`
	RecipientCode string `json:"recipient_code"`
}

type AccountResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    AccountData `json:"data"`
}
