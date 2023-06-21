package token

import "time"

type Maker interface {

	CreateToken(userID int32, duration time.Duration, admin bool) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)

}

