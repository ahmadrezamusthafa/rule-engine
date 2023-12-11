package ruleengine

type Rule struct {
	ID        int       `json:"id"`
	Condition Condition `json:"condition"`
	Actions   []Action  `json:"actions"`
}
