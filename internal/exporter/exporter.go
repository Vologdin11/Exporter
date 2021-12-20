package exporter

import (
	"log"
	"stp-exporter/internal/client"
	"stp-exporter/internal/config"
	"stp-exporter/pkg/replacer"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type collector struct {
	config  config.Config
	metrics map[string]*prometheus.Desc
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	token, err := client.GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	for _, table := range c.config.Tables {
		metric, err := client.GetMetric(table.Name, token)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, value := range metric.Values {
			ch <- createMetric(table, c.metrics[table.Name], value)
		}
	}
}

func createMetric(table config.Table, desc *prometheus.Desc, values []interface{}) prometheus.Metric {
	labelValues := getLabelValues(table, values)
	return prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		values[table.Value_index].(float64),
		labelValues...,
	)
}

func getLabelValues(table config.Table, values []interface{}) []string {
	labelValues := make([]string, len(table.Label_indexes))
	for i, index := range table.Label_indexes {
		labelValues[i] = values[index].(string)
	}
	return labelValues
}

func NewCollector() (*collector, error) {
	config, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	token, err := client.GetToken()
	if err != nil {
		return nil, err
	}
	metrics, err := client.GetAllMetrics(config, token)
	if err != nil {
		return nil, err
	}
	collector := collector{
		config:  config,
		metrics: make(map[string]*prometheus.Desc, len(config.Tables)),
	}
	for i, table := range config.Tables {
		labels := getLabels(table, metrics[i].Columns)
		collector.metrics[table.Name] = prometheus.NewDesc(
			strings.ToLower(table.Name),
			table.Name,
			labels,
			nil,
		)
	}
	return &collector, nil
}

//translate table columns on English and return them
func getLabels(table config.Table, columns []string) []string {
	replacer := replacer.NewReplacer()
	labels := make([]string, len(table.Label_indexes))
	for i, index := range table.Label_indexes {
		labels[i] = replacer.Replace(columns[index])
	}
	return labels
}
