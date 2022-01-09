package database

type Config interface {
	ConnString() string
}
