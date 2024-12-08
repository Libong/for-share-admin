package grpc

import (
	"context"
	"for-share/app/service/record/api"
	"google.golang.org/protobuf/types/known/emptypb"
	qContext "libong/common/context"
)

func (s *Server) AddRecord(ctx context.Context, req *api.AddRecordReq) (*api.AddRecordResp, error) {
	return s.service.AddRecord(ctx.(qContext.Context), req)
}

func (s *Server) UpdateRecord(ctx context.Context, req *api.UpdateRecordReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.UpdateRecord(ctx.(qContext.Context), req)
}

func (s *Server) DeleteRecord(ctx context.Context, req *api.DeleteRecordReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.DeleteRecord(ctx.(qContext.Context), req)
}

func (s *Server) RecordByID(ctx context.Context, req *api.RecordByIDReq) (*api.RecordByIDResp, error) {
	return s.service.RecordByID(ctx.(qContext.Context), req)
}

func (s *Server) SearchRecordsPage(ctx context.Context, req *api.SearchRecordsPageReq) (*api.SearchRecordsPageResp, error) {
	return s.service.SearchRecordsPage(ctx.(qContext.Context), req)
}

func (s *Server) CountRecord(ctx context.Context, req *api.CountRecordReq) (*api.CountRecordResp, error) {
	return s.service.CountRecord(ctx.(qContext.Context), req)
}

func (s *Server) BatchAddRecordCategories(ctx context.Context, req *api.BatchAddRecordCategoriesReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.BatchAddRecordCategories(ctx.(qContext.Context), req)
}

func (s *Server) BatchDeleteRecordCategories(ctx context.Context, req *api.BatchDeleteRecordCategoriesReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.BatchDeleteRecordCategories(ctx.(qContext.Context), req)
}

func (s *Server) SearchRecordCategories(ctx context.Context, req *api.SearchRecordCategoriesReq) (*api.SearchRecordCategoriesResp, error) {
	return s.service.SearchRecordCategories(ctx.(qContext.Context), req)
}
