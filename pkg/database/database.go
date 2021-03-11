package database

type Database interface {
	NewClient() error
}