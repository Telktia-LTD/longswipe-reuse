package interfacesx

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type UserRoles string

const (
	AdminRole UserRoles = "ADMIN"
	UserRole  UserRoles = "USER"
)

type TransactionStatus string

const (
	Pending    TransactionStatus = "PENDING"
	Processing TransactionStatus = "PROCESSING"
	Approved   TransactionStatus = "APPROVED"
	Canceled   TransactionStatus = "CANCELED"
	Hold       TransactionStatus = "HOLD"
	Completed  TransactionStatus = "COMPLETED"
	Failed     TransactionStatus = "FAILED"
	Refunded   TransactionStatus = "REFUNDED"
	Reversed   TransactionStatus = "REVERSED"
)

type KycStatus string

type CryptoNetworks string

const (
	Ethereum CryptoNetworks = "ETH"
	Bitcoin  CryptoNetworks = "BTC"
	Binance  CryptoNetworks = "BNB"
	Tron     CryptoNetworks = "TRX"
	Solana   CryptoNetworks = "SOL"
)

const (
	PendingKyc    KycStatus = "PENDING"
	ApprovedKyc   KycStatus = "APPROVED"
	RejectedKyc   KycStatus = "REJECTED"
	ProcessingKyc KycStatus = "PROCESSING"
)

type TransactionType string

const (
	Deposit      TransactionType = "DEPOSIT"
	Withdrawal   TransactionType = "WITHDRAWAL"
	GameRefund   TransactionType = "GAME_REFUND"
	GameCharge   TransactionType = "GAME_CHARGE"
	GameFee      TransactionType = "GAME_FEE"
	GamePlay     TransactionType = "GAME_PLAY"
	GameLoss     TransactionType = "GAME_LOSS"
	GameWin      TransactionType = "GAME_WIN"
	Transfer     TransactionType = "TRANSFER"
	SystemBonus  TransactionType = "SYSTEM_BONUS"
	ClaimVoucher TransactionType = "CLAIM_VOUCHER"
)

type TransactionChargeType string

const (
	Debit  TransactionChargeType = "DEBIT"
	Credit TransactionChargeType = "CREDIT"
)

type RouteDefinition struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

type SuccessResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type AdminLoginResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Auth LoginResponse `json:"auth"`
	} `json:"data"`
}

type LoginResponse struct {
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expiresAt"`
	User      UserResponse `json:"user"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type CreateAdminRequest struct {
	Email      string `json:"email" validate:"omitempty,email"`
	RegChannel string `json:"regChannel" validate:"omitempty"`
	Password   string `json:"password" validate:"required"`
}

type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateUserRequest struct {
	Email         string `json:"email" validate:"omitempty,email"`
	Username      string `json:"username" validate:"omitempty"`
	Phone         string `json:"phone" validate:"omitempty"`
	Surname       string `json:"surname" validate:"omitempty"`
	OtherNames    string `json:"otherNames" validate:"omitempty"`
	RegChannel    string `json:"regChannel" validate:"omitempty"`
	ExternalID    string `json:"externalID" validate:"omitempty"`
	ReferrralCode string `json:"referralCode" validate:"omitempty"`
	Password      string `json:"password" validate:"omitempty"`
}

type UpdateUserRequest struct {
	Email         string `json:"email" validate:"omitempty,email"`
	Username      string `json:"username" validate:"omitempty"`
	Phone         string `json:"phone" validate:"omitempty"`
	Surname       string `json:"surname" validate:"omitempty"`
	OtherNames    string `json:"otherNames" validate:"omitempty"`
	ExternalID    string `json:"externalID" validate:"omitempty"`
	ReferrralCode string `json:"referralCode" validate:"omitempty"`
	UsersPhoto    string `json:"usersPhoto" validate:"omitempty"`
}

type UpdateMobileUserDetails struct {
	Email        string `json:"email" validate:"omitempty,email"`
	Username     string `json:"username" validate:"omitempty"`
	Phone        string `json:"phone" validate:"omitempty"`
	Surname      string `json:"surname" validate:"omitempty"`
	OtherNames   string `json:"otherNames" validate:"omitempty"`
	ProfileImage string `json:"image" validate:"omitempty"`
}

type FetchUserByEmailPhoneOrIdRequest struct {
	EmailPhoneOrID string `json:"emailPhoneOrID" validate:"required"`
}

type UserResponse struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Surname       string    `json:"surname"`
	Othernames    string    `json:"otherNames"`
	RegChannel    string    `json:"regChannel"`
	ExternalID    string    `json:"externalID"`
	Role          UserRoles `json:"role"`
	IsActive      bool      `json:"isActive"`
	EmailVerified bool      `json:"emailVerified"`
	Avatar        string    `json:"avatar"`
	IsPinSet      bool      `json:"isPinSet"`
}

type DataUserResponse struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Surname    string    `json:"surname"`
	Othernames string    `json:"otherNames"`
	Avatar     string    `json:"avatar"`
	Cover      string    `json:"cover"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
}

