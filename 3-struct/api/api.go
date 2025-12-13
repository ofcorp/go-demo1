package api

type Api struct {
	Key string
}

func NewApi(key string) *Api {
	return &Api{Key: key}
}