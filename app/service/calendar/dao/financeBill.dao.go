package dao

import (
	"for-share/app/service/calendar/api"
	"for-share/app/service/calendar/model"
	"gorm.io/gorm"
	"libong/common/context"
	"libong/common/log"
)

const (
	_financeBillTable = "finance_bill"
)

func (d *Dao) AddFinanceBill(ctx context.Context, financeBill *model.FinanceBill) error {
	db := d.Context(ctx).Table(_financeBillTable).Create(&financeBill)
	if err := db.Error; err != nil {
		log.Error(ctx, "AddFinanceBill(%+v) error(%+v)", financeBill, err)
		return err
	}
	return nil
}
func (d *Dao) UpdateFinanceBill(ctx context.Context, field interface{}, billId string) error {
	db := d.Context(ctx).Table(_financeBillTable).Where("bill_id = ?", billId).Updates(field)
	if err := db.Error; err != nil {
		log.Error(ctx, "UpdateFinanceBill billId(%v) error(%+v)", billId, err)
		return err
	}
	return nil
}
func (d *Dao) DeleteFinanceBill(ctx context.Context, billId string) error {
	db := d.Context(ctx).Table(_financeBillTable).Where("bill_id = ?", billId).Delete(&model.FinanceBill{})
	if err := db.Error; err != nil {
		log.Error(ctx, "DeleteFinanceBill(%+v) error(%+v)", billId, err)
		return err
	}
	return nil
}
func (d *Dao) FinanceBillByID(ctx context.Context, billId string) (*model.FinanceBill, error) {
	var respModel *model.FinanceBill
	db := d.Context(ctx).Table(_financeBillTable).Where("financeBill_id = ?", billId)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "FinanceBillById(%+v) error(%+v)", billId, err)
		return nil, err
	}
	return respModel, nil
}
func (d *Dao) SearchFinanceBillsPage(ctx context.Context, req *api.SearchFinanceBillsPageReq) ([]*model.FinanceBill, error) {
	var respModel []*model.FinanceBill
	db := d.Context(ctx).Table(_financeBillTable).Order("id desc")
	if req.PageNum != 0 && req.PageSize != 0 {
		db.Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize))
	}
	buildFinanceBillConditionDB(db, req.Condition)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "SearchFinanceBillsPage(%+v) error(%+v)", req, err)
		return nil, err
	}
	return respModel, nil
}

func (d *Dao) SearchFinanceBillOwners(ctx context.Context) ([]*model.FinanceBill, error) {
	var respModel []*model.FinanceBill
	db := d.Context(ctx).Table(_financeBillTable)
	db.Group("owner").Select("owner,count(*)")
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "SearchFinanceBillOwners error(%+v)", err)
		return nil, err
	}
	return respModel, nil
}

func buildFinanceBillConditionDB(db *gorm.DB, req *api.FinanceBillCondition) {
	if req == nil {
		return
	}
	//if req.FuzzyGoodsName != "" {
	//	db.Where("goods_name like ?", "%"+req.FuzzyGoodsName+"%")
	//}
	//if len(req.FinanceBillIds) != 0 {
	//	db.Where("financeBill_ids in ?", req.FinanceBillIds)
	//}
	if req.Owner != "" {
		db.Where("owner = ?", req.Owner)
	}
	//if req.BuyStartAt != 0 && req.BuyEndAt != 0 {
	//	db.Where("buy_at between ? and ?", req.BuyStartAt, req.BuyEndAt)
	//}
	//if req.ProduceStartAt != 0 && req.ProduceEndAt != 0 {
	//	db.Where("produce_at between ? and ?", req.ProduceStartAt, req.ProduceEndAt)
	//}
	if req.StartTimestamp != 0 && req.EndTimestamp != 0 {
		db.Where("timestamp between ? and ?", req.StartTimestamp, req.EndTimestamp)
	}
	//if len(req.StarLevels) != 0 {
	//	db.Where("star_level in ?", req.StarLevels)
	//}
	//if req.EstablishStartAt != 0 && req.EstablishEndAt != 0 {
	//	db.Where("created_at between ? and ?", time.Unix(req.EstablishStartAt, 0).Format("2006-01-02 15:04:05"), time.Unix(req.EstablishEndAt, 0).Format("2006-01-02 15:04:05"))
	//}
}

func (d *Dao) CountFinanceBill(ctx context.Context, req *api.FinanceBillCondition) (int64, error) {
	var count int64
	db := d.Context(ctx).Table(_financeBillTable).Model(&model.FinanceBill{})
	buildFinanceBillConditionDB(db, req)
	db.Count(&count)
	if err := db.Error; err != nil {
		log.Error(ctx, "CountPage(%+v) error(%+v)", req, err)
		return 0, err
	}
	return count, nil
}
