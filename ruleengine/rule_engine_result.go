package ruleengine

type EngineResult struct {
	Valid    bool                   `json:"valid"`
	Actions  []ActionResult         `json:"actions,omitempty"`
	Metadata map[string]interface{} `json:"metadata"`
	Error    string                 `json:"error,omitempty"`
}

type ActionResult struct {
	Type   string        `json:"type"`
	Params *ActionParams `json:"params"`
	Result interface{}   `json:"result"`
}
