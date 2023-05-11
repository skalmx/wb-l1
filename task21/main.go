package main

import (
	"fmt"
	"strconv"
)

// Смысл работы этого паттерна в том, что если у вас есть класс и его интерфейс не совместим с кодом вашей системы, то что бы разрешить этот конфликт
// мы не изменяем код этого класса, а пишем для него адаптер. 
//Другими словами Adapter адаптирует существующий код к требуемому интерфейсу (является переходником).
// В данном случае можем использовать числовые данные как строковые если нам это будет нужно

type sendStringData interface {
	sendString()
}


type stringData struct {
	data string
}


func (s *stringData) SendString() {
	fmt.Println("print data from stringData struct method: ",s.data)
}


type intData struct {
	data int
}

type intDataAdapter struct {
	integerData *intData
}


func (adapter *intDataAdapter) SendIntAsString() {
	dataToSend := strconv.Itoa(adapter.integerData.data)
	fmt.Println("print data from adapter struct method: ", dataToSend)
}

func main() {
	strData := &stringData{data: "text data"}
	strData.SendString()

	intData := &intData{data: 12345}

	adapter := &intDataAdapter{intData}
	adapter.SendIntAsString()
}