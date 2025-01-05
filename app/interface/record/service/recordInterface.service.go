package service

import (
	"for-share/app/interface/record/api"
	categoryServiceApi "for-share/app/service/category/api"
	recordServiceApi "for-share/app/service/record/api"
	"for-share/errors"
	"libong/common/context"
	commonTool "libong/common/tool"
)

func (s *Service) AddRecord(ctx context.Context, req *api.AddRecordReq) (*api.AddRecordResp, error) {
	if req.GoodsName == "" || req.BuyAt == 0 || req.ProduceAt == 0 || req.OverdueAt == 0 {
		return nil, errors.REQUEST_PARAM_ERROR
	}
	addRecordResp, err := s.recordService.AddRecord(ctx, &recordServiceApi.AddRecordReq{
		GoodsName: req.GoodsName,
		Status:    uint32(recordServiceApi.RecordStatus_RecordStatusOnline),
		BuyAt:     req.BuyAt,
		ProduceAt: req.ProduceAt,
		OverdueAt: req.OverdueAt,
		StarLevel: req.StarLevel,
		Picture:   req.Picture,
		Comment:   req.Comment,
	})
	if err != nil {
		return nil, err
	}
	return &api.AddRecordResp{
		RecordId: addRecordResp.Id,
	}, nil
}
func (s *Service) UpdateRecord(ctx context.Context, req *api.UpdateRecordReq) error {
	if req.RecordId == "" {
		return errors.REQUEST_PARAM_ERROR
	}
	updateApi := &recordServiceApi.UpdateRecordReq{}
	err := commonTool.CopyByJSON(req, updateApi)
	if err != nil {
		return err
	}
	_, err = s.recordService.UpdateRecord(ctx, updateApi)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteRecord(ctx context.Context, req *api.DeleteRecordReq) error {
	if req.RecordId == "" {
		return errors.REQUEST_PARAM_ERROR
	}
	_, err := s.recordService.DeleteRecord(ctx, &recordServiceApi.DeleteRecordReq{
		Id: req.RecordId,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) RecordByID(ctx context.Context, req *api.RecordByIDReq) (*api.RecordByIDResp, error) {
	if req.RecordId == "" {
		return nil, errors.REQUEST_PARAM_ERROR
	}
	resp := &api.RecordByIDResp{
		Record: &api.Record{},
	}
	recordByIDResp, err := s.recordService.RecordByID(ctx, &recordServiceApi.RecordByIDReq{
		Id: req.RecordId,
	})
	if err != nil {
		return nil, err
	}
	if recordByIDResp.Record == nil {
		return resp, nil
	}
	err = commonTool.CopyByJSON(recordByIDResp.Record, resp.Record)
	if err != nil {
		return nil, err
	}
	searchRecordCategoriesResp, err := s.recordService.SearchRecordCategories(ctx, &recordServiceApi.SearchRecordCategoriesReq{
		RecordIds: []string{req.RecordId},
	})
	if err != nil {
		return nil, err
	}
	if len(searchRecordCategoriesResp.List) != 0 {
		var categoryIds []string
		for _, recordCategory := range searchRecordCategoriesResp.List {
			categoryIds = append(categoryIds, recordCategory.CategoryId)
		}
		categoriesByIdsResp, err := s.categoryService.CategoriesByIds(ctx, &categoryServiceApi.CategoriesByIdsReq{
			Ids: categoryIds,
		})
		if err != nil {
			return nil, err
		}
		for _, category := range categoriesByIdsResp.Map {
			resp.Record.Categories = append(resp.Record.Categories, &api.Category{
				Id:   category.CategoryId,
				Name: category.Name,
			})
		}
	}
	return resp, nil
}
func (s *Service) SearchRecordsPage(ctx context.Context, req *api.SearchRecordsPageReq) (*api.SearchRecordsPageResp, error) {
	condition := &recordServiceApi.RecordCondition{
		FuzzyGoodsName: req.FuzzyGoodsName,
		RecordIds:      req.RecordIds,
		Status:         req.Status,
		StarLevels:     req.StarLevels,
	}
	if req.BuyAtRange != nil {
		condition.BuyStartAt = req.BuyAtRange.StartAt
		condition.BuyEndAt = req.BuyAtRange.EndAt
	}
	if req.OverdueAtRange != nil {
		condition.OverdueStartAt = req.OverdueAtRange.StartAt
		condition.OverdueEndAt = req.OverdueAtRange.EndAt
	}
	if req.ProduceAtRange != nil {
		condition.ProduceStartAt = req.ProduceAtRange.StartAt
		condition.ProduceEndAt = req.ProduceAtRange.EndAt
	}
	if req.EstablishAtRange != nil {
		condition.EstablishStartAt = req.EstablishAtRange.StartAt
		condition.EstablishEndAt = req.EstablishAtRange.EndAt
	}
	reqApi := &recordServiceApi.SearchRecordsPageReq{
		Condition: condition,
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
	}
	resp := &api.SearchRecordsPageResp{}
	countRecordResp, err := s.recordService.CountRecord(ctx, &recordServiceApi.CountRecordReq{
		Condition: condition,
	})
	if err != nil {
		return nil, err
	}
	if countRecordResp.Total == 0 {
		return resp, nil
	}
	resp.Total = countRecordResp.Total
	searchRecordsPageResp, err := s.recordService.SearchRecordsPage(ctx, reqApi)
	if err != nil {
		return nil, err
	}

	var recordIds []string
	for _, record := range searchRecordsPageResp.List {
		recordIds = append(recordIds, record.RecordId)
	}
	var categoryIds []string
	recordIdCategoryIdsMap := make(map[string][]string)
	searchRecordCategoriesResp, err := s.recordService.SearchRecordCategories(ctx, &recordServiceApi.SearchRecordCategoriesReq{
		RecordIds: recordIds,
	})
	if err != nil {
		return nil, err
	}
	for _, recordCategory := range searchRecordCategoriesResp.List {
		categoryIds = append(categoryIds, recordCategory.CategoryId)
		recordIdCategoryIdsMap[recordCategory.RecordId] = append(recordIdCategoryIdsMap[recordCategory.RecordId], recordCategory.CategoryId)
	}
	categoryIdsSet := commonTool.ListRemoveDuplicate(categoryIds)
	categoriesByIdsResp, err := s.categoryService.CategoriesByIds(ctx, &categoryServiceApi.CategoriesByIdsReq{
		Ids: categoryIdsSet,
	})
	if err != nil {
		return nil, err
	}
	categoryMap := categoriesByIdsResp.Map

	for _, record := range searchRecordsPageResp.List {
		respRecord := &api.Record{
			Categories: nil,
		}
		for _, categoryId := range recordIdCategoryIdsMap[record.RecordId] {
			if categoryMap[categoryId] != nil {
				respRecord.Categories = append(respRecord.Categories, &api.Category{
					Id:   categoryId,
					Name: categoryMap[categoryId].Name,
				})
			}
		}
		err = commonTool.CopyByJSON(record, respRecord)
		if err != nil {
			return nil, err
		}
		resp.List = append(resp.List, respRecord)
	}
	return resp, nil
}
