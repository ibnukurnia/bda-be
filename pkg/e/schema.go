package e

type SchemaError struct {
	Status  int
	Message string
}

func (e *SchemaError) Error() string {
	return e.Message
}
