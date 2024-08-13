# Rule Engine

The **rule-engine** project is a simple rule engine implemented in Go that allows you to define rules based on certain conditions and apply those rules to input data.


### Rule Configuration

The rule is defined using a JSON configuration that includes an id, condition, and actions. The condition specifies the logical conditions, and the actions define the actions to be taken when the conditions are met.

 The example below demonstrates how to use this rule engine to define a rule and apply it to a given input.
 
#### Example Usage

```go
package main

import (
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

	re := ruleengine.NewRuleEngine().RegisterRule(ruleConfig)
	output, err := re.Apply(input)
	if err != nil {
		fmt.Println("Error applying rule:", err)
		return
	}

	fmt.Println("Output after applying rule:", output)
}
```
#### Output
