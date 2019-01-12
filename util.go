package core

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

// ToUUID - Converts a string representation of an uuid into a uuid.ToUUID
// Returns a "00000000-0000-0000-0000-000000000000" uuid.UUID and false if it cannot parse the string.
func ToUUID(idStr string) (uuid.UUID, bool) {
	u, err := uuid.FromString(idStr)
	if err != nil {
		u, _ = uuid.FromString("00000000-0000-0000-0000-000000000000")
		return u, false
	}
	return u, true
}

// SampleAccounts - Returns a mock slice of rezerw.Account
func SampleAccounts() []Account {
	accounts := make([]Account, 0)
	if u1, ok := ToUUID("568aabd8-c431-485f-845a-c447083ab287"); ok {
		a1 := Account{ID: u1, Name: "Account1"}
		accounts = append(accounts, a1)
	}
	if u2, ok := ToUUID("338681c2-fb4b-4448-957a-297729eab4a8"); ok {
		a2 := Account{ID: u2, Name: "Account2"}
		accounts = append(accounts, a2)
	}
	return accounts
}

// FindAccountByIDString - Returns an account from mock slice of rezerw.Account
// by its id represented as string.
func FindAccountByIDString(id string) (*Account, error) {
	uid, ok := ToUUID(id)
	if ok {
		return FindAccount(uid)
	}
	return nil, errors.New("Not a valid ID")
}

// FindAccount - Returns an account from mock slice of rezerw.Account
func FindAccount(uid uuid.UUID) (*Account, error) {
	accounts := SampleAccounts()
	for _, a := range accounts {
		if a.ID == uid {
			return &a, nil
		}
	}
	return nil, errors.New("Account not found")
}
