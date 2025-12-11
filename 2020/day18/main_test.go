package main

import "testing"

func TestEvaluate(t *testing.T) {
	tests := []struct {
		expr     Expression
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, tt := range tests {
		result := tt.expr.Evaluate()
		if result != tt.expected {
			t.Errorf("Expression(%q).Evaluate() = %d, expected %d", tt.expr, result, tt.expected)
		}
	}
}

func TestEvaluate2(t *testing.T) {
	tests := []struct {
		expr     Expression
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, tt := range tests {
		result := tt.expr.Evaluate2()
		if result != tt.expected {
			t.Errorf("Expression(%q).Evaluate2() = %d, expected %d", tt.expr, result, tt.expected)
		}
	}
}
