package uc

type Interactor struct {
	Logger Logger
	US UserStorer
}

func NewInteractor(logger Logger) Interactor {
	return Interactor{
		Logger : logger,
	}
}

type Logger interface {
	Log(...interface{})
}

type UserStorer interface {
	StoreUser(name string) error
}