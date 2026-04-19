package main

import (
	"fmt"
	"homework5/cmd/internal/commission"
	"homework5/cmd/internal/receipt"
)

func main() {
	var aName, aSurname string
	var bName, bSurname string
	var card string
	var amount int
	var cardType int

	fmt.Print("Имя отправиьеля: ")
	fmt.Scan(&aName)
	fmt.Print("Фамилия отправителя: ")
	fmt.Scan(&aSurname)

	fmt.Print("Имя получателя: ")
	fmt.Scan(&bName)
	fmt.Print("амилия получателя: ")
	fmt.Scan(&bSurname)

	fmt.Print("Номер карты получателя: ")
	fmt.Scan(&card)

	fmt.Print("Ведите сумму: ")
	fmt.Scan(&amount)

	fmt.Print("Alif карта? (1-да/0-нет): ")
	fmt.Scan(&cardType)

	isAlif := cardType == 1

	comm, err := commission.Calculate(amount, isAlif)
	if err != nil {
		fmt.Println("Ощибка: ", err)
		return
	}

	total := amount + comm

	txId := receipt.TransactionId()
	masked := receipt.Card(card)
	time := receipt.CurrentTime()

	fmt.Println("\n============== Электронный чек ==============")
	fmt.Printf("Отправитель:      %s\n", receipt.FullName(aName, aSurname))
	fmt.Printf("Получатель:       %s\n", receipt.FullName(bName, bSurname))
	fmt.Printf("Номер транзакции: %d\n", txId)
	fmt.Printf("Счёт зачисления:  %s\n", masked)
	fmt.Printf("Дата и время:     %s\n", time)
	fmt.Printf("Сумма:            %d сум\n", receipt.FormatedAmount(amount))
	fmt.Printf("Комиссия:         %d сум\n", receipt.FormatedAmount(comm))
	fmt.Printf("Итого:            %d сум\n", receipt.FormatedAmount(total))
	fmt.Println("AO ALIF TECH · Лицензия ЦБ РУз № 000010")
	fmt.Println("Статус:           Исполнено")
	fmt.Println("================================================")
}
