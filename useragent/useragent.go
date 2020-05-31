package useragent

// UserAgent an user agent
type UserAgent interface {
	Open(address string) error
	Supports(address string) bool
}
