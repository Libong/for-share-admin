package grpc

import (
	"context"
	"for-share/app/service/calendar/api"
	"google.golang.org/protobuf/types/known/emptypb"
	qContext "libong/common/context"
)

func (s *Server) AddFinanceBill(ctx context.Context, req *api.AddFinanceBillReq) (*api.AddFinanceBillResp, error) {
	return s.service.AddFinanceBill(ctx.(qContext.Context), req)
}

func (s *Server) UpdateFinanceBill(ctx context.Context, req *api.UpdateFinanceBillReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.UpdateFinanceBill(ctx.(qContext.Context), req)
}

func (s *Server) DeleteFinanceBill(ctx context.Context, req *api.DeleteFinanceBillReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.DeleteFinanceBill(ctx.(qContext.Context), req)
}

func (s *Server) FinanceBillByID(ctx context.Context, req *api.FinanceBillByIDReq) (*api.FinanceBillByIDResp, error) {
	return s.service.FinanceBillByID(ctx.(qContext.Context), req)
}

func (s *Server) SearchFinanceBillsPage(ctx context.Context, req *api.SearchFinanceBillsPageReq) (*api.SearchFinanceBillsPageResp, error) {
	return s.service.SearchFinanceBillsPage(ctx.(qContext.Context), req)
}

func (s *Server) CountFinanceBill(ctx context.Context, req *api.CountFinanceBillReq) (*api.CountFinanceBillResp, error) {
	return s.service.CountFinanceBill(ctx.(qContext.Context), req)
}
