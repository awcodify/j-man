package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/awcodify/j-man/config"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
)

var cfg, _ = config.New()

type AuthTestSuiteWithRedis struct {
	suite.Suite
	Recorder      *httptest.ResponseRecorder
	HandlerToTest http.Handler
}

type AuthTestSuiteWithoutRedis struct {
	suite.Suite
	Recorder      *httptest.ResponseRecorder
	HandlerToTest http.Handler
}

func (suite *AuthTestSuiteWithoutRedis) SetupTest() {
	cache := cfg.ConnectRedis()
	m := Middleware{Cache: cache}

	suite.Recorder = httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	suite.HandlerToTest = m.Auth(nextHandler)

}

func (suite *AuthTestSuiteWithRedis) SetupTest() {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	s.Set("session_token", "expected")

	cfg.Redis.Host = s.Addr()
	cache := cfg.ConnectRedis()

	fakeCache := redismock.NewNiceMock(cache)
	// We expect every success login wil store session_token as key and email as value in Redis
	fakeCache.On("Get", "thisistoken").
		Return(redis.NewStringResult("mail@example.com", nil))

	fakeCache.On("Get", "not-expected").
		Return(redis.NewStringResult("", nil))

	m := Middleware{Cache: fakeCache}

	suite.Recorder = httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	suite.HandlerToTest = m.Auth(nextHandler)
}

func (suite *AuthTestSuiteWithRedis) TestNoCookie() {
	req := httptest.NewRequest("GET", "http://example", nil)
	suite.HandlerToTest.ServeHTTP(suite.Recorder, req)

	expected := "Unauthorized: Token not found."
	suite.Equal(expected, suite.Recorder.Body.String())
}

func (suite *AuthTestSuiteWithRedis) TestTokenNotValid() {
	http.SetCookie(suite.Recorder, &http.Cookie{Name: "session_token", Value: "not-expected"})
	req := &http.Request{Header: http.Header{"Cookie": suite.Recorder.HeaderMap["Set-Cookie"]}}
	suite.HandlerToTest.ServeHTTP(suite.Recorder, req)

	expected := "Unauthorized: User not found."
	suite.Equal(expected, suite.Recorder.Body.String())
}

func (suite *AuthTestSuiteWithRedis) TestToken() {
	http.SetCookie(suite.Recorder, &http.Cookie{Name: "session_token", Value: "thisistoken"})
	req := &http.Request{Header: http.Header{"Cookie": suite.Recorder.HeaderMap["Set-Cookie"]}}
	suite.HandlerToTest.ServeHTTP(suite.Recorder, req)

	suite.Equal(http.StatusOK, suite.Recorder.Code)
}

func (suite *AuthTestSuiteWithoutRedis) TestRedisUnavailable() {
	http.SetCookie(suite.Recorder, &http.Cookie{Name: "session_token", Value: "not-expected"})
	req := &http.Request{Header: http.Header{"Cookie": suite.Recorder.HeaderMap["Set-Cookie"]}}
	suite.HandlerToTest.ServeHTTP(suite.Recorder, req)

	expected := "Unauthorized: Token invalid or already been expired."
	suite.Equal(expected, suite.Recorder.Body.String())
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuiteWithRedis))
	suite.Run(t, new(AuthTestSuiteWithoutRedis))
}
