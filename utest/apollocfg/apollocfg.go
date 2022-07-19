package apollocfg

//go:generate mockgen -source=$GOFILE -destination=../mock/${GOPACKAGE}_mock.go -package=mock
type IApolloCfg interface {
	GetString(key string) string
}