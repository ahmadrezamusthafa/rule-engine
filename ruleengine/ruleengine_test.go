package ruleengine

import (
	"encoding/json"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/actiontype"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/logicaloperator"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/operator"
	"reflect"
	"testing"
)

func Test_ruleEngine_ApplyRule(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		rule     Rule
		expected interface{}
	}{
		{
			name: "Valid condition and action",
			input: map[string]interface{}{
				"amount":         5000,
				"account_number": "123343242334",
				"remark":         "BFST123456",
			},
			rule: Rule{
				ID: 123,
				Condition: Condition{
					LogicalOperator: logicaloperator.And,
					Conditions: []Condition{
						{
							Name:     "amount",
							Operator: operator.Equals,
							Value:    5000,
						},
						{
							Name:     "account_number",
							Operator: operator.Equals,
							Value:    "123343242334",
						},
						{
							Name:     "remark",
							Operator: operator.Match,
							Value:    "BFST[0-9]+.*",
						},
					},
				},
				Actions: []Action{
					{
						Type: "ReplaceString",
						Params: ActionParams{
							Name:        "remark",
							Pattern:     "BFST([0-9]+).*",
							Replacement: "remark modif",
						},
					},
				},
			},
			expected: "remark modif",
		},
		{
			name: "Valid condition and action 2",
			input: map[string]interface{}{
				"amount":         5000,
				"account_number": "123343242334",
				"remark":         "BFST123456",
			},
			rule: Rule{
				ID: 123,
				Condition: Condition{
					LogicalOperator: logicaloperator.And,
					Conditions: []Condition{
						{
							Name:     "amount",
							Operator: operator.Equals,
							Value:    5000,
						},
						{
							Name:     "account_number",
							Operator: operator.Equals,
							Value:    "123343242334",
						},
						{
							Name:     "remark",
							Operator: operator.Match,
							Value:    "BFST[0-9]+.*",
						},
					},
				},
				Actions: []Action{
					{
						Type: actiontype.ReturnValue,
						Params: ActionParams{
							Value: "overbooking",
						},
					},
				},
			},
			expected: "overbooking",
		},
		{
			name: "Valid condition and action 3",
			input: map[string]interface{}{
				"amount":         5000,
				"account_number": "123343242334",
				"remark":         "BFST123456",
				"bank_id":        "bca",
			},
			rule: Rule{
				ID: 123,
				Condition: Condition{
					LogicalOperator: logicaloperator.And,
					Conditions: []Condition{
						{
							Name:     "amount",
							Operator: operator.Equals,
							Value:    5000,
						},
						{
							Name:     "remark",
							Operator: operator.Match,
							Value:    "BFST[0-9]+.*",
						},
						{
							LogicalOperator: logicaloperator.Or,
							Conditions: []Condition{
								{
									Name:     "bank_id",
									Operator: operator.Equals,
									Value:    "bca",
								},
								{
									Name:     "bank_id",
									Operator: operator.Equals,
									Value:    "bni",
								},
							},
						},
					},
				},
				Actions: []Action{
					{
						Type: actiontype.ReturnValue,
						Params: ActionParams{
							Value: "overbooking",
						},
					},
				},
			},
			expected: "overbooking",
		},
		{
			name: "Valid condition and action 4",
			input: map[string]interface{}{
				"amount":         5000,
				"account_number": "13131",
				"remark":         "BFCST123456",
				"bank_id":        "bri",
			},
			rule: Rule{
				ID: 123,
				Condition: Condition{
					LogicalOperator: logicaloperator.Or,
					Conditions: []Condition{
						{
							Name:     "amount",
							Operator: operator.Equals,
							Value:    5000,
						},
						{
							Name:     "remark",
							Operator: operator.Match,
							Value:    "BFST[0-9]+.*",
						},
						{
							LogicalOperator: logicaloperator.Or,
							Conditions: []Condition{
								{
									Name:     "bank_id",
									Operator: operator.Equals,
									Value:    "bca",
								},
								{
									Name:     "bank_id",
									Operator: operator.Equals,
									Value:    "bni",
								},
							},
						},
					},
				},
				Actions: []Action{
					{
						Type: actiontype.ReturnValue,
						Params: ActionParams{
							Value: "overbooking",
						},
					},
				},
			},
			expected: "overbooking",
		},
		{
			name: "Valid condition and action 5",
			input: map[string]interface{}{
				"amount":         4000,
				"account_number": "13131",
				"remark":         "BFCST123456",
				"bank_id":        "bni",
			},
			rule: Rule{
				ID: 123,
				Condition: Condition{
					LogicalOperator: logicaloperator.Or,
					Conditions: []Condition{
						{
							Name:     "amount",
							Operator: operator.Equals,
							Value:    5000,
						},
						{
							Name:     "remark",
							Operator: operator.Match,
							Value:    "BFST[0-9]+.*",
						},
						{
							LogicalOperator: logicaloperator.Or,
							Conditions: []Condition{
								{
									Name:     "bank_id",
									Operator: operator.Equals,
									Value:    "bca",
								},
								{
									Name:     "bank_id",
									Operator: operator.Equals,
									Value:    "bni",
								},
							},
						},
					},
				},
				Actions: []Action{
					{
						Type: actiontype.ReturnValue,
						Params: ActionParams{
							Value: "overbooking",
						},
					},
				},
			},
			expected: "overbooking",
		},
		{
			name: "Valid condition and action 6",
			input: map[string]interface{}{
				"amount":         4000,
				"account_number": "13131",
				"remark":         "BFCST123456",
				"bank_id":        "bni",
			},
			rule: Rule{
				ID: 123,
				Condition: Condition{
					LogicalOperator: logicaloperator.Or,
					Conditions: []Condition{
						{
							Name:     "amount",
							Operator: operator.Equals,
							Value:    5000,
						},
						{
							Name:     "remark",
							Operator: operator.Match,
							Value:    "BFST[0-9]+.*",
						},
						{
							LogicalOperator: logicaloperator.And,
							Conditions: []Condition{
								{
									Name:     "bank_id",
									Operator: operator.Equals,
									Value:    "bni",
								},
								{
									Name:     "amount",
									Operator: operator.Equals,
									Value:    4000,
								},
							},
						},
					},
				},
				Actions: []Action{
					{
						Type: actiontype.ReturnValue,
						Params: ActionParams{
							Value: "overbooking",
						},
					},
				},
			},
			expected: "overbooking",
		},
		{
			name: "Valid condition without action",
			input: map[string]interface{}{
				"amount":         5000,
				"account_number": "123343242334",
				"remark":         "BFST123456",
			},
			rule: Rule{
				ID: 123,
				Condition: Condition{
					LogicalOperator: logicaloperator.And,
					Conditions: []Condition{
						{
							Name:     "amount",
							Operator: operator.Equals,
							Value:    5000,
						},
						{
							Name:     "account_number",
							Operator: operator.Equals,
							Value:    "123343242334",
						},
						{
							Name:     "remark",
							Operator: operator.Match,
							Value:    "BFST[0-9]+.*",
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Invalid condition without action",
			input: map[string]interface{}{
				"amount":         700,
				"account_number": "123343242334",
				"remark":         "BFST123456",
			},
			rule: Rule{
				ID: 123,
				Condition: Condition{
					LogicalOperator: logicaloperator.And,
					Conditions: []Condition{
						{
							Name:     "amount",
							Operator: operator.Equals,
							Value:    5000,
						},
						{
							Name:     "account_number",
							Operator: operator.Equals,
							Value:    "123343242334",
						},
						{
							Name:     "remark",
							Operator: operator.Match,
							Value:    "BFST[0-9]+.*",
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "Valid one condition with action",
			input: map[string]interface{}{
				"amount":         5000,
				"account_number": "123343242334",
				"remark":         "BFST123456",
			},
			rule: Rule{
				ID: 123,
				Condition: Condition{
					LogicalOperator: logicaloperator.And,
					Conditions: []Condition{
						{
							Name:     "amount",
							Operator: operator.GreaterThan,
							Value:    2000,
						},
					},
				},
				Actions: []Action{
					{
						Type: "ReplaceString",
						Params: ActionParams{
							Name:        "remark",
							Pattern:     "BFST([0-9]+).*",
							Replacement: "remark modif",
						},
					},
				},
			},
			expected: "remark modif",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := NewRuleEngine()
			output, err := re.ApplyRule(tt.input, tt.rule)
			if err != nil {
				t.Fatalf("Error applying rule: %v", err)
			}
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("Unexpected output. Expected: %v, Got: %v", tt.expected, output)
			}
		})
	}
}

func Test_ruleEngine_ApplyRuleSet(t *testing.T) {
	var tests = []struct {
		name            string
		input           map[string]interface{}
		ruleSet         string
		ruleEngine      RuleEngine
		expectedDetails *ruleEngine
		expected        interface{}
	}{
		{
			name: "Valid ruleset",
			input: map[string]interface{}{
				"amount":         5000,
				"account_number": "123343242334",
				"remark":         "BFST123456",
			},
			ruleSet:    `{"logical_operator":"OR","rules":[{"id":1,"condition":{"logical_operator":"AND","conditions":[{"name":"amount","operator":"greater_than","value":2000}]},"actions":[{"type":"ReplaceString","params":{"name":"remark","pattern":"BFST([0-9]+).*","replacement":"remark modif"}}]}]}`,
			ruleEngine: NewRuleEngine(),
			expectedDetails: &ruleEngine{
				OutputDetails: map[string]interface{}{
					"1": "remark modif",
				},
			},
			expected: true,
		},
		{
			name: "Invalid ruleset",
			input: map[string]interface{}{
				"amount":         1000,
				"account_number": "123343242334",
				"remark":         "BFST123456",
			},
			ruleSet:    `{"logical_operator":"OR","rules":[{"id":1,"condition":{"logical_operator":"AND","conditions":[{"name":"amount","operator":"greater_than","value":2000}]},"actions":[{"type":"ReplaceString","params":{"name":"remark","pattern":"BFST([0-9]+).*","replacement":"remark modif"}}]}]}`,
			ruleEngine: NewRuleEngine(),
			expectedDetails: &ruleEngine{
				OutputDetails: map[string]interface{}{
					"1": false,
				},
			},
			expected: false,
		},
		{
			name: "Valid ruleset - complex",
			input: map[string]interface{}{
				"amount":         5000,
				"account_number": "123343242334",
				"remark":         "BFST123456",
				"bank_id":        "bca",
			},
			ruleSet:    `{"logical_operator":"OR","rules":[{"logical_operator":"AND","rules":[{"id":1,"condition":{"logical_operator":"AND","conditions":[{"name":"amount","operator":"greater_than","value":2000}]},"actions":[{"type":"ReplaceString","params":{"name":"remark","pattern":"BFST([0-9]+).*","replacement":"remark modif"}}]},{"logical_operator":"OR","rules":[{"id":4,"condition":{"logical_operator":"AND","conditions":[{"name":"provider","operator":"equals","value":"telkomsel"}]},"actions":null},{"id":5,"condition":{"logical_operator":"AND","conditions":[{"name":"bank_id","operator":"equals","value":"bca"}]},"actions":null}]}]},{"id":3,"condition":{"logical_operator":"AND","conditions":[{"name":"remark","operator":"equals","value":"wkwkwkwk"}]},"actions":null}]}`,
			ruleEngine: NewRuleEngine(),
			expectedDetails: &ruleEngine{
				OutputDetails: map[string]interface{}{
					"1": "remark modif",
					"3": false,
					"4": false,
					"5": true,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ruleSet RuleSet
			_ = json.Unmarshal([]byte(tt.ruleSet), &ruleSet)

			output, err := tt.ruleEngine.ApplyRuleSet(tt.input, ruleSet)
			if err != nil {
				t.Fatalf("Error applying rule: %v", err)
			}
			if !reflect.DeepEqual(tt.ruleEngine, tt.expectedDetails) {
				t.Errorf("Unexpected details. Expected: %v, Got: %v", tt.expectedDetails, tt.ruleEngine)
			}
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("Unexpected output. Expected: %v, Got: %v", tt.expected, output)
			}
		})
	}
}

func convertSliceRuleToSliceInterface(rules []Rule) []interface{} {
	var interfaces []interface{}
	for _, r := range rules {
		interfaces = append(interfaces, r)
	}

	return interfaces
}
