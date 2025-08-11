package service

import (
	"for-share/app/interface/calendar/api"
	financeBillServiceApi "for-share/app/service/calendar/api"
	"for-share/errors"
	"libong/common/context"
	commonTool "libong/common/tool"
	loginError "libong/login/errors"
	accountServiceApi "libong/login/rpc/account/api"
)

func (s *Service) AddFinanceBill(ctx context.Context, req *api.AddFinanceBillReq) (*api.AddFinanceBillResp, error) {
	if ctx.User() == nil {
		return nil, loginError.AuthorizationFail
	}
	uid := ctx.User().UID
	reqApi := &financeBillServiceApi.AddFinanceBillReq{}
	err := commonTool.CopyByJSON(req, reqApi)
	if err != nil {
		return nil, err
	}
	reqApi.Owner = uid
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
	if ctx.User() == nil {
		return nil, loginError.AuthorizationFail
	}
	var uid = ctx.User().UID
	role := ctx.User().CurRole
	if role == "manager" && req.Owner != "" {
		uid = req.Owner
	}
	condition := &financeBillServiceApi.FinanceBillCondition{
		StartTimestamp: req.StartTimestamp,
		EndTimestamp:   req.EndTimestamp,
		Owner:          uid,
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
func (s *Service) SearchFinanceBillAccounts(ctx context.Context, req *api.SearchFinanceBillAccountsReq) (*api.SearchFinanceBillAccountsResp, error) {
	if ctx.User() == nil {
		return nil, loginError.AuthorizationFail
	}
	curRole := ctx.User().CurRole
	if curRole != "manager" {
		return nil, errors.PERMISSION_FAIL
	}
	resp := &api.SearchFinanceBillAccountsResp{}
	searchFinanceBillsPageResp, err := s.calendarService.SearchFinanceBillsPage(ctx, &financeBillServiceApi.SearchFinanceBillsPageReq{
		Condition: &financeBillServiceApi.FinanceBillCondition{
			GroupByOwner: true,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(searchFinanceBillsPageResp.List) == 0 {
		return resp, nil
	}
	var accountIds []string
	for _, financeBill := range searchFinanceBillsPageResp.List {
		accountIds = append(accountIds, financeBill.Owner)
	}
	accountsByIdsResp, err := s.accountService.AccountsByIds(ctx, &accountServiceApi.AccountsByIdsReq{
		Ids: accountIds,
	})
	if err != nil {
		return nil, err
	}
	for _, account := range accountsByIdsResp.Map {
		resp.List = append(resp.List, &api.FinanceBillAccount{
			Avatar:    account.Avatar,
			AccountId: account.AccountId,
			Account:   account.Account,
		})
	}
	return resp, nil
}
