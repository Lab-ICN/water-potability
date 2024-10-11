package mqtt

import (
	"encoding/json"

	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
)

func parseWaterPotabilityJsonData(data []byte) (domain.WaterPotability, error) {
	var waterPotability domain.WaterPotability
	err := json.Unmarshal(data, &waterPotability)
	if err != nil {
		return domain.WaterPotability{}, err
	}

	return waterPotability, nil
}
