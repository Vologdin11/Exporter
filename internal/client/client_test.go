package client

import (
	"net/http"
	"os"
	"stp-exporter/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllMetrics(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		env := map[string]string{
			"URL": "http://localhost:8080",
			"DB":  "test",
		}
		for key, value := range env {
			if err := os.Setenv(key, value); err != nil {
				t.Fatal(err)
			}
		}
		defer func() {
			for key := range env {
				if err := os.Unsetenv(key); err != nil {
					t.Fatal(err)
				}
			}
		}()
		token := token{
			Access_token: "eyJhbGciOiJSUzI1NiIsImtpZCI6IkI4QTNCM…..",
			Expires_in:   3600,
			Token_type:   "Bearer",
			Scope:        "viqube_api viqubeadmin_api",
		}
		expected := make([]metrics, 4)
		for i := range expected {
			expected[i] = metrics{
				Columns: []string{
					"Код региона",
					"Услуга",
					"Этап",
					"Количество обращений",
					"Дата выгрузки",
				},
				Values: [][]interface{}{
					{
						"0",
						"Обращения от ЦА ФНС по вопросам АИС ФНС России",
						"Уточнение",
						1.0,
						"2021-10-24",
					},
					{
						"0",
						"Федеральная информационная адресная система (ФИАС)",
						"Внешняя линия",
						1.0,
						"2021-10-24",
					},
					{
						"0",
						"Федеральная информационная адресная система (ФИАС)",
						"Уточнение",
						2.0,
						"2021-10-24",
					},
				}}
		}
		config := config.Config{
			Tables: []config.Table{
				{
					Name:          "BM_Dinamika_postupleniya_obrasch",
					Value_index:   3,
					Label_indexes: []int{0, 1, 2},
				},
				{
					Name:          "BM_Massovie_problemi",
					Value_index:   3,
					Label_indexes: []int{0, 1, 2},
				},
				{
					Name:          "BM_Obrascheniya_v_rabote",
					Value_index:   3,
					Label_indexes: []int{0, 1, 2},
				},
				{
					Name:          "BM_Srednee_vremya_ras_obr_STP",
					Value_index:   3,
					Label_indexes: []int{0, 1, 2},
				},
			},
		}

		metrics, err := GetAllMetrics(config, token)

		assert.NoError(t, err)
		assert.Equal(t, expected, metrics)
	})
}

func TestGetMetric(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		env := map[string]string{
			"URL": "http://localhost:8080",
			"DB":  "test",
		}
		for key, value := range env {
			if err := os.Setenv(key, value); err != nil {
				t.Fatal(err)
			}
		}
		defer func() {
			for key := range env {
				if err := os.Unsetenv(key); err != nil {
					t.Fatal(err)
				}
			}
		}()
		token := token{
			Access_token: "eyJhbGciOiJSUzI1NiIsImtpZCI6IkI4QTNCM…..",
			Expires_in:   3600,
			Token_type:   "Bearer",
			Scope:        "viqube_api viqubeadmin_api",
		}
		expected := metrics{
			Columns: []string{
				"Код региона",
				"Услуга",
				"Этап",
				"Количество обращений",
				"Дата выгрузки",
			},
			Values: [][]interface{}{
				{
					"0",
					"Обращения от ЦА ФНС по вопросам АИС ФНС России",
					"Уточнение",
					1.0,
					"2021-10-24",
				},
				{
					"0",
					"Федеральная информационная адресная система (ФИАС)",
					"Внешняя линия",
					1.0,
					"2021-10-24",
				},
				{
					"0",
					"Федеральная информационная адресная система (ФИАС)",
					"Уточнение",
					2.0,
					"2021-10-24",
				},
			}}

		metric, err := GetMetric("test", token)

		assert.NoError(t, err)
		assert.Equal(t, expected, metric)
	})
}

func TestCreateMetricRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		env := map[string]string{
			"URL": "http://localhost:8080",
			"DB":  "test",
		}
		for key, value := range env {
			if err := os.Setenv(key, value); err != nil {
				t.Fatal(err)
			}
		}
		defer func() {
			for key := range env {
				if err := os.Unsetenv(key); err != nil {
					t.Fatal(err)
				}
			}
		}()
		token := token{
			Access_token: "eyJhbGciOiJSUzI1NiIsImtpZCI6IkI4QTNCM…..",
			Expires_in:   3600,
			Token_type:   "Bearer",
			Scope:        "viqube_api viqubeadmin_api",
		}

		request, err := createMetricRequest("test", token)

		assert.NoError(t, err)
		assert.Equal(t, "GET", request.Method)
		assert.Equal(t, "3.5", request.Header.Get("X-API-VERSION"))
		assert.Equal(t, "BearereyJhbGciOiJSUzI1NiIsImtpZCI6IkI4QTNCM…..", request.Header.Get("Authorization")) //уточнить
		assert.Equal(t, "http://localhost:8080/viqube/databases/test/tables/test/records", request.URL.String())
	})
}

func TestGetToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		env := map[string]string{
			"URL":           "http://localhost:8080",
			"DB":            "test",
			"LOGIN":         "login",
			"PASSWORD":      "pass",
			"AUTHORIZATION": "qweasd",
		}
		for key, value := range env {
			if err := os.Setenv(key, value); err != nil {
				t.Fatal(err)
			}
		}
		defer func() {
			for key := range env {
				if err := os.Unsetenv(key); err != nil {
					t.Fatal(err)
				}
			}
		}()
		expectedToken := token{
			Access_token: "eyJhbGciOiJSUzI1NiIsImtpZCI6IkI4QTNCM…..",
			Expires_in:   3600,
			Token_type:   "Bearer",
			Scope:        "viqube_api viqubeadmin_api",
		}

		token, err := GetToken()

		assert.NoError(t, err)
		assert.Equal(t, expectedToken, token)
	})
}

func TestCreateTokenRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) { //спросить про Body
		env := map[string]string{
			"URL":           "http://localhost:8080",
			"DB":            "test",
			"LOGIN":         "login",
			"PASSWORD":      "pass",
			"AUTHORIZATION": "qweasd",
		}
		for key, value := range env {
			if err := os.Setenv(key, value); err != nil {
				t.Fatal(err)
			}
		}
		defer func() {
			for key := range env {
				if err := os.Unsetenv(key); err != nil {
					t.Fatal(err)
				}
			}
		}()

		request, err := createTokenRequest()

		assert.NoError(t, err)
		assert.Equal(t, "application/x-www-form-urlencoded", request.Header.Get("Content-Type"))
		assert.Equal(t, "qweasd", request.Header.Get("Authorization"))
		assert.Equal(t, "POST", request.Method)
		assert.Equal(t, "http://localhost:8080", request.URL.String())

		// expected := "data-raw:grant_type=password&scope=viqube_api+viqubeadmin_api&response_type=id_token+token&username=login&password=pass"
		// bodyRequest, err := ioutil.ReadAll(request.Body)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// defer request.Body.Close()
		// assert.Equal(t, expected, string(bodyRequest))
	})

	t.Run("Not env", func(t *testing.T) {
		env := []string{"URL", "LOGIN", "PASSWORD", "AUTHORIZATION"}
		for _, envName := range env {
			if err := os.Unsetenv(envName); err != nil {
				t.Fatal(err)
			}
		}

		request, err := createTokenRequest()

		assert.Error(t, err)
		assert.Equal(t, (*http.Request)(nil), request)
	})
}

func TestGetEnv(t *testing.T) {
	t.Run("success all env", func(t *testing.T) {
		expected := map[string]string{
			"URL":           "http://localhost:8080",
			"DB":            "test",
			"LOGIN":         "login",
			"PASSWORD":      "pass",
			"AUTHORIZATION": "qweasd",
		}
		for key, value := range expected {
			if err := os.Setenv(key, value); err != nil {
				t.Fatal(err)
			}
		}
		defer func() {
			for key := range expected {
				if err := os.Unsetenv(key); err != nil {
					t.Fatal(err)
				}
			}
		}()
		envNames := make([]string, 0, len(expected))
		for key := range expected {
			envNames = append(envNames, key)
		}

		env, err := getEnv(envNames...)

		assert.NoError(t, err)
		assert.Equal(t, expected, env)
	})

	t.Run("Not env", func(t *testing.T) {
		if err := os.Unsetenv("URL"); err != nil {
			t.Fatal(err)
		}

		env, err := getEnv("URL")

		assert.Error(t, err)
		assert.Equal(t, map[string]string(nil), env) //спросить почему так возвращает
	})
}
