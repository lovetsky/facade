package account

import (
	"fmt"
	"github.com/labstack/gommon/color"
)

var done = color.Green("[ OK ]")
var fail = color.Red("[Fail]")

type accounter interface {
	Check(uidAccount string) (status bool)
	GetUid() (accountUid string)
}

type account struct {
	uid     string
	balance int
}

// Check реализует метод проверки пользователя
func (a *account) Check(uidAccount string) (status bool) {
	if a.uid != uidAccount {
		fmt.Println(fail, " Аккаунт не найден")
		return false
	}

	fmt.Println(done, " Аккаунт подтвержден")
	return true
}

// GetUid реализует метод возврата идентификатора пользователя
func (a *account) GetUid() (accountUid string) {
	return a.uid
}

// NewAccount создает реализацию интерефейса accounter
func NewAccount(uidAccount string) accounter {
	return &account{
		uid:     uidAccount,
		balance: 100,
	}
}
