package ruleengine

import (
	"errors"
	"fmt"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/actiontype"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/logicaloperator"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/operator"
	"github.com/mitchellh/mapstructure"
	"log"
	"reflect"
	"regexp"
)

type RuleEngine interface {
	ApplyRule(input map[string]interface{}, rule Rule) (interface{}, error)
	ApplyRuleSet(input map[string]interface{}, ruleSet RuleSet) (result bool, err error)
}

type ruleEngine struct {
	OutputDetails map[string]interface{} `json:"output_details"`
}

func NewRuleEngine() RuleEngine {
	return &ruleEngine{
		OutputDetails: make(map[string]interface{}, 0),
	}
}

func (re *ruleEngine) ApplyRuleSet(input map[string]interface{}, ruleSet RuleSet) (result bool, err error) {
	if ruleSet.LogicalOperator == "" {
		ruleSet.LogicalOperator = logicaloperator.And
	}

	if ruleSet.LogicalOperator == logicaloperator.And {
		result, err = applyLogicalAnd(input, ruleSet.Rules, re)
	} else if ruleSet.LogicalOperator == logicaloperator.Or {
		result, err = applyLogicalOr(input, ruleSet.Rules, re)
	}

	return result, err
}

func applyLogicalAnd(input map[string]interface{}, rules []interface{}, re *ruleEngine) (bool, error) {
	result := true
	for _, nestedRule := range rules {
		switch r := nestedRule.(type) {
		case map[string]interface{}:
			if _, ok := r["rules"]; ok {
				ruleSetResult, err := applyMapRuleSet(input, r, re)
				if err != nil {
					return false, err
				}
				result = result && ruleSetResult
			} else {
				ruleResult, err := applyMapRule(input, r, re)
				if err != nil {
					return false, err
				}
				result = result && ruleResult
			}
		case Rule:
			ruleResult, ruleErr := re.ApplyRule(input, r)
			if ruleErr != nil {
				return false, ruleErr
			}

			switch r := ruleResult.(type) {
			case bool:
				result = result && r
			}
		default:
			return false, errors.New(fmt.Sprintf("invalid nested rule type: %s", reflect.TypeOf(nestedRule)))
		}
	}

	return result, nil
}

func applyLogicalOr(input map[string]interface{}, rules []interface{}, re *ruleEngine) (bool, error) {
	result := false
	for _, nestedRule := range rules {
		switch r := nestedRule.(type) {
		case map[string]interface{}:
			if _, ok := r["rules"]; ok {
				ruleSetResult, err := applyMapRuleSet(input, r, re)
				if err != nil {
					return false, err
				}
				result = result || ruleSetResult
			} else {
				ruleResult, err := applyMapRule(input, r, re)
				if err != nil {
					return false, err
				}
				result = result || ruleResult
			}
		case Rule:
			ruleResult, ruleErr := re.ApplyRule(input, r)
			if ruleErr != nil {
				return false, ruleErr
			}

			switch r := ruleResult.(type) {
			case bool:
				result = result || r
			}
		default:
			return false, errors.New(fmt.Sprintf("invalid nested rule type: %s", reflect.TypeOf(nestedRule)))
		}
	}

	return result, nil
}

func applyMapRuleSet(input map[string]interface{}, ruleMap map[string]interface{}, re *ruleEngine) (bool, error) {
	var ruleSet RuleSet
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &ruleSet,
		TagName:  "json",
	}

	decoder, _ := mapstructure.NewDecoder(cfg)
	err := decoder.Decode(ruleMap)
	if err != nil {
		return false, err
	}

	return re.ApplyRuleSet(input, ruleSet)
}

func applyMapRule(input map[string]interface{}, ruleMap map[string]interface{}, re *ruleEngine) (bool, error) {
	var rule Rule
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &rule,
		TagName:  "json",
	}

	decoder, _ := mapstructure.NewDecoder(cfg)
	err := decoder.Decode(ruleMap)
	if err != nil {
		return false, err
	}

	ruleResult, err := re.ApplyRule(input, rule)
	if err != nil {
		return false, err
	}

	switch r := ruleResult.(type) {
	case bool:
		return r, nil
	default:
		return true, nil
	}
}

func (re *ruleEngine) ApplyRule(input map[string]interface{}, rule Rule) (result interface{}, err error) {
	if rule.Condition.LogicalOperator == "" {
		rule.Condition.LogicalOperator = logicaloperator.And
	}

	result = false
	if evaluateConditions(input, rule.Condition) {
		result = true
		for _, action := range rule.Actions {
			result = applyAction(input, action)
		}
	}

	id := fmt.Sprint(rule.ID)
	if _, ok := re.OutputDetails[id]; !ok {
		re.OutputDetails[id] = result
	}

	return
}

func applyAction(input map[string]interface{}, action Action) (result interface{}) {
	switch action.Type {
	case actiontype.ReplaceString:
		params := action.Params
		result = regexp.MustCompile(params.Pattern).ReplaceAllString(input[params.Name].(string), params.Replacement)
	case actiontype.ReturnValue:
		params := action.Params
		result = params.Value
	default:
		// Ignore unknown action
	}

	return
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
	default:
		log.Fatal("Invalid condition operator: ", condition.Operator)
	}

	return false
}

func isEqual(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		if val, ok := b.(int); ok {
			return a == val
		}
		if val, ok := b.(float64); ok {
			return a == int(val)
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
		if val, ok := b.(float64); ok {
			return a > int(val)
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
		if val, ok := b.(float64); ok {
			return a >= int(val)
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
		if val, ok := b.(float64); ok {
			return a < int(val)
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
		if val, ok := b.(float64); ok {
			return a <= int(val)
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
		if val, ok := b.(float64); ok {
			return a != int(val)
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
