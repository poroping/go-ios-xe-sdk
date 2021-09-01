package models

type BgpNeighbor struct {
	Neighbor Neighbor `json:"ned:neighbor"`
}

type BgpNeighborList struct {
	Collection struct {
		Neighbor []Neighbor `json:"ned:neighbor"`
	} `json:"collection"`
}

type Neighbor struct {
	ID       string `json:"id,omitempty"`
	RemoteAs int    `json:"remote-as"`
}
