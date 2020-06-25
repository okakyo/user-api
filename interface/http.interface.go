package interface


/*
	Definition of the Interface in Success and Error Response 
*/
type SuccessResponse struct {
	code: int
	response interface 
}

type ErrorResponse struct {
	code:int
	message: string 
}