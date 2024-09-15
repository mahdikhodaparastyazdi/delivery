package constants

import (
	"errors"
	"net/http"
)

const (
	providerIsNotActiveErrMsg          = "provider is not active"
	templateParamsIsRequiredErrMes     = "template params are required"
	couriorProviderErrMsg              = "internal request in provider"
	paramsAreNotProvidedErrMsg         = "params are not provided"
	templateOrProviderNotDefinedErrMsg = "provider of template not defined"
	templateNameMustBeUniqueErrMsg     = "template name is not unique"
	unsupportedCouriorTypeErrMsg       = "unsupported courior type"
	templateNotFoundErrMsg             = "template not found"
	wrongApiKeyErrMsg                  = "wrong api key"
	templateProviderNotDefinedErrMsg   = "provider template not defined"
	deviceInfoNotfoundErrMsg           = "device info not found"
	expiryDateTimeErrMsg               = "expires at can't be for a past time"
	pushCouriorNotFoundErrMsg          = "push courior not found"
	couriorNotFoundErrMsg              = "courior not found"
	wrongStatus                        = "wrong status"
)

var (
	ErrProviderIsNotActive          = errors.New(providerIsNotActiveErrMsg)
	ErrTemplateParamsIsRequired     = errors.New(templateParamsIsRequiredErrMes)
	ErrInternalCouriorProviderError = errors.New(couriorProviderErrMsg)
	ErrParamsAreNotProvided         = errors.New(paramsAreNotProvidedErrMsg)
	ErrTemplateOrProviderNotDefined = errors.New(templateOrProviderNotDefinedErrMsg)
	ErrTemplateNameIsNotUnique      = errors.New(templateNameMustBeUniqueErrMsg)
	ErrUnsupportedCouriorType       = errors.New(unsupportedCouriorTypeErrMsg)
	ErrTemplateNotFound             = errors.New(templateNotFoundErrMsg)
	ErrWrongApiKey                  = errors.New(wrongApiKeyErrMsg)
	ErrTemplateProviderNotDefined   = errors.New(templateProviderNotDefinedErrMsg)
	ErrDeviceInfoNotFound           = errors.New(deviceInfoNotfoundErrMsg)
	ErrExpiryDateTime               = errors.New(expiryDateTimeErrMsg)
	ErrPushCouriorNotfound          = errors.New(pushCouriorNotFoundErrMsg)
	ErrCouriorNotFound              = errors.New(couriorNotFoundErrMsg)
	ErrWrongStatus                  = errors.New(wrongStatus)
)

const (
	validationErrMsg       = "body request is not valid"
	internalServerErrMsg   = "internal server error"
	backOffRetryErrMsg     = "backoff from retrying"
	expiryReachedErrMsg    = "courior start date expires"
	providerNotFoundErrMsg = "provider not found"
)

var (
	ErrValidation       = errors.New(validationErrMsg)
	ErrBackOffRetry     = errors.New(backOffRetryErrMsg)
	ErrExpiryReached    = errors.New(expiryReachedErrMsg)
	ErrProviderNotFound = errors.New(providerNotFoundErrMsg)

	ErrValidationCode     = http.StatusBadRequest
	ErrInternalServer     = errors.New(internalServerErrMsg)
	ErrInternalServerCode = http.StatusInternalServerError
)
