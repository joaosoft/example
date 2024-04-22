package test

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type Call struct {
	Arguments []interface{}
	Expected  []interface{}
	Other     MapSomething
}

type CallList []Call

type BaseTestCase struct {
	Test        int
	Description string
	Call
}

type MapCall map[string]CallList

type MapSomething map[string]interface{}

func (b BaseTestCase) Log(t *testing.T) {
	t.Log("::", t.Name(), "Running test", b.Test)
	t.Log("  ", b.Description)
}

type IMock interface {
	On(methodName string, arguments ...interface{}) *mock.Call
}

func (m MapCall) Setup(mock IMock) {
	for key, list := range m {
		for _, value := range list {
			mock.On(key, value.Arguments...).Return(value.Expected...)
		}
	}
}
