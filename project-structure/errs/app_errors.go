package errs

import (
	"fmt"
	"net/http"
)

type Exception struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func (e *Exception) Error() string {
	return e.Detail
}

func newError(code int, defaultMsg string, args ...any) error {
	if len(args) > 0 {
		return &Exception{
			Code:   code,
			Detail: fmt.Sprint(args...),
		}
	}
	return &Exception{
		Code:   code,
		Detail: defaultMsg,
	}
}

func newErrorf(code int, format string, params ...any) error {
	return &Exception{
		Code:   code,
		Detail: fmt.Sprintf(format, params...),
	}
}

// --- Error constructors ---

func UnexpectedError(args ...any) error {
	return newError(http.StatusInternalServerError, "something went wrong", args...)
}
func UnexpectedErrorf(format string, params ...any) error {
	return newErrorf(http.StatusInternalServerError, format, params...)
}

func NotFoundError(args ...any) error {
	return newError(http.StatusNotFound, "data not found", args...)
}
func NotFoundErrorf(format string, params ...any) error {
	return newErrorf(http.StatusNotFound, format, params...)
}

func NoPermissionError(args ...any) error {
	return newError(http.StatusForbidden, "you have no permission to access this resource", args...)
}
func NoPermissionErrorf(format string, params ...any) error {
	return newErrorf(http.StatusForbidden, format, params...)
}

func BadRequestError(args ...any) error {
	return newError(http.StatusBadRequest, "bad request", args...)
}
func BadRequestErrorf(format string, params ...any) error {
	return newErrorf(http.StatusBadRequest, format, params...)
}

func ValidationError(args ...any) error {
	return newError(http.StatusBadRequest, "validation error", args...)
}
func ValidationErrorf(format string, params ...any) error {
	return newErrorf(http.StatusBadRequest, format, params...)
}

func UnauthorizedError(args ...any) error {
	return newError(http.StatusUnauthorized, "unauthorized", args...)
}
func UnauthorizedErrorf(format string, params ...any) error {
	return newErrorf(http.StatusUnauthorized, format, params...)
}

func ForbiddenError(args ...any) error {
	return newError(http.StatusForbidden, "forbidden", args...)
}
func ForbiddenErrorf(format string, params ...any) error {
	return newErrorf(http.StatusForbidden, format, params...)
}

func TokenExpiredError(args ...any) error {
	return newError(http.StatusUnauthorized, "token has expired", args...)
}
func TokenExpiredErrorf(format string, params ...any) error {
	return newErrorf(http.StatusUnauthorized, format, params...)
}

func InvalidTokenError(args ...any) error {
	return newError(http.StatusUnauthorized, "invalid authentication token", args...)
}
func InvalidTokenErrorf(format string, params ...any) error {
	return newErrorf(http.StatusUnauthorized, format, params...)
}

func MissingAuthHeaderError(args ...any) error {
	return newError(http.StatusUnauthorized, "authorization header is required", args...)
}
func MissingAuthHeaderErrorf(format string, params ...any) error {
	return newErrorf(http.StatusUnauthorized, format, params...)
}

func ProxyAuthRequiredError(args ...any) error {
	return newError(http.StatusProxyAuthRequired, "proxy authentication required", args...)
}
func ProxyAuthRequiredErrorf(format string, params ...any) error {
	return newErrorf(http.StatusProxyAuthRequired, format, params...)
}
