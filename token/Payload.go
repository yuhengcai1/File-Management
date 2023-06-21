package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID    uuid.UUID `json:id"`
	Admin  bool `json:admin"`
	UserID int32 `json:"userid"`
	IssuedAt time.Time `json:"issued"`
	ExpiredAt time.Time `json:"expired"`
}

func NewPayload(userID int32, admin bool, duration time.Duration) (*Payload,error) {
	takenID, err := uuid.NewRandom()
	if err != nil {
		return nil,err
	}

	payload := &Payload{
		ID: takenID,
		UserID: userID,
		Admin: admin,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload,nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt){
		return errors.New("payload expired")
	}
	return nil
}

