package models

import "assignment/utils"

type User struct {
	Id           int            `gorm:"primary_key"`
	Balance      utils.Decimal2 `gorm:"default:0; not null"`
	Name         string         `gorm:"not null; unique_index"`
	Transactions []*Transaction `gorm:"foreignkey:UserId"`
}
