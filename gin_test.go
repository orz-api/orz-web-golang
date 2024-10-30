package orzweb

import (
	"errors"
	"github.com/gin-gonic/gin"
	"testing"
)

type testReq struct {
	Test string `json:"test" binding:"required"`
}

type testRsp struct {
	Test string `json:"test,omitempty"`
}

func TestCase1(t *testing.T) {
	PropsObj.ExposeErrorReason = true
	PropsObj.ExposeErrorTraces = true

	testApi := &ApiDef[testReq, testRsp]{
		Scope:    "WebV1",
		Domain:   "Test",
		Resource: "Case1",
		Action:   "Query",
		Query:    true,
		Variant:  1,
		Errors: map[string]*ApiErrorDef{
			"1": {
				Reason: "test reason 1",
				Notice: "eng test / 测试中文通知",
			},
		},
		Handler: func(req *testReq) (*testRsp, error) {
			if req.Test == "1" {
				return &testRsp{Test: "test"}, nil
			} else if req.Test == "2" {
				return nil, NewApiError("1")
			} else if req.Test == "3" {
				return nil, NewApiError("2")
			}
			return nil, errors.New("unknown error")
		},
	}

	engine := gin.Default()
	testApi.Gin(engine)
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
