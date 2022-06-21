package database

import (
	"testing"
)

func TestConnect(t *testing.T) {
	err := ConnectDB()
	if err != nil {
		panic(err)
	}
}
