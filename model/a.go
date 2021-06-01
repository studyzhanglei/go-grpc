package model

type A struct {
	Name    string      `json:"name" gorm:"comment:用户名"`
	Gender  uint8       `json:"gender"  gorm:"comment:性别"`
	Age     uint8       `json:"age" gorm:"default:0;comment:年龄" `
	Paycode string 		`json:"paycode" gorm:"comment:订单号"`
	ID      uint 		`gorm:"primarykey"`
}

func (A) TableName() string {
	return "a"
}