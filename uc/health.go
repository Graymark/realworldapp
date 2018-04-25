package uc

import (
	"errors"
	"math/rand"

	"github.com/Graymark/realworldapp/domain"
)

func (i Interactor) Health() error {
	i.Logger.Log(domain.NewError(domain.ErrDebug, domain.ErrLogical, "Health", "has been called"))

	if rand.Intn(2) == 0 {
		return errors.New("hey an error")
	}

	return nil
}
