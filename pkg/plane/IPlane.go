package plane

type IPlane interface {
	GetMfr() string
	SetDetail(key string, value string)
	GetSpec(query string) string
	GetMach() float32
}
