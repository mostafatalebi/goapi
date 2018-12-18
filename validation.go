package main

import (
	"fmt"
	"net/url"
)

type Validation struct {
	errors         []string
	requiredFields []string
	errorsCnt      int
	status         bool
}

//HasAnyError if there is any error
func (v *Validation) HasAnyError() bool {
	return v.errorsCnt > 0
}

//LastError the last entry of the error
func (v *Validation) LastError() string {
	if v.HasAnyError() {
		return v.errors[v.errorsCnt-1]
	}
	return ""
}

//FirstError the first entry of the error
func (v *Validation) FirstError() string {
	if v.HasAnyError() {
		return v.errors[0]
	}
	return ""
}

//AddError adds an error
func (v *Validation) AddError(msg string) {
	v.errors = append(v.errors, msg)
	v.errorsCnt = v.errorsCnt + 1
}

//Required defines a custom list of required fields to be checked by HasRequiredFields
func (v *Validation) Required(fields ...string) {
	if len(fields) > 0 {
		for _, val := range fields {
			v.requiredFields = append(v.requiredFields, val)
		}
	}
}

//HasRequiredFields checks to see if required fields exists
func (v *Validation) HasRequiredFields(formParams *url.Values) bool {
	if len(v.requiredFields) > 0 {
		for _, vv := range v.requiredFields {
			if formParams.Get(vv) == "" {
				v.AddError(fmt.Sprintf("Field %s is required", vv))
				return false
			}
		}
	}
	return true
}
