package main

type UserValidation struct {
	Validator Validation
	max       int
	min       int
}

func (v *UserValidation) init() {
	v.Validator = Validation{}
}

func (v *UserValidation) UserName(inp string) *UserValidation {
	if len(inp) < v.min {
		v.Validator.AddError("Username less then the minimum required length")
	} else if len(inp) > v.max {
		v.Validator.AddError("Username exceeding the maximum allowed length")
	}
	return v
}

func (v *UserValidation) Password(inp string) *UserValidation {
	if len(inp) < v.min {
		v.Validator.AddError("Password exceeding the maximum allowed length")
	} else if len(inp) > v.max {
		v.Validator.AddError("Password less then the minimum required length")
	}
	return v
}

func (v *UserValidation) Compare(inp string, inp2 string) *UserValidation {
	if inp != inp2 {
		v.Validator.AddError("Passwords do not match")
	}
	return v
}

func (v *UserValidation) Max(val int) *UserValidation {
	v.max = val
	return v
}

func (v *UserValidation) Min(val int) *UserValidation {
	v.min = val
	return v
}

//Succeed checks to see if there isn't any error after validations
func (v *UserValidation) Succeed() bool {
	return (v.Validator.HasAnyError() == false)
}
