// This file is part of conscheme
// Automatically generated by compiler/primitives.scm
package vm
import "fmt"
import "os"
var primitive [103]Obj
func init() {
	primitive[98] = wrap(&Procedure{name:"stop-cpu-profile",required:0,apply:nil,label:98})
	primitive[97] = wrap(&Procedure{name:"start-cpu-profile",required:1,apply:nil,label:97})
	primitive[96] = wrap(&Procedure{name:"current-thread",required:0,apply:nil,label:96})
	primitive[95] = wrap(&Procedure{name:"$receive",required:1,apply:nil,label:95})
	primitive[94] = wrap(&Procedure{name:"send",required:2,apply:nil,label:94})
	primitive[93] = wrap(&Procedure{name:"thread-link!",required:2,apply:nil,label:93})
	primitive[92] = wrap(&Procedure{name:"thread-start!",required:1,apply:nil,label:92})
	primitive[91] = wrap(&Procedure{name:"thread-yield!",required:0,apply:nil,label:91})
	primitive[102] = wrap(&Procedure{name:"thread-queue-set!",required:2,apply:nil,label:102})
	primitive[101] = wrap(&Procedure{name:"thread-queue",required:1,apply:nil,label:101})
	primitive[90] = wrap(&Procedure{name:"thread-specific-set!",required:2,apply:nil,label:90})
	primitive[89] = wrap(&Procedure{name:"thread-specific",required:1,apply:nil,label:89})
	primitive[88] = wrap(&Procedure{name:"thread-name",required:1,apply:nil,label:88})
	primitive[87] = wrap(&Procedure{name:"thread?",required:1,apply:nil,label:87})
	primitive[86] = wrap(&Procedure{name:"$make-thread",required:2,apply:nil,label:86})
	primitive[85] = wrap(&Procedure{name:"$write",required:2,apply:nil,label:85})
	primitive[84] = wrap(&Procedure{name:"$display",required:2,apply:nil,label:84})
	primitive[83] = wrap(&Procedure{name:"lookahead-u8",required:1,apply:nil,label:83})
	primitive[82] = wrap(&Procedure{name:"put-bytevector",required:2,apply:nil,label:82})
	primitive[81] = wrap(&Procedure{name:"put-u8",required:2,apply:nil,label:81})
	primitive[80] = wrap(&Procedure{name:"get-u8",required:1,apply:nil,label:80})
	primitive[79] = wrap(&Procedure{name:"$write-char",required:2,apply:nil,label:79})
	primitive[78] = wrap(&Procedure{name:"$peek-char",required:1,apply:nil,label:78})
	primitive[77] = wrap(&Procedure{name:"$read-char",required:1,apply:nil,label:77})
	primitive[76] = wrap(&Procedure{name:"close-port",required:1,apply:nil,label:76})
	primitive[75] = wrap(&Procedure{name:"close-output-port",required:1,apply:nil,label:75})
	primitive[74] = wrap(&Procedure{name:"close-input-port",required:1,apply:nil,label:74})
	primitive[73] = wrap(&Procedure{name:"open-file-output-port",required:1,apply:nil,label:73})
	primitive[99] = wrap(&Procedure{name:"open-output-file",required:1,apply:nil,label:99})
	primitive[72] = wrap(&Procedure{name:"open-input-file",required:1,apply:nil,label:72})
	primitive[71] = wrap(&Procedure{name:"delete-file",required:1,apply:nil,label:71})
	primitive[70] = wrap(&Procedure{name:"file-exists?",required:1,apply:nil,label:70})
	primitive[69] = wrap(&Procedure{name:"current-output-port",required:0,apply:nil,label:69})
	primitive[68] = wrap(&Procedure{name:"current-input-port",required:0,apply:nil,label:68})
	primitive[67] = wrap(&Procedure{name:"output-port?",required:1,apply:nil,label:67})
	primitive[66] = wrap(&Procedure{name:"input-port?",required:1,apply:nil,label:66})
	primitive[65] = wrap(&Procedure{name:"port?",required:1,apply:nil,label:65})
	primitive[64] = wrap(&Procedure{name:"$bytevector-output-port-extract",required:1,apply:nil,label:64})
	primitive[63] = wrap(&Procedure{name:"$open-bytevector-output-port",required:0,apply:nil,label:63})
	primitive[62] = wrap(&Procedure{name:"string->utf8",required:1,apply:nil,label:62})
	primitive[61] = wrap(&Procedure{name:"u8-list->bytevector",required:1,apply:nil,label:61})
	primitive[60] = wrap(&Procedure{name:"bytevector-length",required:1,apply:nil,label:60})
	primitive[59] = wrap(&Procedure{name:"bytevector?",required:1,apply:nil,label:59})
	primitive[58] = wrap(&Procedure{name:"$global-set!",required:2,apply:nil,label:58})
	primitive[57] = wrap(&Procedure{name:"$global-ref",required:1,apply:nil,label:57})
	primitive[56] = wrap(&Procedure{name:"$cell-set!",required:2,apply:nil,label:56})
	primitive[55] = wrap(&Procedure{name:"$cell-ref",required:1,apply:nil,label:55})
	primitive[54] = wrap(&Procedure{name:"$make-cell",required:1,apply:nil,label:54})
	primitive[53] = wrap(&Procedure{name:"$bytecode-run",required:3,apply:nil,label:53})
	primitive[52] = wrap(&Procedure{name:"command-line",required:0,apply:nil,label:52})
	primitive[51] = wrap(&Procedure{name:"exit",required:1,apply:nil,label:51})
	primitive[50] = wrap(&Procedure{name:"eq?",required:2,apply:nil,label:50})
	primitive[49] = wrap(&Procedure{name:"eof-object",required:0,apply:nil,label:49})
	primitive[48] = wrap(&Procedure{name:"unspecified",required:0,apply:nil,label:48})
	primitive[47] = wrap(&Procedure{name:"procedure?",required:1,apply:nil,label:47})
	primitive[46] = wrap(&Procedure{name:"apply",required:1,apply:nil,label:46})
	primitive[45] = wrap(&Procedure{name:"make-string",required:1,apply:nil,label:45})
	primitive[44] = wrap(&Procedure{name:"string-set!",required:3,apply:nil,label:44})
	primitive[43] = wrap(&Procedure{name:"string-ref",required:2,apply:nil,label:43})
	primitive[42] = wrap(&Procedure{name:"string-length",required:1,apply:nil,label:42})
	primitive[41] = wrap(&Procedure{name:"string?",required:1,apply:nil,label:41})
	primitive[40] = wrap(&Procedure{name:"greatest-fixnum",required:0,apply:nil,label:40})
	primitive[39] = wrap(&Procedure{name:"least-fixnum",required:0,apply:nil,label:39})
	primitive[100] = wrap(&Procedure{name:"bitwise-arithmetic-shift-left",required:2,apply:nil,label:100})
	primitive[38] = wrap(&Procedure{name:"bitwise-arithmetic-shift-right",required:2,apply:nil,label:38})
	primitive[37] = wrap(&Procedure{name:"$bitwise-and",required:2,apply:nil,label:37})
	primitive[36] = wrap(&Procedure{name:"$bitwise-ior",required:2,apply:nil,label:36})
	primitive[35] = wrap(&Procedure{name:"denominator",required:1,apply:nil,label:35})
	primitive[34] = wrap(&Procedure{name:"$cmp",required:2,apply:nil,label:34})
	primitive[33] = wrap(&Procedure{name:"$-",required:2,apply:nil,label:33})
	primitive[32] = wrap(&Procedure{name:"$*",required:2,apply:nil,label:32})
	primitive[31] = wrap(&Procedure{name:"$/",required:2,apply:nil,label:31})
	primitive[30] = wrap(&Procedure{name:"$+",required:2,apply:nil,label:30})
	primitive[29] = wrap(&Procedure{name:"$string->number",required:2,apply:nil,label:29})
	primitive[28] = wrap(&Procedure{name:"$number->string",required:2,apply:nil,label:28})
	primitive[27] = wrap(&Procedure{name:"integer?",required:1,apply:nil,label:27})
	primitive[26] = wrap(&Procedure{name:"number?",required:1,apply:nil,label:26})
	primitive[25] = wrap(&Procedure{name:"vector",required:0,apply:nil,label:25})
	primitive[24] = wrap(&Procedure{name:"vector-set!",required:3,apply:nil,label:24})
	primitive[23] = wrap(&Procedure{name:"vector-ref",required:2,apply:nil,label:23})
	primitive[22] = wrap(&Procedure{name:"vector-length",required:1,apply:nil,label:22})
	primitive[21] = wrap(&Procedure{name:"make-vector",required:2,apply:nil,label:21})
	primitive[20] = wrap(&Procedure{name:"vector?",required:1,apply:nil,label:20})
	primitive[19] = wrap(&Procedure{name:"char-downcase",required:1,apply:nil,label:19})
	primitive[18] = wrap(&Procedure{name:"char-upcase",required:1,apply:nil,label:18})
	primitive[17] = wrap(&Procedure{name:"char-whitespace?",required:1,apply:nil,label:17})
	primitive[16] = wrap(&Procedure{name:"integer->char",required:1,apply:nil,label:16})
	primitive[15] = wrap(&Procedure{name:"char->integer",required:1,apply:nil,label:15})
	primitive[14] = wrap(&Procedure{name:"char?",required:1,apply:nil,label:14})
	primitive[13] = wrap(&Procedure{name:"string->symbol",required:1,apply:nil,label:13})
	primitive[12] = wrap(&Procedure{name:"symbol->string",required:1,apply:nil,label:12})
	primitive[11] = wrap(&Procedure{name:"symbol?",required:1,apply:nil,label:11})
	primitive[10] = wrap(&Procedure{name:"null?",required:1,apply:nil,label:10})
	primitive[9] = wrap(&Procedure{name:"set-cdr!",required:2,apply:nil,label:9})
	primitive[8] = wrap(&Procedure{name:"set-car!",required:2,apply:nil,label:8})
	primitive[7] = wrap(&Procedure{name:"length",required:1,apply:nil,label:7})
	primitive[6] = wrap(&Procedure{name:"floyd",required:1,apply:nil,label:6})
	primitive[5] = wrap(&Procedure{name:"cdr",required:1,apply:nil,label:5})
	primitive[4] = wrap(&Procedure{name:"car",required:1,apply:nil,label:4})
	primitive[3] = wrap(&Procedure{name:"cons",required:2,apply:nil,label:3})
	primitive[2] = wrap(&Procedure{name:"pair?",required:1,apply:nil,label:2})
	primitive[1] = wrap(&Procedure{name:"not",required:1,apply:nil,label:1})
	primitive[0] = wrap(&Procedure{name:"boolean?",required:1,apply:nil,label:0})
}

