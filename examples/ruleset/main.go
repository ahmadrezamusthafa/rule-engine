package main

import (
	"encoding/json"
	"fmt"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine"
)

func main() {
	input := map[string]interface{}{
		"amount":         5000,
		"account_number": "123343242334",
		"remark":         "BFST123456",
	}

	ruleConfig := `
		{
		  "logical_operator": "OR",
		  "rules": [
			{
			  "logical_operator": "AND",
			  "rules": [
				{
				  "id": 1,
				  "condition": {
					"logical_operator": "AND",
					"conditions": [
					  {
						"name": "amount",
						"operator": "equals",
						"value": 5000
					  },
					  {
						"name": "account_number",
						"operator": "equals",
						"value": "123343242334"
					  },
					  {
						"name": "remark",
						"operator": "match",
						"value": "BFST[0-9]+.*"
					  }
					]
				  },
				  "actions": [
					{
					  "type": "ReplaceString",
					  "params": {
						"name": "remark",
						"pattern": "BFST([0-9]+).*",
						"replacement": "replace_1"
					  }
					}
				  ]
				},
				{
				  "id": 2,
				  "condition": {
					"logical_operator": "AND",
					"conditions": [
					  {
						"name": "amount",
						"operator": "greater_than",
						"value": 2000
					  }
					]
				  },
				  "actions": [
					{
					  "type": "ReplaceString",
					  "params": {
						"name": "remark",
						"pattern": "BFST([0-9]+).*",
						"replacement": "replace_2"
					  }
					}
				  ]
				}
			  ]
			},
			{
			  "id": 3,
			  "condition": {
				"logical_operator": "AND",
				"conditions": [
				  {
					"name": "amount",
					"operator": "greater_than",
					"value": 1000
				  }
				]
			  },
			  "actions": [
				{
				  "type": "ReplaceString",
				  "params": {
					"name": "remark",
					"pattern": "BFST([0-9]+).*",
					"replacement": "replace_3"
				  }
				}
			  ]
			}
		  ]
		}
	`

	var ruleSet ruleengine.RuleSet
	_ = json.Unmarshal([]byte(ruleConfig), &ruleSet)

	// Create a new rule engine
	re := ruleengine.NewRuleEngine()

	// Apply the ruleSet to the input data
	result, err := re.ApplyRuleSet(input, ruleSet)
	if err != nil {
		fmt.Println("Error applying rule set:", err)
		return
	}

	js, _ := json.Marshal(re)
	fmt.Println("Detail rule output:", string(js))

	// Display the result after applying the ruleSet
	fmt.Println("Output after applying rule set:", result)
}
