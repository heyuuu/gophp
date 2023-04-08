package file

/**
 * FILE
 */
type FILE struct{}

const BUFSIZ = 1024
const EOF = -1

/* must be == _POSIX_STREAM_MAX <limits.h> */

const SEEK_SET = 0
const SEEK_CUR = 1
const SEEK_END = 2

var Fclose func(*FILE) int
var Feof func(*FILE) int
var Fflush func(*FILE) int
var Fgetc func(*FILE) int
var Fopen func(__filename string, __mode string) *FILE
var Fprintf func(*FILE, *byte, ...any) int
var Fread func(__ptr any, __size int, __nitems int, __stream *FILE) int
var Fseek func(*FILE, int64, int) int
var Ftell func(*FILE) int64
var Fwrite func(__ptr any, __size int, __nitems int, __stream *FILE) int
var Getc func(*FILE) int
var Perror func(*byte)
var Printf func(*byte, ...any) int
var Rename func(__old *byte, __new *byte) int
var Setvbuf func(*FILE, *byte, int, int) int
