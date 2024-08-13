package ruleengine

import (
	"encoding/json"
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
			name: "Valid condition",
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
			name: "Valid condition 2",
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
			name: "Valid condition 3",
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
			},
			expected: true,
		},
		{
			name: "Valid condition 4",
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
			},
			expected: true,
		},
		{
			name: "Valid condition 5",
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
			},
			expected: true,
		},
		{
			name: "Valid condition 6",
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
			},
			expected: true,
		},
		{
			name: "Valid condition 7",
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
			name: "Invalid condition 8",
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
			name: "Valid one condition 9",
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
			},
			expected: true,
		},
		{
			name: "Valid one condition",
			input: map[string]interface{}{
				"bank_id":        "gv",
				"account_number": "2193038077",
				"credit":         "150000",
				"debit":          0,
				"remark":         "04515301",
				"description":    "Transfer saldo dari user flip1352297860ed88e3adb4f",
				"transferred_at": "1630038687",
				"va_bank_id":     "bni",
				"va_number":      "8558040402472870",
			},
			rule: Rule{
				ID: 123,
				Condition: Condition{
					LogicalOperator: logicaloperator.And,
					Conditions: []Condition{
						{
							Name:     "credit",
							Operator: operator.GreaterThan,
							Value:    2000,
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Valid without condition",
			input: map[string]interface{}{
				"bank_id":        "gv",
				"account_number": "2193038077",
				"credit":         "150000",
				"debit":          0,
				"remark":         "04515301",
				"description":    "Transfer saldo dari user flip1352297860ed88e3adb4f",
				"transferred_at": "1630038687",
				"va_bank_id":     "bni",
				"va_number":      "8558040402472870",
			},
			rule: Rule{
				ID: 123,
			},
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := NewRuleEngine()
			output, err := re.applyRule(tt.input, tt.rule)
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
		expectedDetails *engine
		expected        interface{}
	}{
		{
			name: "Valid ruleset",
			input: map[string]interface{}{
				"amount":         5000,
				"account_number": "123343242334",
				"remark":         "BFST123456",
			},
			ruleSet:    `{"logical_operator":"OR","rules":[{"id":1,"condition":{"logical_operator":"AND","conditions":[{"name":"amount","operator":"greater_than","value":2000}]}}],"actions":[{"type":"ReplaceString","params":{"name":"remark","pattern":"BFST([0-9]+).*","replacement":"remark modif"}}]}`,
			ruleEngine: NewRuleEngine(),
			expectedDetails: &engine{
				OutputDetails: map[string]interface{}{
					"1": true,
				},
			},
			expected: "remark modif",
		},
		{
			name: "Invalid ruleset",
			input: map[string]interface{}{
				"amount":         1000,
				"account_number": "123343242334",
				"remark":         "BFST123456",
			},
			ruleSet:    `{"logical_operator":"OR","rules":[{"id":1,"condition":{"logical_operator":"AND","conditions":[{"name":"amount","operator":"greater_than","value":2000}]}}],"actions":[{"type":"ReplaceString","params":{"name":"remark","pattern":"BFST([0-9]+).*","replacement":"remark modif 2"}}]}`,
			ruleEngine: NewRuleEngine(),
			expectedDetails: &engine{
				OutputDetails: map[string]interface{}{
					"1": false,
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ruleSet RuleSet
			_ = json.Unmarshal([]byte(tt.ruleSet), &ruleSet)

			output, err := tt.ruleEngine.applyRuleSet(tt.input, ruleSet)
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
