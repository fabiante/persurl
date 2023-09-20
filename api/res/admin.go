package res

type SavePURL struct {
	Target string `json:"target"`
}

type SavePURLResponse struct {
	Path string `json:"path"`
}

func NewSavePURLResponse(path string) *SavePURLResponse {
	return &SavePURLResponse{
		Path: path,
	}
}