type SuccessResponseWithPayload struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type CreateWalletRequest struct {
	EmailPhoneOrExternalId string    `json:"emailPhoneOrExternalId" validate:"required"`
	CurrencyID             uuid.UUID `json:"currencyID" validate:"omitempty"`
}

type CreateVirtualAccount struct {
	Email     string  `json:"email"`
	Amount    float64 `json:"amount"`
	Narration string  `json:"narration"`
	TxtRef    string  `json:"tx_ref"`
}

type CreateTransactionRequest struct {
	ReferenceID   string                `json:"referenceId"`
	Amount        float64               `json:"amount"`
	Title         string                `json:"title"`
	UserID        uuid.UUID             `json:"userId"`
	ChargedAmount float64               `json:"chargedAmount"`
	Status        TransactionStatus     `json:"status"`
	ChargeType    TransactionChargeType `json:"chargeType"`
	Type          TransactionType       `json:"type"`
	CurrencyID    uuid.UUID             `json:"currencyID"`
}

type AwardBonusRequest struct {
	EmailPhoneOrID string  `json:"emailPhoneOrID" validate:"required"`
	Amount         float64 `json:"amount" validate:"required"`
}

type WebhookPayload struct {
	Event     string `json:"event"`
	Data      Data   `json:"data"`
	EventType string `json:"event.type"`
}

type Data struct {
	ID                int       `json:"id"`
	TxRef             string    `json:"tx_ref"`
	FlwRef            string    `json:"flw_ref"`
	DeviceFingerprint string    `json:"device_fingerprint"`
	Amount            float64   `json:"amount"`
	Currency          string    `json:"currency"`
	ChargedAmount     float64   `json:"charged_amount"`
	AppFee            float64   `json:"app_fee"`
	MerchantFee       float64   `json:"merchant_fee"`
	ProcessorResponse string    `json:"processor_response"`
	AuthModel         string    `json:"auth_model"`
	IP                string    `json:"ip"`
	Narration         string    `json:"narration"`
	Status            string    `json:"status"`
	PaymentType       string    `json:"payment_type"`
	CreatedAt         time.Time `json:"created_at"`
	AccountID         int       `json:"account_id"`
	Customer          Customer  `json:"customer"`
}

type Customer struct {
	ID          int       `json:"id"`
	FullName    string    `json:"fullname"`
	PhoneNumber *string   `json:"phone_number"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateTransactionStatusRequest struct {
	ChargedAmount float64   `json:"chargedAmount"`
	TransactionID uuid.UUID `json:"transactionId"`
	Status        string    `json:"status"`
}

type UserBalanceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		Balance     float64         `json:"balance"`
		GameBalance float64         `json:"game_balance"`
		Currency    CurrencyDetails `json:"currency"`
	} `json:"data"`
}

type BusinessBalanceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		Balance  float64         `json:"balance"`
		Currency CurrencyDetails `json:"currency"`
	} `json:"data"`
}

type GameHerosRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GameHerosResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Heroes     []GameHeroes   `json:"heroes"`
		Pagination PaginationInfo `json:"pagination"`
	} `json:"data"`
}

type GameHeroes struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type CachedGameHeroes struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Option      string    `json:"option"`
	Description string    `json:"description"`
}

type PaginationInfo struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

type AddFundsToPoolRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}

type GamePlayResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		GamPlayFeedback GamPlayFeedback `json:"gamePlayFeedback"`
	} `json:"data"`
}

type GamPlayFeedback struct {
	AmountPlaced  float64     `json:"amountPlaced"`
	AmountWon     float64     `json:"amountWon"`
	FoundGem      bool        `json:"foundGem"`
	GameCharge    float64     `json:"gameCharge"`
	RefundAmount  float64     `json:"refundValue"`
	WinPercentage float64     `json:"winPercentage"`
	LossCharge    float64     `json:"lossCharge"`
	GameWinHero   GameWinHero `json:"gameWinHero"`
}

type GameWinHero struct {
	Option      string `json:"option"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GamePlanRequest struct {
	MinAmount              float64 `json:"minAmount"`
	MaxAmount              float64 `json:"maxAmount"`
	MinWinPercentage       float64 `json:"minWinPercentage"`
	MaxWinPercentage       float64 `json:"maxWinPercentage"`
	GameChargePercentage   float64 `json:"gameChargePercentage"`
	LossChargePercentage   float64 `json:"lossChargePercentage"`
	RefundChargePercentage float64 `json:"refundChargePercentage"`
	Currency               string  `json:"currency"`
}

type UpdateGamePlanRequest struct {
	ID                     uuid.UUID `json:"id" validate:"required"`
	MinAmount              float64   `json:"minAmount"`
	MaxAmount              float64   `json:"maxAmount"`
	MinWinPercentage       float64   `json:"minWinPercentage"`
	MaxWinPercentage       float64   `json:"maxWinPercentage"`
	GameChargePercentage   float64   `json:"gameChargePercentage"`
	LossChargePercentage   float64   `json:"lossChargePercentage"`
	RefundChargePercentage float64   `json:"refundChargePercentage"`
	Currency               string    `json:"currency" validate:"required"`
}

type GamePlansResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Plans []GamePlans `json:"plans"`
	} `json:"data"`
}

