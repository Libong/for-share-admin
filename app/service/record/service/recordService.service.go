package service

import (
	"for-share/app/service/record/api"
	"for-share/app/service/record/model"
	"libong/common/context"
	"libong/common/snowflake"
	commonTool "libong/common/tool"
)

func (s *Service) AddRecord(ctx context.Context, req *api.AddRecordReq) (*api.AddRecordResp, error) {
	addModel := &model.Record{}
	err := commonTool.CopyByJSON(req, addModel)
	if err != nil {
		return nil, err
	}
	addModel.RecordId = snowflake.SnowflakeWorker.NextID().String()
	err = s.dao.AddRecord(ctx, addModel)
	if err != nil {
		return nil, err
	}
	return &api.AddRecordResp{
		Id: addModel.RecordId,
	}, nil
}

func (s *Service) UpdateRecord(ctx context.Context, req *api.UpdateRecordReq) error {
	//先更新非零值 只有map才能更新零值
	updateModel := &model.UpdateRecordReq{}
	err := commonTool.CopyByJSON(req, updateModel)
	if err != nil {
		return err
	}
	err = s.dao.UpdateRecord(ctx, updateModel, req.RecordId)
	if err != nil {
		return err
	}
	//再更新零值
	if req.HandleZero {
		updateMap := make(map[string]interface{})
		if req.ZeroFields == nil || len(req.ZeroFields) == 0 {
			return nil
		}
		valueMap := commonTool.StructToFilterEmptyMap(*updateModel)
		zeroMap := commonTool.StructToDefaultValueMap(updateModel)
		for _, field := range req.ZeroFields {
			if zeroMap[field] != nil && valueMap[field] != nil {
				updateMap[field] = zeroMap[field]
			}
		}
		err = s.dao.UpdateRecord(ctx, updateMap, req.RecordId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) DeleteRecord(ctx context.Context, req *api.DeleteRecordReq) error {
	return s.dao.DeleteRecord(ctx, req.Id)
}

func (s *Service) RecordByID(ctx context.Context, req *api.RecordByIDReq) (*api.RecordByIDResp, error) {
	m, err := s.dao.RecordByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	respApi := &api.RecordByIDResp{Record: &api.Record{}}
	if m == nil {
		return respApi, nil
	}
	err = commonTool.CopyByJSON(m, respApi.Record)
	respApi.Record.EstablishAt = m.CreatedAt.Unix()
	if err != nil {
		return nil, err
	}
	return respApi, nil
}

func (s *Service) SearchRecordsPage(ctx context.Context, req *api.SearchRecordsPageReq) (*api.SearchRecordsPageResp, error) {
	models, err := s.dao.SearchRecordsPage(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := &api.SearchRecordsPageResp{}
	if len(models) == 0 {
		return resp, nil
	}
	for _, m := range models {
		respModel := &api.Record{}
		err = commonTool.CopyByJSON(m, respModel)
		if err != nil {
			return nil, err
		}
		respModel.EstablishAt = m.CreatedAt.Unix()
		resp.List = append(resp.List, respModel)
	}
	return resp, nil
}

func (s *Service) CountRecord(ctx context.Context, req *api.CountRecordReq) (*api.CountRecordResp, error) {
	count, err := s.dao.CountRecord(ctx, req.Condition)
	if err != nil {
		return nil, err
	}
	return &api.CountRecordResp{
		Total: count,
	}, nil
}
