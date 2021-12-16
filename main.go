// ============================================================================
// = main.go																  =
// = 	Description: The file from which to run any of the sequences		  =
// = 	Date: October 08, 2021												  =
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
	seqlen := flag.Int64("seqlen", 0, "How many elements to generate. Most sequences will have restrictions on the # of elements to generate.")
	comptime := flag.Bool("time", false, "True if you want approximate time-of-computation information printed. False otherwise")
	
	flag.Parse()		// remember to parse!

	_, exists := StubStorage[*seqid]

	if *seqid == "" {				// user must specify a sequence to generate
		utils.HandleError(errors.New("you need to specify a sequence to run! "))
	} else if *seqlen <= 0 {		// check for invalid lengths
		utils.HandleError(errors.New("you need to specify a positive sequence length! "))		
	} else if !exists {				// user must specify a sequence that exists
		utils.HandleError(errors.New("either this sequence has not been implemented yet, or your id is invalid! "))
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
	"A000013": seq.A000013,
	"A000018": seq.A000018,
	"A000021": seq.A000021,
	"A000024": seq.A000024,
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
	"A000047": seq.A000047,
	"A000049": seq.A000049,
	"A000050": seq.A000050,
	"A000051": seq.A000051,
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
	"A000079": seq.A000079,
	"A000082": seq.A000082,
	"A000086": seq.A000086,
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
	"A000114": seq.A000114,
	"A000115": seq.A000115,
	"A000116": seq.A000116,
	"A000117": seq.A000117,
	"A000118": seq.A000118,
	"A000120": seq.A000120,
	"A000123": seq.A000123,
	"A000124": seq.A000124,
	"A000125": seq.A000125,
	"A000126": seq.A000126,
	"A000127": seq.A000127,
	"A000128": seq.A000128,
	"A000129": seq.A000129,
	"A000133": seq.A000133,
	"A000138": seq.A000138,
	"A000139": seq.A000139,
	"A000142": seq.A000142,
	"A000149": seq.A000149,
	"A000150": seq.A000150,
	"A000153": seq.A000153,
	"A000158": seq.A000158,
	"A000160": seq.A000160,
	"A000161": seq.A000161,
	"A000164": seq.A000164,
	"A000165": seq.A000165,
	"A000166": seq.A000166,
	"A000168": seq.A000168,
	"A000169": seq.A000169,
	"A000172": seq.A000172,
	"A000174": seq.A000174,
	"A000178": seq.A000178,
	"A000179": seq.A000179,
	"A000182": seq.A000182,
	"A000184": seq.A000184,
	"A000188": seq.A000188,
	"A000189": seq.A000189,
	"A000190": seq.A000190,
	"A000193": seq.A000193,
	"A000194": seq.A000194,
	"A000195": seq.A000195,
	"A000196": seq.A000196,
	"A000197": seq.A000197,
	// thru300.go
	"A000201": seq.A000201,
	"A000202": seq.A000202,
	"A000203": seq.A000203,
	"A000204": seq.A000204,
	"A000205": seq.A000205,
	"A000207": seq.A000207,
	"A000208": seq.A000208,
	"A000209": seq.A000209,
	"A000210": seq.A000210,
	"A000211": seq.A000211,
	"A000212": seq.A000212,
	"A000213": seq.A000213,
	"A000215": seq.A000215,
	"A000216": seq.A000216,
	"A000217": seq.A000217,
	"A000218": seq.A000218,
	"A000219": seq.A000219,
	"A000221": seq.A000221,
	"A000225": seq.A000225,
	"A000227": seq.A000227,
	"A000230": seq.A000230,
	"A000231": seq.A000231,
	"A000240": seq.A000240,
	"A000244": seq.A000244,
	"A000245": seq.A000245,
	"A000246": seq.A000246,
	"A000247": seq.A000247,
	"A000248": seq.A000248,
	"A000252": seq.A000252,
	"A000253": seq.A000253,
	"A000254": seq.A000254,
	"A000255": seq.A000255,
	"A000256": seq.A000256,
	"A000257": seq.A000257,
	"A000259": seq.A000259,
	"A000260": seq.A000260,
	"A000261": seq.A000261,
	"A000262": seq.A000262,
	"A000263": seq.A000263,
	"A000265": seq.A000265,
	"A000266": seq.A000266,
	"A000267": seq.A000267,
	"A000270": seq.A000270,
	"A000271": seq.A000271,
	"A000272": seq.A000272,
	"A000274": seq.A000274,
	"A000275": seq.A000275,
	"A000276": seq.A000276,
	"A000277": seq.A000277,
	"A000278": seq.A000278,
	"A000279": seq.A000279,
	"A000280": seq.A000280,
	"A000283": seq.A000283,
	"A000284": seq.A000284,
	"A000285": seq.A000285,
	"A000286": seq.A000286,
	"A000287": seq.A000287,
	"A000288": seq.A000288,
	"A000289": seq.A000289,
	"A000290": seq.A000290,
	"A000291": seq.A000291,
	"A000292": seq.A000292,
	"A000294": seq.A000294,
	"A000295": seq.A000295,
	"A000296": seq.A000296,
	"A000297": seq.A000297,
	// otherseq.go
	"A001065": seq.A001065,
	"A001223": seq.A001223,
	"A001622": seq.A001622,
	"A001840": seq.A001840,
	"A002061": seq.A002061,
	"A002386": seq.A002386,
	"A003048": seq.A003048,
	"A007947": seq.A007947,
	"A011848": seq.A011848,
	"A011858": seq.A011858,
	"A027641": seq.A027641,
	"A027642": seq.A027642,
	"A032346": seq.A032346,
	"A038040": seq.A038040,
	"A052614": seq.A052614,
	"A088218": seq.A088218,
	"A128422": seq.A128422,
	"A132269": seq.A132269,
	"A164514": seq.A164514,
	"A168014": seq.A168014,
}