type GamePlans struct {
	ID               uuid.UUID `json:"id"`
	MinAmount        float64   `json:"minAmount"`
	MaxAmount        float64   `json:"maxAmount"`
	MinWinPercentage float64   `json:"minWinPercentage"`
	MaxWinPercentage float64   `json:"maxWinPercentage"`
	GameCharge       float64   `json:"gameCharge"`
	LossCharge       float64   `json:"lossCharge"`
	RefundCharge     float64   `json:"refundCharge"`
	Currency         string    `json:"currency"`
	CreatedAt        time.Time `json:"createdAt"`
}

type Statistics struct {
	TotalUsers            int64   `json:"total_users"`
	TotalHeroesAdded      int64   `json:"total_heroes_added"`
	TotalGameParticipants int64   `json:"total_game_participants"`
	TotalAmountInGameWin  float64 `json:"total_amount_in_game_win"`
	TotalWinPoolBalance   float64 `json:"total_win_pool_balance"`
	TotalWinPoolThreshold float64 `json:"total_win_pool_threshold"`
	TotalNumberOfDeposits int64   `json:"total_number_of_deposits"`
	TotalGameCharges      float64 `json:"total_game_charges"`
	TotalGameFee          float64 `json:"total_game_fee"`
	TotalPendingWithdraws float64 `json:"total_pending_withdraws"`
}

type StatisticsResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Statistics Statistics `json:"statistics"`
	} `json:"data"`
}

type TransactionListResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Transactions []Transactions `json:"transactions"`
		Pagination   PaginationInfo `json:"pagination"`
	} `json:"data"`
}

type VoucherListResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Vouchers   []VoucherResponse `json:"vouchers"`
		Pagination PaginationInfo    `json:"pagination"`
	} `json:"data"`
}

type Transactions struct {
	ID            uuid.UUID         `json:"id"`
	ReferenceID   string            `json:"referenceId"`
	Amount        float64           `json:"amount"`
	Title         string            `json:"title"`
	ChargedAmount float64           `json:"chargedAmount"`
	ChargeType    TransactionType   `json:"chargeType"`
	Type          TransactionType   `json:"type"`
	Status        TransactionStatus `json:"status"`
	Currency      CurrencyDetails   `json:"currency"`
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updatedAt"`
}

type FetchAllUsersResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Users      []UserResponse `json:"users"`
		Pagination PaginationInfo `json:"pagination"`
	} `json:"data"`
}

type CreatePinRequest struct {
	EmailPhoneOrID string `json:"emailPhoneOrID" validate:"omitempty"`
	Pin            string `json:"pin" validate:"required"`
}

type UpdatePinRequest struct {
	EmailPhoneOrID string `json:"emailPhoneOrID" validate:"omitempty"`
	OldPin         string `json:"oldPin" validate:"required"`
	Pin            string `json:"pin" validate:"required"`
}

type AddBankAccountRequest struct {
	EmailPhoneOrID string `json:"emailPhoneOrID" validate:"required"`
	BankCode       string `json:"bankCode" validate:"required"`
	AccountNumber  string `json:"accountNumber" validate:"required"`
	BankName       string `json:"bankName" validate:"optional"`
	AccountName    string `json:"accountName" validate:"required"`
	RecipientCode  string `json:"recipient_code" validate:"required"`
}

