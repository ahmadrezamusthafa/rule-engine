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

	result := ruleengine.NewRuleEngine().
		RegisterJsonRuleSet(ruleSet).
		Apply(input).GetResult()

	js, _ := json.Marshal(result)
	fmt.Println("Result:", string(js))
}
