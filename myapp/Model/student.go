package Model

import (
	"myapp/dataStore/postgres"
)

type Student struct {
	StdId     int64  `json:"stdid"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
}

const queryInsertUser = "INSERT INTO student(stdid, firstname, lastname, email) VALUES ($1, $2, $3, $4);"

func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)

	return err
}

var queryGetUser = "SELECT stdid, firstname, lastname, email FROM student WHERE stdid=$1;"

func (s *Student) Read() error {
	return postgres.Db.QueryRow(queryGetUser, s.StdId).Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
}

var queryUpdateUser = "UPDATE student SET stdid = $1, firstname = $2, lastname=$3, email=$4 WHERE stdid=$5;"

func (s *Student) Update(oldID int64) error {
	_, err := postgres.Db.Exec(queryUpdateUser, s.StdId, s.FirstName, s.LastName, s.Email, oldID)
	return err
}

var queryDeleteUser = "DELETE FROM student WHERE stdid=$1;"

func (s *Student) Delete() error {
	if _, err := postgres.Db.Exec(queryDeleteUser, s.StdId); err != nil {
		return err
	}
	return nil
}

func GetAllStudents() ([]Student, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM student;")
	if getErr != nil {
		return nil, getErr
	}
	students := []Student{}

	for rows.Next() {
		var s Student
		dbErr := rows.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
		if dbErr != nil {
			return nil, dbErr
		}
		students = append(students, s)
	}
	rows.Close()
	return students, nil
}

// course
type Course struct {
	Courseid   string `json:"courseid"`
	CourseName string `json:"coursename"`
}

const courseinsert = "INSERT INTO course(courseid, coursename)VALUES($1,$2);"

func (c *Course) Create() error {
	_, err := postgres.Db.Exec(courseinsert, c.Courseid, c.CourseName)
	return err
	// return postgres.Db.Exec(queryInsertUser, s.stdId, s.FirstName, s.LastName, s.email)
}

var queryGetCourse = "SELECT courseid, coursename FROM course WHERE courseid=$1;"

func (c *Course) Read() error {
	return postgres.Db.QueryRow(queryGetCourse, c.Courseid).Scan(&c.Courseid, &c.CourseName)
}

var queryUpdateCourse = "UPDATE course SET courseid = $1, coursename = $2 WHERE courseid = $3;"

func (c *Course) Update(oldID string) error {
	_, err := postgres.Db.Exec(queryUpdateCourse, c.Courseid, c.CourseName, oldID)
	return err
}

var queryDeleteCourse = "DELETE FROM course WHERE courseid=$1;"

func (c *Course) Delete() error {
	if _, err := postgres.Db.Exec(queryDeleteCourse, c.Courseid); err != nil {
		return err
	}
	return nil
}

func GetAllCourses() ([]Course, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM course;")
	if getErr != nil {
		return nil, getErr
	}
	courses := []Course{}

	for rows.Next() {
		var c Course
		dbErr := rows.Scan(&c.Courseid, &c.CourseName)
		if dbErr != nil {
			return nil, dbErr
		}
		courses = append(courses, c)
	}
	rows.Close()
	return courses, nil
}

// signup
type Admin struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

const queryInsertAdmin = "INSERT into admin(firstname, lastname, email, password) VALUES ($1, $2, $3, $4);"

func (adm *Admin) Create() error {
	_, err := postgres.Db.Exec(queryInsertAdmin, adm.Firstname, adm.Lastname, adm.Email, adm.Password)
	return err
}

// login
type Login struct {
	Email    string
	Password string
}

const queryGetAdmin = "SELECT email, password FROM admin WHERE email = $1 and password = $2;"

func (adm *Login) Get() error {
	return postgres.Db.QueryRow(queryGetAdmin, adm.Email, adm.Password).Scan(&adm.Email, &adm.Password)
}