type ResolveBankAccountRequest struct {
	EmailPhoneOrID string `json:"emailPhoneOrID" validate:"omitempty"`
	AccountNumber  string `json:"account_number" validate:"required"`
	BankCode       string `json:"bank_code" validate:"required"`
}

type ConfirmBankAccountRequest struct {
	EmailPhoneOrID string `json:"emailPhoneOrID" validate:"required"`
	IsCorrect      bool   `json:"isCorrect" validate:"required"`
}

type BankAccountResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		BankAccount BankAccount `json:"bankAccount"`
	} `json:"data"`
}

type BankAccount struct {
	ID            uuid.UUID `json:"id"`
	BankName      string    `json:"bankName"`
	AccountNumber string    `json:"accountNumber"`
	BankCode      string    `json:"bankCode"`
	AccountName   string    `json:"accountName"`
	IsConfirmed   bool      `json:"isConfirmed"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type BankAccountListResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		BankAccounts []BankAccount  `json:"bankAccounts"`
		Pagination   PaginationInfo `json:"pagination"`
	} `json:"data"`
}

type PlayGameRequest struct {
	EmailPhoneOrID string  `json:"emailPhoneOrID" validate:"required"`
	Amount         float64 `json:"amount" validate:"required,min=1"`
	Option         string  `json:"option" validate:"required"`
	Currency       string  `json:"currency" validate:"required"`
}

type WithdrawRequest struct {
	EmailPhoneOrID string  `json:"emailPhoneOrID" validate:"required"`
	Amount         float64 `json:"amount" validate:"required,min=1"`
	TransactionPin string  `json:"transactionPin" validate:"required"`
}

type TransactionResponse struct {
	ID            uuid.UUID         `json:"id"`
	ReferenceID   string            `json:"referenceId"`
	Amount        float64           `json:"amount"`
	Title         string            `json:"title"`
	ChargedAmount float64           `json:"chargedAmount"`
	ChargeType    TransactionType   `json:"chargeType"` // Adjust type based on your actual model
	Type          TransactionType   `json:"type"`       // Adjust type based on your actual model
	Status        TransactionStatus `json:"status"`     // Adjust type based on your actual model
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updatedAt"`
	UserDetails   UserResponse      `json:"userDetails"`
	BankAccount   BankAccount       `json:"bankAccount"`
	Currency      CurrencyDetails   `json:"currency"`
}

type TransactionDetailsResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Transactions TransactionResponse `json:"transactions"`
	} `json:"data"`
}

type CreateAdminActivityRequest struct {
	AdminID  uuid.UUID `json:"adminID" validate:"required"`
	Activity string    `json:"activity" validate:"required"`
}

type EmailMessage struct {
	Recipient    string                 `json:"recipient"`
	Subject      string                 `json:"subject"`
	TemplateName string                 `json:"template_name"`
	TemplateVars map[string]interface{} `json:"template_vars"`
	Body         string                 `json:"body"`
}

type ReturnLatestGameHeroes struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Winners []LatestWinnersResponse `json:"winners"`
	} `json:"data"`
}

type LatestWinnersResponse struct {
	PseudoName   string    `json:"pseudoName"`
	AmountPlaced float64   `json:"amountPlaced"`
	AmountWon    float64   `json:"amountWon"`
	FoundGem     bool      `json:"foundGem"`
	WonAt        time.Time `json:"wonAt"`
}

type VerifyEmailRequest struct {
	ValidationString string `json:"emailAndCode" validate:"required"`
}

type ActivateEmailOrPhoneRequest struct {
	EmailOrPhone string `json:"emailOrPhone" validate:"required"`
	IsEmail      string `json:"isEmail" validate:"required"`
	Code         string `json:"code" validate:"required"`
}

type CreatePaystackVirtualAccountRequest struct {
	Customer      string `json:"customer" validate:"required"`
	PreferredBank string `json:"preferred_bank" validate:"required"`
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
}

type CreateVoucherRequest struct {
	Amount                    float64   `json:"amount" validate:"required,min=1"`
	CurrencyID                uuid.UUID `json:"currencyID" validate:"required"`
	CreatedForMerchant        bool      `json:"createdForMerchant" validate:"required"`
	IsUtilityVoucher          bool      `json:"isUtilityVoucher" validate:"required"`
	CreatedForExistingUser    bool      `json:"createdForExistingUser" validate:"required"`
	CreatedForNonExistingUser bool      `json:"createdForNonExistingUser" validate:"required"`
	RecipientId               uuid.UUID `json:"recipientId" validate:"omitempty"`
}

