package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
	// fn("I still can't believe south korea beat Germany 2-0 to put them last")
}
