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
	}

	builder := rulebuilder.NewRuleSetBuilder().
		RegisterParentOperator(logicaloperators.Or).
		RegisterSubRule(1, logicaloperators.And, []ruleengine.Condition{
			ruleengine.NewCondition("amount", operators.Equals, 5000),
			ruleengine.NewCondition("account_number", operators.Equals, "123343242334"),
		}).
		RegisterAction("ReplaceString", "remark", "BFST([0-9]+).*", "$1")

	result := ruleengine.NewRuleEngine().
		RegisterRuleSet(builder.Build()).
		Apply(input).GetResult()

	js, _ := json.Marshal(result)
	fmt.Println("Result:", string(js))
}