type CreateVoucherForBusinessRequest struct {
	Amount             float64   `json:"amount" validate:"required,min=1"`
	CurrencyID         uuid.UUID `json:"currencyID" validate:"required"`
	CreatedForMerchant bool      `json:"createdForMerchant" validate:"required"`
	BusinessID         string    `json:"businessID" validate:"omitempty"`
}
type CreateVoucherForNonUserRequest struct {
	Amount     float64   `json:"amount" validate:"required,min=1"`
	CurrencyID uuid.UUID `json:"currencyID" validate:"required"`
	Email      string    `json:"email" validate:"email,required"`
}
type VoucherResponse struct {
	ID                        uuid.UUID       `json:"id"`
	Amount                    float64         `json:"amount"`
	Balance                   float64         `json:"balance"`
	GeneratedCurrency         CurrencyDetails `json:"generatedCurrency"`
	Code                      string          `json:"code"`
	WasPaidFor                bool            `json:"wasPaidFor"`
	IsUsed                    bool            `json:"isUsed"`
	CreatedAt                 time.Time       `json:"createdAt"`
	CreatedForMerchant        bool            `json:"createdForMerchant"`
	CreatedForExistingUser    bool            `json:"createdForExistingUser"`
	CreatedForNonExistingUser bool            `json:"createdForNonExistingUser"`
	IsLocked                  bool            `json:"isLocked"`
}

type FetchVoucherByIDResponse struct {
	ID                        uuid.UUID                `json:"id"`
	Amount                    float64                  `json:"amount"`
	Balance                   float64                  `json:"balance"`
	GeneratedCurrency         CurrencyDetails          `json:"generatedCurrency"`
	Code                      string                   `json:"code"`
	WasPaidFor                bool                     `json:"wasPaidFor"`
	IsUsed                    bool                     `json:"isUsed"`
	CreatedAt                 time.Time                `json:"createdAt"`
	CreatedForMerchant        bool                     `json:"createdForMerchant"`
	CreatedForExistingUser    bool                     `json:"createdForExistingUser"`
	CreatedForNonExistingUser bool                     `json:"createdForNonExistingUser"`
	IsLocked                  bool                     `json:"isLocked"`
	RedeemedBy                []RedeemedUserDetails    `json:"redeemedBy"`
	RedeemedBusiness          []RedeemdBusinessDetails `json:"redeemedBusiness"`
}

type RedeemedUserDetails struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Surname      string    `json:"surname"`
	OtherNames   string    `json:"otherNames"`
	Amount       float64   `json:"amount"`
	Avatar       string    `json:"avatar"`
	RedeemedDate time.Time `json:"redeemedDate"`
}

type RedeemdBusinessDetails struct {
	ID           uuid.UUID           `json:"id"`
	BusinessName string              `json:"businessName"`
	TradingName  string              `json:"tradingName"`
	Amount       float64             `json:"amount"`
	Logo         string              `json:"logo"`
	RedeemedDate time.Time           `json:"redeemedDate"`
	RedeemedBy   RedeemedUserDetails `json:"redeemedBy"`
}

type LoginRequest struct {
	EmailOrPhone string `json:"emailOrPhone" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type RequestPhoneVerificationRequest struct {
	EmailOrPhone string `json:"emailOrPhone" validate:"required"`
}

type RequestForgotPasswordOTPRequest struct {
	Email string `json:"email" validate:"required"`
}

type RestForgotPasswordRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Code     string `json:"code" validate:"required"`
}

type ResetForgotPinRequest struct {
	Code string `json:"code" validate:"required"`
	Pin  string `json:"pin" validate:"required"`
}

type ResetPasswordRequest struct {
	RecentPassword string `json:"recentPassword" validate:"required"`
	NewPassword    string `json:"newPassword" validate:"required"`
}

type RequestEmailVerificationOTPRequest struct {
	Email string `json:"email" validate:"required"`
}

type AddSupportedCurrencyRequest struct {
	Image        string `json:"image" validate:"required"`
	Currency     string `json:"currency" validate:"required"`
	Symbol       string `json:"symbol" validate:"required"`
	Abbreviation string `json:"abbreviation" validate:"required"`
	IsActive     bool   `json:"isActive" validate:"required"`
	CurrencyType string `json:"currencyType" validate:"required"`
}

type FetchCurrenciesResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Currencies []Currencies `json:"currencies"`
	} `json:"data"`
}

type Currencies struct {
	ID           uuid.UUID `json:"id"`
	Image        string    `json:"image"`
	Currency     string    `json:"currency"`
	Symbol       string    `json:"symbol"`
	Abbreviation string    `json:"abbreviation"`
	IsActive     bool      `json:"isActive"`
	CurrencyType string    `json:"currencyType"`
	CreatedAt    time.Time `json:"createdAt"`
}

type UserAccounts struct {
	UserID      uuid.UUID       `json:"userID"`
	Balance     float64         `json:"balance"`
	GameBalance float64         `json:"gameBalance"`
	Currency    CurrencyDetails `json:"currency"`
}

type CurrencyDetails struct {
	ID           uuid.UUID `json:"id"`
	Image        string    `json:"image"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	Abbrev       string    `json:"abbrev"`
	CurrencyType string    `json:"currencyType"`
	IsActive     bool      `json:"isActive"`
}

type FetchUserAccountsResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		UserAccounts []UserAccounts `json:"userAccounts"`
	} `json:"data"`
}

type AddNewAccountRequest struct {
	CurrencyID uuid.UUID `json:"currencyID" validate:"required"`
}

type GetHomeFeedRquest struct {
	CurrencyID uuid.UUID `json:"currencyID" validate:"required"`
}

type HomeFeedResponse struct {
	Message string   `json:"message"`
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Data    HomeFeed `json:"data"`
}

type HomeFeed struct {
	SelectedAccount     UserAccounts              `json:"selectedAccount"`
	AddedAccounts       []UserAccounts            `json:"addedAccounts"`
	AvailableCurrencies []CurrencyDetails         `json:"availableCurrencies"`
	Transactions        []FeedTransactionResponse `json:"transactions"`
}

type FeedTransactionResponse struct {
	ID            uuid.UUID         `json:"id"`
	ReferenceID   string            `json:"referenceId"`
	Amount        float64           `json:"amount"`
	Title         string            `json:"title"`
	ChargedAmount float64           `json:"chargedAmount"`
	ChargeType    TransactionType   `json:"chargeType"` // Adjust type based on your actual model
	Type          TransactionType   `json:"type"`       // Adjust type based on your actual model
	Status        TransactionStatus `json:"status"`     // Adjust type based on your actual model
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updatedAt"`
	BankAccount   BankAccount       `json:"bankAccount"`
	Currency      CurrencyDetails   `json:"currency"`
}

type UpdateKYCWorkFlowRequest struct {
	WorkFlowID string `json:"workFlowID" validate:"required"`
}

