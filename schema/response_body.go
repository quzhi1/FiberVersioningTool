package schema

type ResponseBodyVersion1_0 struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CreatedTime int64  `json:"created_time"`
}

type ResponseBodyVersion1_1 struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
