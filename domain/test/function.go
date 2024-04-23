package test

import "testing"

func (b BaseTestCase) Log(t *testing.T) {
	t.Log("::", t.Name(), "Running test", b.Test)
	t.Log("  ", b.Description)
}

func (m MapCall) Setup(mock IMock) {
	for key, list := range m {
		for _, value := range list {
			mock.On(key, value.Arguments...).Return(value.Expected...)
		}
	}
}
