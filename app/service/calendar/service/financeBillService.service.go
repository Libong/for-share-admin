package service

import (
	"for-share/app/service/calendar/api"
	"for-share/app/service/calendar/model"
	"libong/common/context"
	"libong/common/snowflake"
	commonTool "libong/common/tool"
)

func (s *Service) AddFinanceBill(ctx context.Context, req *api.AddFinanceBillReq) (*api.AddFinanceBillResp, error) {
	addModel := &model.FinanceBill{}
	err := commonTool.CopyByJSON(req, addModel)
	if err != nil {
		return nil, err
	}
	addModel.BillId = snowflake.SnowflakeWorker.NextID().String()
	err = s.dao.AddFinanceBill(ctx, addModel)
	if err != nil {
		return nil, err
	}
	return &api.AddFinanceBillResp{
		Id: addModel.BillId,
	}, nil
}

func (s *Service) UpdateFinanceBill(ctx context.Context, req *api.UpdateFinanceBillReq) error {
	//先更新非零值 只有map才能更新零值
	updateModel := &model.UpdateFinanceBillReq{}
	err := commonTool.CopyByJSON(req, updateModel)
	if err != nil {
		return err
	}
	err = s.dao.UpdateFinanceBill(ctx, updateModel, req.BillId)
	if err != nil {
		return err
	}
	//再更新零值
	if req.HandleZero {
		updateMap := make(map[string]interface{})
		if req.ZeroFields == nil || len(req.ZeroFields) == 0 {
			return nil
		}
		valueMap := commonTool.StructToKeepEmptyMap(*updateModel)
		zeroMap := commonTool.StructToDefaultValueMap(updateModel)
		for _, field := range req.ZeroFields {
			if zeroMap[field] != nil && valueMap[field] != nil {
				updateMap[field] = zeroMap[field]
			}
		}
		err = s.dao.UpdateFinanceBill(ctx, updateMap, req.BillId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) DeleteFinanceBill(ctx context.Context, req *api.DeleteFinanceBillReq) error {
	return s.dao.DeleteFinanceBill(ctx, req.Id)
}

func (s *Service) FinanceBillByID(ctx context.Context, req *api.FinanceBillByIDReq) (*api.FinanceBillByIDResp, error) {
	m, err := s.dao.FinanceBillByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	respApi := &api.FinanceBillByIDResp{FinanceBill: &api.FinanceBill{}}
	if m == nil {
		return respApi, nil
	}
	err = commonTool.CopyByJSON(m, respApi.FinanceBill)
	respApi.FinanceBill.EstablishAt = m.CreatedAt.Unix()
	if err != nil {
		return nil, err
	}
	return respApi, nil
}

func (s *Service) SearchFinanceBillsPage(ctx context.Context, req *api.SearchFinanceBillsPageReq) (*api.SearchFinanceBillsPageResp, error) {
	var models []*model.FinanceBill
	var err error
	if req.Condition != nil && req.Condition.GroupByOwner {
		models, err = s.dao.SearchFinanceBillOwners(ctx)
	} else {
		models, err = s.dao.SearchFinanceBillsPage(ctx, req)
	}
	if err != nil {
		return nil, err
	}
	resp := &api.SearchFinanceBillsPageResp{}
	if len(models) == 0 {
		return resp, nil
	}
	for _, m := range models {
		respModel := &api.FinanceBill{}
		err = commonTool.CopyByJSON(m, respModel)
		if err != nil {
			return nil, err
		}
		respModel.EstablishAt = m.CreatedAt.Unix()
		resp.List = append(resp.List, respModel)
	}
	return resp, nil
}

func (s *Service) CountFinanceBill(ctx context.Context, req *api.CountFinanceBillReq) (*api.CountFinanceBillResp, error) {
	count, err := s.dao.CountFinanceBill(ctx, req.Condition)
	if err != nil {
		return nil, err
	}
	return &api.CountFinanceBillResp{
		Total: count,
	}, nil
}
