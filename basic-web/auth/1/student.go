package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

// func init() adalah fungsi yang secara otomatis dipanggil ketika package-dimana-fungsi-ini-berada diimport atau di-run
func init() {
	students = append(students, &Student{Id: "S001", Name: "Bourne", Grade: 7})
	students = append(students, &Student{Id: "S002", Name: "Ethan", Grade: 3})
	students = append(students, &Student{Id: "S003", Name: "Wick", Grade: 4})
}
