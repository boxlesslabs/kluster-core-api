//=============================================================================
// developer: boxlesslabsng@gmail.com
// General utility library
//=============================================================================
 
/**
 **
 * @struct Result
 **
 * @ReturnErrorResult() Returns an error Result objects
 * @ReturnSuccessResult() Returns a success result with a message
 * @ReturnSuccessMessage() Returns only a success message
 * @ReturnValidateError() Returns a validation error message
 * @ReturnBasicResult() Returns an object only
 * @ReturnAuthResult() Returns an authenticated object with token
**/

package utils

type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Error   interface{}      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
	Count   int64       `json:"count,omitempty"`
}

func (util *Result) ReturnErrorResult(error string) Result {
	return Result{
		Success: false,
		Error:   error,
	}
}

func (util *Result) ReturnSuccessResult(data interface{}, message string) Result {
	return Result{
		Success: true,
		Data:    data,
		Message: message,
	}
}

func (util *Result) ReturnSuccessMessage(message string) Result {
	return Result{
		Success: true,
		Message: message,
	}
}

func (util *Result) ReturnValidateError(error interface{}) Result {
	return Result{
		Success: false,
		Error:   error,
	}
}

func (util *Result) ReturnBasicResult(data interface{}) Result {
	return Result{
		Success: true,
		Data:    data,
	}
}

func (util *Result) ReturnAuthResult(data interface{}, token string) Result {
	return Result{
		Success: true,
		Token:   token,
		Data:    data,
	}
}