package status_test

import (
	"errors"
	"github.com/go-leo/status"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/genproto/googleapis/type/phone_number"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"testing"
)

func TestCode(t *testing.T) {
	st := status.Internal()
	assert.Equal(t, codes.Internal, st.Code())
}

func TestGRPCStatus(t *testing.T) {
	text := "password is invalid"
	domain := "account"
	metadata := map[string]string{"username": "leo"}
	number := "13013013013"
	key := "WWW-Authenticate"
	value := "Basic realm=xxx"
	st := status.Unauthenticated(
		status.Message(text),
		status.ErrorInfo(text, domain, metadata),
		status.Headers(http.Header{key: {value}}),
		status.Detail(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)
	grpcStatus := st.GRPCStatus()
	assert.Equal(t, codes.Unauthenticated, grpcStatus.Code())
	assert.Equal(t, text, grpcStatus.Message())
	details := grpcStatus.Details()
	assert.Equal(t, 3, len(details))

	errorInfo := details[0].(*errdetails.ErrorInfo)
	assert.Equal(t, text, errorInfo.GetReason())
	assert.Equal(t, domain, errorInfo.GetDomain())
	assert.Equal(t, metadata, errorInfo.GetMetadata())

	phoneNumber := details[1].(*phone_number.PhoneNumber)
	assert.Equal(t, number, phoneNumber.GetE164Number())

	response := details[2].(*httpstatus.HttpResponse)
	assert.Equal(t, http.StatusUnauthorized, int(response.GetStatus()))
	assert.Equal(t, key, response.GetHeaders()[0].GetKey())
	assert.Equal(t, value, response.GetHeaders()[0].GetValue())
}

func TestHTTPStatus(t *testing.T) {
	text := "password is invalid"
	domain := "account"
	metadata := map[string]string{"username": "leo"}
	number := "13013013013"
	key := "WWW-Authenticate"
	value := "Basic realm=xxx"
	st := status.Unauthenticated(
		status.Message(text),
		status.ErrorInfo(text, domain, metadata),
		status.Headers(http.Header{key: {value}}),
		status.Detail(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)
	httpStatus := st.HTTPStatus()
	assert.Equal(t, http.StatusUnauthorized, int(httpStatus.GetStatus()))
	assert.Equal(t, http.StatusText(http.StatusUnauthorized), httpStatus.GetReason())
	assert.Equal(t, key, httpStatus.GetHeaders()[0].GetKey())
	assert.Equal(t, value, httpStatus.GetHeaders()[0].GetValue())

	body := &status.HttpBody{}
	err := protojson.Unmarshal(httpStatus.GetBody(), body)
	assert.NoErrorf(t, err, "unmarshal http body")
	assert.Equal(t, text, body.GetError().GetMessage())
	assert.Equal(t, int(codes.Unauthenticated), int(body.GetError().GetStatus()))
	assert.Equal(t, http.StatusUnauthorized, int(body.GetError().GetCode()))

	details := body.GetError().GetDetails()
	d1, err := details[0].UnmarshalNew()
	assert.NoError(t, err, "UnmarshalNew")
	errorInfo := d1.(*errdetails.ErrorInfo)
	assert.Equal(t, text, errorInfo.GetReason())
	assert.Equal(t, domain, errorInfo.GetDomain())
	assert.Equal(t, metadata, errorInfo.GetMetadata())

	d2, err := details[1].UnmarshalNew()
	assert.NoError(t, err, "UnmarshalNew")
	phoneNumber := d2.(*phone_number.PhoneNumber)
	assert.Equal(t, number, phoneNumber.GetE164Number())

}

func TestIs(t *testing.T) {
	text := "password is invalid"
	domain := "account"
	metadata := map[string]string{"username": "leo"}
	number := "13013013013"
	key := "WWW-Authenticate"
	value := "Basic realm=xxx"
	st1 := status.Unauthenticated(
		status.Message(text),
		status.ErrorInfo(text, domain, metadata),
		status.Headers(http.Header{key: {value}}),
		status.Detail(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)

	st2 := status.Unauthenticated(
		status.Message(text),
		status.ErrorInfo(text, domain, metadata),
		status.Headers(http.Header{key: {value}}),
		status.Detail(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)

	assert.True(t, errors.Is(st1, st2))

	st3 := status.Internal(
		status.Message(text),
		status.ErrorInfo(text, domain, metadata),
		status.Headers(http.Header{key: {value}}),
		status.Detail(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)
	assert.False(t, errors.Is(st1, st3))
}

func TestCodeEquals(t *testing.T) {
	text := "password is invalid"
	domain := "account"
	metadata := map[string]string{"username": "leo"}
	number := "13013013013"
	key := "WWW-Authenticate"
	value := "Basic realm=xxx"
	st1 := status.Unauthenticated(
		status.Message(text),
		status.ErrorInfo(text, domain, metadata),
		status.Headers(http.Header{key: {value}}),
		status.Detail(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)

	st2 := status.Unauthenticated()
	assert.True(t, st1.CodeEquals(st2))

	st3 := status.Internal(
		status.Message(text),
		status.ErrorInfo(text, domain, metadata),
		status.Headers(http.Header{key: {value}}),
		status.Detail(&phone_number.PhoneNumber{
			Kind: &phone_number.PhoneNumber_E164Number{
				E164Number: number,
			},
		}),
	)
	assert.False(t, st1.CodeEquals(st3))
}
