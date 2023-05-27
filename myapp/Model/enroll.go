package Model

import "myapp/dataStore/postgres"

type Enroll struct {
	StdId         int64  `json:"stdid"`
	CourseID      string `json:"courseid"`
	Date_Enrolled string `json:"date"`
}

const queryEnrollStd = "INSERT INTO enroll(std_id,course_id,date_enrolled) VALUES($1, $2, $3);"

func (e *Enroll) EnrollStud() error {
	if _, err := postgres.Db.Exec(queryEnrollStd, e.StdId, e.CourseID,
		e.Date_Enrolled); err != nil {
		return err
	}
	return nil
}

const queryGetEnroll = "SELECT std_id, course_id, date_enrolled FROM enroll WHERE std_id=$1 and course_id=$2"

func (e *Enroll) Get() error {
	return postgres.Db.QueryRow(queryGetEnroll, e.StdId, e.CourseID).Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
}

func GetAllEnrolls() ([]Enroll, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM enroll;")
	if getErr != nil {
		return nil, getErr
	}
	enrolls := []Enroll{}

	for rows.Next() {
		var e Enroll
		dbErr := rows.Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
		if dbErr != nil {
			return nil, dbErr
		}
		enrolls = append(enrolls, e)
	}
	rows.Close()
	return enrolls, nil
}

const queryDeleteEnroll = "DELETE FROM enroll WHERE std_id = $1 and course_id=$2;"

func (e *Enroll) Delete() error {
	if _, err := postgres.Db.Exec(queryDeleteEnroll, e.StdId, e.CourseID); err != nil {
		return err
	}
	return nil
}
