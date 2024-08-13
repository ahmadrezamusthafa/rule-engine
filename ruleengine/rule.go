package ruleengine

type Rule struct {
	ID        int       `json:"id"`
	Condition Condition `json:"condition"`
}
