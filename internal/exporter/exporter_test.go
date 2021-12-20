package exporter

import (
	"stp-exporter/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLabelValues(t *testing.T) {
	table := config.Table{
		Name:          "Test",
		Value_index:   3,
		Label_indexes: []int{0, 1, 2},
	}
	values := []interface{}{
		"0",
		"test",
		"example",
		5,
	}
	expected := []string{"0", "test", "example"}

	labelValues := getLabelValues(table, values)

	assert.Equal(t, expected, labelValues)
}

func TestNewCollector(t *testing.T) {

}

func TestGetLabels(t *testing.T) {
	table := config.Table{
		Name:          "Test",
		Value_index:   3,
		Label_indexes: []int{0, 1, 2},
	}
	columns := []string{"Тест", "Перевода", "Колонок"}
	expected := []string{"test", "perevoda", "kolonok"}

	labels := getLabels(table, columns)

	assert.Equal(t, expected, labels)
}
