package http

import (
	"for-share/app/interface/record/api"
	"libong/common/context"
	"libong/common/server/http"
)

func addRecord(ctx *http.Context) error {
	var (
		req  = &api.AddRecordReq{}
		resp *api.AddRecordResp
		err  error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if resp, err = svc.AddRecord(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}

func updateRecord(ctx *http.Context) error {
	var (
		req = &api.UpdateRecordReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.UpdateRecord(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func deleteRecord(ctx *http.Context) error {
	var (
		req = &api.DeleteRecordReq{}
		err error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if err = svc.DeleteRecord(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(nil)
	return nil
}

func recordByID(ctx *http.Context) error {
	var (
		req  = &api.RecordByIDReq{}
		resp *api.RecordByIDResp
		err  error
	)
	err = ctx.MarshalGet(req)
	if err != nil {
		return err
	}
	if resp, err = svc.RecordByID(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}

func searchRecordsPage(ctx *http.Context) error {
	var (
		req  = &api.SearchRecordsPageReq{}
		resp *api.SearchRecordsPageResp
		err  error
	)
	err = ctx.MarshalPost(req)
	if err != nil {
		return err
	}
	if resp, err = svc.SearchRecordsPage(context.FromHTTPContext(ctx), req); err != nil {
		return err
	}
	ctx.ResponseData(resp)
	return nil
}
