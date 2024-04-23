package test

import "github.com/stretchr/testify/mock"

type IMock interface {
	On(methodName string, arguments ...interface{}) *mock.Call
}