type UsersKYC struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"userID"`
	Status KycStatus `json:"status"`
}

type DeleteUserRequest struct {
	TransactionPin string `json:"transactionPin" validate:"required"`
}

type UserServiceResponse struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Surname       string    `json:"surname"`
	Othernames    string    `json:"otherNames"`
	RegChannel    string    `json:"regChannel"`
	ExternalID    string    `json:"externalID"`
	Role          UserRoles `json:"role"`
	IsActive      bool      `json:"isActive"`
	EmailVerified bool      `json:"emailVerified"`
	Avatar        string    `json:"avatar"`
	IsPinSet      bool      `json:"isPinSet"`
	PaystackCode  string    `json:"paystackCode"`
	PaystackId    string    `json:"paystackId"`
	DeletedAt     time.Time `json:"deletedAt"`
	ReferralCode  string    `json:"referralCode"`
	IsVerified    bool      `json:"isVerified"`
}

type FetchBusinessByResponse struct {
	ID                  uuid.UUID `json:"id"`
	UserID              uuid.UUID `json:"userID"`
	BusinessName        string    `json:"businessName"`
	BusinessDescription string    `json:"businessDescription"`
	BusinessAddress     string    `json:"businessAddress"`
	BusinessEmail       string    `json:"businessEmail"`
	TradingName         string    `json:"tradingName"`
	Logo                string    `json:"logo"`
	BusinessCertificate string    `json:"businessCertificate"`
	MerchantCode        string    `json:"merchantCode"`
}

type FetchCollectorResponse struct {
	ID               uuid.UUID               `json:"id"`
	BusinessID       uuid.UUID               `json:"businessID"`
	CollectorID      uuid.UUID               `json:"CollectorID"`
	MerchantCode     string                  `json:"merchantCode"`
	BusinessOwner    uuid.UUID               `json:"businessOwner"`
	CollectorDetails DataUserResponse        `json:"collectorDetails"`
	BusinessDetails  FetchBusinessByResponse `json:"businessDetails"`
}

type AccountDetails struct {
	BankName      string          `json:"bank"`
	AccountName   string          `json:"account_name"`
	AccountNumber string          `json:"account_number"`
	Currency      CurrencyDetails `json:"currency"`
}

type ForeignAccountDetails struct {
	BankName    string          `json:"bank"`
	AccountName string          `json:"account_name"`
	SwiftCode   string          `json:"swift_code"`
	IBAN        string          `json:"iban"`
	BankAddress string          `json:"bank_address"`
	Currency    CurrencyDetails `json:"currency"`
}

type InternationalAccountDetails struct {
	LocalAccount   AccountDetails        `json:"local_account"`
	ForeignAccount ForeignAccountDetails `json:"foreign_account"`
}

type DepositAddressResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type WithdrawNairaRequest struct {
	Amount         float64 `json:"amount" validate:"required"`
	ReceipientCode string  `json:"receipientCode" validate:"required"`
	TransactionPin string  `json:"transactionPin" validate:"required"`
}

type CryptoDepositAddress struct {
	Address  string          `json:"address"`
	Currency CurrencyDetails `json:"currency"`
}

type CryptoDepositAddressResponse struct {
	Address string `json:"address"`
	Network string `json:"network"`
}

type UpdateOpenBalanceRequest struct {
	Auth              MicroServiceAuth `json:"auth" validate:"required"`
	Amount            float64          `json:"amount" validate:"required"`
	CurrencyID        uuid.UUID        `json:"currencyID" validate:"required"`
	IsDeposit         bool             `json:"isDeposit" validate:"omitempty"`
	IsGameBalance     bool             `json:"isGameBalance" validate:"omitempty"`
	IsBusinessBalance bool             `json:"IsBusinessBalance" validate:"omitempty"`
}

type UpdateOpenBusinessBalanceRequest struct {
	Amount     float64   `json:"amount" validate:"required"`
	CurrencyID uuid.UUID `json:"currencyID" validate:"required"`
	IsDeposit  bool      `json:"isDeposit" validate:"omitempty"`
}
type ServiceUserBalance struct {
	Auth       MicroServiceAuth `json:"auth" validate:"required"`
	CurrencyID uuid.UUID        `json:"currencyID" validate:"required"`
}

type MicroServiceAuth struct {
	AuthToken string              `json:"authToken" validate:"omitempty"`
	User      UserServiceResponse `json:"user" validate:"required"`
}

type CreditGameBalanceRequest struct {
	Amount     float64   `json:"amount" validate:"required"`
	CurrencyID uuid.UUID `json:"currencyID" validate:"required"`
}

type CreditSocialGameBalanceRequest struct {
	Amount         float64   `json:"amount" validate:"required"`
	CurrencyID     uuid.UUID `json:"currencyID" validate:"required"`
	EmailOrPhoneID string    `json:"emailOrPhoneID" validate:"required"`
}

type SendFundsRequest struct {
	Amount                   float64   `json:"amount" validate:"required"`
	RecipientUsernameOrEmail string    `json:"recipientUsernameOrEmail" validate:"required"`
	TransactionPin           string    `json:"transactionPin" validate:"required"`
	CurrencyID               uuid.UUID `json:"currencyID" validate:"required"`
}

type RecipientsResponse struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Status  string         `json:"status"`
	Data    RecipientsData `json:"data"`
}

type RecipientsData struct {
	NairaRecipients    []NairaRecipient    `json:"nairaRecipients"`
	UsernameRecipients []UsernameRecipient `json:"usernameRecipients"`
}

type NairaRecipient struct {
	ID            uuid.UUID `json:"id"`
	RecipientCode string    `json:"recipientCode"`
	BankName      string    `json:"bankName"`
	AccountNumber string    `json:"accountNumber"`
	AccountName   string    `json:"accountName"`
}

type UsernameRecipient struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
}

type CreateSelfVoucherRequest struct {
	Amount     float64   `json:"amount" validate:"required,min=1"`
	CurrencyID uuid.UUID `json:"currencyID" validate:"required"`
}

type CreateMerchantVoucherRequest struct {
	Amount     float64   `json:"amount" validate:"required,min=1"`
	CurrencyID uuid.UUID `json:"currencyID" validate:"required"`
	BusinessID string    `json:"businessID" validate:"required"`
}

type CreateExistingUserVoucherRequest struct {
	Amount      float64   `json:"amount" validate:"required,min=1"`
	CurrencyID  uuid.UUID `json:"currencyID" validate:"required"`
	RecipientID uuid.UUID `json:"recipientID" validate:"required"`
	Email       string    `json:"email" validate:"email,min=1"`
}

type CreateThirdPrtyVoucherRequest struct {
	Amount          float64   `json:"amount" validate:"required,min=1"`
	CurrencyID      uuid.UUID `json:"currencyID" validate:"required"`
	UsernameOrPhone uuid.UUID `json:"usernameOrPhone" validate:"required"`
}

type CreateNonUserVoucherRequest struct {
	Amount     float64   `json:"amount" validate:"required,min=1"`
	CurrencyID uuid.UUID `json:"currencyID" validate:"required"`
	Email      string    `json:"email" validate:"required"`
}

type SetUpAccountDetails struct {
	EmailPhoneOrID string    `json:"emailPhoneOrID" validate:"required"`
	AccountId      uuid.UUID `json:"accountId" validate:"required"`
	IsDefault      bool      `json:"isDefault" validate:"required"`
}

type ClaimVoucherRequest struct {
	Code    string  `json:"code" validate:"required"`
	Amount  float64 `json:"amount" validate:"required"`
	LockPin string  `json:"lockPin" validate:"omitempty"`
}

type FetchRecentMerchantsResponse struct {
	Message string                    `json:"message"`
	Code    int                       `json:"code"`
	Status  string                    `json:"status"`
	Data    []FetchBusinessByResponse `json:"data"`
}

type LockVoucherRequest struct {
	VoucherID uuid.UUID `json:"voucherId" validate:"required"`
	LockPin   string    `json:"lockPin" validate:"required"`
}

type ResetLockVoucherRequest struct {
	VoucherID      uuid.UUID `json:"voucherId" validate:"required"`
	TransactionPin string    `json:"transactionPin" validate:"required"`
	NewLockPin     string    `json:"newLockPin" validate:"required"`
}

type RedeemedVoucherData struct {
	RedeemedUserID     uuid.UUID `json:"redeemedUserID"`
	RedeemedMerchantID uuid.UUID `json:"redeemedMerchantID"`
	VoucherID          uuid.UUID `json:"voucherID"`
	Amount             float64   `json:"amount"`
	IsMerchant         bool      `json:"isMerchant"`
}

type LoginWithFingerprintRequest struct {
	Token string `json:"token" validate:"required"`
}

type LoginWithPinRequest struct {
	Pin   string `json:"pin" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type BusinessAccountRequest struct {
	BusinessName        string `json:"businessName" validate:"required"`
	BusinessDescription string `json:"businessDescription" validate:"required"`
	BusinessAddress     string `json:"businessAddress" validate:"required"`
	BusinessEmail       string `json:"businessEmail" validate:"required,email"`
	TradingName         string `json:"tradingName" validate:"omitempty"`
	Logo                string `json:"logo" validate:"required"`
	BusinessCertificate string `json:"businessCertificate" validate:"omitempty"`
}

