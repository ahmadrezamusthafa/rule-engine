package ruleengine

type ActionParams struct {
	Name        string `json:"name,omitempty"`
	Pattern     string `json:"pattern,omitempty"`
	Replacement string `json:"replacement,omitempty"`
}
