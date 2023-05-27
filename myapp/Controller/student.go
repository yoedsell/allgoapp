package Controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/Model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func Addstudent(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "add student handler")
	var stud Model.Student
	_ = stud

	decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)

	if err := decoder.Decode(&stud); err != nil {
		// w.Write([]byte("Invalid json data"))
		response, _ := json.Marshal(map[string]string{"error": "Invalid Json Type"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	defer r.Body.Close()

	saveErr := stud.Create()
	fmt.Println(saveErr)
	if saveErr != nil {
		// w.Write([]byte("Database error"))
		response, _ := json.Marshal(map[string]string{"error": saveErr.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}
	//no error
	// w.Write([]byte("response success"))
	response, _ := json.Marshal(map[string]string{"status": "student added"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func GetStud(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	fmt.Println(sid)
	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := Model.Student{StdId: stdId}
	getErr := s.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Student not found")
		default:
			httpResp.RespondWithError(w, http.StatusNotFound, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, s)
}

func getUserId(userIdParam string) (int64, error) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, userErr
	}
	return userId, nil
}

func UpdateStud(w http.ResponseWriter, r *http.Request) {
	old_sid := mux.Vars(r)["sid"]
	old_stdId, idErr := getUserId(old_sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	var stud Model.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid jason body")
		return
	}
	defer r.Body.Close()

	err := stud.Update(old_stdId)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, stud)
}

func DeleteStud(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := Model.Student{StdId: stdId}
	if err := s.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllStuds(w http.ResponseWriter, r *http.Request) {
	students, getErr := Model.GetAllStudents()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, students)
}

// course
func Addcourse(w http.ResponseWriter, r *http.Request) {
	var cour Model.Course
	_ = cour

	decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)

	if err := decoder.Decode(&cour); err != nil {

		response, _ := json.Marshal(map[string]string{"error": "Invalid Json Type"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	defer r.Body.Close()

	saveErr := cour.Create()
	fmt.Println(saveErr)
	if saveErr != nil {
		response, _ := json.Marshal(map[string]string{"error": saveErr.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}
	//no error
	response, _ := json.Marshal(map[string]string{"status": "course added"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func Getcour(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	fmt.Println(cid)
	courseid, idErr := getCourseId(cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	c := Model.Course{Courseid: courseid}
	getErr := c.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Course not found")
		default:
			httpResp.RespondWithError(w, http.StatusNotFound, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, c)
}

func getCourseId(courseIdParam string) (string, error) {
	return courseIdParam, nil
}

// func getCourseId(courseIdParam string) (int64, error) {
// 	courseId, courseErr := strconv.ParseInt(courseIdParam, 10, 64)
// 	if courseErr != nil {
// 		return 0, courseErr
// 	}
// 	return courseId, nil
// }

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	old_cid := mux.Vars(r)["cid"]
	old_courseId, idErr := getCourseId(old_cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	var cours Model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cours); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid jason body")
		return
	}
	defer r.Body.Close()

	err := cours.Update(old_courseId)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, cours)

}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	courseid, idErr := getCourseId(cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	c := Model.Course{Courseid: courseid}
	if err := c.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, getErr := Model.GetAllCourses()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)
}

// signup
func Signup(w http.ResponseWriter, r *http.Request) {

	var admin Model.Admin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&admin)

	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid Json Typer")
		return
	}
	defer r.Body.Close()
	saveErr := admin.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "login success"})
}

// login
func Login(w http.ResponseWriter, r *http.Request) {

	//setting cookie
	cookie := http.Cookie{
		Name:    "my-cookie",
		Value:   "my-value",
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}
	http.SetCookie(w, &cookie)

	var admin Model.Login

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&admin)

	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}
	defer r.Body.Close()

	getErr := admin.Get()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusUnauthorized, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Login success"})
}

// logout
func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-cookie",
		Expires: time.Now(),
	})
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "cookie deleted"})
}

func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			httpResp.RespondWithError(w, http.StatusSeeOther, "cookie no found")

			return false
		}
		httpResp.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return false
	}

	if cookie.Value != "my-value" {
		httpResp.RespondWithError(w, http.StatusSeeOther, "cookie does not match")
		return false
	}
	return true
}
