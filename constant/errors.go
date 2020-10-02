package constant

import (
	"gitlab.com/RezaOptic/go-utils/errorsHandler"
	"net/http"
)

const (
	NewError = "new_error"
)

var ErrorsMap = errorsHandler.MapErrors{
	errorsHandler.FA: FAErrorsMap,
	errorsHandler.EN: ENErrorsMap,
}

var FAErrorsMap = errorsHandler.LangErrorMap{
	NewError: "ارور جدید",
}

var ENErrorsMap = errorsHandler.LangErrorMap{
	NewError: "new error",
}

var CodeErrorsMap = errorsHandler.CodeErrorMap{
	NewError: http.StatusBadRequest,
}
