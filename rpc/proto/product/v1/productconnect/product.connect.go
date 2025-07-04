// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/product/v1/product.proto

package productconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "example/pb/proto/product/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ProductServiceName is the fully-qualified name of the ProductService service.
	ProductServiceName = "proto.product.v1.ProductService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProductServiceCreateProductProcedure is the fully-qualified name of the ProductService's
	// CreateProduct RPC.
	ProductServiceCreateProductProcedure = "/proto.product.v1.ProductService/CreateProduct"
	// ProductServiceGetProductProcedure is the fully-qualified name of the ProductService's GetProduct
	// RPC.
	ProductServiceGetProductProcedure = "/proto.product.v1.ProductService/GetProduct"
	// ProductServiceListProductProcedure is the fully-qualified name of the ProductService's
	// ListProduct RPC.
	ProductServiceListProductProcedure = "/proto.product.v1.ProductService/ListProduct"
	// ProductServiceUpdateProductProcedure is the fully-qualified name of the ProductService's
	// UpdateProduct RPC.
	ProductServiceUpdateProductProcedure = "/proto.product.v1.ProductService/UpdateProduct"
	// ProductServiceDeleteProductProcedure is the fully-qualified name of the ProductService's
	// DeleteProduct RPC.
	ProductServiceDeleteProductProcedure = "/proto.product.v1.ProductService/DeleteProduct"
)

// ProductServiceClient is a client for the proto.product.v1.ProductService service.
type ProductServiceClient interface {
	CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error)
	GetProduct(context.Context, *connect.Request[v1.GetProductRequest]) (*connect.Response[v1.GetProductResponse], error)
	ListProduct(context.Context, *connect.Request[v1.ListProductRequest]) (*connect.Response[v1.ListProductResponse], error)
	UpdateProduct(context.Context, *connect.Request[v1.UpdateProductRequest]) (*connect.Response[v1.UpdateProductResponse], error)
	DeleteProduct(context.Context, *connect.Request[v1.DeleteProductRequest]) (*connect.Response[v1.DeleteProductResponse], error)
}

// NewProductServiceClient constructs a client for the proto.product.v1.ProductService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProductServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ProductServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	productServiceMethods := v1.File_proto_product_v1_product_proto.Services().ByName("ProductService").Methods()
	return &productServiceClient{
		createProduct: connect.NewClient[v1.CreateProductRequest, v1.CreateProductResponse](
			httpClient,
			baseURL+ProductServiceCreateProductProcedure,
			connect.WithSchema(productServiceMethods.ByName("CreateProduct")),
			connect.WithClientOptions(opts...),
		),
		getProduct: connect.NewClient[v1.GetProductRequest, v1.GetProductResponse](
			httpClient,
			baseURL+ProductServiceGetProductProcedure,
			connect.WithSchema(productServiceMethods.ByName("GetProduct")),
			connect.WithClientOptions(opts...),
		),
		listProduct: connect.NewClient[v1.ListProductRequest, v1.ListProductResponse](
			httpClient,
			baseURL+ProductServiceListProductProcedure,
			connect.WithSchema(productServiceMethods.ByName("ListProduct")),
			connect.WithClientOptions(opts...),
		),
		updateProduct: connect.NewClient[v1.UpdateProductRequest, v1.UpdateProductResponse](
			httpClient,
			baseURL+ProductServiceUpdateProductProcedure,
			connect.WithSchema(productServiceMethods.ByName("UpdateProduct")),
			connect.WithClientOptions(opts...),
		),
		deleteProduct: connect.NewClient[v1.DeleteProductRequest, v1.DeleteProductResponse](
			httpClient,
			baseURL+ProductServiceDeleteProductProcedure,
			connect.WithSchema(productServiceMethods.ByName("DeleteProduct")),
			connect.WithClientOptions(opts...),
		),
	}
}

// productServiceClient implements ProductServiceClient.
type productServiceClient struct {
	createProduct *connect.Client[v1.CreateProductRequest, v1.CreateProductResponse]
	getProduct    *connect.Client[v1.GetProductRequest, v1.GetProductResponse]
	listProduct   *connect.Client[v1.ListProductRequest, v1.ListProductResponse]
	updateProduct *connect.Client[v1.UpdateProductRequest, v1.UpdateProductResponse]
	deleteProduct *connect.Client[v1.DeleteProductRequest, v1.DeleteProductResponse]
}

// CreateProduct calls proto.product.v1.ProductService.CreateProduct.
func (c *productServiceClient) CreateProduct(ctx context.Context, req *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error) {
	return c.createProduct.CallUnary(ctx, req)
}

// GetProduct calls proto.product.v1.ProductService.GetProduct.
func (c *productServiceClient) GetProduct(ctx context.Context, req *connect.Request[v1.GetProductRequest]) (*connect.Response[v1.GetProductResponse], error) {
	return c.getProduct.CallUnary(ctx, req)
}

