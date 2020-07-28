package messages

import "fmt"

type messager interface {
	SendMessageReserved()
	SendMessageReturned()
}

type message struct{}

// SendMessageReserved реализует метод отобажения сообщения о завершении набора операций резервирования
func (m *message) SendMessageReserved() {
	fmt.Println("---- Операция резервирования завершена ----")
}

// SendMessageReturned реализует метод отобажения сообщения о завершении набора операций возврата товара
func (m *message) SendMessageReturned() {
	fmt.Println("---- Операция возврата завершена ----")
}

// NewMessage создает реализацию интерефейса messager
func NewMessage() messager {
	return &message{}
}
