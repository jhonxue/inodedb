// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: collection.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Collection_CreateCollection_FullMethodName = "/collection.Collection/CreateCollection"
	Collection_DeleteCollection_FullMethodName = "/collection.Collection/DeleteCollection"
	Collection_CreateShard_FullMethodName      = "/collection.Collection/CreateShard"
	Collection_DeleteShard_FullMethodName      = "/collection.Collection/DeleteShard"
	Collection_Metrics_FullMethodName          = "/collection.Collection/Metrics"
	Collection_Cluster_FullMethodName          = "/collection.Collection/Cluster"
	Collection_AddItems_FullMethodName         = "/collection.Collection/AddItems"
	Collection_DelItems_FullMethodName         = "/collection.Collection/DelItems"
	Collection_AddLinks_FullMethodName         = "/collection.Collection/AddLinks"
	Collection_DelLinks_FullMethodName         = "/collection.Collection/DelLinks"
	Collection_Search_FullMethodName           = "/collection.Collection/Search"
)

// CollectionClient is the client API for Collection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectionClient interface {
	CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error)
	DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error)
	CreateShard(ctx context.Context, in *CreateShardRequest, opts ...grpc.CallOption) (*CreateShardResponse, error)
	DeleteShard(ctx context.Context, in *DeleteShardRequest, opts ...grpc.CallOption) (*DeleteShardResponse, error)
	Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
	Cluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterResponse, error)
	AddItems(ctx context.Context, in *AddItemsRequest, opts ...grpc.CallOption) (*AddItemsResponse, error)
	DelItems(ctx context.Context, in *DelItemsRequest, opts ...grpc.CallOption) (*DelItemsResponse, error)
	AddLinks(ctx context.Context, in *AddLinksRequest, opts ...grpc.CallOption) (*AddLinksResponse, error)
	DelLinks(ctx context.Context, in *DelLinksRequest, opts ...grpc.CallOption) (*DelLinksResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type collectionClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionClient(cc grpc.ClientConnInterface) CollectionClient {
	return &collectionClient{cc}
}

func (c *collectionClient) CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error) {
	out := new(CreateCollectionResponse)
	err := c.cc.Invoke(ctx, Collection_CreateCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error) {
	out := new(DeleteCollectionResponse)
	err := c.cc.Invoke(ctx, Collection_DeleteCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) CreateShard(ctx context.Context, in *CreateShardRequest, opts ...grpc.CallOption) (*CreateShardResponse, error) {
	out := new(CreateShardResponse)
	err := c.cc.Invoke(ctx, Collection_CreateShard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) DeleteShard(ctx context.Context, in *DeleteShardRequest, opts ...grpc.CallOption) (*DeleteShardResponse, error) {
	out := new(DeleteShardResponse)
	err := c.cc.Invoke(ctx, Collection_DeleteShard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, Collection_Metrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) Cluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterResponse, error) {
	out := new(ClusterResponse)
	err := c.cc.Invoke(ctx, Collection_Cluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) AddItems(ctx context.Context, in *AddItemsRequest, opts ...grpc.CallOption) (*AddItemsResponse, error) {
	out := new(AddItemsResponse)
	err := c.cc.Invoke(ctx, Collection_AddItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) DelItems(ctx context.Context, in *DelItemsRequest, opts ...grpc.CallOption) (*DelItemsResponse, error) {
	out := new(DelItemsResponse)
	err := c.cc.Invoke(ctx, Collection_DelItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) AddLinks(ctx context.Context, in *AddLinksRequest, opts ...grpc.CallOption) (*AddLinksResponse, error) {
	out := new(AddLinksResponse)
	err := c.cc.Invoke(ctx, Collection_AddLinks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) DelLinks(ctx context.Context, in *DelLinksRequest, opts ...grpc.CallOption) (*DelLinksResponse, error) {
	out := new(DelLinksResponse)
	err := c.cc.Invoke(ctx, Collection_DelLinks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, Collection_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionServer is the server API for Collection service.
// All implementations must embed UnimplementedCollectionServer
// for forward compatibility
type CollectionServer interface {
	CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error)
	DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error)
	CreateShard(context.Context, *CreateShardRequest) (*CreateShardResponse, error)
	DeleteShard(context.Context, *DeleteShardRequest) (*DeleteShardResponse, error)
	Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
	Cluster(context.Context, *ClusterRequest) (*ClusterResponse, error)
	AddItems(context.Context, *AddItemsRequest) (*AddItemsResponse, error)
	DelItems(context.Context, *DelItemsRequest) (*DelItemsResponse, error)
	AddLinks(context.Context, *AddLinksRequest) (*AddLinksResponse, error)
	DelLinks(context.Context, *DelLinksRequest) (*DelLinksResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	mustEmbedUnimplementedCollectionServer()
}

// UnimplementedCollectionServer must be embedded to have forward compatible implementations.
type UnimplementedCollectionServer struct {
}

func (UnimplementedCollectionServer) CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedCollectionServer) DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollection not implemented")
}
func (UnimplementedCollectionServer) CreateShard(context.Context, *CreateShardRequest) (*CreateShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShard not implemented")
}
func (UnimplementedCollectionServer) DeleteShard(context.Context, *DeleteShardRequest) (*DeleteShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShard not implemented")
}
func (UnimplementedCollectionServer) Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metrics not implemented")
}
func (UnimplementedCollectionServer) Cluster(context.Context, *ClusterRequest) (*ClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cluster not implemented")
}
func (UnimplementedCollectionServer) AddItems(context.Context, *AddItemsRequest) (*AddItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItems not implemented")
}
func (UnimplementedCollectionServer) DelItems(context.Context, *DelItemsRequest) (*DelItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelItems not implemented")
}
func (UnimplementedCollectionServer) AddLinks(context.Context, *AddLinksRequest) (*AddLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLinks not implemented")
}
func (UnimplementedCollectionServer) DelLinks(context.Context, *DelLinksRequest) (*DelLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelLinks not implemented")
}
func (UnimplementedCollectionServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedCollectionServer) mustEmbedUnimplementedCollectionServer() {}

// UnsafeCollectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectionServer will
// result in compilation errors.
type UnsafeCollectionServer interface {
	mustEmbedUnimplementedCollectionServer()
}

func RegisterCollectionServer(s grpc.ServiceRegistrar, srv CollectionServer) {
	s.RegisterService(&Collection_ServiceDesc, srv)
}

func _Collection_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_CreateCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).CreateCollection(ctx, req.(*CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_DeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).DeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_DeleteCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).DeleteCollection(ctx, req.(*DeleteCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_CreateShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).CreateShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_CreateShard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).CreateShard(ctx, req.(*CreateShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_DeleteShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).DeleteShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_DeleteShard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).DeleteShard(ctx, req.(*DeleteShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_Metrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).Metrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_Metrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).Metrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_Cluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).Cluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_Cluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).Cluster(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_AddItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).AddItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_AddItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).AddItems(ctx, req.(*AddItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_DelItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).DelItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_DelItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).DelItems(ctx, req.(*DelItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_AddLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).AddLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_AddLinks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).AddLinks(ctx, req.(*AddLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_DelLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).DelLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_DelLinks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).DelLinks(ctx, req.(*DelLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Collection_ServiceDesc is the grpc.ServiceDesc for Collection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Collection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collection.Collection",
	HandlerType: (*CollectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollection",
			Handler:    _Collection_CreateCollection_Handler,
		},
		{
			MethodName: "DeleteCollection",
			Handler:    _Collection_DeleteCollection_Handler,
		},
		{
			MethodName: "CreateShard",
			Handler:    _Collection_CreateShard_Handler,
		},
		{
			MethodName: "DeleteShard",
			Handler:    _Collection_DeleteShard_Handler,
		},
		{
			MethodName: "Metrics",
			Handler:    _Collection_Metrics_Handler,
		},
		{
			MethodName: "Cluster",
			Handler:    _Collection_Cluster_Handler,
		},
		{
			MethodName: "AddItems",
			Handler:    _Collection_AddItems_Handler,
		},
		{
			MethodName: "DelItems",
			Handler:    _Collection_DelItems_Handler,
		},
		{
			MethodName: "AddLinks",
			Handler:    _Collection_AddLinks_Handler,
		},
		{
			MethodName: "DelLinks",
			Handler:    _Collection_DelLinks_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Collection_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collection.proto",
}
