package dao

import (
	"for-share/app/service/record/api"
	"for-share/app/service/record/model"
	"libong/common/context"
	"libong/common/log"
)

const (
	_RecordCategoryTable = "record_category"
)

func (d *Dao) BatchAddRecordCategory(ctx context.Context, recordCategories []*model.RecordCategory) error {
	db := d.Context(ctx).Table(_RecordCategoryTable).CreateInBatches(recordCategories, 10)
	if err := db.Error; err != nil {
		log.Error(ctx, "BatchAddRecordCategory  error.  recordCategories:(%+v);error:(%+v)", recordCategories, err)
		return err
	}
	return nil
}
func (d *Dao) BatchDeleteRecordCategory(ctx context.Context, recordIds []string, categoryIds []string) error {
	db := d.Context(ctx).Table(_RecordCategoryTable)
	if len(recordIds) != 0 {
		db.Where("record_id in ?", recordIds)
	}
	if len(categoryIds) != 0 {
		db.Where("category_id in ?", categoryIds)
	}
	db.Delete(&model.RecordCategory{})
	if err := db.Error; err != nil {
		log.Error(ctx, "BatchDeleteRecordCategory  error.  recordIds:(%+v);categoryIds:(%+v);error:(%+v)", recordIds, categoryIds, err)
		return err
	}
	return nil
}
func (d *Dao) SearchRecordCategoriesPage(ctx context.Context, req *api.SearchRecordCategoriesReq) ([]*model.RecordCategory, error) {
	var respModel []*model.RecordCategory
	db := d.Context(ctx).Table(_RecordCategoryTable).Order("id desc")
	if len(req.RecordIds) != 0 {
		db.Where("record_id in ?", req.RecordIds)
	}
	if len(req.CategoryIds) != 0 {
		db.Where("category_id in ?", req.CategoryIds)
	}
	db.Find(&respModel)
	if err := db.Error; err != nil {
		log.Error(ctx, "SearchRecordCategoriesPage(%+v) error(%+v)", req, err)
		return nil, err
	}
	return respModel, nil
}
func (d *Dao) CountRecordCategory(ctx context.Context, req *api.SearchRecordCategoriesReq) (int64, error) {
	var count int64
	db := d.Context(ctx).Table(_RecordCategoryTable).Model(&model.RecordCategory{})
	if len(req.RecordIds) != 0 {
		db.Where("record_id in ?", req.RecordIds)
	}
	if len(req.CategoryIds) != 0 {
		db.Where("category_id in ?", req.CategoryIds)
	}
	db.Count(&count)
	if err := db.Error; err != nil {
		log.Error(ctx, "CountRecordCategory(%+v) error(%+v)", req, err)
		return 0, err
	}
	return count, nil
}
