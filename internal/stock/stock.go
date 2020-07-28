package stock

import (
	"fmt"
	"github.com/labstack/gommon/color"
)

var done = color.Green("[ OK ]")
var fail = color.Red("[Fail]")

type stock struct {
	product map[string]int
}

type stocker interface {
	Check(uidFood string) (status bool)
	ReserveGood(uidFood string) (status bool)
	ReturnGood(uidFood string) (status bool)
}

// Check реализует метод проверки товара на складе
func (s *stock) Check(uidFood string) (status bool) {

	if _, found := s.product[uidFood]; !found {
		fmt.Println(fail, " Товара с такой наменклатурой не найден на складе")
		return false
	}

	if s.product[uidFood] == 0 {
		fmt.Println(fail, " Товара закончился!")
		return false
	}

	fmt.Println(done, " Товар найден на складе")
	return true
}

// ReserveGood реализует метод резервирования товара на складе
func (s *stock) ReserveGood(uidFood string) (status bool) {
	if !s.Check(uidFood) {
		return false
	}

	s.product[uidFood] = s.product[uidFood] - 1 // уменьшаем количество

	fmt.Println(done, " Товар успешно зарезервирован")
	return true
}

// ReturnGood реализует метод возврата товара на склад
func (s *stock) ReturnGood(uidFood string) (status bool) {
	s.product[uidFood] = s.product[uidFood] + 1 // уменьшаем количество

	fmt.Println(done, " Товар успешно возвращен на склад")
	return true
}

// NewStock создает экземпляр типа stock
func NewStock(uids map[string]int) *stock {
	return &stock{
		product: uids,
	}
}