func evprimn(primop uint32, args []Obj, ct Obj) Obj {
	switch primop {
	case 98: // stop-cpu-profile
		return stop_cpu_profile()
	case 97: // start-cpu-profile
		return start_cpu_profile(args[0])
	case 96: // current-thread
		return ct
	case 95: // $receive
		return _receive(args[0])
	case 94: // send
		return send(args[0], args[1])
	case 93: // thread-link!
		return thread_link_ex(args[0], args[1])
	case 92: // thread-start!
		return thread_start_ex(args[0])
	case 91: // thread-yield!
		return thread_yield_ex()
	case 102: // thread-queue-set!
		return thread_queue_set_ex(args[0], args[1])
	case 101: // thread-queue
		return thread_queue(args[0])
	case 90: // thread-specific-set!
		return thread_specific_set_ex(args[0], args[1])
	case 89: // thread-specific
		return thread_specific(args[0])
	case 88: // thread-name
		return thread_name(args[0])
	case 87: // thread?
		return thread_p(args[0])
	case 86: // $make-thread
		return _make_thread(args[0], args[1])
	case 85: // $write
		return write(args[0], args[1])
	case 84: // $display
		return display(args[0], args[1])
	case 83: // lookahead-u8
		return lookahead_u8(args[0])
	case 82: // put-bytevector
		return put_bytevector(args[0], args[1])
	case 81: // put-u8
		return put_u8(args[0], args[1])
	case 80: // get-u8
		return get_u8(args[0])
	case 79: // $write-char
		return _write_char(args[0], args[1])
	case 78: // $peek-char
		return _peek_char(args[0])
	case 77: // $read-char
		return _read_char(args[0])
	case 76: // close-port
		return close_port(args[0])
	case 75: // close-output-port
		return close_output_port(args[0])
	case 74: // close-input-port
		return close_input_port(args[0])
	case 73: // open-file-output-port
		return open_file_output_port(args[0])
	case 99: // open-output-file
		return open_output_file(args[0])
	case 72: // open-input-file
		return open_input_file(args[0])
	case 71: // delete-file
		return delete_file(args[0])
	case 70: // file-exists?
		return file_exists_p(args[0])
	case 69: // current-output-port
		return current_output_port()
	case 68: // current-input-port
		return current_input_port()
	case 67: // output-port?
		return output_port_p(args[0])
	case 66: // input-port?
		return input_port_p(args[0])
	case 65: // port?
		return port_p(args[0])
	case 64: // $bytevector-output-port-extract
		return _bytevector_output_port_extract(args[0])
	case 63: // $open-bytevector-output-port
		return _open_bytevector_output_port()
	case 62: // string->utf8
		return string_to_utf8(args[0])
	case 61: // u8-list->bytevector
		return u8_list_to_bytevector(args[0])
	case 60: // bytevector-length
		return bytevector_length(args[0])
	case 59: // bytevector?
		return bytevector_p(args[0])
	case 58: // $global-set!
		name := args[0]
		value := args[1]
		sname := (name).(string)
		env[sname] = value
		return Void
	case 57: // $global-ref
		name := args[0]
		sname := (name).(string)
		v, def := env[sname]
		if !def { panic(fmt.Sprintf("undefined top-level variable: %s",sname)) }
		return v
	case 56: // $cell-set!
		v := args[0]
		vv := (v).(*[1]Obj)
		vv[0] = args[1]
		return Void
	case 55: // $cell-ref
		v := args[0]
		vv := (v).(*[1]Obj)
		return vv[0]
	case 54: // $make-cell
		var v [1]Obj
		v[0] = args[0]
		var vv interface{} = &v
		return Obj(vv)
	case 53: // $bytecode-run
		return _bytecode_run(args[0], args[1], args[2])
	case 52: // command-line
		return Command_line()
	case 51: // exit
		os.Exit(number_to_int(args[0]))
	case 50: // eq?
		if args[0] == args[1] {
			return True
		} else {
			return False
		}
	case 49: // eof-object
		return Eof
	case 48: // unspecified
		return Void
	case 47: // procedure?
		return procedure_p(args[0])
	case 46: // apply
		return apply(args,ct)
	case 45: // make-string
		switch len(args) {
		default: return Make_string(args[0],args[1])
		case 1: return Make_string(args[0],Make_char(32))
		}
	case 44: // string-set!
		return String_set_ex(args[0], args[1], args[2])
	case 43: // string-ref
		return String_ref(args[0], args[1])
	case 42: // string-length
		return String_length(args[0])
	case 41: // string?
		return string_p(args[0])
	case 40: // greatest-fixnum
		return Make_fixnum(fixnum_max)
	case 39: // least-fixnum
		return Make_fixnum(fixnum_min)
	case 100: // bitwise-arithmetic-shift-left
		return bitwise_arithmetic_shift_left(args[0], args[1])
	case 38: // bitwise-arithmetic-shift-right
		return bitwise_arithmetic_shift_right(args[0], args[1])
	case 37: // $bitwise-and
		return bitwise_and(args[0], args[1])
	case 36: // $bitwise-ior
		return bitwise_ior(args[0], args[1])
	case 35: // denominator
		return denominator(args[0])
	case 34: // $cmp
		return number_cmp(args[0], args[1])
	case 33: // $-
		return number_subtract(args[0], args[1])
	case 32: // $*
		return number_multiply(args[0], args[1])
	case 31: // $/
		return number_divide(args[0], args[1])
	case 30: // $+
		return number_add(args[0], args[1])
	case 29: // $string->number
		return _string_to_number(args[0], args[1])
	case 28: // $number->string
		return _number_to_string(args[0], args[1])
	case 27: // integer?
		return integer_p(args[0])
	case 26: // number?
		return number_p(args[0])
	case 25: // vector
		return wrap(args)
	case 24: // vector-set!
		return Vector_set_ex(args[0], args[1], args[2])
	case 23: // vector-ref
		return Vector_ref(args[0], args[1])
	case 22: // vector-length
		return vector_length(args[0])
	case 21: // make-vector
		return Make_vector(args[0], args[1])
	case 20: // vector?
		return vector_p(args[0])
	case 19: // char-downcase
		return char_downcase(args[0])
	case 18: // char-upcase
		return char_upcase(args[0])
	case 17: // char-whitespace?
		return char_whitespace_p(args[0])
	case 16: // integer->char
		return integer_to_char(args[0])
	case 15: // char->integer
		return char_to_integer(args[0])
	case 14: // char?
		return char_p(args[0])
	case 13: // string->symbol
		return String_to_symbol(args[0])
	case 12: // symbol->string
		return Symbol_to_string(args[0])
	case 11: // symbol?
		return symbol_p(args[0])
	case 10: // null?
		if args[0] ==  Eol {
			return True
		} else {
			return False
		}
	case 9: // set-cdr!
		return set_cdr_ex(args[0], args[1])
	case 8: // set-car!
		return set_car_ex(args[0], args[1])
	case 7: // length
		return Length(args[0])
	case 6: // floyd
		return Floyd(args[0])
	case 5: // cdr
		return cdr(args[0])
	case 4: // car
		return car(args[0])
	case 3: // cons
		return Cons(args[0], args[1])
	case 2: // pair?
		return pair_p(args[0])
	case 1: // not
		return not(args[0])
	case 0: // boolean?
		return boolean_p(args[0])
	default:
		fmt.Fprintf(os.Stderr, "Please regenerate primitives.go\n")
		panic(fmt.Sprintf("Unimplemented primitive: %s",primop))
	}
	panic(fmt.Sprintf("Fell off the edge in evprimn(): %s",primop))
}

