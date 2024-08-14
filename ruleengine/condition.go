package ruleengine

type Condition struct {
	LogicalOperator string      `json:"logical_operator,omitempty"`
	Conditions      []Condition `json:"conditions,omitempty"`
	Name            string      `json:"name,omitempty"`
	Operator        string      `json:"operator,omitempty"`
	Value           interface{} `json:"value,omitempty"`
}

func NewCondition(name string, operator string, value interface{}) Condition {
	return Condition{Name: name, Operator: operator, Value: value}
}

func NewGroupCondition(logicalOperator string, conditions ...Condition) Condition {
	return Condition{LogicalOperator: logicalOperator, Conditions: conditions}
}
