package utils

import "github.com/vektah/gqlparser/v2/gqlerror"

func ConvertErrorsToGqlError(errs []error) gqlerror.List {
	errorList := gqlerror.List{}

	for _, err := range errs {
		errorList = append(errorList, gqlerror.Wrap(err))
	}

	return errorList
}
