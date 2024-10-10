package helperfuncx

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/Telktia-LTD/longswipe-reuse/securityx"
)

func ShuffleOptions(providedOption string) string {
	options := []string{"A", "B", "C", "D"}

	var remainingOptions []string
	for _, opt := range options {
		if opt != providedOption {
			remainingOptions = append(remainingOptions, opt)
		}
	}

	shuffledOptions := shuffle(remainingOptions)
	return shuffledOptions[0]
}

func shuffle(options []string) []string {
	shuffled := make([]string, len(options))
	copy(shuffled, options)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}

func EncryptEmailAndCode(email, code, key string) (string, string, error) {
	encryptedEmail, err := securityx.Encrypt(email, key)
	if err != nil {
		return "", "", err
	}

	encryptedCode, err := securityx.Encrypt(code, key)
	if err != nil {
		return "", "", err
	}

	return encryptedEmail, encryptedCode, nil
}

func GetEncryptedCode(data string) (string, error) {
	parts := strings.Split(data, "|")
	if len(parts) != 2 {
		return "", fmt.Errorf("data does not contain two parts")
	}

	return parts[1], nil
}

func EncryptCode(code, key string) (string, error) {
	encryptedCode, err := securityx.Encrypt(code, key)
	if err != nil {
		return "", err
	}

	return encryptedCode, nil
}

func DecryptCode(data, key string) (string, error) {
	decryptedCode, err := securityx.Decrypt(data, key)
	if err != nil {
		return "", err
	}

	return decryptedCode, nil
}

func DecryptEmailAndCode(data, key string) (string, string, error) {
	parts := strings.Split(data, "|")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("data does not contain two parts")
	}

	decryptedEmail, err := securityx.Decrypt(parts[0], key)
	if err != nil {
		return "", "", err
	}

	decryptedCode, err := securityx.Decrypt(parts[1], key)
	if err != nil {
		return "", "", err
	}

	return decryptedEmail, decryptedCode, nil
}

func GenerateCode() string {
	return fmt.Sprintf("%04d", rand.Intn(10000))
}

func TruncateAndInsert(input string, splitIndex int) string {
	if splitIndex < 0 || splitIndex >= len(input) {
		return input // If splitIndex is out of bounds, return the original string
	}
	part1 := input[:splitIndex]
	part2 := input[splitIndex:]
	return fmt.Sprintf("%s*****%s", part1, part2)
}

func generateReferralCode(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Use base64 URL encoding to get a string without padding characters
	code := base64.URLEncoding.EncodeToString(bytes)
	// Ensure the length of the code is as expected
	if len(code) > length {
		code = code[:length]
	}

	return code, nil
}

func GetUniqueReferralCode(existingCodes map[string]struct{}) (string, error) {
	var code string
	var err error

	for {
		code, err = generateReferralCode(8)
		if err != nil {
			return "", fmt.Errorf("failed to generate referral code: %v", err)
		}

		// Check for uniqueness
		if _, exists := existingCodes[code]; !exists {
			break // Code is unique, break out of the loop
		}
	}

	// Add the newly generated code to the existingCodes map
	existingCodes[code] = struct{}{}
	return code, nil
}

func IsValidNigerianPhoneNumber(phoneNumber string) bool {
	re := regexp.MustCompile(`^(?:\+234|234|0)?[789]\d{9}$`)
	return re.MatchString(phoneNumber)
}

var CountryCurrencyMap = map[string]string{
	"254": "KES", // Kenya
	"91":  "INR", // India
	"972": "ILS", // Israel
	"61":  "AUD", // Australia
	"971": "AED", // United Arab Emirates
	"1":   "USD", // United States/Canada
	"234": "NGN", // Nigeria
	"256": "UGX", // Uganda
	"27":  "ZAR", // South Africa
	"65":  "SGD", // Singapore
	"90":  "TRY", // Turkey
	"47":  "NOK", // Norway
	"44":  "GBP", // United Kingdom
	"49":  "EUR", // Germany (EUR for simplicity)
	"33":  "EUR", // France (EUR for simplicity)
	"34":  "EUR", // Spain (EUR for simplicity)
}

func NormalizePhoneNumber(phoneNumber string) string {
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
	phoneNumber = strings.TrimPrefix(phoneNumber, "+")
	return phoneNumber
}

func GetCountryCode(phoneNumber string) string {
	phoneNumber = NormalizePhoneNumber(phoneNumber)
	if len(phoneNumber) >= 3 {
		return phoneNumber[:3]
	}
	if len(phoneNumber) >= 2 {
		return phoneNumber[:2]
	}
	return ""
}

func GetCurrencyName(phoneNumber string) (string, bool) {
	countryCode := GetCountryCode(phoneNumber)
	currency, exists := CountryCurrencyMap[countryCode]
	return currency, exists
}

func TruncateWithAsterisks(input string, length int) string {
	if len(input) > length {
		return input[:length] + "***"
	}
	return input
}
