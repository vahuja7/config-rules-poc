package main

import rule_engine "github.com/swiggy-private/rule-engine"

//meta - contains information to be displayed on UI

type Rule struct {
	Type string
	Rule []rule_engine.ASTNode
	Meta string
}

// will be evaluating my rule on this struct

type InputData struct {
	RestaurantId      int
	CouponName        string
	MinimumOrderValue int
}

// create new rule input

type NewRuleInput struct {
	Type       string
	CouponCode string
	MOV        int
}
