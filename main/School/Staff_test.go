package School

import (
	"testing"
)

func TestStaff_StatusofEmployment(t *testing.T) {
	test := []struct {
		input4 Staff

		wanted4 string
	}{
		{Staff{Attendance: 17}, "Teacher on probation"},
		{Staff{Attendance: 14}, "Terminate Employment"},
		{Staff{Attendance: 20}, "Employment Active"},
	}
	for _, v := range test {
		got := v.input4.StatusofEmployment()
		if got != v.wanted4 {
			t.Errorf("Expected %v, got %v", v.wanted4, got)
		}
	}
}
func TestNonteaching(t *testing.T) {
	test := []struct {
		input5 Nonteaching

		wanted5 string
	}{
		{Nonteaching{Power_supply_hours: 4}, "Determine cause and fix"},
		{Nonteaching{Power_supply_hours: 5}, "maintained"},
	}
	for _, v := range test {
		got := v.input5.Maintain()
		if got != v.wanted5 {
			t.Errorf("Expected %v, got %v", v.wanted5, got)
		}
	}
}
