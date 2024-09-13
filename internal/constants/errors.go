package constants

import (
	"errors"
	"net/http"
)

const (
	providerNotFoundErrMsg             = "provider not found"
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
	userIsNotInWhiteListMsg            = "user not allowed to receive courior"
	expiryDateTimeErrMsg               = "expires at can't be for a past time"
	expiryReachedErrMsg                = "message expires"
	pushCouriorNotFoundErrMsg          = "push courior not found"
	couriorNotFoundErrMsg              = "courior not found"
	backOffRetryErrMsg                 = "backoff from retrying"
	wrongStatus                        = "wrong status"
)

var (
	ErrProviderNotFound             = errors.New(providerNotFoundErrMsg)
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
	ErrUserIsNotInWhiteList         = errors.New(userIsNotInWhiteListMsg)
	ErrExpiryDateTime               = errors.New(expiryDateTimeErrMsg)
	ErrExpiryReached                = errors.New(expiryReachedErrMsg)
	ErrPushCouriorNotfound          = errors.New(pushCouriorNotFoundErrMsg)
	ErrCouriorNotFound              = errors.New(couriorNotFoundErrMsg)
	ErrBackOffRetry                 = errors.New(backOffRetryErrMsg)
	ErrWrongStatus                  = errors.New(wrongStatus)
)

const (
	validationErrMsg     = "body request is not valid"
	internalServerErrMsg = "internal server error"
)

var (
	ErrValidation         = errors.New(validationErrMsg)
	ErrValidationCode     = http.StatusBadRequest
	ErrInternalServer     = errors.New(internalServerErrMsg)
	ErrInternalServerCode = http.StatusInternalServerError
)
