# Rule Engine

The **rule-engine** project is a lightweight rule engine implemented in Go. It allows you to define rules with
conditions and actions, then apply these rules to input data to determine if they should trigger specific actions.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Rule Configuration](#rule-configuration)
    - [Rule Structure](#rule-structure)
    - [Single Rule Example](#single-rule-example)
    - [Multiple Rules with Actions Example](#multiple-rules-with-actions-example)
3. [How It Works](#how-it-works)
4. [Contributing](#contributing)

## Getting Started

To use the rule engine, you'll need to define rules in a JSON configuration. Each rule consists of conditions and
optional actions. The engine evaluates these rules against input data to decide whether the actions should be executed.

## Rule Configuration

### Rule Structure

A rule is defined using a JSON configuration that includes:

- **id**: Unique identifier for the rule.
- **condition**: Specifies the logical conditions that must be met.
- **actions** (optional): Defines what actions to take if the conditions are met.

### Single Rule Example

**Input**

```
input := map[string]interface{}{
		"amount":         5000,
		"account_number": "123343242334",
		"remark":         "BFST123456",
	}
```

Here’s an example of a JSON configuration for a single rule:

```json
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
    }
  ]
}
```

#### Result

The result of applying the above rule might look like this:

```json
{
  "metadata": {
    "description": "Rule id #1 result is true."
  },
  "valid": true
}
```

## Actions

### Supported Actions

The rule engine supports the following types of actions:

| Action Type     | Description                                                                   |
|-----------------|-------------------------------------------------------------------------------|
| `ReplaceString` | Replaces parts of a string based on a regular expression pattern.             |
| `ReturnValue`   | Returns a value based on the name from the input or uses a replacement value. |

### Action Parameters

| Action Type     | Required Parameters              | Optional Parameters | Description                                                                                                                |
|-----------------|----------------------------------|---------------------|----------------------------------------------------------------------------------------------------------------------------|
| `ReplaceString` | `name`, `pattern`, `replacement` | N/A                 | Replaces occurrences in the `name` value based on the `pattern` with the `replacement` value.                              |
| `ReturnValue`   | `name`                           | `replacement`       | Returns the value associated with `name` from the input. If `replacement` is provided, it will return `replacement` value. |

### Multiple Rules with Actions Example

**Input**

```
input := map[string]interface{}{
		"amount":         5000,
		"account_number": "123343242334",
		"remark":         "BFST123456",
	}
```

You can also define multiple rules and specify actions to be taken if the conditions are met. Here’s an example:

```json
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
```

#### Result

The result of applying the above rules and actions might look like this:

```json
{
  "actions": [
    {
      "params": {
        "name": "remark",
        "pattern": "BFST([0-9]+).*",
        "replacement": "$1"
      },
      "result": "123456",
      "type": "ReplaceString"
    }
  ],
  "metadata": {
    "description": "Rule id #1 result is true. Rule id #2 result is false."
  },
  "valid": true
}
```

## How It Works

1. **Define Rules**: Create a JSON configuration with your rules and actions.
2. **Apply Rules**: Pass input data through the rule engine.
3. **Evaluate**: The engine evaluates the conditions and executes actions if the conditions are met.
4. **Receive Results**: Get the results of the evaluation along with any applied actions.

## Rule Builder

The rule builder feature simplifies the creation and management of rules using a fluent builder pattern. Below are detailed steps for creating different types of rules using the `rule-builder` feature.

### Creating a Simple Rule

**Steps:**

1. **Initialize the Rule Set Builder**

   ```go
   builder := rulebuilder.NewRuleSetBuilder()
   ```

2. **Register the Parent Logical Operator**

   Specify the logical operator for combining rules, e.g., `Or` or `And`.

   ```go
   builder.RegisterParentOperator(logicaloperators.Or)
   ```

3. **Register a Sub-Rule**

   Define a sub-rule with conditions using the `RegisterSubRule` method.

   ```go
   builder.RegisterSubRule(1, logicaloperators.And, []ruleengine.Condition{
       ruleengine.NewCondition("amount", operators.Equals, 5000),
       ruleengine.NewCondition("account_number", operators.Equals, "123343242334"),
   })
   ```

### Creating a Simple Rule with Actions

**Steps:**

1. **Initialize the Rule Set Builder**

   ```go
   builder := rulebuilder.NewRuleSetBuilder()
   ```

2. **Register the Parent Logical Operator**

   ```go
   builder.RegisterParentOperator(logicaloperators.Or)
   ```

3. **Register a Sub-Rule**

   ```go
   builder.RegisterSubRule(1, logicaloperators.And, []ruleengine.Condition{
       ruleengine.NewCondition("amount", operators.Equals, 5000),
       ruleengine.NewCondition("account_number", operators.Equals, "123343242334"),
   })
   ```

4. **Register an Action**

   Define actions to be executed if the conditions are met.

   ```go
   builder.RegisterAction("ReplaceString", "remark", "BFST([0-9]+).*", "$1")
   ```

### Creating Nested Sub-Rules

**Steps:**

1. **Initialize the Rule Set Builder**

   ```go
   builder := rulebuilder.NewRuleSetBuilder()
   ```

2. **Register the Parent Logical Operator**

   ```go
   builder.RegisterParentOperator(logicaloperators.Or)
   ```

3. **Register Nested Sub-Rules**

   ```go
   builder.RegisterSubRule(1, logicaloperators.And, []ruleengine.Condition{
       ruleengine.NewCondition("amount", operators.Equals, 5000),
       ruleengine.NewCondition("account_number", operators.Equals, "123343242334"),
       ruleengine.NewCondition("remark", operators.Match, "BFST[0-9]+.*"),
       ruleengine.NewGroupCondition(logicaloperators.Or,
           ruleengine.NewCondition("amount1", operators.Equals, 123),
           ruleengine.NewCondition("amount2", operators.Equals, 567),
           ruleengine.NewGroupCondition(logicaloperators.Or,
               ruleengine.NewCondition("amount1", operators.Equals, 123),
               ruleengine.NewCondition("amount2", operators.Equals, 78978))),
   })
   ```

4. **Register an Action**

   ```go
   builder.RegisterAction("ReplaceString", "remark", "BFST([0-9]+).*", "$1")
   ```

### Creating Multiple Sub-Rules

**Steps:**

1. **Initialize the Rule Set Builder**

   ```go
   builder := rulebuilder.NewRuleSetBuilder()
   ```

2. **Register the Parent Logical Operator**

   ```go
   builder.RegisterParentOperator(logicaloperators.Or)
   ```

3. **Register Multiple Sub-Rules**

   ```go
   builder.RegisterSubRule(1, logicaloperators.And, []ruleengine.Condition{
       ruleengine.NewCondition("amount", operators.Equals, 5000),
       ruleengine.NewCondition("account_number", operators.Equals, "123343242334"),
       ruleengine.NewGroupCondition(logicaloperators.Or,
          ruleengine.NewCondition("amount1", operators.Equals, 123),
          ruleengine.NewCondition("amount2", operators.Equals, 567)),
   })
   
   builder.RegisterSubRule(2, logicaloperators.And, []ruleengine.Condition{
	   ruleengine.NewCondition("remark", operators.Equals, "hahaha"),
   })
   ```

4. **Register an Action**

   ```go
   builder.RegisterAction("ReplaceString", "remark", "BFST([0-9]+).*", "$1")
   ```

### Example
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine"
	logicaloperators "github.com/ahmadrezamusthafa/rule-engine/ruleengine/logical-operator"
	"github.com/ahmadrezamusthafa/rule-engine/ruleengine/operators"
	rulebuilder "github.com/ahmadrezamusthafa/rule-engine/ruleengine/rule-builder"
)

func main() {
	input := map[string]interface{}{
		"amount":         5000,
		"account_number": "123343242334",
		"remark":         "BFST123456",
		"amount1":        123,
	}

	builder := rulebuilder.NewRuleSetBuilder().
		RegisterParentOperator(logicaloperators.Or).
		RegisterSubRule(1, logicaloperators.And, []ruleengine.Condition{
			ruleengine.NewCondition("amount", operators.Equals, 5000),
			ruleengine.NewCondition("account_number", operators.Equals, "123343242334"),
			ruleengine.NewCondition("remark", operators.Match, "BFST[0-9]+.*"),
			ruleengine.NewGroupCondition(logicaloperators.Or,
				ruleengine.NewCondition("amount1", operators.Equals, 123),
				ruleengine.NewCondition("amount2", operators.Equals, 567),
				ruleengine.NewGroupCondition(logicaloperators.Or,
					ruleengine.NewCondition("amount1", operators.Equals, 123),
					ruleengine.NewCondition("amount2", operators.Equals, 78978))),
		}).
		RegisterSubRule(2, logicaloperators.And, []ruleengine.Condition{
			ruleengine.NewCondition("remark", operators.Equals, "hahaha"),
		}).
		RegisterAction("ReplaceString", "remark", "BFST([0-9]+).*", "$1")

	result := ruleengine.NewRuleEngine().
		RegisterRuleSet(builder.Build()).
		Apply(input).GetResult()

	js, _ := json.Marshal(result)
	fmt.Println("Result:", string(js))
}
```
#### Result
```
{
  "valid": true,
  "actions": [
    {
      "type": "ReplaceString",
      "params": {
        "name": "remark",
        "pattern": "BFST([0-9]+).*",
        "replacement": "$1"
      },
      "result": "123456"
    }
  ],
  "metadata": {
    "description": "Rule id #1 result is true. Rule id #2 result is false."
  }
}
```

## Contributing

Feel free to contribute to this project by opening issues or submitting pull requests.

---

This structure should help users quickly find the information they need and understand how to use the rule engine
effectively.