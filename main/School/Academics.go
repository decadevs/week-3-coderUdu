package School

type Academics struct {
	Basic    Registeredpupil
	Name     string
	Age      string
	Subjects string
	Classes  string
	Payment  float32
}

func (Acd *Academics) Teachsubjects() string {
	if Acd.Payment >= 43500 {

		return "Approved to be taught subjects"
	}
	return "not approved to be taught subjects"
}

func (Acd *Academics) PupilAllocated() string {
	pupilInterest := Acd.Basic.SubjectInterest

	if pupilInterest == "Fineart" || pupilInterest == "English" {
		return "Allocate Primary1 art"
	}
	if pupilInterest == "maths" || pupilInterest == "healtheducation" {
		return "Primary1 Science"
	}

	return ""
}
