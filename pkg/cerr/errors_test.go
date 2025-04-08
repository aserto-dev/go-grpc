package cerr_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/aserto-dev/go-grpc/pkg/cerr"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

var (
	ErrBoom = errors.New("boom")
	ErrPow  = errors.New("pow")
)

func TestDoubleCerr(t *testing.T) {
	assert := require.New(t)

	err := cerr.ErrAccountNotFound.Err(cerr.ErrRepoAlreadyConnected)

	assert.Contains(err.Error(), "already been connected")
	assert.Contains(err.Error(), "account not found")
}

func TestDoubleCerrWithMsg(t *testing.T) {
	assert := require.New(t)

	err := cerr.ErrAccountNotFound.Err(cerr.ErrRepoAlreadyConnected).Msg("failed to setup")

	assert.Contains(err.Error(), "already been connected")
	assert.Contains(err.Error(), "account not found")
}

func TestWithEmptyMsg(t *testing.T) {
	assert := require.New(t)

	err := cerr.ErrAccountNotFound.Msg("")

	fields := err.Fields()
	assert.Nil(fields["msg"])

	err = cerr.ErrAccountNotFound.Msg("bla")

	fields = err.Fields()
	assert.NotNil(fields["msg"])
}

func TestError(t *testing.T) {
	assert := require.New(t)

	err := cerr.ErrAccountNotFound.Msg("bla").Err(ErrBoom)
	err2 := cerr.ErrAccountNotFound.Msg("bla").Msg("ala")
	err3 := cerr.ErrAccountNotFound.Err(ErrBoom).Msg("bla").Msg("ala")
	err4 := cerr.ErrAccountNotFound.Err(ErrBoom).Err(ErrPow).Msg("bla").Msg("ala")
	err5 := cerr.ErrAccountNotFound.Err(ErrBoom)
	err6 := cerr.ErrAccountNotFound.Err(ErrBoom).Err(ErrPow)
	err7 := cerr.ErrAccountNotFound.Msg("bla")

	assert.ErrorContains(err, "E10012 account not found: bla: boom")
	assert.ErrorContains(err2, "E10012 account not found: bla: ala")
	assert.ErrorContains(err3, "E10012 account not found: bla: ala: boom")
	assert.ErrorContains(err4, "E10012 account not found: bla: ala: boom: pow")
	assert.ErrorContains(err5, "E10012 account not found: boom")
	assert.ErrorContains(err6, "E10012 account not found: boom: pow")
	assert.ErrorContains(err7, "E10012 account not found: bla")
}

func TestWithGrpcStatusCode(t *testing.T) {
	assert := require.New(t)
	err := cerr.ErrAccountNotFound.WithGRPCStatus(codes.Canceled)
	assert.Equal(codes.Canceled, err.StatusCode)
}

func TestWithHttpStatusCode(t *testing.T) {
	assert := require.New(t)
	err := cerr.ErrAccountNotFound.WithHTTPStatus(http.StatusAccepted)
	assert.Equal(http.StatusAccepted, err.HTTPCode)
}
