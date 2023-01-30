package dtos

type SearchFile struct {
	FileName string `json:"fileName" validate:"required"`
}
