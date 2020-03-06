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
	cfg := config.New()
	cache := cfg.ConnectRedis()
	h := Handler{Config: cfg, Cache: cache}

	suite.Recorder = httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	suite.HandlerToTest = h.Auth(nextHandler)

}

func (suite *AuthTestSuiteWithRedis) SetupTest() {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	s.Set("session_token", "expected")

	cfg := config.New()
	cfg.Redis.Host = s.Addr()
	cache := cfg.ConnectRedis()

	fakeCache := redismock.NewNiceMock(cache)
	fakeCache.On("Get", "session_token").
		Return(redis.NewStringResult("expected", nil))

	h := Handler{Config: cfg, Cache: fakeCache}

	suite.Recorder = httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	suite.HandlerToTest = h.Auth(nextHandler)
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

	expected := "Unauthorized: Token invalid."
	suite.Equal(expected, suite.Recorder.Body.String())
}

func (suite *AuthTestSuiteWithRedis) TestToken() {
	http.SetCookie(suite.Recorder, &http.Cookie{Name: "session_token", Value: "expected"})
	req := &http.Request{Header: http.Header{"Cookie": suite.Recorder.HeaderMap["Set-Cookie"]}}
	suite.HandlerToTest.ServeHTTP(suite.Recorder, req)

	suite.Equal(http.StatusOK, suite.Recorder.Code)
}

func (suite *AuthTestSuiteWithoutRedis) TestRedisUnavailable() {
	http.SetCookie(suite.Recorder, &http.Cookie{Name: "session_token", Value: "not-expected"})
	req := &http.Request{Header: http.Header{"Cookie": suite.Recorder.HeaderMap["Set-Cookie"]}}
	suite.HandlerToTest.ServeHTTP(suite.Recorder, req)

	expected := "Opps... Something went wrong."
	suite.Equal(expected, suite.Recorder.Body.String())
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuiteWithRedis))
	suite.Run(t, new(AuthTestSuiteWithoutRedis))
}
