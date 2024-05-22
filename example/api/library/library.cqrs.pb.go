// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package library

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// LibraryServiceAssembler responsible for completing the transformation between domain model objects and DTOs
type LibraryServiceAssembler interface {
}

// LibraryServiceCQRSService implement the LibraryService service with CQRS pattern
type LibraryServiceCQRSService struct {
	bus       cqrs.Bus
	assembler LibraryServiceAssembler
}

func NewLibraryServiceCQRSService(bus cqrs.Bus, assembler LibraryServiceAssembler) *LibraryServiceCQRSService {
	return &LibraryServiceCQRSService{bus: bus, assembler: assembler}
}

func (svc *LibraryServiceCQRSService) CreateShelf(ctx context.Context, request *CreateShelfRequest) (*Shelf, error) {
	args, ctx, err := svc.assembler.FromCreateShelfRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func (svc *LibraryServiceCQRSService) GetShelf(ctx context.Context, request *GetShelfRequest) (*Shelf, error) {
	args, ctx, err := svc.assembler.FromGetShelfRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func (svc *LibraryServiceCQRSService) ListShelves(ctx context.Context, request *ListShelvesRequest) (*ListShelvesResponse, error) {
	args, ctx, err := svc.assembler.FromListShelvesRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func (svc *LibraryServiceCQRSService) DeleteShelf(ctx context.Context, request *DeleteShelfRequest) (*emptypb.Empty, error) {
	args, ctx, err := svc.assembler.FromDeleteShelfRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func (svc *LibraryServiceCQRSService) MergeShelves(ctx context.Context, request *MergeShelvesRequest) (*Shelf, error) {
	args, ctx, err := svc.assembler.FromMergeShelvesRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func (svc *LibraryServiceCQRSService) CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error) {
	args, ctx, err := svc.assembler.FromCreateBookRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func (svc *LibraryServiceCQRSService) GetBook(ctx context.Context, request *GetBookRequest) (*Book, error) {
	args, ctx, err := svc.assembler.FromGetBookRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func (svc *LibraryServiceCQRSService) ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error) {
	args, ctx, err := svc.assembler.FromListBooksRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func (svc *LibraryServiceCQRSService) DeleteBook(ctx context.Context, request *DeleteBookRequest) (*emptypb.Empty, error) {
	args, ctx, err := svc.assembler.FromDeleteBookRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func (svc *LibraryServiceCQRSService) UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error) {
	args, ctx, err := svc.assembler.FromUpdateBookRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func (svc *LibraryServiceCQRSService) MoveBook(ctx context.Context, request *MoveBookRequest) (*Book, error) {
	args, ctx, err := svc.assembler.FromMoveBookRequest(ctx, request)
	if err != nil {
		return nil, err
	}
}

func NewLibraryServiceBus(
	opts ...cqrs.Option,
) (cqrs.Bus, error) {
	bus := cqrs.NewBus(opts...)
	return bus, nil
}
