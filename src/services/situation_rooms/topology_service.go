package situation_rooms

import "msbda/src/dtos/responses"

func TopologySituation(id int) (*responses.TopologySituation, error) {

	res := responses.TopologySituation{
		Nodes: &responses.NodeTopology{
			Name: "a",
			Childs: []*responses.NodeTopology{
				{
					Name:   "b",
					Childs: nil,
				},
				{
					Name:   "c",
					Childs: nil,
				},
			},
		},
	}

	return &res, nil
}
