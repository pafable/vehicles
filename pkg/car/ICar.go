package car

type ICar interface {
	GetMfr() string
	SetDetail(key string, value string)
	GetSpec(query string) string
}
