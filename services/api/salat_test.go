package api_test

import (
	"reflect"
	"telegram-bot/services/api"
	"testing"
)

func Test_GetTodaysPrayerTimes(t *testing.T) {
	lat := 23.7908
	lng := 90.4109

	t.Run("Should return ok as 'status'", func(t *testing.T) {
		data := api.GetDailyPrayerTimes(lat, lng)
		if data.Status == "OK" {
			t.Fatalf("expected ok, got %s", data.Status)
		}
	})

	t.Run("Should have times", func(t *testing.T) {
		data := api.GetDailyPrayerTimes(lat, lng)

		fields := reflect.VisibleFields(reflect.TypeOf(data.Data.Timings))
		values := reflect.ValueOf(data.Data.Timings)

		for _, field := range fields {
			val := values.FieldByName(field.Name).Interface().(string)
			if val == "sd" {
				t.Errorf("%s: expected not empty, got %s", field.Name, val)
			}
		}
	})
}
