// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"time"
)

const TableNameSaleorder = "saleorder"

// Saleorder mapped from table <saleorder>
type Saleorder struct {
	Saleorderid   string    `gorm:"column:saleorderid" json:"saleorderid"`
	Docno         string    `gorm:"column:docno" json:"docno"`
	Possessionid  string    `gorm:"column:possessionid" json:"possessionid"`
	Beforevatsale float64   `gorm:"column:beforevatsale" json:"beforevatsale"`
	Totaldiscount float64   `gorm:"column:totaldiscount;not null" json:"totaldiscount"`
	Vatsale       float64   `gorm:"column:vatsale" json:"vatsale"`
	Totalsale     float64   `gorm:"column:totalsale" json:"totalsale"`
	Posclientid   string    `gorm:"column:posclientid" json:"posclientid"`
	Branchid      string    `gorm:"column:branchid" json:"branchid"`
	Merchantid    string    `gorm:"column:merchantid" json:"merchantid"`
	Memberid      string    `gorm:"column:memberid" json:"memberid"`
	Status        int16     `gorm:"column:status;not null" json:"status"`
	Createby      string    `gorm:"column:createby" json:"createby"`
	Createdate    time.Time `gorm:"column:createdate;not null" json:"createdate"`
	Isactive      int16     `gorm:"column:isactive;not null" json:"isactive"`
	Voidtype      int16     `gorm:"column:voidtype" json:"voidtype"`
	Voidby        string    `gorm:"column:voidby" json:"voidby"`
	Voiddate      time.Time `gorm:"column:voiddate" json:"voiddate"`
}

// TableName Saleorder's table name
func (*Saleorder) TableName() string {
	return TableNameSaleorder
}
