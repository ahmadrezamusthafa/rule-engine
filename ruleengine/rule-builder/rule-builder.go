package rulebuilder

import (
	"encoding/json"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine"
)

type Builder struct {
	ruleSet ruleengine.RuleSet
}

func NewRuleSetBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) RegisterParentOperator(logicalOperator string) *Builder {
	b.ruleSet.LogicalOperator = logicalOperator
	return b
}

func (b *Builder) RegisterSubRule(id int, logicalOperator string, conditions []ruleengine.Condition, subConditions ...ruleengine.Condition) *Builder {
	condition := ruleengine.Condition{
		LogicalOperator: logicalOperator,
		Conditions:      conditions,
	}
	condition.Conditions = append(condition.Conditions, subConditions...)
	subRule := ruleengine.Rule{
		ID:        id,
		Condition: condition,
	}
	b.ruleSet.Rules = append(b.ruleSet.Rules, subRule)
	return b
}

func (b *Builder) RegisterAction(actionType string, name, pattern, replacement string) *Builder {
	action := ruleengine.Action{
		Type: actionType,
		Params: &ruleengine.ActionParams{
			Name:        name,
			Pattern:     pattern,
			Replacement: replacement,
		},
	}
	b.ruleSet.Actions = append(b.ruleSet.Actions, action)
	return b
}

func (b *Builder) Build() ruleengine.RuleSet {
	return b.ruleSet
}

func (b *Builder) BuildJson() string {
	data, err := json.MarshalIndent(b.ruleSet, "", "  ")
	if err != nil {
		return ""
	}
	return string(data)
}
