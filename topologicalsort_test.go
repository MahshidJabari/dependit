package topologicalsort

import (
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		name         string
		dependencies map[string][]string
		expected     []string
		expectError  bool
	}{
		{
			name: "Simple DAG",
			dependencies: map[string][]string{
				"ServiceA": {"ServiceB", "ServiceC"},
				"ServiceB": {"ServiceC"},
				"ServiceC": {},
			},
			expected:    []string{"ServiceC", "ServiceB", "ServiceA"},
			expectError: false,
		},
		{
			name: "DAG with multiple valid topological sorts",
			dependencies: map[string][]string{
				"ServiceA": {"ServiceB"},
				"ServiceB": {},
				"ServiceC": {"ServiceB"},
			},
			expected:    []string{"ServiceB", "ServiceA", "ServiceC"},
			expectError: false,
		},
		{
			name: "Cycle detected",
			dependencies: map[string][]string{
				"ServiceA": {"ServiceB"},
				"ServiceB": {"ServiceA"},
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "Unresolved dependencies",
			dependencies: map[string][]string{
				"ServiceA": {"ServiceB"},
				"ServiceB": {"ServiceC"},
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DefineServices(tt.dependencies)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, but got: %v", err)
				}
				// Check if the result matches the expected output.
				for i, service := range result {
					if service != tt.expected[i] {
						t.Errorf("expected %v, but got %v", tt.expected[i], service)
					}
				}
			}
		})
	}
}

func TestIsDAG(t *testing.T) {
	tests := []struct {
		name         string
		dependencies map[string][]string
		expected     bool
		expectError  bool
	}{
		{
			name: "Simple DAG",
			dependencies: map[string][]string{
				"ServiceA": {"ServiceB", "ServiceC"},
				"ServiceB": {"ServiceC"},
				"ServiceC": {},
			},
			expected:    true,
			expectError: false,
		},
		{
			name: "DAG with multiple valid topological sorts",
			dependencies: map[string][]string{
				"ServiceA": {"ServiceB"},
				"ServiceB": {},
				"ServiceC": {"ServiceB"},
			},
			expected:    true,
			expectError: false,
		},
		{
			name: "Cycle detected",
			dependencies: map[string][]string{
				"ServiceA": {"ServiceB"},
				"ServiceB": {"ServiceA"},
			},
			expected:    false,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CreateAndCheckDAG(tt.dependencies)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, but got: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %v, but got %v", tt.expected, result)
				}
			}
		})
	}
}
