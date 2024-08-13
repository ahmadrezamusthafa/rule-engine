package ruleengine

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/actiontype"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/logicaloperator"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/operator"
	"github.com/mitchellh/mapstructure"
	"log"
	"reflect"
	"regexp"
	"strconv"
)

type RuleEngine interface {
	RegisterRule(ruleStr string) Processor
	RegisterRuleSet(ruleSetStr string) Processor
	applyRule(input map[string]interface{}, rule Rule) (interface{}, error)
	applyRuleSet(input map[string]interface{}, ruleSet RuleSet) (result interface{}, err error)
}

type Processor interface {
	Apply(input map[string]interface{}) (interface{}, error)
	GetOutputDetails() map[string]interface{}
}

type engine struct {
	rule          *Rule
	ruleSet       *RuleSet
	OutputDetails map[string]interface{} `json:"output_details"`
}

type processor struct {
	ruleEngine *engine
}

func NewRuleEngine() RuleEngine {
	return &engine{
		OutputDetails: make(map[string]interface{}),
	}
}

func newRuleEngineProcessor(ruleEngine *engine) Processor {
	return &processor{
		ruleEngine: ruleEngine,
	}
}

func (re *engine) RegisterRule(ruleStr string) Processor {
	var rule Rule
	err := json.Unmarshal([]byte(ruleStr), &rule)
	if err != nil {
		return nil
	}
	re.rule = &rule
	return newRuleEngineProcessor(re)
}

func (re *engine) RegisterRuleSet(ruleSetStr string) Processor {
	var ruleSet RuleSet
	err := json.Unmarshal([]byte(ruleSetStr), &ruleSet)
	if err != nil {
		return nil
	}
	re.ruleSet = &ruleSet
	return newRuleEngineProcessor(re)
}

func (p *processor) Apply(input map[string]interface{}) (result interface{}, err error) {
	if p.ruleEngine.rule != nil {
		return p.ruleEngine.applyRule(input, *p.ruleEngine.rule)
	} else if p.ruleEngine.ruleSet != nil {
		return p.ruleEngine.applyRuleSet(input, *p.ruleEngine.ruleSet)
	}
	return nil, errors.New("rule and rule set are empty, please specify one")
}

func (p *processor) GetOutputDetails() map[string]interface{} {
	return p.ruleEngine.OutputDetails
}

func (re *engine) applyRuleSet(input map[string]interface{}, ruleSet RuleSet) (result interface{}, err error) {
	if ruleSet.LogicalOperator == "" {
		ruleSet.LogicalOperator = logicaloperator.And
	}

	if ruleSet.LogicalOperator == logicaloperator.And {
		result, err = applyLogicalAnd(input, ruleSet.Rules, re)
	} else if ruleSet.LogicalOperator == logicaloperator.Or {
		result, err = applyLogicalOr(input, ruleSet.Rules, re)
	}

	if result == true {
		for _, action := range ruleSet.Actions {
			result = applyAction(input, action)
		}
	}

	return result, err
}

func (re *engine) applyRule(input map[string]interface{}, rule Rule) (result interface{}, err error) {
	if rule.Condition.LogicalOperator == "" {
		rule.Condition.LogicalOperator = logicaloperator.And
	}

	result = evaluateConditions(input, rule.Condition)

	id := fmt.Sprint(rule.ID)
	if _, ok := re.OutputDetails[id]; !ok {
		re.OutputDetails[id] = result
	}

	return
}

func applyLogicalAnd(input map[string]interface{}, rules []interface{}, re *engine) (bool, error) {
	result := true
	for _, nestedRule := range rules {
		switch r := nestedRule.(type) {
		case map[string]interface{}:
			if _, ok := r["rules"]; ok {
				ruleSetResult, err := applyMapRuleSet(input, r, re)
				if err != nil {
					return false, err
				}
				switch r := ruleSetResult.(type) {
				case bool:
					result = result && r
				}
			} else {
				ruleResult, err := applyMapRule(input, r, re)
				if err != nil {
					return false, err
				}
				result = result && ruleResult
			}
		case Rule:
			ruleResult, ruleErr := re.applyRule(input, r)
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

func applyLogicalOr(input map[string]interface{}, rules []interface{}, re *engine) (bool, error) {
	result := false
	for _, nestedRule := range rules {
		switch r := nestedRule.(type) {
		case map[string]interface{}:
			if _, ok := r["rules"]; ok {
				ruleSetResult, err := applyMapRuleSet(input, r, re)
				if err != nil {
					return false, err
				}
				switch r := ruleSetResult.(type) {
				case bool:
					result = result || r
				}
			} else {
				ruleResult, err := applyMapRule(input, r, re)
				if err != nil {
					return false, err
				}
				result = result || ruleResult
			}
		case Rule:
			ruleResult, ruleErr := re.applyRule(input, r)
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

func applyMapRuleSet(input map[string]interface{}, ruleMap map[string]interface{}, re *engine) (interface{}, error) {
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

	return re.applyRuleSet(input, ruleSet)
}

func applyMapRule(input map[string]interface{}, ruleMap map[string]interface{}, re *engine) (bool, error) {
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

	ruleResult, err := re.applyRule(input, rule)
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

func applyAction(input map[string]interface{}, action Action) (result interface{}) {
	switch action.Type {
	case actiontype.ReplaceString:
		params := action.Params
		result = regexp.MustCompile(params.Pattern).ReplaceAllString(input[params.Name].(string), params.Replacement)
		input[params.Name] = result
	case actiontype.ReturnValue:
		params := action.Params
		if v, ok := input[params.Name]; ok {
			result = v
		} else {
			result = params.Replacement
		}
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
	if val, ok := a.(string); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			a = int(intVal)
		}
	}

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
	if val, ok := a.(string); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			a = int(intVal)
		}
	}

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
	if val, ok := a.(string); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			a = int(intVal)
		}
	}

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
	if val, ok := a.(string); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			a = int(intVal)
		}
	}

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
