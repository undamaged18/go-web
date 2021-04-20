package errors

import (
	"net/http"
	"runtime"
)

// Error - If updating Error you must amend the type assertion in the recovery middleware
// For example: if Code changes to Status, you must change the below line in the recovery middleware
// rec.(utils.Error).Code to rec.(utils.Error).Status
type Error struct {
	Code     int
	Function string
	Line     int
	Message  interface{}
}

// Panic is a custom panic function to pass in the HTTP Status code, function name and error message
//
// Example usage is below
//
//  err = someFunction()
//	if err != nil {
//		errors.Panic(http.StatusBadRequest, errors.FuncTrace(), errors.Trace(), err)
//	}
//
func Panic(code int, function string, line int, message interface{}) {
	e := Error{
		code,
		function,
		line,
		message,
	}
	panic(e)
}

// GetStatusCode converts the HTTP Status Code from a string to an int
//
// The HTTP Status Codes available are defined in Internet Engineering Task Force (IETF) RFC 2616 / RFC 7231.
//
// If code is passed through as an empty string it will default to http.StatusInternalServerError (500)
func GetStatusCode(code string) int {
	errorCode := map[string]int{
		"": http.StatusInternalServerError,
		// Status 2XX codes
		"200": http.StatusOK,
		"201": http.StatusCreated,
		"202": http.StatusAccepted,
		"203": http.StatusNonAuthoritativeInfo,
		"204": http.StatusNoContent,
		"205": http.StatusResetContent,
		"206": http.StatusPartialContent,
		"207": http.StatusMultiStatus,
		"208": http.StatusAlreadyReported,
		"226": http.StatusIMUsed,
		// Status 3XX codes
		"300": http.StatusMultipleChoices,
		"301": http.StatusMovedPermanently,
		"302": http.StatusFound,
		"303": http.StatusSeeOther,
		"304": http.StatusNotModified,
		"305": http.StatusUseProxy,
		"307": http.StatusTemporaryRedirect,
		"308": http.StatusPermanentRedirect,
		// Status 4XX codes
		"400": http.StatusBadRequest,
		"401": http.StatusUnauthorized,
		"402": http.StatusPaymentRequired,
		"403": http.StatusForbidden,
		"404": http.StatusNotFound,
		"405": http.StatusMethodNotAllowed,
		"406": http.StatusNotAcceptable,
		"407": http.StatusProxyAuthRequired,
		"408": http.StatusRequestTimeout,
		"409": http.StatusConflict,
		"410": http.StatusGone,
		"411": http.StatusLengthRequired,
		"412": http.StatusPreconditionFailed,
		"413": http.StatusRequestEntityTooLarge,
		"414": http.StatusRequestURITooLong,
		"415": http.StatusUnsupportedMediaType,
		"416": http.StatusRequestedRangeNotSatisfiable,
		"417": http.StatusExpectationFailed,
		"418": http.StatusTeapot,
		"421": http.StatusMisdirectedRequest,
		"422": http.StatusUnprocessableEntity,
		"423": http.StatusLocked,
		"424": http.StatusFailedDependency,
		"426": http.StatusUpgradeRequired,
		"428": http.StatusPreconditionRequired,
		"429": http.StatusTooManyRequests,
		"431": http.StatusRequestHeaderFieldsTooLarge,
		"451": http.StatusUnavailableForLegalReasons,
		// Status 5XX codes
		"500": http.StatusInternalServerError,
		"501": http.StatusNotImplemented,
		"502": http.StatusBadGateway,
		"503": http.StatusServiceUnavailable,
		"504": http.StatusGatewayTimeout,
		"505": http.StatusHTTPVersionNotSupported,
		"506": http.StatusVariantAlsoNegotiates,
		"507": http.StatusInsufficientStorage,
		"508": http.StatusLoopDetected,
		"510": http.StatusNotExtended,
		"511": http.StatusNetworkAuthenticationRequired,
	}
	return errorCode[code]
}

// Trace returns the file location, line number and function name where it is called
//
// Example: file = /path/to/file/jwt.go line = 20 function = app.controller.IndexController
func Trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	return file, line, fn.Name()
}

// FuncTrace returns the function name where it is called
//
// Example:
//
//Function: app.controller.IndexController Line: 20
func FuncTrace() (string, int) {
	pc, _, line, ok := runtime.Caller(1)
	if !ok {
		return "?", 0
	}

	fn := runtime.FuncForPC(pc)
	return fn.Name(), line
}
