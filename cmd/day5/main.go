package main

import (
	"fmt"
	"slices"
)

func main() {
	orderingRules, updates, err := parseFile("./challenge_data.txt")
	if err != nil {
		panic(err)
	}

	compliantUpdates := []Update{}
	nonCompliantUpdates := []Update{}
	for _, update := range updates {
		if update.isCompliantWithAll(orderingRules) {
			compliantUpdates = append(compliantUpdates, update)
		} else {
			nonCompliantUpdates = append(nonCompliantUpdates, update)
		}
	}

	sumOfMiddlesOfCompliantUpdates := 0
	for _, update := range compliantUpdates {
		middle := update.Pages[len(update.Pages)/2]
		sumOfMiddlesOfCompliantUpdates += middle
	}

	fmt.Printf("sumOfMiddlesOfCompliantUpdates: %v\n", sumOfMiddlesOfCompliantUpdates)

	//===========================

	for _, update := range nonCompliantUpdates {
		update.makeCompliantWithAll(orderingRules)
	}

	sumOfMiddlesOfMadeCompliantUpdates := 0
	for _, update := range nonCompliantUpdates {
		middle := update.Pages[len(update.Pages)/2]
		sumOfMiddlesOfMadeCompliantUpdates += middle
	}

	fmt.Printf("sumOfMiddlesOfMadeCompliantUpdates: %v\n", sumOfMiddlesOfMadeCompliantUpdates)
}

type OrderingRule struct {
	PageBefore int
	PageAfter  int
}

type Update struct {
	Pages []int
}

func (u *Update) isCompliantWith(orderingRule *OrderingRule) bool {
	beforeIndex := slices.Index(u.Pages, orderingRule.PageBefore)
	afterIndex := slices.Index(u.Pages, orderingRule.PageAfter)

	if beforeIndex == -1 || afterIndex == -1 {
		return true
	}

	if beforeIndex < afterIndex {
		return true
	}

	return false
}

func (u *Update) isCompliantWithAll(orderingRules []OrderingRule) bool {
	for _, orderingRule := range orderingRules {
		if !u.isCompliantWith(&orderingRule) {
			return false
		}
	}

	return true
}

func (u *Update) makeCompliantWith(orderingRule *OrderingRule) {
	beforeIndex := slices.Index(u.Pages, orderingRule.PageBefore)
	afterIndex := slices.Index(u.Pages, orderingRule.PageAfter)

	u.Pages = slices.Delete(u.Pages, beforeIndex, beforeIndex+1)
	u.Pages = slices.Insert(u.Pages, afterIndex, orderingRule.PageBefore)
}

func (u *Update) makeCompliantWithAll(orderingRules []OrderingRule) {
	for _, orderingRule := range orderingRules {
		if !u.isCompliantWith(&orderingRule) {
			u.makeCompliantWith(&orderingRule)
			u.makeCompliantWithAll(orderingRules)
			break
		}
	}
}
