package School

type Principal struct {
	Name          string
	Age           int
	Qualification string
	Gender        string
}

//type manager struct {
//	Name               string
//	age                int
//	gender             string
//	power_supply_hours int
//}

type Nonteaching struct {
	Principal          Principal
	Power_supply_hours int
}
type Teacher struct {
	Teacherattendance int
}

type Staff struct {
	Nonteachingstaff Nonteaching
	Teacher          Teacher
	Attendance       int
}

func (s *Staff) StatusofEmployment() string {
	if s.Attendance >= 20 {
		return "Employment Active"
	}
	if s.Attendance >= 16 && s.Attendance < 20 {
		return "Teacher on probation"
	}
	if s.Attendance < 16 {
		return "Terminate Employment"
	}
	return ""
}

func (man *Nonteaching) Maintain() string {
	if man.Power_supply_hours >= 5 {
		return "maintained"
	}
	if man.Power_supply_hours < 5 {
		return "Determine cause and fix"
	}
	return ""
}
