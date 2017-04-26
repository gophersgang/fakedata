package fakedata_test

import (
	"reflect"
	"testing"

	"github.com/lucapette/fakedata/pkg/fakedata"
)

func TestNewColumns(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected fakedata.Columns
	}{
		{name: "one column", input: []string{"email"}, expected: fakedata.Columns{{Key: "email", Name: "email"}}},
		{name: "two columns", input: []string{"email", "domain"}, expected: fakedata.Columns{{Key: "email", Name: "email"}, {Key: "domain", Name: "domain"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := fakedata.NewColumns(tt.input); !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("NewColumns() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestNewColumnsWithName(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected fakedata.Columns
	}{
		{name: "one column", input: []string{"login=email"}, expected: fakedata.Columns{{Key: "email", Name: "login"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := fakedata.NewColumns(tt.input); !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("NewColumns() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestNewColumnsWithSpec(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected fakedata.Columns
	}{
		{name: "int full range", input: []string{"int,1..100"}, expected: fakedata.Columns{{Key: "int", Name: "int", Min: "1", Max: "100"}}},
		{name: "int lower bound", input: []string{"int,1.."}, expected: fakedata.Columns{{Key: "int", Name: "int", Min: "1"}}},
		{name: "int lower bound no range syntax", input: []string{"int,10"}, expected: fakedata.Columns{{Key: "int", Name: "int", Min: "10"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := fakedata.NewColumns(tt.input); !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("NewColumns() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}