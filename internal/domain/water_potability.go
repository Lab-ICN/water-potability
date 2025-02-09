package domain

type WaterPotability struct {
	Node                 string
	PH                   float64 `json:"ph"`
	TotalDissolvedSolids float64 `json:"totalDissolvedSolids"`
	Turbidity            float64 `json:"turbidity"`
}

type WaterPotabilityWithPrediction struct {
	Node                 string  `json:"node"`
	PH                   float64 `json:"ph"`
	TotalDissolvedSolids float64 `json:"totalDissolvedSolids"`
	Turbidity            float64 `json:"turbidity"`
	Prediction           float64 `json:"prediction"`
}
