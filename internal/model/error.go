package model

// NotFound will return an error when a resource is not found
type NotFound string
// NotFound implements the Error interface
func (e NotFound) Error() string {
	return string(e)
}

// InternalServerError will return an error when a server error occurs
type InternalServerError string
// InternalServerError implements the Error interface
func (e InternalServerError) Error() string{
	return string(e)
}

// Forbidden will return an error when trying to access a forbidden resource
type Forbidden string 
// InternalServerError implements the Error interface
func (e Forbidden) Error() string {
	return string(e)
}

// StatusBadRequest will return a error when the client makes a mistakes
type StatusBadRequest string 
// StatusBadRequest implements the Error interface
func (e StatusBadRequest) Error() string {
	return string(e)
}


