package mock_uc

import (
	"github.com/Graymark/realworldapp/uc"
	"github.com/golang/mock/gomock"
)

type MockedInteractor struct {
	Logger *MockLogger
	US     *MockUserStorer
}

func NewMockedInteractor(mockCtrl *gomock.Controller) MockedInteractor {
	return MockedInteractor{
		Logger: NewMockLogger(mockCtrl),
		US:     NewMockUserStorer(mockCtrl),
	}
}

func (i MockedInteractor) GetInteractor() uc.Interactor {
	return uc.Interactor{
		Logger: i.Logger,
		US:     i.US,
	}
}
