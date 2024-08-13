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

	ruleSet := `
	{
	  "logical_operator": "OR",
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
		  }
		},
		{
		  "id": 2,
		  "condition": {
			"logical_operator": "AND",
			"conditions": [
			  {
				"name": "remark",
				"operator": "equals",
				"value": "hahaha"
			  }
			]
		  }
		}
	  ],
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

	re := ruleengine.NewRuleEngine().RegisterRuleSet(ruleSet)
	result, err := re.Apply(input)
	if err != nil {
		fmt.Println("Error applying rule set:", err)
		return
	}

	js, _ := json.Marshal(re.GetOutputDetails())
	fmt.Println("Detail rule output:", string(js))

	// Display the result after applying the ruleSet
	fmt.Println("Output after applying rule set:", result)
}
