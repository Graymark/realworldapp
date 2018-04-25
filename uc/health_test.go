package uc_test

import (
	"testing"

	"github.com/Graymark/realworldapp/uc/mock_uc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("withError", func(t *testing.T) {
		assert.Error(t, withError(mockCtrl))
	})

	t.Run("withoutError", func(t *testing.T) {
		assert.NoError(t, withoutError(mockCtrl))
	})
}

func withError(mockCtrl *gomock.Controller) error {
	i := mock_uc.NewMockedInteractor(mockCtrl)
	i.Logger.EXPECT().Log(gomock.Any()).Return().Times(0)
	return i.GetInteractor().Health()
}

func withoutError(mockCtrl *gomock.Controller) error {
	i := mock_uc.NewMockedInteractor(mockCtrl)
	i.Logger.EXPECT().Log(gomock.Any()).Return().AnyTimes()
	return i.GetInteractor().Health()
}
