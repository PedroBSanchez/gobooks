package authorHandler

import "fmt"



func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}



//Create Author Request
type CreateAuthorRequest struct {

	Name string `json:"name"`
	Age int64 `json:"age"`

}


func (r *CreateAuthorRequest) Validate() error {


	if r.Age <= 0 && r.Name == "" {
		return fmt.Errorf("param body is empty or malformed")
	}

	if r.Age <= 0 {
		return errParamIsRequired("age", "int64")
	}

	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}


	return nil
}

