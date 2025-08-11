package http

import (
	"for-share/app/interface/calendar/api"
	"libong/common/context"
	"libong/common/server/http"
)

func addFinanceBill(ctx *http.Context) error {
	var (
		req  = &api.AddFinanceBillReq{}
		resp *api.AddFinanceBillResp
		err  error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if resp, err = svc.AddFinanceBill(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}

func updateFinanceBill(ctx *http.Context) error {
	var (
		req = &api.UpdateFinanceBillReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.UpdateFinanceBill(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func deleteFinanceBill(ctx *http.Context) error {
	var (
		req = &api.DeleteFinanceBillReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.DeleteFinanceBill(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func financeBillByID(ctx *http.Context) error {
	var (
		req  = &api.FinanceBillByIDReq{}
		resp *api.FinanceBillByIDResp
		err  error
	)
	err = ctx.MarshalGet(req)
	if err != nil {
		return err
	}
	if resp, err = svc.FinanceBillByID(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}

func searchFinanceBillsPage(ctx *http.Context) error {
	var (
		req  = &api.SearchFinanceBillsPageReq{}
		resp *api.SearchFinanceBillsPageResp
		err  error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if resp, err = svc.SearchFinanceBillsPage(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}
func searchFinanceBillAccounts(ctx *http.Context) error {
	var (
		req  = &api.SearchFinanceBillAccountsReq{}
		resp *api.SearchFinanceBillAccountsResp
		err  error
	)
	err = ctx.MarshalGet(req)
	if err != nil {
		return err
	}
	if resp, err = svc.SearchFinanceBillAccounts(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}
