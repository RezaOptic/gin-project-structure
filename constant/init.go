package constant

import "gitlab.com/RezaOptic/go-utils/errorsHandler"

func init() {
	for k, v := range errorsHandler.DefaultFAErrorsMap {
		FAErrorsMap[k] = v
	}
	for k, v := range errorsHandler.DefaultENErrorsMap {
		ENErrorsMap[k] = v
	}
	for k, v := range errorsHandler.DefaultCodeErrorsMap {
		CodeErrorsMap[k] = v
	}
	errors := errorsHandler.ErrorsModel{
		ErrorsMap :ErrorsMap,
		CodeErrorsMap :CodeErrorsMap,
	}
	errorsHandler.Init(errors)
}
