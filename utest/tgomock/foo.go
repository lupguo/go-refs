package tgomock

//go:generate mockgen -source=$GOFILE -destination=foo_mock.go -package=tgomock
type IApolloConfig interface {
	GetConfig(key string) bool
}

type Apollo struct {
	ID int
}

func (a *Apollo) GetConfig(string) bool {
	if a.ID % 2 == 0 {
		return true
	}
	return false
}

func OrderHandle(x string, a IApolloConfig) string {
	if ok := a.GetConfig("key"); ok {
		return "ok," + x
	}

	return "nok," + x
}
