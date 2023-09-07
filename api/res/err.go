package res

type Error = string

type ErrorList struct {
	Errors []Error `json:"errors"`
}
