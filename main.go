// ============================================================================
// = main.go
// = 	Description		The file from which to run any of the sequences
// = 	Date			October 08, 2021
// ============================================================================

package main

import (
	"OEIS/seq"
	"OEIS/utils"
	"errors"
	"flag"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	// program initialization (flags)
	seqid := flag.String("seq", "", "Which sequence to run. Example: -seq A000042")
	seqlen := flag.Int64("seqlen", 5, "How many elements to generate. Most sequences will have restrictions on the # of elements to generate.")
	comptime := flag.Bool("time", true, "True if you want approximate time-of-computation information printed. False otherwise")

	flag.Parse() // remember to parse!

	// check for invalid inputs
	if *seqid == "" { // user must specify a sequence to generate
		utils.HandleError(errors.New("you need to specify a sequence to generate! "))
	} else if *seqlen <= 0 { // check for invalid lengths
		utils.HandleError(errors.New("you need to specify a positive sequence length! "))
	} else if *seqlen < 5 {
		utils.HandleError(errors.New("sequence length should be at least 5. "))
	}

	// warn about long computation times if seqlen is greater than 500
	if *seqlen >= 500 {
		utils.PrintWarning("Depending on your system, this large of a sequence length will probably take a while to compute.")
	}

	start := time.Now()
	temp, offset := handler(strings.ToUpper(*seqid), *seqlen)
	duration := time.Since(start)

	// convert & act accordingly
	if reflect.TypeOf(temp).String() == "[]int64" {
		utils.PrintSequence(*seqid, temp.([]int64), offset)
	} else if reflect.TypeOf(temp).String() == "[]*big.Int" {
		utils.PrintBigSequence(*seqid, temp.([]*big.Int), offset)
	}

	// output time if requested
	if *comptime {
		utils.PrintInfo("Computed " + strconv.FormatInt(*seqlen, 10) + " terms of sequence " + *seqid + " in " + duration.String())
	}
}

// this handles the call to make life easier
func handler(name string, params ...interface{}) (interface{}, int64) {
	out1, out2, err := call(name, params...)
	utils.HandleError(err)
	idx := out2.(int64)
	return out1, idx
}

// this is based upon:
// https://medium.com/@vicky.kurniawan/go-call-a-function-from-string-name-30b41dcb9e12
func call(name string, params ...interface{}) (result interface{}, arg2 interface{}, err error) {
	s := seq.SeqFuncs{}
	f := reflect.ValueOf(s).MethodByName(name)
	
	if !f.IsValid() {
		err = errors.New("either this sequence has not been implemented yet, or your id is invalid")
		return
	}

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
