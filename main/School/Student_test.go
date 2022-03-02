package School

import (
	"testing"
)

func TestApplicant_Admit(t *testing.T) {
	test := []struct {
		input1 Applicant

		output string
	}{

		{Applicant{Name: "Chucks", Examscore: 250, Age: 3}, "Not Admitted"},
		{Applicant{Name: "Jimmy", Examscore: 180, Age: 3}, "Not Admitted"},
		{Applicant{Name: "Steven Strange", Examscore: 241, Age: 6}, "Admitted"},
	}
	for _, v := range test {
		got := v.input1.Admit()
		if got != v.output {
			t.Errorf("Expected %v, got %v", v.output, got)
		}
	}
}

func TestPupil_Takeexam(t *testing.T) {
	test := []struct {
		input0 Pupil

		wanted string
	}{
		{input0: Pupil{Registeredpupil{"English"}, "Victor", 9, 87}, wanted: "Pass"},
		{Pupil{Registered: Registeredpupil{"Math"}, Name: "Steven re", Age: 7, Grade: 30}, "Fail"},
	}
	for _, v := range test {
		got := v.input0.Takeexam()
		if got != v.wanted {
			t.Errorf("Expected %v, got %v", v.wanted, got)
		}
	}
}

func TestPupil_PrincipalExpel(t *testing.T) {
	test := []struct {
		input1 Pupil

		wanted1 string
	}{
		{Pupil{Registered: Registeredpupil{"maths"}, Name: "Charles Branson", Age: 10, Grade: 70}, "Promoted"},
		{Pupil{Registered: Registeredpupil{"English"}, Name: "Camson Adams", Age: 9, Grade: 29}, "Expel"},
	}
	for _, v := range test {
		got := v.input1.PrincipalExpel()
		if got != v.wanted1 {
			t.Errorf("Expected %v, got %v", v.wanted1, got)
		}
	}
}

func TestAcademics_Teachsubjects(t *testing.T) {
	test := []struct {
		input2 Academics

		wanted2 string
	}{
		{Academics{
			Basic:    Registeredpupil{"civic studies"},
			Name:     "Johnson John",
			Age:      "8",
			Subjects: "maths, english, civic studies, quantitative reasoning, verbal reasoning, fine art",
			Classes:  "Primary1Science",
			Payment:  30000,
		}, "not approved to be taught subjects"},
		{Academics{
			Basic:    Registeredpupil{"civic studies"},
			Name:     "Johnson John",
			Age:      "8",
			Subjects: "maths, english, civic studies, quantitative reasoning, verbal reasoning, fine art",
			Classes:  "Primary1Science",
			Payment:  43500,
		}, "Approved to be taught subjects"},
	}
	for _, v := range test {
		got := v.input2.Teachsubjects()
		if got != v.wanted2 {
			t.Errorf("Expected %v, got %v", v.wanted2, got)
		}
	}
}

func TestAcademics_PupilAllocated(t *testing.T) {
	test := []struct {
		input3 Academics

		wanted3 string
	}{
		{Academics{
			Basic:    Registeredpupil{"Fineart"},
			Name:     "Rufai Bala",
			Age:      "8",
			Subjects: "maths, english, civic studies, quantitative reasoning, verbal reasoning, fine art",
			Payment:  43500},
			"Allocate Primary1 art"},
		{Academics{
			Basic:    Registeredpupil{"English"},
			Name:     "Chioma Jesus",
			Age:      "7",
			Subjects: "maths, english, civic studies, quantitative reasoning, verbal reasoning, fine art",
			Payment:  43500},
			"Allocate Primary1 art"},

		{Academics{
			Basic:    Registeredpupil{"healtheducation"},
			Name:     "Rufus Giwa",
			Age:      "7",
			Subjects: "maths, english, civic studies, quantitative reasoning, verbal reasoning, fine art",
			Payment:  43500},
			"Primary1 Science"},

		{Academics{
			Basic:    Registeredpupil{"maths"},
			Name:     "James Chidi",
			Age:      "9",
			Subjects: "maths, english, civic studies, quantitative reasoning, verbal reasoning, fine art",
			Payment:  43500},
			"Primary1 Science"},
	}
	for _, v := range test {
		got := v.input3.PupilAllocated()
		if got != v.wanted3 {
			t.Errorf("Expected %v, got %v", v.wanted3, got)
		}
	}
}
