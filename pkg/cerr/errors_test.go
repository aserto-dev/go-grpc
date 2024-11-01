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

	assert.Equal(err.Error(), "E10012 account not found: boom: bla")
	assert.Equal(err2.Error(), "E10012 account not found: bla: ala")
	assert.Equal(err3.Error(), "E10012 account not found: boom: bla: ala")
	assert.Equal(err4.Error(), "E10012 account not found: boom: pow: bla: ala")
	assert.Equal(err5.Error(), "E10012 account not found: boom")
	assert.Equal(err6.Error(), "E10012 account not found: boom: pow")
	assert.Equal(err7.Error(), "E10012 account not found: bla")
}

func TestWithGrpcStatusCode(t *testing.T) {
	assert := require.New(t)
	err := cerr.ErrAccountNotFound.WithGRPCStatus(codes.Canceled)
	assert.Equal(err.StatusCode, codes.Canceled)
}

func TestWithHttpStatusCode(t *testing.T) {
	assert := require.New(t)
	err := cerr.ErrAccountNotFound.WithHTTPStatus(http.StatusAccepted)
	assert.Equal(err.HTTPCode, http.StatusAccepted)
}
