package domain

type WaterPotability struct {
	ID                   string  `json:"id"`
	PH                   float64 `json:"ph"`
	TotalDissolvedSolids float64 `json:"totalDissolvedSolids"`
	Turbidity            float64 `json:"turbidity"`
}

type WaterPotabilityWithPrediction struct {
	ID                   string  `json:"id"`
	PH                   float64 `json:"ph"`
	TotalDissolvedSolids float64 `json:"totalDissolvedSolids"`
	Turbidity            float64 `json:"turbidity"`
	Prediction           float64 `json:"prediction"`
	Level                string  `json:"level"`
}
