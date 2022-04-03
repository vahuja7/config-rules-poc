package main

import (
	rule_engine "github.com/swiggy-private/rule-engine"
	"strconv"
)

//this file contains the structure/template of the rules

//Rule 1 Sample Rule -> prefix notation : array of AST NODES
// sharingRule_couponCode == input_couponName && sharingRule_minimumOrderValue == input_minimumOrderValue
//test expression := "input_couponName == \"SWIGGYIT\" && input_minimumOrderValue >= 120"

//var sampleRule = []rule_engine.ASTNode {
//	{
//		NodeType: rule_engine.ASTNodeType(1),
//		Operator: "==",
//	},
//	{
//		NodeType: rule_engine.ASTNodeType(2),
//		Name : "input_couponCode",
//	},
//	{
//		NodeType: rule_engine.ASTNodeType(3),
//		Kind: "STRING",
//		Value: "\"SWIGGYIT\"",
//	},
//}

func generateSampleRule(couponCode string, mov int) []rule_engine.ASTNode {
	var sampleRule = []rule_engine.ASTNode {
		{
			NodeType: rule_engine.ASTNodeType(1),
			Operator: "&&",
		},
		{
			NodeType: rule_engine.ASTNodeType(1),
			Operator: "==",
		},
		{
			NodeType: rule_engine.ASTNodeType(2),
			Name : "input_couponCode",
		},
		{
			NodeType: rule_engine.ASTNodeType(3),
			Kind: "STRING",
			Value: "\"" + couponCode + "\"",
		},
		{
			NodeType: rule_engine.ASTNodeType(1),
			Operator: ">=",
		},
		{
			NodeType: rule_engine.ASTNodeType(2),
			Name : "input_minimumOrderValue",
		},
		{
			NodeType: rule_engine.ASTNodeType(3),
			Kind: "INT",
			Value: strconv.Itoa(mov),
		},
	}
	return sampleRule
}