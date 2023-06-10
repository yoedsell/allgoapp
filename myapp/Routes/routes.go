package Routes

import (
	"log"
	"myapp/Controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {

	router := mux.NewRouter()

	//student
	router.HandleFunc("/student", Controller.Addstudent).Methods("POST")
	router.HandleFunc("/student/{sid}", Controller.GetStud).Methods("GET")
	router.HandleFunc("/student/{sid}", Controller.UpdateStud).Methods("PUT")
	router.HandleFunc("/student/{sid}", Controller.DeleteStud).Methods("DELETE")
	router.HandleFunc("/students", Controller.GetAllStuds)

	//course
	router.HandleFunc("/course", Controller.Addcourse).Methods("POST")
	router.HandleFunc("/course/{cid}", Controller.Getcour).Methods("GET")
	router.HandleFunc("/course/{cid}", Controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", Controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/courses", Controller.GetAllCourses)

	//enroll
	router.HandleFunc("/enroll", Controller.Enroll).Methods("POST")
	router.HandleFunc("/enroll/{sid}/{cid}", Controller.GetEnroll).Methods("GET")
	router.HandleFunc("/enroll/{sid}/{cid}", Controller.DeleteEnroll).Methods("DELETE")
	router.HandleFunc("/enrolls", Controller.GetEnrolls)

	// router.HandleFunc("/enroll", Controller.GetAllEnrolls)

	//signup
	router.HandleFunc("/signup", Controller.Signup).Methods("POST")

	//login
	router.HandleFunc("/login", Controller.Login).Methods("POST")

	//logout
	router.HandleFunc("/logout", Controller.Logout).Methods("GET")

	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("Application running on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
