package teacherportal

import "html/template"

var rootTemplate *template.Template

func ImportTemplares() (err error) {
	rootTemplate, err = template.ParseFiles(
		"teacherportal/students.gohtml",
		"teacherportal/student.gohtml",
	)

	return err
}
