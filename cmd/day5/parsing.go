package main

import (
	"strconv"
	"strings"

	"github.com/gornius/aoc24/pkg/fileutils"
)

func parseFile(filePath string) ([]OrderingRule, []Update, error) {
	orderingRules := []OrderingRule{}
	updates := []Update{}

	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, nil, err
	}

	// ensure there is empty line at the end so buf will be always flushed to parts
	lines = append(lines, "")

	buf := []string{}
	parts := [][]string{}

	for _, line := range lines {
		if line == "" {
			parts = append(parts, buf)
			buf = []string{}
			continue
		}
		buf = append(buf, line)
	}

	for _, orderingRuleLine := range parts[0] {
		orderingRule, err := parseOrderingRule(orderingRuleLine)
		if err != nil {
			return nil, nil, err
		}
		orderingRules = append(orderingRules, *orderingRule)
	}

	for _, updateLine := range parts[1] {
		update, err := parseUpdate(updateLine)
		if err != nil {
			return nil, nil, err
		}
		updates = append(updates, *update)
	}

	return orderingRules, updates, nil
}

func parseOrderingRule(orderingRuleLine string) (*OrderingRule, error) {
	parts := strings.Split(orderingRuleLine, "|")

	before, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	after, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &OrderingRule{
		PageBefore: before,
		PageAfter:  after,
	}, nil
}

func parseUpdate(updateLine string) (*Update, error) {
	pages := []int{}
	parts := strings.Split(updateLine, ",")
	for _, part := range parts {
		intVal, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		pages = append(pages, intVal)

	}

	return &Update{
		Pages: pages,
	}, nil
}
