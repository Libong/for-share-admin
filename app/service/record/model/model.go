package model

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	RecordId  string `gorm:"column:record_id" json:"record_id"`   //记录编号
	GoodName  string `gorm:"column:good_name" json:"good_name"`   //物品名称
	Status    uint32 `gorm:"column:status" json:"status"`         //记录状态 1使用中 2历史
	BuyAt     int64  `gorm:"column:buy_at" json:"buy_at"`         //购买时间
	ProductAt int64  `gorm:"column:product_at" json:"product_at"` //生产日期
	OverdueAt int64  `gorm:"column:overdue_at" json:"overdue_at"` //过期时间
	StarLevel int64  `gorm:"column:star_level" json:"star_level"` //星级
	Picture   string `gorm:"column:picture" json:"picture"`       //图片
	Comment   string `gorm:"column:comment" json:"comment"`       //评价
}
type UpdateRecordReq struct {
	GoodName  string `json:"good_name"`  //物品名称
	Status    uint32 `json:"status"`     //记录状态 1使用中 2历史
	BuyAt     int64  `json:"buy_at"`     //购买时间
	ProductAt int64  `json:"product_at"` //生产日期
	OverdueAt int64  `json:"overdue_at"` //过期时间
	StarLevel int64  `json:"star_level"` //星级
	Picture   string `json:"picture"`    //图片
	Comment   string `json:"comment"`    //评价
}
type RecordCategory struct {
	gorm.Model
	RecordId   string `gorm:"column:record_id" json:"record_id"`     //记录编号
	CategoryId string `gorm:"column:category_id" json:"category_id"` //分类编号
}
