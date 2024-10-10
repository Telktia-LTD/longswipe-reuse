package validationx

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func IsDuplicateKeyError(err error) bool {
	const uniqueViolation = "23505"
	return strings.Contains(err.Error(), uniqueViolation)
}

func GetDuplicateKey(err error) string {
	logrus.Error("err.Error() ", err.Error())
	parts := strings.Split(err.Error(), "\"")
	if len(parts) >= 2 {
		return parts[1]
	}
	return ""
}
