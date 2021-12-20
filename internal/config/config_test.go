package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		file, err := os.Create("config.yml")
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()
		defer os.Remove("config.yml")
		_, err = file.Write([]byte(
			`tables:
  - name: test
    value_index: 3
    label_indexes: [0,1,2]`))
		if err != nil {
			t.Fatal(err)
		}
		expected := Config{
			Tables: []Table{
				{
					Name:          "test",
					Value_index:   3,
					Label_indexes: []int{0, 1, 2},
				},
			},
		}

		config, err := GetConfig()

		assert.NoError(t, err)
		assert.Equal(t, expected, config)
	})

	t.Run("error unmarshal config.yml", func(t *testing.T) {
		file, err := os.Create("config.yml")
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()
		defer os.Remove("config.yml")
		_, err = file.Write([]byte(
			`tables:
  - name: test
	value_index: 3`))
		if err != nil {
			t.Fatal(err)
		}
		expected := Config{}

		config, err := GetConfig()

		assert.Error(t, err)
		assert.Equal(t, expected, config)
	})

	//переписать
	t.Run("error write file", func(t *testing.T) {
		file, err := os.Create("config.yml")
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()
		defer os.Remove("config.yml")
		err = file.Chmod(0000)
		if err != nil {
			t.Fatal(err)
		}
		expected := Config{}

		config, err := GetConfig()

		assert.Error(t, err)
		assert.Equal(t, expected, config)
	})

	t.Run("error no config", func(t *testing.T) {
		os.Remove("config.yml")
		expected := Config{}

		config, err := GetConfig()

		assert.Error(t, err)
		assert.Equal(t, expected, config)
	})
}
