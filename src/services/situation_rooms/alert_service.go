package situation_rooms

import "msbda/src/dtos/responses"

func AlertService(id int) ([]responses.SituationRoomAlert, error) {
	alerts := []responses.SituationRoomAlert{
		{
			Type:      "major",
			Title:     "Error occurred during VM initialization in bri_ats_N8jBB92uijlK_vm",
			AlertTime: "11/01/2024 13:54:02 P.M",
		},
		{
			Type:      "minor",
			Title:     "CPU limit has reached 98.12%!",
			AlertTime: "11/01/2024 13:54:02 P.M",
		},
		{
			Type:      "minor",
			Title:     "CPU limit has reached 86.34%!",
			AlertTime: "11/01/2024 13:54:02 P.M",
		},
	}

	return alerts, nil
}
