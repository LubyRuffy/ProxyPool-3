package database

import (
	"testing"
	"fmt"
)

func TestGetSession(t *testing.T) {
	session := GetSession()
	fmt.Println(*session)
}

