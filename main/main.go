package main

import (
	"Alex/School/main/School"
	"fmt"
)

//This is models the actions/functions of an elementary school
//chronologically from the beginning of term to end of term.

//This is models the actions/functions of an elementary school
//chronologically from the beginning of term to end of term.

func main() {
	//An applicant or prospective student/pupil will take an exam to determine if he/she will be admitted or not.
	//Applicant will also be granted admission on the criteria of age.
	a := School.Applicant{Name: "Alex", Examscore: 250, Age: 9}
	fmt.Println(a.Admit())

	//Pupils/Students are required to Pay a school fee at the beginning of a term.
	//A teacher will then teach pupils who have met the condition for the payment of the school fee.
	b := School.Academics{
		Basic:    School.Registeredpupil{"maths"},
		Name:     "Shuaib Ahmed",
		Age:      "8",
		Subjects: "maths, english, civic studies, quantitative reasoning, verbal reasoning, fine art",
		Classes:  "Primary1Science",
		Payment:  30000,
	}
	fmt.Println(b.Teachsubjects())

	//Subject/course interest will determine what class the pupil/student will be allocated.
	//A teacher will be responsible for allocating classes based on subject interest of Pupil.
	c := School.Academics{
		Basic:    School.Registeredpupil{"Fineart"},
		Name:     "Rufai Bala",
		Age:      "6",
		Subjects: "maths, english, civic studies, quantitative reasoning, verbal reasoning, fine art",
		Classes:  "Primary1art",
		Payment:  43500,
	}
	fmt.Println(c.PupilAllocated())

	//Pupils/students will take exams at end of term.
	//Teachers will administer exams at end of term.
	d := School.Pupil{
		Registered: School.Registeredpupil{"English"},
		Name:       "Victor Ikagu",
		Age:        9,
		Grade:      87,
	}

	fmt.Println(d.Takeexam())

	//The Principal will expel those that fail and promote those that pass the end-of-term exam.
	e := School.Pupil{
		Registered: School.Registeredpupil{"math"},
		Name:       "Charles Branson",
		Age:        10,
		Grade:      70,
	}
	fmt.Println(e.PrincipalExpel())

	//The principal will either terminate if attendance days in a month period is less than 16,
	//Put on probation, if attendance days greater than or equal to 16 and less than 20,
	//or keep active the employment status if attendance days are greater than 20.
	f := School.Staff{
		Attendance: 21,
	}
	fmt.Println(f.StatusofEmployment())

	//The manager(Non-academic staff) has to ensure
	//that there is power in the school premises for at least 5 hrs daily.
	g := School.Nonteaching{Power_supply_hours: 4}
	fmt.Println(g.Maintain())
}
