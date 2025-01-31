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
	CategoryService_AddCategory_FullMethodName          = "/com.libong.forShare.service.category.CategoryService/addCategory"
	CategoryService_UpdateCategory_FullMethodName       = "/com.libong.forShare.service.category.CategoryService/updateCategory"
	CategoryService_DeleteCategory_FullMethodName       = "/com.libong.forShare.service.category.CategoryService/deleteCategory"
	CategoryService_CategoryByID_FullMethodName         = "/com.libong.forShare.service.category.CategoryService/categoryByID"
	CategoryService_SearchCategoriesPage_FullMethodName = "/com.libong.forShare.service.category.CategoryService/searchCategoriesPage"
	CategoryService_CountCategory_FullMethodName        = "/com.libong.forShare.service.category.CategoryService/countCategory"
	CategoryService_CategoriesByIds_FullMethodName      = "/com.libong.forShare.service.category.CategoryService/categoriesByIds"
)

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	// 添加分类
	AddCategory(ctx context.Context, in *AddCategoryReq, opts ...grpc.CallOption) (*AddCategoryResp, error)
	// 修改分类
	UpdateCategory(ctx context.Context, in *UpdateCategoryReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除分类
	DeleteCategory(ctx context.Context, in *DeleteCategoryReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 分类详情
	CategoryByID(ctx context.Context, in *CategoryByIDReq, opts ...grpc.CallOption) (*CategoryByIDResp, error)
	// 分页查询
	SearchCategoriesPage(ctx context.Context, in *SearchCategoriesPageReq, opts ...grpc.CallOption) (*SearchCategoriesPageResp, error)
	// 查询分类总数
	CountCategory(ctx context.Context, in *CountCategoryReq, opts ...grpc.CallOption) (*CountCategoryResp, error)
	// 分类详情
	CategoriesByIds(ctx context.Context, in *CategoriesByIdsReq, opts ...grpc.CallOption) (*CategoriesByIdsResp, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) AddCategory(ctx context.Context, in *AddCategoryReq, opts ...grpc.CallOption) (*AddCategoryResp, error) {
	out := new(AddCategoryResp)
	err := c.cc.Invoke(ctx, CategoryService_AddCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CategoryService_UpdateCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) DeleteCategory(ctx context.Context, in *DeleteCategoryReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CategoryService_DeleteCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CategoryByID(ctx context.Context, in *CategoryByIDReq, opts ...grpc.CallOption) (*CategoryByIDResp, error) {
	out := new(CategoryByIDResp)
	err := c.cc.Invoke(ctx, CategoryService_CategoryByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) SearchCategoriesPage(ctx context.Context, in *SearchCategoriesPageReq, opts ...grpc.CallOption) (*SearchCategoriesPageResp, error) {
	out := new(SearchCategoriesPageResp)
	err := c.cc.Invoke(ctx, CategoryService_SearchCategoriesPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CountCategory(ctx context.Context, in *CountCategoryReq, opts ...grpc.CallOption) (*CountCategoryResp, error) {
	out := new(CountCategoryResp)
	err := c.cc.Invoke(ctx, CategoryService_CountCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CategoriesByIds(ctx context.Context, in *CategoriesByIdsReq, opts ...grpc.CallOption) (*CategoriesByIdsResp, error) {
	out := new(CategoriesByIdsResp)
	err := c.cc.Invoke(ctx, CategoryService_CategoriesByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	// 添加分类
	AddCategory(context.Context, *AddCategoryReq) (*AddCategoryResp, error)
	// 修改分类
	UpdateCategory(context.Context, *UpdateCategoryReq) (*emptypb.Empty, error)
	// 删除分类
	DeleteCategory(context.Context, *DeleteCategoryReq) (*emptypb.Empty, error)
	// 分类详情
	CategoryByID(context.Context, *CategoryByIDReq) (*CategoryByIDResp, error)
	// 分页查询
	SearchCategoriesPage(context.Context, *SearchCategoriesPageReq) (*SearchCategoriesPageResp, error)
	// 查询分类总数
	CountCategory(context.Context, *CountCategoryReq) (*CountCategoryResp, error)
	// 分类详情
	CategoriesByIds(context.Context, *CategoriesByIdsReq) (*CategoriesByIdsResp, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) AddCategory(context.Context, *AddCategoryReq) (*AddCategoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedCategoryServiceServer) UpdateCategory(context.Context, *UpdateCategoryReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedCategoryServiceServer) DeleteCategory(context.Context, *DeleteCategoryReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedCategoryServiceServer) CategoryByID(context.Context, *CategoryByIDReq) (*CategoryByIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryByID not implemented")
}
func (UnimplementedCategoryServiceServer) SearchCategoriesPage(context.Context, *SearchCategoriesPageReq) (*SearchCategoriesPageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCategoriesPage not implemented")
}
func (UnimplementedCategoryServiceServer) CountCategory(context.Context, *CountCategoryReq) (*CountCategoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCategory not implemented")
}
func (UnimplementedCategoryServiceServer) CategoriesByIds(context.Context, *CategoriesByIdsReq) (*CategoriesByIdsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoriesByIds not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_AddCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).AddCategory(ctx, req.(*AddCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_UpdateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).DeleteCategory(ctx, req.(*DeleteCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CategoryByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CategoryByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CategoryByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CategoryByID(ctx, req.(*CategoryByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_SearchCategoriesPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCategoriesPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).SearchCategoriesPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_SearchCategoriesPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).SearchCategoriesPage(ctx, req.(*SearchCategoriesPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CountCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CountCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CountCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CountCategory(ctx, req.(*CountCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CategoriesByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoriesByIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CategoriesByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CategoriesByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CategoriesByIds(ctx, req.(*CategoriesByIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.libong.forShare.service.category.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addCategory",
			Handler:    _CategoryService_AddCategory_Handler,
		},
		{
			MethodName: "updateCategory",
			Handler:    _CategoryService_UpdateCategory_Handler,
		},
		{
			MethodName: "deleteCategory",
			Handler:    _CategoryService_DeleteCategory_Handler,
		},
		{
			MethodName: "categoryByID",
			Handler:    _CategoryService_CategoryByID_Handler,
		},
		{
			MethodName: "searchCategoriesPage",
			Handler:    _CategoryService_SearchCategoriesPage_Handler,
		},
		{
			MethodName: "countCategory",
			Handler:    _CategoryService_CountCategory_Handler,
		},
		{
			MethodName: "categoriesByIds",
			Handler:    _CategoryService_CategoriesByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}
