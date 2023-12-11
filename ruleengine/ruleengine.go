package ruleengine

import (
	"fmt"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/actiontype"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/logicaloperator"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/operator"
	"regexp"
)

type RuleEngine interface {
	ApplyRule(input map[string]interface{}, rule Rule) (map[string]interface{}, error)
}

type ruleEngine struct{}

func NewRuleEngine() RuleEngine {
	return &ruleEngine{}
}

func (re *ruleEngine) ApplyRule(input map[string]interface{}, rule Rule) (map[string]interface{}, error) {
	if evaluateConditions(input, rule.Condition) {
		for _, action := range rule.Actions {
			applyAction(input, action)
		}
	}

	return input, nil
}

func applyAction(input map[string]interface{}, action Action) {
	switch action.Type {
	case actiontype.ReplaceString:
		params := action.Params
		input[params.Name] = regexp.MustCompile(params.Pattern).ReplaceAllString(input[params.Name].(string), params.Replacement)
	case actiontype.ReturnValue:
		params := action.Params
		input[params.Name] = params.Value
	}
}

func evaluateConditions(input map[string]interface{}, condition Condition) bool {
	if condition.LogicalOperator == logicaloperator.And {
		for _, subCondition := range condition.Conditions {
			if !evaluateConditions(input, subCondition) {
				return false
			}
		}
		return true
	} else if condition.LogicalOperator == logicaloperator.Or {
		for _, subCondition := range condition.Conditions {
			if evaluateConditions(input, subCondition) {
				return true
			}
		}
		return false
	}

	switch condition.Operator {
	case operator.Equals:
		return isEqual(input[condition.Name], condition.Value)
	case operator.GreaterThan:
		return isGreaterThan(input[condition.Name], condition.Value)
	case operator.GreaterThanEquals:
		return isGreaterThanOrEqual(input[condition.Name], condition.Value)
	case operator.LessThan:
		return isLessThan(input[condition.Name], condition.Value)
	case operator.LessThanEquals:
		return isLessThanOrEqual(input[condition.Name], condition.Value)
	case operator.NotEquals:
		return isNotEqual(input[condition.Name], condition.Value)
	case operator.Match:
		match, _ := regexp.MatchString(condition.Value.(string), fmt.Sprintf("%v", input[condition.Name]))
		return match
	}

	return false
}

func isEqual(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		if val, ok := b.(int); ok {
			return a == val
		}
	case float64:
		if val, ok := b.(float64); ok {
			return a == val
		}
	case string:
		if val, ok := b.(string); ok {
			return a == val
		}
	}
	return false
}

func isGreaterThan(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		if val, ok := b.(int); ok {
			return a > val
		}
	case float64:
		if val, ok := b.(float64); ok {
			return a > val
		}
	}
	return false
}

func isGreaterThanOrEqual(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		if val, ok := b.(int); ok {
			return a >= val
		}
	case float64:
		if val, ok := b.(float64); ok {
			return a >= val
		}
	}
	return false
}

func isLessThan(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		if val, ok := b.(int); ok {
			return a < val
		}
	case float64:
		if val, ok := b.(float64); ok {
			return a < val
		}
	}
	return false
}

func isLessThanOrEqual(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		if val, ok := b.(int); ok {
			return a <= val
		}
	case float64:
		if val, ok := b.(float64); ok {
			return a <= val
		}
	}
	return false
}

func isNotEqual(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		if val, ok := b.(int); ok {
			return a != val
		}
	case float64:
		if val, ok := b.(float64); ok {
			return a != val
		}
	case string:
		if val, ok := b.(string); ok {
			return a != val
		}
	}
	return false
}
