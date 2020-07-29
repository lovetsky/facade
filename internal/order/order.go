package order

type accounter interface {
	Check(uidAccount string) (status bool)
	GetUid() (accountUid string)
}

type stocker interface {
	Check(uidFood string) (status bool)
	ReserveGood(uidFood string) (status bool)
	ReturnGood(uidFood string) (status bool)
}

type messager interface {
	SendMessageReturned()
	SendMessageReserved()
}

type orderer interface {
	Reserved(uidGood string) (status bool)
	Returned(uidGood string) (status bool)
}

type orderFacade struct {
	stock   stocker
	message messager
	account accounter
}

// Reserved реализует метод резервирования товара по идентификатору
func (o *orderFacade) Reserved(uidGood string) (status bool) {
	if !o.account.Check(o.account.GetUid()) {
		return false
	}

	o.stock.ReserveGood(uidGood)
	o.message.SendMessageReserved()

	return true
}

// Returned реализует метод возврата товара по идентификатору
func (o *orderFacade) Returned(uidGood string) (status bool) {
	o.stock.ReturnGood(uidGood)
	o.message.SendMessageReturned()

	return true
}

// NewOrderFacade создает экземпляр типа orderFacade
func NewOrderFacade(account accounter, message messager, stock stocker) *orderFacade {
	return &orderFacade{
		stock:   stock,
		account: account,
		message: message,
	}
}
