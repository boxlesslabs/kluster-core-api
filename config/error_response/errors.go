//=============================================================================
// developer: boxlesslabsng@gmail.com
// Define custom error messages
//=============================================================================
 
/**
 * Define package secrets
 **
 * @struct 
 	NotCreated
	ErrorGetting
	ErrorUpdating
	ErrorDeleting
	DuplicateError
	InvalidLoginCredentials
	NotFound
	ErrorSaving
	ErrorProcessing

 * @return formatted error string
**/

package error_response

import "fmt"

type NotCreated struct {
	Resource string
}

type ErrorGetting struct {
	Resource string
}

type ErrorUpdating struct {
	Resource string
}

type ErrorDeleting struct {
	Resource string
}

type DuplicateError struct {
	Resource string
}

type InvalidLoginCredentials struct {
	Resource string
}

type NotFound struct {
	Resource string
}

type ErrorSaving struct {
	Resource string
}

type ErrorProcessing struct {
	Action string
}

func (e NotCreated) Error() string {
	err := fmt.Sprintf("unable to create %s at this time", e.Resource)
	return err
}

func (e ErrorGetting) Error() string {
	return fmt.Sprintf("unable to get %s at this time", e.Resource)
}

func (e ErrorUpdating) Error() string {
	return fmt.Sprintf("unable to update %s at this time", e.Resource)
}

func (e ErrorDeleting) Error() string {
	return fmt.Sprintf("unable to delete %s at this time", e.Resource)
}

func (e DuplicateError) Error() string {
	return fmt.Sprintf("%s already exsist", e.Resource)
}

func (e InvalidLoginCredentials) Error() string {
	return fmt.Sprintf("invalid email or password")
}

func (e NotFound) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}

func (e ErrorSaving) Error() string {
	return fmt.Sprintf("unable to save %v at this time", e.Resource)
}

func (e ErrorProcessing) Error() string {
	return fmt.Sprintf("unable to %s at this time", e.Action)
}
