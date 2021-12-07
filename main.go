// ============================================================================
// = main.go																  =
// = 	Description: The file from which to run any of the sequences		  =
// = 	Date: October 08, 2021												  =
// ============================================================================

package main

import (
	seq "OEIS/sequences"
	"OEIS/utils"
	"errors"
	"reflect"
)

func main() {
	test, startidx := handler("A000041", int64(25))
	utils.PrintSequence("", test, startidx)
}

// this handles the call and conversion of the returns from call()
// to use: seq, idx := handler(seq_name, int64(seq_len_to_generate))
func handler(name string, params ...interface{}) ([]int64, int64) {
	out1, out2, err := call(name, params...)
	utils.HandleError(err)
	seq := out1.([]int64)
	idx := out2.(int64)
	return seq, idx
}

// this is based upon: 
// https://medium.com/@vicky.kurniawan/go-call-a-function-from-string-name-30b41dcb9e12
func call(name string, params ...interface{}) (result interface{}, arg2 interface{}, err error) {
	f := reflect.ValueOf(StubStorage[name])

	if len(params) != f.Type().NumIn() {
		err = errors.New("param count is out of bounds")
		return
	}

	// build parameter list for call
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	// build result interface
	var res []reflect.Value = f.Call(in)
	result = res[0].Interface()
	arg2 = res[1].Interface()
	return
}

// the following is a (large) mapping from strings to the corresponding function
var StubStorage = map[string]interface{}{
	"A000002": seq.A000002,
	"A000004": seq.A000004,
	"A000005": seq.A000005,
	"A000006": seq.A000006,
	"A000007": seq.A000007,
	"A000008": seq.A000008,
	"A000010": seq.A000010,
	"A000011": seq.A000011,
	"A000012": seq.A000012,
	"A000027": seq.A000027,
	"A000030": seq.A000030,
	"A000032": seq.A000032,
	"A000034": seq.A000034,
	"A000035": seq.A000035,
	"A000037": seq.A000037,
	"A000038": seq.A000038,
	"A000040": seq.A000040,
	"A000041": seq.A000041,
	"A000042": seq.A000042,
	"A007947": seq.A007947,
}
