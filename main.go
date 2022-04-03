package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	ruleengine "github.com/swiggy-private/rule-engine"
)


func main() {

	//sample input to be used to create rules
	input := &NewRuleInput{
		Type:       "SampleRule",
		CouponCode: "SWIGGYIT",
		MOV:        130,
	}

	// create the rule -> template is defined
	myRule := createNewRule(input)


	//input Data -> data on which rule is to be evaluated on
	inputData := InputData{
		RestaurantId:      9990,
		MinimumOrderValue: 131,
		CouponName:        "SWIGGYIT",
	}

	evaluateRule(inputData, myRule)

	//update rule? -> delete existing rule and use createNewRule to create on the basis of existing template
}

// evaluates the rule and returns a result

func evaluateRule(inputData InputData, rule Rule) ruleengine.Result {
	astTree, err := ruleengine.Decode(rule.Rule)
	if err != nil {
		log.Infof("Error while decoding rule %v", err)
	}

	preparedExpression := ruleengine.PreparedExpr{Value: astTree}

	ruleContext := map[string]ruleengine.AttributeContext{
		"input_couponCode": {Value: inputData.CouponName},
		"input_minimumOrderValue": {Value: inputData.MinimumOrderValue},
	}

	log.Infof("Context for prepared expression is : %v", ruleContext)

	res, err := preparedExpression.Execute(ruleContext)
	if err != nil {
		log.Infof("Error : %v", err)
	}
	log.Infof("Result of executing rule is %v", res)

	return res
}

func createNewRule(input *NewRuleInput) Rule {
	var newRule Rule
	newRule.Type = input.Type

	//generate rule on the basis of input
	sampleRule := generateSampleRule(input.CouponCode, input.MOV)

	//contains the serialized input object of the rule
	inputData, err := json.Marshal(input)
	if err != nil {
		log.Infof("Error: %s", err)
	}
	newRule.Meta = string(inputData)

	newRule.Rule = sampleRule

	log.Infof("Generated new rule for INPUT : %v, is RULE : %s", input, newRule)
	return newRule
}

func evaluateTestExpression() {
	exp := "a <= b"
	preparedExpr, _ := ruleengine.PrepareExpr(exp)
	ctx := map[string]ruleengine.AttributeContext{"a": {Value: 6}, "b": {Value: 5}}
	evalResult, _ := preparedExpr.Execute(ctx)
	log.Infof("Result : %v", evalResult)
}
