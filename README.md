# Rule Engine
Author: Ahmad Reza Musthafa

The **rule-engine** project is a simple rule engine implemented in Go that allows you to define rules based on certain conditions and apply those rules to input data.


### Rule Configuration

The rule is defined using a JSON configuration that includes an id, condition, and actions. The condition specifies the logical conditions, and the actions define the actions to be taken when the conditions are met.

 The example below demonstrates how to use this rule engine to define a rule and apply it to a given input.
 
#### Example Usage

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/ahmadrezamusthafa/rule-engine/ruleengine"
)

func main() {
    // Input data
    input := map[string]interface{}{
        "amount":         5000,
        "account_number": "123343242334",
        "remark":         "BFST123456",
    }

    // Rule configuration in JSON format
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

    // Unmarshal rule configuration
    var rule ruleengine.Rule
    _ = json.Unmarshal([]byte(ruleConfig), &rule)

    // Create a new rule engine
    re := ruleengine.NewRuleEngine()

    // Apply the rule to the input data
    output, err := re.ApplyRule(input, rule)
    if err != nil {
        fmt.Println("Error applying rule:", err)
        return
    }

    // Display the output after applying the rule
    fmt.Println("Output after applying rule:", output)
}
```
