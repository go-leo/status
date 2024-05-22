// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package library

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	errorx "github.com/go-leo/gox/errorx"
	urlx "github.com/go-leo/gox/netx/urlx"
	strconvx "github.com/go-leo/gox/strconvx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	mux "github.com/gorilla/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http1 "net/http"
	url "net/url"
	strings "strings"
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
	opts []http.ServerOption,
	mdw ...endpoint.Middleware,
) http1.Handler {
	router := mux.NewRouter()
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/CreateShelf").
		Methods("POST").
		Path("/v1/shelves").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.CreateShelf(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &CreateShelfRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(req.Shelf); err != nil {
					return nil, err
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Shelf)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/GetShelf").
		Methods("GET").
		Path("/v1/shelves/{shelf}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.GetShelf(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &GetShelfRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Name = fmt.Sprintf("shelves/%s", vars.Get("shelf"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Shelf)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/ListShelves").
		Methods("GET").
		Path("/v1/shelves").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.ListShelves(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &ListShelvesRequest{}
				queries := r.URL.Query()
				var queryErr error
				req.PageSize, queryErr = errorx.Break[int32](queryErr)(urlx.GetInt[int32](queries, "page_size"))
				req.PageToken = queries.Get("page_token")
				if queryErr != nil {
					return nil, queryErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*ListShelvesResponse)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/DeleteShelf").
		Methods("DELETE").
		Path("/v1/shelves/{shelf}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.DeleteShelf(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &DeleteShelfRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Name = fmt.Sprintf("shelves/%s", vars.Get("shelf"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/MergeShelves").
		Methods("POST").
		Path("/v1/shelves/{shelf}:merge").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.MergeShelves(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &MergeShelvesRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
					return nil, err
				}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Name = fmt.Sprintf("shelves/%s", vars.Get("shelf"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Shelf)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/CreateBook").
		Methods("POST").
		Path("/v1/shelves/{shelf}/books").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.CreateBook(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &CreateBookRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(req.Book); err != nil {
					return nil, err
				}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Parent = fmt.Sprintf("shelves/%s", vars.Get("shelf"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Book)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/GetBook").
		Methods("GET").
		Path("/v1/shelves/{shelf}/books/{book}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.GetBook(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &GetBookRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Name = fmt.Sprintf("shelves/%s/books/%s", vars.Get("shelf"), vars.Get("book"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Book)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/ListBooks").
		Methods("GET").
		Path("/v1/shelves/{shelf}/books").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.ListBooks(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &ListBooksRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Parent = fmt.Sprintf("shelves/%s", vars.Get("shelf"))
				if varErr != nil {
					return nil, varErr
				}
				queries := r.URL.Query()
				var queryErr error
				req.PageSize, queryErr = errorx.Break[int32](queryErr)(urlx.GetInt[int32](queries, "page_size"))
				req.PageToken = queries.Get("page_token")
				if queryErr != nil {
					return nil, queryErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*ListBooksResponse)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/DeleteBook").
		Methods("DELETE").
		Path("/v1/shelves/{shelf}/books/{book}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.DeleteBook(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &DeleteBookRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Name = fmt.Sprintf("shelves/%s/books/%s", vars.Get("shelf"), vars.Get("book"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/UpdateBook").
		Methods("PATCH").
		Path("/v1/shelves/{shelf}/books/{book}").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.UpdateBook(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &UpdateBookRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(req.Book); err != nil {
					return nil, err
				}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				if req.Book == nil {
					req.Book = &Book{}
				}
				req.Book.Name = fmt.Sprintf("shelves/%s/books/%s", vars.Get("shelf"), vars.Get("book"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Book)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/MoveBook").
		Methods("POST").
		Path("/v1/shelves/{shelf}/books/{book}:move").
		Handler(http.NewServer(
			endpointx.Chain(endpoints.MoveBook(), mdw...),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &MoveBookRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
					return nil, err
				}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Name = fmt.Sprintf("shelves/%s/books/%s", vars.Get("shelf"), vars.Get("book"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Book)
				w.WriteHeader(http1.StatusOK)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			opts...,
		))
	return router
}

type libraryServiceHTTPClient struct {
	createShelf  endpoint.Endpoint
	getShelf     endpoint.Endpoint
	listShelves  endpoint.Endpoint
	deleteShelf  endpoint.Endpoint
	mergeShelves endpoint.Endpoint
	createBook   endpoint.Endpoint
	getBook      endpoint.Endpoint
	listBooks    endpoint.Endpoint
	deleteBook   endpoint.Endpoint
	updateBook   endpoint.Endpoint
	moveBook     endpoint.Endpoint
}

func (c *libraryServiceHTTPClient) CreateShelf(ctx context.Context, request *CreateShelfRequest) (*Shelf, error) {
	rep, err := c.createShelf(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Shelf), nil
}

func (c *libraryServiceHTTPClient) GetShelf(ctx context.Context, request *GetShelfRequest) (*Shelf, error) {
	rep, err := c.getShelf(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Shelf), nil
}

func (c *libraryServiceHTTPClient) ListShelves(ctx context.Context, request *ListShelvesRequest) (*ListShelvesResponse, error) {
	rep, err := c.listShelves(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListShelvesResponse), nil
}

func (c *libraryServiceHTTPClient) DeleteShelf(ctx context.Context, request *DeleteShelfRequest) (*emptypb.Empty, error) {
	rep, err := c.deleteShelf(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *libraryServiceHTTPClient) MergeShelves(ctx context.Context, request *MergeShelvesRequest) (*Shelf, error) {
	rep, err := c.mergeShelves(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Shelf), nil
}

func (c *libraryServiceHTTPClient) CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error) {
	rep, err := c.createBook(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func (c *libraryServiceHTTPClient) GetBook(ctx context.Context, request *GetBookRequest) (*Book, error) {
	rep, err := c.getBook(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func (c *libraryServiceHTTPClient) ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error) {
	rep, err := c.listBooks(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListBooksResponse), nil
}

func (c *libraryServiceHTTPClient) DeleteBook(ctx context.Context, request *DeleteBookRequest) (*emptypb.Empty, error) {
	rep, err := c.deleteBook(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *libraryServiceHTTPClient) UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error) {
	rep, err := c.updateBook(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func (c *libraryServiceHTTPClient) MoveBook(ctx context.Context, request *MoveBookRequest) (*Book, error) {
	rep, err := c.moveBook(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func NewLibraryServiceHTTPClient(
	scheme string,
	instance string,
	opts []http.ClientOption,
	mdw ...endpoint.Middleware,
) interface {
	CreateShelf(ctx context.Context, request *CreateShelfRequest) (*Shelf, error)
	GetShelf(ctx context.Context, request *GetShelfRequest) (*Shelf, error)
	ListShelves(ctx context.Context, request *ListShelvesRequest) (*ListShelvesResponse, error)
	DeleteShelf(ctx context.Context, request *DeleteShelfRequest) (*emptypb.Empty, error)
	MergeShelves(ctx context.Context, request *MergeShelvesRequest) (*Shelf, error)
	CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error)
	GetBook(ctx context.Context, request *GetBookRequest) (*Book, error)
	ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error)
	DeleteBook(ctx context.Context, request *DeleteBookRequest) (*emptypb.Empty, error)
	UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error)
	MoveBook(ctx context.Context, request *MoveBookRequest) (*Book, error)
} {
	router := mux.NewRouter()
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/CreateShelf").
		Methods("POST").
		Path("/v1/shelves")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/GetShelf").
		Methods("GET").
		Path("/v1/shelves/{shelf}")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/ListShelves").
		Methods("GET").
		Path("/v1/shelves")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/DeleteShelf").
		Methods("DELETE").
		Path("/v1/shelves/{shelf}")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/MergeShelves").
		Methods("POST").
		Path("/v1/shelves/{shelf}:merge")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/CreateBook").
		Methods("POST").
		Path("/v1/shelves/{shelf}/books")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/GetBook").
		Methods("GET").
		Path("/v1/shelves/{shelf}/books/{book}")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/ListBooks").
		Methods("GET").
		Path("/v1/shelves/{shelf}/books")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/DeleteBook").
		Methods("DELETE").
		Path("/v1/shelves/{shelf}/books/{book}")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/UpdateBook").
		Methods("PATCH").
		Path("/v1/shelves/{shelf}/books/{book}")
	router.NewRoute().
		Name("/google.example.library.v1.LibraryService/MoveBook").
		Methods("POST").
		Path("/v1/shelves/{shelf}/books/{book}:move")
	return &libraryServiceHTTPClient{
		createShelf: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*CreateShelfRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var bodyBuf bytes.Buffer
					if err := jsonx.NewEncoder(&bodyBuf).Encode(req.GetShelf()); err != nil {
						return nil, err
					}
					body = &bodyBuf
					contentType := "application/json; charset=utf-8"
					var pairs []string
					path, err := router.Get("/google.example.library.v1.LibraryService/CreateShelf").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "POST", target.String(), body)
					if err != nil {
						return nil, err
					}
					r.Header.Set("Content-Type", contentType)
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &Shelf{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		getShelf: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*GetShelfRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetName()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 2 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1])
					path, err := router.Get("/google.example.library.v1.LibraryService/GetShelf").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &Shelf{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		listShelves: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*ListShelvesRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					path, err := router.Get("/google.example.library.v1.LibraryService/ListShelves").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					queries["page_size"] = append(queries["page_size"], strconvx.FormatInt(req.GetPageSize(), 10))
					queries["page_token"] = append(queries["page_token"], req.GetPageToken())
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &ListShelvesResponse{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		deleteShelf: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*DeleteShelfRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetName()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 2 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1])
					path, err := router.Get("/google.example.library.v1.LibraryService/DeleteShelf").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "DELETE", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &emptypb.Empty{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		mergeShelves: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*MergeShelvesRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var bodyBuf bytes.Buffer
					if err := jsonx.NewEncoder(&bodyBuf).Encode(req); err != nil {
						return nil, err
					}
					body = &bodyBuf
					contentType := "application/json; charset=utf-8"
					var pairs []string
					namedPathParameter := req.GetName()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 2 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1])
					path, err := router.Get("/google.example.library.v1.LibraryService/MergeShelves").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "POST", target.String(), body)
					if err != nil {
						return nil, err
					}
					r.Header.Set("Content-Type", contentType)
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &Shelf{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		createBook: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*CreateBookRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var bodyBuf bytes.Buffer
					if err := jsonx.NewEncoder(&bodyBuf).Encode(req.GetBook()); err != nil {
						return nil, err
					}
					body = &bodyBuf
					contentType := "application/json; charset=utf-8"
					var pairs []string
					namedPathParameter := req.GetParent()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 2 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1])
					path, err := router.Get("/google.example.library.v1.LibraryService/CreateBook").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "POST", target.String(), body)
					if err != nil {
						return nil, err
					}
					r.Header.Set("Content-Type", contentType)
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &Book{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		getBook: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*GetBookRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetName()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 4 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1], "book", namedPathValues[3])
					path, err := router.Get("/google.example.library.v1.LibraryService/GetBook").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &Book{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		listBooks: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*ListBooksRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetParent()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 2 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1])
					path, err := router.Get("/google.example.library.v1.LibraryService/ListBooks").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					queries["page_size"] = append(queries["page_size"], strconvx.FormatInt(req.GetPageSize(), 10))
					queries["page_token"] = append(queries["page_token"], req.GetPageToken())
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &ListBooksResponse{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		deleteBook: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*DeleteBookRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var pairs []string
					namedPathParameter := req.GetName()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 4 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1], "book", namedPathValues[3])
					path, err := router.Get("/google.example.library.v1.LibraryService/DeleteBook").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "DELETE", target.String(), body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &emptypb.Empty{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		updateBook: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*UpdateBookRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var bodyBuf bytes.Buffer
					if err := jsonx.NewEncoder(&bodyBuf).Encode(req.GetBook()); err != nil {
						return nil, err
					}
					body = &bodyBuf
					contentType := "application/json; charset=utf-8"
					var pairs []string
					namedPathParameter := req.GetBook().GetName()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 4 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1], "book", namedPathValues[3])
					path, err := router.Get("/google.example.library.v1.LibraryService/UpdateBook").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "PATCH", target.String(), body)
					if err != nil {
						return nil, err
					}
					r.Header.Set("Content-Type", contentType)
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &Book{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
		moveBook: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					if obj == nil {
						return nil, errors.New("request object is nil")
					}
					req, ok := obj.(*MoveBookRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					_ = req
					var body io.Reader
					var bodyBuf bytes.Buffer
					if err := jsonx.NewEncoder(&bodyBuf).Encode(req); err != nil {
						return nil, err
					}
					body = &bodyBuf
					contentType := "application/json; charset=utf-8"
					var pairs []string
					namedPathParameter := req.GetName()
					namedPathValues := strings.Split(namedPathParameter, "/")
					if len(namedPathValues) != 4 {
						return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
					}
					pairs = append(pairs, "shelf", namedPathValues[1], "book", namedPathValues[3])
					path, err := router.Get("/google.example.library.v1.LibraryService/MoveBook").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					queries := url.Values{}
					target := &url.URL{
						Scheme:   scheme,
						Host:     instance,
						Path:     path.Path,
						RawQuery: queries.Encode(),
					}
					r, err := http1.NewRequestWithContext(ctx, "POST", target.String(), body)
					if err != nil {
						return nil, err
					}
					r.Header.Set("Content-Type", contentType)
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					resp := &Book{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
	}
}
