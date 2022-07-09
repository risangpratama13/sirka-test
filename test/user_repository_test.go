package test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/risangpratama13/sirka-test/model"
	"github.com/risangpratama13/sirka-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type UserRepoSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository repository.UserRepository
	user       *model.User
}

func (s *UserRepoSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	s.repository = repository.NewUserRepository(s.DB)
}

func (s *UserRepoSuite) TestFindById() {
	var (
		userid = "01"
		name   = "budi"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE userid = $1 ORDER BY "users"."userid" LIMIT 1`)).
		WithArgs(userid).
		WillReturnRows(s.mock.NewRows([]string{"userid", "name"}).
			AddRow(userid, name))

	userRepository := repository.NewUserRepository(s.DB)
	user, err := userRepository.FindById(context.TODO(), userid)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), user)
	assert.Equal(s.T(), userid, user.Userid)
	assert.Equal(s.T(), name, user.Name)
}

func (s *UserRepoSuite) TestFindAll() {
	dataUsers := []model.User{
		{Userid: "01", Name: "budi"},
		{Userid: "02", Name: "Nano"},
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"`)).
		WillReturnRows(s.mock.NewRows([]string{"userid", "name"}).
			AddRow(dataUsers[0].Userid, dataUsers[0].Name).
			AddRow(dataUsers[1].Userid, dataUsers[1].Name))

	userRepository := repository.NewUserRepository(s.DB)
	users := userRepository.FindAll(context.TODO())
	assert.NotNil(s.T(), users)
	assert.Len(s.T(), users, len(dataUsers))
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepoSuite))
}
