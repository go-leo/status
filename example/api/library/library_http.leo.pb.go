// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package library

import (
	context "context"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	mux "github.com/gorilla/mux"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	io "io"
	http1 "net/http"
	strconv "strconv"
)

func NewLibraryServiceHTTPServer(
	endpoints interface {
		CreateShelf() endpoint.Endpoint
		GetShelf() endpoint.Endpoint
		ListShelves() endpoint.Endpoint
		DeleteShelf() endpoint.Endpoint
		MergeShelves() endpoint.Endpoint
		CreateBook() endpoint.Endpoint
		GetBook() endpoint.Endpoint
		ListBooks() endpoint.Endpoint
		DeleteBook() endpoint.Endpoint
		UpdateBook() endpoint.Endpoint
		MoveBook() endpoint.Endpoint
	},
	mdw []endpoint.Middleware,
	opts ...http.ServerOption,
) http1.Handler {
	r := mux.NewRouter()
	r.Methods("POST").
		Path("/v1/shelves").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.CreateShelf(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &CreateShelfRequest{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				if err := protojson.Unmarshal(body, req.Shelf); err != nil {
					return nil, err
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	r.Methods("GET").
		Path("/v1/shelves/{shelf}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.GetShelf(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &GetShelfRequest{}
				vars := mux.Vars(r)
				req.Name = fmt.Sprintf("shelves/%s", vars["shelf"])
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	r.Methods("GET").
		Path("/v1/shelves").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.ListShelves(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &ListShelvesRequest{}
				queries := r.URL.Query()
				if v, err := strconv.ParseInt(queries.Get("page_size"), 10, 32); err != nil {
					return nil, err
				} else {
					req.PageSize = int32(v)
				}
				req.PageToken = queries.Get("page_token")
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	r.Methods("DELETE").
		Path("/v1/shelves/{shelf}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.DeleteShelf(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &DeleteShelfRequest{}
				vars := mux.Vars(r)
				req.Name = fmt.Sprintf("shelves/%s", vars["shelf"])
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	r.Methods("POST").
		Path("/v1/shelves/{shelf}:merge").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.MergeShelves(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &MergeShelvesRequest{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				if err := protojson.Unmarshal(body, req); err != nil {
					return nil, err
				}
				vars := mux.Vars(r)
				req.Name = fmt.Sprintf("shelves/%s", vars["shelf"])
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	r.Methods("POST").
		Path("/v1/shelves/{shelf}/books").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.CreateBook(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &CreateBookRequest{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				if err := protojson.Unmarshal(body, req.Book); err != nil {
					return nil, err
				}
				vars := mux.Vars(r)
				req.Parent = fmt.Sprintf("shelves/%s", vars["shelf"])
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	r.Methods("GET").
		Path("/v1/shelves/{shelf}/books/{book}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.GetBook(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &GetBookRequest{}
				vars := mux.Vars(r)
				req.Name = fmt.Sprintf("shelves/%s/books/%s", vars["shelf"], vars["book"])
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	r.Methods("GET").
		Path("/v1/shelves/{shelf}/books").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.ListBooks(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &ListBooksRequest{}
				vars := mux.Vars(r)
				req.Parent = fmt.Sprintf("shelves/%s", vars["shelf"])
				queries := r.URL.Query()
				if v, err := strconv.ParseInt(queries.Get("page_size"), 10, 32); err != nil {
					return nil, err
				} else {
					req.PageSize = int32(v)
				}
				req.PageToken = queries.Get("page_token")
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	r.Methods("DELETE").
		Path("/v1/shelves/{shelf}/books/{book}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.DeleteBook(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &DeleteBookRequest{}
				vars := mux.Vars(r)
				req.Name = fmt.Sprintf("shelves/%s/books/%s", vars["shelf"], vars["book"])
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	r.Methods("PATCH").
		Path("/v1/shelves/{shelf}/books/{book}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.UpdateBook(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &UpdateBookRequest{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				if err := protojson.Unmarshal(body, req.Book); err != nil {
					return nil, err
				}
				vars := mux.Vars(r)
				if req.Book == nil {
					req.Book = &Book{}
				}
				req.Book.Name = fmt.Sprintf("shelves/%s/books/%s", vars["shelf"], vars["book"])
				queries := r.URL.Query()
				mask, err := fieldmaskpb.New(req.Book, queries["update_mask"]...)
				if err != nil {
					return nil, err
				}
				req.UpdateMask = mask
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	r.Methods("POST").
		Path("/v1/shelves/{shelf}/books/{book}:move").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.MoveBook(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &MoveBookRequest{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				if err := protojson.Unmarshal(body, req); err != nil {
					return nil, err
				}
				vars := mux.Vars(r)
				req.Name = fmt.Sprintf("shelves/%s/books/%s", vars["shelf"], vars["book"])
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, resp any) error {
				return nil
			},
			opts...,
		))
	return r
}
