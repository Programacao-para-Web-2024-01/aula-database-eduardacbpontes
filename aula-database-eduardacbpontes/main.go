package main

import (
	"aula-database/student"
	"aula-database/subject"
	"aula-database/user"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	if err := createServer(); err != nil {
		log.Panic(err)
	}
}

func connectDB() *sql.DB {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "uniceub"
	config.DBName = "web"
	conn, err := mysql.NewConnector(config)
	if err != nil {
		panic(err)
	}
	return sql.OpenDB(conn)
}

func createServer() error {
	db := connectDB()

	subjectRepository := subject.NewRepository(db)
	subjectService := subject.NewService(subjectRepository)

	studentRepository := student.NewStudentRepository(db)
	studentService := student.NewStudentService(studentRepository, subjectService)
	studentController := student.NewStudentController(studentService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /students/", studentController.List)
	mux.HandleFunc("GET /students/{id}", studentController.Get)
	mux.HandleFunc("POST /students/", appendMiddlewares(studentController.Create, authentication))
	mux.HandleFunc("PUT /students/{id}", appendMiddlewares(studentController.Update, authentication))
	mux.HandleFunc("DELETE /students/{id}", appendMiddlewares(studentController.Delete, authentication))
	mux.HandleFunc("PUT /students/{id}/subjects", appendMiddlewares(studentController.AssociateSubjects, authentication))

	mux.HandleFunc("POST /auth/register", userController.Register)
	mux.HandleFunc("POST /auth/login", userController.Login)

	return http.ListenAndServe("localhost:8080", mux)
}

func appendMiddlewares(
	handler func(w http.ResponseWriter, req *http.Request),
	mw ...func(w http.ResponseWriter, req *http.Request) error,
) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		for _, middleware := range mw {
			err := middleware(w, req)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
		}

		handler(w, req)
	}
}

func authentication(w http.ResponseWriter, req *http.Request) error {
	authorization := req.Header.Get("Authorization")
	_, err := validateToken(authorization)
	if err != nil {
		w.WriteHeader(401)
		return err
	}

	return nil
}

var key = []byte("TOKEN_SECRETO")
var jwtManager = jwt.New(jwt.SigningMethodHS256)

func createToken() (string, error) {
	return jwtManager.SignedString(key)
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.NewParser().Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
}
