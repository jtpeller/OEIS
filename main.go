// ============================================================================
// = main.go																  =
// = 	Description: The file from which to run any of the sequences		  =
// = 	Date: October 08, 2021												  =
// ============================================================================

package main

import (
	seq "OEIS/seq"
	"OEIS/utils"
	"errors"
	"flag"
	"reflect"
	"strings"
	"time"
)

func main() {
	// program initialization (flags)
	seqid := flag.String("seq", "", "Which sequence to run. Example: -seq A000042")
	seqlen := flag.Int64("seqlen", 0, "How many elements to generate. Most sequences will have restrictions on the # of elements to generate.")
	comptime := flag.Bool("time", false, "True if you want approximate time-of-computation information printed. False otherwise")
	
	flag.Parse()		// remember to parse!

	_, exists := StubStorage[*seqid]

	if *seqid == "" {				// user must specify a sequence to generate
		utils.HandleError(errors.New("you need to specify a sequence to run! "))
	} else if *seqlen <= 0 {		// check for invalid lengths
		utils.HandleError(errors.New("you need to specify a positive sequence length! "))		
	} else if !exists {				// user must specify a sequence that exists
		utils.HandleError(errors.New("either this sequence has not been implemented yet, or your id is improper! "))
	}

	start := time.Now()
	seq, offset := handler(strings.ToUpper(*seqid), *seqlen)
	duration := time.Since(start)
	utils.PrintSequence("", seq, offset)
	if *comptime {
		utils.PrintInfo("Computed sequence " + *seqid + " in " + duration.String())
	}
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
		err = errors.New("error in call(): param count is out of bounds")
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
	// thru100.go
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
	"A000043": seq.A000043,
	"A000044": seq.A000044,
	"A000045": seq.A000045,
	"A000058": seq.A000058,
	"A000059": seq.A000059,
	"A000062": seq.A000062,
	"A000064": seq.A000064,
	"A000065": seq.A000065,
	"A000068": seq.A000068,
	"A000069": seq.A000069,
	"A000070": seq.A000070,
	"A000071": seq.A000071,
	"A000073": seq.A000073,
	"A000078": seq.A000078,
	"A000082": seq.A000082,
	"A000093": seq.A000093,
	"A000094": seq.A000094,
	"A000096": seq.A000096,
	"A000097": seq.A000097,
	"A000098": seq.A000098,
	"A000100": seq.A000100,
	// thru200.go
	"A000101": seq.A000101,
	"A000102": seq.A000102,
	"A000108": seq.A000108,
	"A000110": seq.A000110,
	"A000111": seq.A000111,
	"A000115": seq.A000115,
	"A000116": seq.A000116,
	"A000120": seq.A000120,
	"A007947": seq.A007947,
}
