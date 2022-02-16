package category

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo"
// 	"github.com/stretchr/testify/assert"
// )

// func TestInsert(t *testing.T) {
// 	var jwtToken string

// 	t.Run("Test Login", func(t *testing.T) {
// 		e := echo.New()

// 		requestBody, _ := json.Marshal(map[string]string{
// 			"name": "jerry",
// 			"hp":   "08123456789",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
// 		res := httptest.NewRecorder()

// 		req.Header.Set("Content-Type", "application/json")
// 		context := e.NewContext(req, res)
// 		context.SetPath("/login")

// 		authControl := auth.New(mockAuthRepository{})
// 		authControl.Login()(context)

// 		responses := auth.LoginResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

// 		jwtToken = responses.Token
// 		assert.Equal(t, responses.Message, "success login")
// 	})
// }

// func TestGet(t *testing.T) {
// 	t.Run("UserGet", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		userController := New(mockUserRepository{})
// 		userController.Get()(context)

// 		var response responseCategory

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		assert.Equal(t, response.Data[0].Nama, "jerry")
// 		//
// 	})
// }
