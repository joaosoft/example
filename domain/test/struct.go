package test

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
