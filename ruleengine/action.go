package ruleengine

type Action struct {
	Type   string       `json:"type"`
	Params ActionParams `json:"params"`
}
