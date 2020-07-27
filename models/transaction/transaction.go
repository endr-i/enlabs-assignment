package transaction

import "assignment/utils"

type Transaction struct {
	Id       int
	State    int
	Type     int
	UserId   int
	Amount   utils.Decimal2
	ActionId string
}
