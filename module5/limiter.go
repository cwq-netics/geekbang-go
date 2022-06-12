package limiter

type Limiter interface {
	AllowRequest()
}
