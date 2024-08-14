package ruleengine

import "testing"

func BenchmarkApplyRule(b *testing.B) {
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
		  "id": 2,
		  "condition": {
			"logical_operator": "AND",
			"conditions": [
			  {
				"name": "remark",
				"operator": "equals",
				"value": "BFST123456"
			  }
			]
		  }
		}
	  ]
	}
	`

	for i := 0; i < b.N; i++ {
		NewRuleEngine().RegisterRuleSet(ruleSet).Apply(input).GetResult()
	}
}
