Here’s an updated version of your `README.md` with a Table of Contents:

---

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
    "description": "Rule id #1 result is true.",
    "timestamp": "2024-08-14T08:55:52+07:00"
  },
  "valid": true
}
```

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
    "description": "Rule id #1 result is true. Rule id #2 result is false.",
    "timestamp": "2024-08-14T08:52:42+07:00"
  },
  "valid": true
}
```

## How It Works

1. **Define Rules**: Create a JSON configuration with your rules and actions.
2. **Apply Rules**: Pass input data through the rule engine.
3. **Evaluate**: The engine evaluates the conditions and executes actions if the conditions are met.
4. **Receive Results**: Get the results of the evaluation along with any applied actions.

## Contributing

Feel free to contribute to this project by opening issues or submitting pull requests.

---

This structure should help users quickly find the information they need and understand how to use the rule engine
effectively.