package constants

import (
	"errors"
	"net/http"
)

const (
	alreadyExistErrMsg     = "already exist"
	unexpectedErrMsg       = "unexpected error"
	wrongStatusCodeMsg     = "wrong status code"
	wrongApiKeyErrMsg      = "wrong api key"
	validationErrMsg       = "body request is not valid"
	internalServerErrMsg   = "internal server error"
	backOffRetryErrMsg     = "backoff from retrying"
	expiryReachedErrMsg    = "courior start date expires"
	providerNotFoundErrMsg = "provider not found"
)

var (
	ErrAlreadyExist       = errors.New(alreadyExistErrMsg)
	ErrUnexpected         = errors.New(unexpectedErrMsg)
	ErrWrongApiKey        = errors.New(wrongApiKeyErrMsg)
	ErrValidation         = errors.New(validationErrMsg)
	ErrBackOffRetry       = errors.New(backOffRetryErrMsg)
	ErrExpiryReached      = errors.New(expiryReachedErrMsg)
	ErrProviderNotFound   = errors.New(providerNotFoundErrMsg)
	ErrWrongStatus        = errors.New(wrongStatusCodeMsg)
	ErrValidationCode     = http.StatusBadRequest
	ErrInternalServer     = errors.New(internalServerErrMsg)
	ErrInternalServerCode = http.StatusInternalServerError
)
