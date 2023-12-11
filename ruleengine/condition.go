package ruleengine

type Condition struct {
	LogicalOperator string      `json:"logical_operator,omitempty"`
	Conditions      []Condition `json:"conditions,omitempty"`
	Name            string      `json:"name,omitempty"`
	Operator        string      `json:"operator,omitempty"`
	Value           interface{} `json:"value,omitempty"`
}
