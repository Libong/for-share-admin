package service

import (
	"for-share/app/interface/calendar/api"
	financeBillServiceApi "for-share/app/service/calendar/api"
	"for-share/errors"
	"libong/common/context"
	commonTool "libong/common/tool"
)

func (s *Service) AddFinanceBill(ctx context.Context, req *api.AddFinanceBillReq) (*api.AddFinanceBillResp, error) {
	reqApi := &financeBillServiceApi.AddFinanceBillReq{}
	err := commonTool.CopyByJSON(req, reqApi)
	if err != nil {
		return nil, err
	}
	addFinanceBillResp, err := s.calendarService.AddFinanceBill(ctx, reqApi)
	if err != nil {
		return nil, err
	}
	return &api.AddFinanceBillResp{
		BillId: addFinanceBillResp.Id,
	}, nil
}
func (s *Service) UpdateFinanceBill(ctx context.Context, req *api.UpdateFinanceBillReq) error {
	if req.BillId == "" {
		return errors.REQUEST_PARAM_ERROR
	}
	updateApi := &financeBillServiceApi.UpdateFinanceBillReq{}
	err := commonTool.CopyByJSON(req, updateApi)
	if err != nil {
		return err
	}
	_, err = s.calendarService.UpdateFinanceBill(ctx, updateApi)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteFinanceBill(ctx context.Context, req *api.DeleteFinanceBillReq) error {
	if req.BillId == "" {
		return errors.REQUEST_PARAM_ERROR
	}
	_, err := s.calendarService.DeleteFinanceBill(ctx, &financeBillServiceApi.DeleteFinanceBillReq{
		Id: req.BillId,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) FinanceBillByID(ctx context.Context, req *api.FinanceBillByIDReq) (*api.FinanceBillByIDResp, error) {
	if req.BillId == "" {
		return nil, errors.REQUEST_PARAM_ERROR
	}
	resp := &api.FinanceBillByIDResp{
		FinanceBill: &api.FinanceBill{},
	}
	financeBillByIDResp, err := s.calendarService.FinanceBillByID(ctx, &financeBillServiceApi.FinanceBillByIDReq{
		Id: req.BillId,
	})
	if err != nil {
		return nil, err
	}
	if financeBillByIDResp.FinanceBill == nil {
		return resp, nil
	}
	err = commonTool.CopyByJSON(financeBillByIDResp.FinanceBill, resp.FinanceBill)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s *Service) SearchFinanceBillsPage(ctx context.Context, req *api.SearchFinanceBillsPageReq) (*api.SearchFinanceBillsPageResp, error) {
	condition := &financeBillServiceApi.FinanceBillCondition{
		StartTimestamp: req.StartTimestamp,
		EndTimestamp:   req.EndTimestamp,
	}
	reqApi := &financeBillServiceApi.SearchFinanceBillsPageReq{
		Condition: condition,
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
	}
	resp := &api.SearchFinanceBillsPageResp{}
	countFinanceBillResp, err := s.calendarService.CountFinanceBill(ctx, &financeBillServiceApi.CountFinanceBillReq{
		Condition: condition,
	})
	if err != nil {
		return nil, err
	}
	if countFinanceBillResp.Total == 0 {
		return resp, nil
	}
	resp.Total = countFinanceBillResp.Total
	searchFinanceBillsPageResp, err := s.calendarService.SearchFinanceBillsPage(ctx, reqApi)
	if err != nil {
		return nil, err
	}

	for _, financeBill := range searchFinanceBillsPageResp.List {
		respFinanceBill := &api.FinanceBill{}
		err = commonTool.CopyByJSON(financeBill, respFinanceBill)
		if err != nil {
			return nil, err
		}
		resp.List = append(resp.List, respFinanceBill)
	}
	return resp, nil
}
