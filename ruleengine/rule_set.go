package ruleengine

type RuleSet struct {
	LogicalOperator string        `json:"logical_operator,omitempty"`
	Rules           []interface{} `json:"rules,omitempty"`
	Actions         []Action      `json:"actions,omitempty"`
}
