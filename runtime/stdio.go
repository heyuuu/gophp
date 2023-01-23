// <<generate>>

package runtime

// Source: <runtime/stdio.h>

type __sFILE = FILE

// #define BUFSIZ       1024

// #define EOF       ( - 1 )

/* must be == _POSIX_STREAM_MAX <limits.h> */

// #define FOPEN_MAX       20

// #define FILENAME_MAX       1024

// #define SEEK_SET       0

// #define SEEK_CUR       1

// #define SEEK_END       2

var Clearerr func(*FILE)
var Fclose func(*FILE) int
var Feof func(*FILE) int
var Ferror func(*FILE) int
var Fflush func(*FILE) int
var Fgetc func(*FILE) int
var Fgetpos func(*FILE, *fpos_t) int
var Fgets func(*byte, int, *FILE) *byte
var Fopen func(__filename *byte, __mode *byte) *FILE
var Fprintf func(*FILE, *byte, ...any) int
var Fputc func(int, *FILE) int
var Fputs func(*byte, *FILE) int
var Fread func(__ptr any, __size int, __nitems int, __stream *FILE) int
var Freopen func(*byte, *byte, *FILE) *FILE
var Fscanf func(*FILE, *byte, ...any) int
var Fseek func(*FILE, long, int) int
var Fsetpos func(*FILE, *fpos_t) int
var Ftell func(*FILE) long
var Fwrite func(__ptr any, __size int, __nitems int, __stream *FILE) int
var Getc func(*FILE) int
var Getchar func() int
var Gets func(*byte) *byte
var Perror func(*byte)
var Printf func(*byte, ...any) int
var Putc func(int, *FILE) int
var Putchar func(int) int
var Puts func(*byte) int
var Remove func(*byte) int
var Rename func(__old *byte, __new *byte) int
var Rewind func(*FILE)
var Scanf func(*byte, ...any) int
var Setbuf func(*FILE, *byte)
var Setvbuf func(*FILE, *byte, int, int) int
