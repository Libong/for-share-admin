package dao

import (
	"for-share/app/service/record/api"
	"for-share/app/service/record/model"
	"gorm.io/gorm"
	"libong/common/context"
	"libong/common/log"
	"time"
)

const (
	_recordTable = "record"
)

func (d *Dao) AddRecord(ctx context.Context, record *model.Record) error {
	db := d.Context(ctx).Table(_recordTable).Create(&record)
	if err := db.Error; err != nil {
		log.Error(ctx, "AddRecord(%+v) error(%+v)", record, err)
		return err
	}
	return nil
}
func (d *Dao) UpdateRecord(ctx context.Context, field interface{}, recordId string) error {
	db := d.Context(ctx).Table(_recordTable).Where("record_id = ?", recordId).Updates(field)
	if err := db.Error; err != nil {
		log.Error(ctx, "UpdateRecord recordId(%v) error(%+v)", recordId, err)
		return err
	}
	return nil
}
func (d *Dao) DeleteRecord(ctx context.Context, recordId string) error {
	db := d.Context(ctx).Table(_recordTable).Where("record_id = ?", recordId).Delete(&model.Record{})
	if err := db.Error; err != nil {
		log.Error(ctx, "DeleteRecord(%+v) error(%+v)", recordId, err)
		return err
	}
	return nil
}
func (d *Dao) RecordByID(ctx context.Context, recordId string) (*model.Record, error) {
	var respModel *model.Record
	db := d.Context(ctx).Table(_recordTable).Where("record_id = ?", recordId)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "RecordById(%+v) error(%+v)", recordId, err)
		return nil, err
	}
	return respModel, nil
}
func (d *Dao) SearchRecordsPage(ctx context.Context, req *api.SearchRecordsPageReq) ([]*model.Record, error) {
	var respModel []*model.Record
	db := d.Context(ctx).Table(_recordTable).Order("id desc")
	if req.PageNum != 0 && req.PageSize != 0 {
		db.Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize))
	}
	buildRecordConditionDB(db, req.Condition)
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "SearchRecordsPage(%+v) error(%+v)", req, err)
		return nil, err
	}
	return respModel, nil
}

func buildRecordConditionDB(db *gorm.DB, req *api.RecordCondition) {
	if req.FuzzyGoodName != "" {
		db.Where("fuzzy_good_name like ?", "%"+req.FuzzyGoodName+"%")
	}
	if len(req.RecordIds) != 0 {
		db.Where("record_ids in ?", req.RecordIds)
	}
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	if req.BuyStartAt != 0 && req.BuyEndAt != 0 {
		db.Where("buy_at between ? and ?", req.BuyStartAt, req.BuyEndAt)
	}
	if req.ProductStartAt != 0 && req.ProductEndAt != 0 {
		db.Where("product_at between ? and ?", req.ProductStartAt, req.ProductEndAt)
	}
	if req.OverdueStartAt != 0 && req.OverdueEndAt != 0 {
		db.Where("overdue_at between ? and ?", req.OverdueStartAt, req.OverdueEndAt)
	}
	if len(req.StarLevels) != 0 {
		db.Where("star_level in ?", req.StarLevels)
	}
	if req.EstablishStartAt != 0 && req.EstablishEndAt != 0 {
		db.Where("created_at between ? and ?", time.Unix(req.EstablishStartAt, 0).Format("2006-01-02 15:04:05"), time.Unix(req.EstablishEndAt, 0).Format("2006-01-02 15:04:05"))
	}
}

func (d *Dao) CountRecord(ctx context.Context, req *api.RecordCondition) (int64, error) {
	var count int64
	db := d.Context(ctx).Table(_recordTable).Model(&model.Record{})
	buildRecordConditionDB(db, req)
	db.Count(&count)
	if err := db.Error; err != nil {
		log.Error(ctx, "CountPage(%+v) error(%+v)", req, err)
		return 0, err
	}
	return count, nil
}
