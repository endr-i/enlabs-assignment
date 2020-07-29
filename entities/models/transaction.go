package models

import (
	"assignment/utils"
	"time"
)

var (
	transactionStates = [2]string{"win", "loss"}
	transactionTypes  = [3]string{"game", "server", "payment"}
)

const (
	TransactionStateUnknown = iota
	TransactionStateWin
	TransactionStateLose
)

const (
	TransactionTypeUnknown = iota
	TransactionTypeGame
	TransactionTypeServer
	TransactionTypePayment
)

const (
	TransactionStatusCreated = iota
	TransactionStatusSuccess
	TransactionStatusFailure
	TransactionStatusCancelled
)

type Transaction struct {
	Id            int            `gorm:"primary_key"`
	State         int            `gorm:"default:0"`
	Type          int            `gorm:"default:0"`
	UserId        int            `gorm:"not null"`
	Status        int            `gorm:"default:0"`
	Amount        utils.Decimal2 `gorm:"default:0"`
	TransactionId string         `gorm:"unique_index"`
	DateCreate    *time.Time     `gorm:"default:CURRENT_TIMESTAMP"`
	User          *User          `gorm:"foreignkey=UserId"`
}

func ConvertTransactionStateStoI(state string) int {
	for i, s := range transactionStates {
		if s == state {
			return i + 1
		}
	}
	return TransactionStateUnknown
}

func ConvertTransactionStateIToS(state int) string {
	if 0 < state && state < len(transactionStates) {
		return transactionStates[state]
	}
	return ""
}

func ConvertTransactionTypeStoI(t string) int {
	for i, s := range transactionTypes {
		if s == t {
			return i + 1
		}
	}
	return TransactionTypeUnknown
}

func ConvertTransactionTypeIToS(t int) string {
	if 0 < t && t < len(transactionStates) {
		return transactionStates[t]
	}
	return ""
}
