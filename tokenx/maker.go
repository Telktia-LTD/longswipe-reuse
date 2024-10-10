package tokenx

import (
	"time"

	"github.com/Telktia-LTD/longswipe-reuse/interfacesx"
)

type Maker interface {
	CreateToken(user interfacesx.UserResponse, duration time.Duration) (string, *AuthPayload, error)

	VerifyToken(token string) (*AuthPayload, error)
}
