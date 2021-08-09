package model

import (
	"gorm.io/gorm"
)

type Requestjsonmodel struct {
	Model  `bson:"-"`
	MID_   string `bson:"_id" gorm:"-" json:"-" xml:"mid_"`
	Opcode string `bson:"opcode" json:"opcode" xml:"opcode"`
	Data   string `bson:"data" json:"data" xml:"data"`
}

func (Requestjsonmodel) TableName() string {
	return "requestjsonmodel"
}
func (m *Requestjsonmodel) PreloadRequestjsonmodel(db *gorm.DB) *gorm.DB {
	return db
}
