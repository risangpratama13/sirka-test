package test

import (
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/risangpratama13/sirka-test/app"
	"github.com/risangpratama13/sirka-test/controller"
	"github.com/risangpratama13/sirka-test/model"
	"github.com/risangpratama13/sirka-test/repository"
	"github.com/risangpratama13/sirka-test/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

type UserControllerSuite struct {
	suite.Suite
	DB     *gorm.DB
	router *gin.Engine
	mock   sqlmock.Sqlmock
}

func (s *UserControllerSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	s.router = SetUpRouter()

	routerGroup := s.router.Group("/MyWeb")
	userRepository := repository.NewUserRepository(s.DB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserControllerImpl(userService)
	app.UserRouter(routerGroup, userController)
}

func (s *UserControllerSuite) TestDisplayAllUsers() {
	dataUsers := []model.User{
		{Userid: "01", Name: "Nino"},
		{Userid: "02", Name: "Nano"},
	}

	request, err := http.NewRequest(http.MethodGet, "/MyWeb/DisplayAllUsers", nil)
	require.NoError(s.T(), err)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(s.mock.NewRows([]string{"userid", "name"}).
			AddRow(dataUsers[0].Userid, dataUsers[0].Name).
			AddRow(dataUsers[1].Userid, dataUsers[1].Name))

	response := httptest.NewRecorder()
	s.router.ServeHTTP(response, request)

	var users []model.User
	err = json.Unmarshal(response.Body.Bytes(), &users)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), 200, response.Code)
	assert.Len(s.T(), users, 2)
	assert.Equal(s.T(), users[0], dataUsers[0])
	assert.Equal(s.T(), users[1], dataUsers[1])
}

func (s *UserControllerSuite) TestDisplayUser() {
	var (
		userid = "01"
		name   = "Nino"
	)

	data := url.Values{}
	data.Set("userid", "01")
	request, err := http.NewRequest(http.MethodPost, "/MyWeb/DisplayUser", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	require.NoError(s.T(), err)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE userid = $1 ORDER BY "users"."userid" LIMIT 1`)).
		WithArgs(userid).
		WillReturnRows(s.mock.NewRows([]string{"userid", "name"}).
			AddRow(userid, name))

	response := httptest.NewRecorder()
	s.router.ServeHTTP(response, request)

	var user model.User
	err = json.Unmarshal(response.Body.Bytes(), &user)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), 200, response.Code)
	assert.Equal(s.T(), userid, user.Userid)
	assert.Equal(s.T(), name, user.Name)
}

func (s *UserControllerSuite) TestFailedDisplayUser() {
	data := url.Values{}
	data.Set("userid", "01")
	request, err := http.NewRequest(http.MethodPost, "/MyWeb/DisplayUser", strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	require.NoError(s.T(), err)

	response := httptest.NewRecorder()
	s.router.ServeHTTP(response, request)

	assert.Equal(s.T(), 404, response.Code)
}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
