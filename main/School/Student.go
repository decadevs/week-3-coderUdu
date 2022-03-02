package School

type Registeredpupil struct {
	SubjectInterest string
}

type Applicant struct {
	Name      string
	Examscore float32
	Age       int
}

type Pupil struct {
	Registered Registeredpupil
	Name       string
	Age        int
	Grade      float32
}

func (Appl *Applicant) Admit() string {
	if Appl.Examscore > 200 && Appl.Age >= 4 && Appl.Age <= 11 {
		return "Admitted"
	}
	return "Not Admitted"

}

func (ppl *Pupil) Takeexam() string {
	if ppl.Grade < 40 {
		return "Fail"
	}
	if ppl.Grade >= 40 {
		return "Pass"
	}
	return ""
}

func (ppl *Pupil) PrincipalExpel() string {
	if ppl.Takeexam() == "Fail" {
		return "Expel"
	}

	if ppl.Takeexam() == "Pass" {
		return "Promoted"
	}
	return ""
}