type AddCollectorRequest struct {
	CollectorID string `json:"collectorID" validate:"required"`
	BusinessID  string `json:"businessID" validate:"required"`
}

type FetchBusinessResponse struct {
	ID                  uuid.UUID      `json:"id"`
	UserID              uuid.UUID      `json:"userID"`
	BusinessName        string         `json:"businessName"`
	BusinessDescription string         `json:"businessDescription"`
	BusinessAddress     string         `json:"businessAddress"`
	BusinessEmail       string         `json:"businessEmail"`
	TradingName         string         `json:"tradingName"`
	BusinessCertificate string         `json:"businessCertificate"`
	Logo                string         `json:"logo"`
	MerchantCode        string         `json:"merchantCode"`
	Collectors          []UserResponse `json:"collectors"`
}

type JWTTokenStoreRequest struct {
	TelegramID string `json:"telegramID" validate:"required"`
	Token      string `json:"token" validate:"required"`
}

type UpdatePushNotificationTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

type KYCWebhook struct {
	ID             string   `json:"id"`
	URL            string   `json:"url"`
	Enabled        bool     `json:"enabled"`
	Href           string   `json:"href"`
	Token          string   `json:"token"`
	Environments   []string `json:"environments"`
	Events         []string `json:"events"`
	PayloadVersion int      `json:"payload_version"`
}

type KYCWebhooksResponse struct {
	Webhooks []KYCWebhook `json:"webhooks"`
}
