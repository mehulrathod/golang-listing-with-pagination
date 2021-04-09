package api

type EmployeeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Filter struct {
	Id    string
	Value string
}
