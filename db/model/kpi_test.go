package model

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
)

// JSON 字段名与建表列名一致；特别检查带数字、缩写的字段不会被 GORM 错误映射。
func TestKPIColumnNames(t *testing.T) {
	for _, record := range []interface{}{KPIMemberModel{}, KPIMemberMonthConfigModel{}, KPIMonthTargetModel{}} {
		typ := reflect.TypeOf(record)
		t.Run(typ.Name(), func(t *testing.T) {
			for i := 0; i < typ.NumField(); i++ {
				field := typ.Field(i)
				column := gorm.ToDBName(field.Name)
				for _, setting := range strings.Split(field.Tag.Get("gorm"), ";") {
					if strings.HasPrefix(setting, "column:") {
						column = strings.TrimPrefix(setting, "column:")
					}
				}
				if want := field.Tag.Get("json"); column != want {
					t.Errorf("%s maps to %q, want %q", field.Name, column, want)
				}
			}
		})
	}
}

func TestKPIMemberMonthConfigNullableMetrics(t *testing.T) {
	fields := []string{
		"initial_lead_cnt", "valid_lead_cnt", "silent_over_7d_pct",
		"first_lesson_on_time_pct", "lesson_consumption_pct",
		"four_lessons_in_30d_pct", "renewal_pct",
	}
	for _, value := range []string{"null", "0"} {
		t.Run(value, func(t *testing.T) {
			input := make(map[string]json.RawMessage)
			for _, field := range fields {
				input[field] = json.RawMessage(value)
			}
			data, err := json.Marshal(input)
			if err != nil {
				t.Fatal(err)
			}
			var record KPIMemberMonthConfigModel
			if err := json.Unmarshal(data, &record); err != nil {
				t.Fatal(err)
			}
			data, err = json.Marshal(record)
			if err != nil {
				t.Fatal(err)
			}
			var output map[string]json.RawMessage
			if err := json.Unmarshal(data, &output); err != nil {
				t.Fatal(err)
			}
			for _, field := range fields {
				if got := string(output[field]); got != value {
					t.Errorf("%s = %s, want %s", field, got, value)
				}
			}
		})
	}
}
