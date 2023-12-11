package ruleengine

import (
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/actiontype"
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
					LogicalOperator: "AND",
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
					LogicalOperator: "AND",
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
					LogicalOperator: "AND",
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
							LogicalOperator: "OR",
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
					LogicalOperator: "OR",
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
							LogicalOperator: "OR",
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
					LogicalOperator: "OR",
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
							LogicalOperator: "OR",
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
					LogicalOperator: "OR",
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
							LogicalOperator: "AND",
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
					LogicalOperator: "AND",
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
					LogicalOperator: "AND",
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
