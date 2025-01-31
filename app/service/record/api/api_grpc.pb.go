// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: proto/api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RecordService_AddRecord_FullMethodName                   = "/com.libong.forShare.service.record.RecordService/addRecord"
	RecordService_UpdateRecord_FullMethodName                = "/com.libong.forShare.service.record.RecordService/updateRecord"
	RecordService_DeleteRecord_FullMethodName                = "/com.libong.forShare.service.record.RecordService/deleteRecord"
	RecordService_RecordByID_FullMethodName                  = "/com.libong.forShare.service.record.RecordService/recordByID"
	RecordService_SearchRecordsPage_FullMethodName           = "/com.libong.forShare.service.record.RecordService/searchRecordsPage"
	RecordService_CountRecord_FullMethodName                 = "/com.libong.forShare.service.record.RecordService/countRecord"
	RecordService_BatchAddRecordCategories_FullMethodName    = "/com.libong.forShare.service.record.RecordService/batchAddRecordCategories"
	RecordService_BatchDeleteRecordCategories_FullMethodName = "/com.libong.forShare.service.record.RecordService/batchDeleteRecordCategories"
	RecordService_SearchRecordCategories_FullMethodName      = "/com.libong.forShare.service.record.RecordService/searchRecordCategories"
)

// RecordServiceClient is the client API for RecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordServiceClient interface {
	// 添加记录
	AddRecord(ctx context.Context, in *AddRecordReq, opts ...grpc.CallOption) (*AddRecordResp, error)
	// 修改记录
	UpdateRecord(ctx context.Context, in *UpdateRecordReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除记录
	DeleteRecord(ctx context.Context, in *DeleteRecordReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 记录详情
	RecordByID(ctx context.Context, in *RecordByIDReq, opts ...grpc.CallOption) (*RecordByIDResp, error)
	// 分页查询
	SearchRecordsPage(ctx context.Context, in *SearchRecordsPageReq, opts ...grpc.CallOption) (*SearchRecordsPageResp, error)
	// 查询记录总数
	CountRecord(ctx context.Context, in *CountRecordReq, opts ...grpc.CallOption) (*CountRecordResp, error)
	// 批量添加记录分类
	BatchAddRecordCategories(ctx context.Context, in *BatchAddRecordCategoriesReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 批量删除记录分类
	BatchDeleteRecordCategories(ctx context.Context, in *BatchDeleteRecordCategoriesReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 查询记录分类
	SearchRecordCategories(ctx context.Context, in *SearchRecordCategoriesReq, opts ...grpc.CallOption) (*SearchRecordCategoriesResp, error)
}

type recordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordServiceClient(cc grpc.ClientConnInterface) RecordServiceClient {
	return &recordServiceClient{cc}
}

func (c *recordServiceClient) AddRecord(ctx context.Context, in *AddRecordReq, opts ...grpc.CallOption) (*AddRecordResp, error) {
	out := new(AddRecordResp)
	err := c.cc.Invoke(ctx, RecordService_AddRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) UpdateRecord(ctx context.Context, in *UpdateRecordReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RecordService_UpdateRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) DeleteRecord(ctx context.Context, in *DeleteRecordReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RecordService_DeleteRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) RecordByID(ctx context.Context, in *RecordByIDReq, opts ...grpc.CallOption) (*RecordByIDResp, error) {
	out := new(RecordByIDResp)
	err := c.cc.Invoke(ctx, RecordService_RecordByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) SearchRecordsPage(ctx context.Context, in *SearchRecordsPageReq, opts ...grpc.CallOption) (*SearchRecordsPageResp, error) {
	out := new(SearchRecordsPageResp)
	err := c.cc.Invoke(ctx, RecordService_SearchRecordsPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) CountRecord(ctx context.Context, in *CountRecordReq, opts ...grpc.CallOption) (*CountRecordResp, error) {
	out := new(CountRecordResp)
	err := c.cc.Invoke(ctx, RecordService_CountRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) BatchAddRecordCategories(ctx context.Context, in *BatchAddRecordCategoriesReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RecordService_BatchAddRecordCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) BatchDeleteRecordCategories(ctx context.Context, in *BatchDeleteRecordCategoriesReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RecordService_BatchDeleteRecordCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) SearchRecordCategories(ctx context.Context, in *SearchRecordCategoriesReq, opts ...grpc.CallOption) (*SearchRecordCategoriesResp, error) {
	out := new(SearchRecordCategoriesResp)
	err := c.cc.Invoke(ctx, RecordService_SearchRecordCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServiceServer is the server API for RecordService service.
// All implementations must embed UnimplementedRecordServiceServer
// for forward compatibility
type RecordServiceServer interface {
	// 添加记录
	AddRecord(context.Context, *AddRecordReq) (*AddRecordResp, error)
	// 修改记录
	UpdateRecord(context.Context, *UpdateRecordReq) (*emptypb.Empty, error)
	// 删除记录
	DeleteRecord(context.Context, *DeleteRecordReq) (*emptypb.Empty, error)
	// 记录详情
	RecordByID(context.Context, *RecordByIDReq) (*RecordByIDResp, error)
	// 分页查询
	SearchRecordsPage(context.Context, *SearchRecordsPageReq) (*SearchRecordsPageResp, error)
	// 查询记录总数
	CountRecord(context.Context, *CountRecordReq) (*CountRecordResp, error)
	// 批量添加记录分类
	BatchAddRecordCategories(context.Context, *BatchAddRecordCategoriesReq) (*emptypb.Empty, error)
	// 批量删除记录分类
	BatchDeleteRecordCategories(context.Context, *BatchDeleteRecordCategoriesReq) (*emptypb.Empty, error)
	// 查询记录分类
	SearchRecordCategories(context.Context, *SearchRecordCategoriesReq) (*SearchRecordCategoriesResp, error)
	mustEmbedUnimplementedRecordServiceServer()
}

// UnimplementedRecordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecordServiceServer struct {
}

func (UnimplementedRecordServiceServer) AddRecord(context.Context, *AddRecordReq) (*AddRecordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRecord not implemented")
}
func (UnimplementedRecordServiceServer) UpdateRecord(context.Context, *UpdateRecordReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecord not implemented")
}
func (UnimplementedRecordServiceServer) DeleteRecord(context.Context, *DeleteRecordReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecord not implemented")
}
func (UnimplementedRecordServiceServer) RecordByID(context.Context, *RecordByIDReq) (*RecordByIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordByID not implemented")
}
func (UnimplementedRecordServiceServer) SearchRecordsPage(context.Context, *SearchRecordsPageReq) (*SearchRecordsPageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRecordsPage not implemented")
}
func (UnimplementedRecordServiceServer) CountRecord(context.Context, *CountRecordReq) (*CountRecordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountRecord not implemented")
}
func (UnimplementedRecordServiceServer) BatchAddRecordCategories(context.Context, *BatchAddRecordCategoriesReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchAddRecordCategories not implemented")
}
func (UnimplementedRecordServiceServer) BatchDeleteRecordCategories(context.Context, *BatchDeleteRecordCategoriesReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteRecordCategories not implemented")
}
func (UnimplementedRecordServiceServer) SearchRecordCategories(context.Context, *SearchRecordCategoriesReq) (*SearchRecordCategoriesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRecordCategories not implemented")
}
func (UnimplementedRecordServiceServer) mustEmbedUnimplementedRecordServiceServer() {}

// UnsafeRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServiceServer will
// result in compilation errors.
type UnsafeRecordServiceServer interface {
	mustEmbedUnimplementedRecordServiceServer()
}

func RegisterRecordServiceServer(s grpc.ServiceRegistrar, srv RecordServiceServer) {
	s.RegisterService(&RecordService_ServiceDesc, srv)
}

func _RecordService_AddRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).AddRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_AddRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).AddRecord(ctx, req.(*AddRecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_UpdateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).UpdateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_UpdateRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).UpdateRecord(ctx, req.(*UpdateRecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_DeleteRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).DeleteRecord(ctx, req.(*DeleteRecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_RecordByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).RecordByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_RecordByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).RecordByID(ctx, req.(*RecordByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_SearchRecordsPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRecordsPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).SearchRecordsPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_SearchRecordsPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).SearchRecordsPage(ctx, req.(*SearchRecordsPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_CountRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).CountRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_CountRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).CountRecord(ctx, req.(*CountRecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_BatchAddRecordCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchAddRecordCategoriesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).BatchAddRecordCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_BatchAddRecordCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).BatchAddRecordCategories(ctx, req.(*BatchAddRecordCategoriesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_BatchDeleteRecordCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteRecordCategoriesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).BatchDeleteRecordCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_BatchDeleteRecordCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).BatchDeleteRecordCategories(ctx, req.(*BatchDeleteRecordCategoriesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_SearchRecordCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRecordCategoriesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).SearchRecordCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_SearchRecordCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).SearchRecordCategories(ctx, req.(*SearchRecordCategoriesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RecordService_ServiceDesc is the grpc.ServiceDesc for RecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.libong.forShare.service.record.RecordService",
	HandlerType: (*RecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addRecord",
			Handler:    _RecordService_AddRecord_Handler,
		},
		{
			MethodName: "updateRecord",
			Handler:    _RecordService_UpdateRecord_Handler,
		},
		{
			MethodName: "deleteRecord",
			Handler:    _RecordService_DeleteRecord_Handler,
		},
		{
			MethodName: "recordByID",
			Handler:    _RecordService_RecordByID_Handler,
		},
		{
			MethodName: "searchRecordsPage",
			Handler:    _RecordService_SearchRecordsPage_Handler,
		},
		{
			MethodName: "countRecord",
			Handler:    _RecordService_CountRecord_Handler,
		},
		{
			MethodName: "batchAddRecordCategories",
			Handler:    _RecordService_BatchAddRecordCategories_Handler,
		},
		{
			MethodName: "batchDeleteRecordCategories",
			Handler:    _RecordService_BatchDeleteRecordCategories_Handler,
		},
		{
			MethodName: "searchRecordCategories",
			Handler:    _RecordService_SearchRecordCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}
