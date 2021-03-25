package service

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	fmt.Println("test begin")
	m.Run()
	fmt.Println("test end")
}
func TestAddCustomerRewardPointsSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{gin.Param{Key: "k", Value: "v"}}

	AddCustomerRewardPoints(c)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}
}
