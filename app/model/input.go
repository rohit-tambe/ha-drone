package model

// CalculateRequestDTO coordinate request parameter
type CalculateRequestDTO struct {
	X   string `json:"x" validate:"required"`
	Y   string `json:"y" validate:"required"`
	Z   string `json:"z" validate:"required"`
	Vel string `json:"vel" validate:"required"`
}
