package schema

type RequestBodyVersion1_0 struct {
	Name string `json:"name"`
}

type RequestBodyVersion1_1 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