// ListProduct calls proto.product.v1.ProductService.ListProduct.
func (c *productServiceClient) ListProduct(ctx context.Context, req *connect.Request[v1.ListProductRequest]) (*connect.Response[v1.ListProductResponse], error) {
	return c.listProduct.CallUnary(ctx, req)
}

// UpdateProduct calls proto.product.v1.ProductService.UpdateProduct.
func (c *productServiceClient) UpdateProduct(ctx context.Context, req *connect.Request[v1.UpdateProductRequest]) (*connect.Response[v1.UpdateProductResponse], error) {
	return c.updateProduct.CallUnary(ctx, req)
}

// DeleteProduct calls proto.product.v1.ProductService.DeleteProduct.
func (c *productServiceClient) DeleteProduct(ctx context.Context, req *connect.Request[v1.DeleteProductRequest]) (*connect.Response[v1.DeleteProductResponse], error) {
	return c.deleteProduct.CallUnary(ctx, req)
}

// ProductServiceHandler is an implementation of the proto.product.v1.ProductService service.
type ProductServiceHandler interface {
	CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error)
	GetProduct(context.Context, *connect.Request[v1.GetProductRequest]) (*connect.Response[v1.GetProductResponse], error)
	ListProduct(context.Context, *connect.Request[v1.ListProductRequest]) (*connect.Response[v1.ListProductResponse], error)
	UpdateProduct(context.Context, *connect.Request[v1.UpdateProductRequest]) (*connect.Response[v1.UpdateProductResponse], error)
	DeleteProduct(context.Context, *connect.Request[v1.DeleteProductRequest]) (*connect.Response[v1.DeleteProductResponse], error)
}

// NewProductServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProductServiceHandler(svc ProductServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	productServiceMethods := v1.File_proto_product_v1_product_proto.Services().ByName("ProductService").Methods()
	productServiceCreateProductHandler := connect.NewUnaryHandler(
		ProductServiceCreateProductProcedure,
		svc.CreateProduct,
		connect.WithSchema(productServiceMethods.ByName("CreateProduct")),
		connect.WithHandlerOptions(opts...),
	)
	productServiceGetProductHandler := connect.NewUnaryHandler(
		ProductServiceGetProductProcedure,
		svc.GetProduct,
		connect.WithSchema(productServiceMethods.ByName("GetProduct")),
		connect.WithHandlerOptions(opts...),
	)
	productServiceListProductHandler := connect.NewUnaryHandler(
		ProductServiceListProductProcedure,
		svc.ListProduct,
		connect.WithSchema(productServiceMethods.ByName("ListProduct")),
		connect.WithHandlerOptions(opts...),
	)
	productServiceUpdateProductHandler := connect.NewUnaryHandler(
		ProductServiceUpdateProductProcedure,
		svc.UpdateProduct,
		connect.WithSchema(productServiceMethods.ByName("UpdateProduct")),
		connect.WithHandlerOptions(opts...),
	)
	productServiceDeleteProductHandler := connect.NewUnaryHandler(
		ProductServiceDeleteProductProcedure,
		svc.DeleteProduct,
		connect.WithSchema(productServiceMethods.ByName("DeleteProduct")),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.product.v1.ProductService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ProductServiceCreateProductProcedure:
			productServiceCreateProductHandler.ServeHTTP(w, r)
		case ProductServiceGetProductProcedure:
			productServiceGetProductHandler.ServeHTTP(w, r)
		case ProductServiceListProductProcedure:
			productServiceListProductHandler.ServeHTTP(w, r)
		case ProductServiceUpdateProductProcedure:
			productServiceUpdateProductHandler.ServeHTTP(w, r)
		case ProductServiceDeleteProductProcedure:
			productServiceDeleteProductHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedProductServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProductServiceHandler struct{}

func (UnimplementedProductServiceHandler) CreateProduct(context.Context, *connect.Request[v1.CreateProductRequest]) (*connect.Response[v1.CreateProductResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.product.v1.ProductService.CreateProduct is not implemented"))
}

func (UnimplementedProductServiceHandler) GetProduct(context.Context, *connect.Request[v1.GetProductRequest]) (*connect.Response[v1.GetProductResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.product.v1.ProductService.GetProduct is not implemented"))
}

func (UnimplementedProductServiceHandler) ListProduct(context.Context, *connect.Request[v1.ListProductRequest]) (*connect.Response[v1.ListProductResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.product.v1.ProductService.ListProduct is not implemented"))
}

func (UnimplementedProductServiceHandler) UpdateProduct(context.Context, *connect.Request[v1.UpdateProductRequest]) (*connect.Response[v1.UpdateProductResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.product.v1.ProductService.UpdateProduct is not implemented"))
}

func (UnimplementedProductServiceHandler) DeleteProduct(context.Context, *connect.Request[v1.DeleteProductRequest]) (*connect.Response[v1.DeleteProductResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.product.v1.ProductService.DeleteProduct is not implemented"))
}
