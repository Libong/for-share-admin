package service

import (
	"for-share/app/service/record/api"
	"for-share/app/service/record/model"
	"libong/common/context"
	commonTool "libong/common/tool"
)

func (s *Service) BatchAddRecordCategories(ctx context.Context, req *api.BatchAddRecordCategoriesReq) error {
	var models []*model.RecordCategory
	for _, item := range req.List {
		m := &model.RecordCategory{}
		err := commonTool.CopyByJSON(item, m)
		if err != nil {
			return err
		}
		models = append(models, m)
	}
	err := s.dao.BatchAddRecordCategory(ctx, models)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) BatchDeleteRecordCategories(ctx context.Context, req *api.BatchDeleteRecordCategoriesReq) error {
	return s.dao.BatchDeleteRecordCategory(ctx, req.RecordIds, req.CategoryIds)
}

func (s *Service) SearchRecordCategories(ctx context.Context, req *api.SearchRecordCategoriesReq) (*api.SearchRecordCategoriesResp, error) {
	models, err := s.dao.SearchRecordCategoriesPage(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := &api.SearchRecordCategoriesResp{}
	if len(models) == 0 {
		return resp, nil
	}
	for _, m := range models {
		respModel := &api.RecordCategory{}
		err = commonTool.CopyByJSON(m, respModel)
		if err != nil {
			return nil, err
		}
		resp.List = append(resp.List, respModel)
	}
	return resp, nil
}
