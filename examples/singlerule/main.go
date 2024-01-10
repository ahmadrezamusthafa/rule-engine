package main

import (
	"encoding/json"
	"fmt"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine"
)

func main(){
	input := map[string]interface{}{
		"amount":         5000,
		"account_number": "123343242334",
		"remark":         "BFST123456",
	}

	ruleConfig := `
		{
		  "id": 12345,
		  "condition": {
			"logical_operator": "OR",
			"conditions": [
			  {
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
			  }
			]
		  },
		  "actions": [
			{
			  "type": "ReplaceString",
			  "params": {
				"name": "remark",
				"pattern": "BFST([0-9]+).*",
				"replacement": "$1"
			  }
			}
		  ]
		}
	`

	var rule ruleengine.Rule
	_ = json.Unmarshal([]byte(ruleConfig), &rule)

	re := ruleengine.NewRuleEngine()

	output, err := re.ApplyRule(input, rule)
	if err != nil {
		fmt.Println("Error applying rule:", err)
		return
	}

	fmt.Println("Output after applying rule:", output)
}