// <<generate>>

package cli

import (
	b "sik/builtin"
	"sik/zend"
)

func SavePsArgs(argc int, argv **byte) **byte {
	SaveArgc = argc
	SaveArgv = argv

	/*
	 * If we're going to overwrite the argv area, count the available space.
	 * Also move the environment to make additional room.
	 */

	var end_of_area *byte = nil
	var non_contiguous_area int = 0
	var i int

	/*
	 * check for contiguous argv strings
	 */

	for i = 0; non_contiguous_area == 0 && i < argc; i++ {
		if i != 0 && end_of_area+1 != argv[i] {
			non_contiguous_area = 1
		}
		end_of_area = argv[i] + strlen(argv[i])
	}

	/*
	 * check for contiguous environ strings following argv
	 */

	for i = 0; non_contiguous_area == 0 && Environ[i] != nil; i++ {
		if end_of_area+1 != Environ[i] {
			non_contiguous_area = 1
		}
		end_of_area = Environ[i] + strlen(Environ[i])
	}
	if non_contiguous_area != 0 {
		goto clobber_error
	}
	PsBuffer = argv[0]
	PsBufferSize = end_of_area - argv[0]

	/*
	 * move the environment out of the way
	 */

	NewEnviron = (**byte)(zend.Malloc((i + 1) * b.SizeOf("char *")))
	FrozenEnviron = (**byte)(zend.Malloc((i + 1) * b.SizeOf("char *")))
	if NewEnviron == nil || FrozenEnviron == nil {
		goto clobber_error
	}
	for i = 0; Environ[i] != nil; i++ {
		NewEnviron[i] = strdup(Environ[i])
		if NewEnviron[i] == nil {
			goto clobber_error
		}
	}
	NewEnviron[i] = nil
	Environ = NewEnviron
	memcpy((*byte)(FrozenEnviron), (*byte)(NewEnviron), b.SizeOf("char *")*(i+1))

	/*
	 * If we're going to change the original argv[] then make a copy for
	 * argument parsing purposes.
	 *
	 * (NB: do NOT think to remove the copying of argv[]!
	 * On some platforms, getopt() keeps pointers into the argv array, and
	 * will get horribly confused when it is re-called to analyze a subprocess'
	 * argument string if the argv storage has been clobbered meanwhile.
	 * Other platforms have other dependencies on argv[].)
	 */

	var new_argv **byte
	var i int
	new_argv = (**byte)(zend.Malloc((argc + 1) * b.SizeOf("char *")))
	if new_argv == nil {
		goto clobber_error
	}
	for i = 0; i < argc; i++ {
		new_argv[i] = strdup(argv[i])
		if new_argv[i] == nil {
			zend.Free(new_argv)
			goto clobber_error
		}
	}
	new_argv[argc] = nil

	/*
	 * Darwin (and perhaps other NeXT-derived platforms?) has a static
	 * copy of the argv pointer, which we may fix like so:
	 */

	(*_NSGetArgv)() = new_argv
	argv = new_argv

	/* make extra argv slots point at end_of_area (a NUL) */

	var i int
	for i = 1; i < SaveArgc; i++ {
		SaveArgv[i] = PsBuffer + PsBufferSize
	}
	return argv
clobber_error:

	/* probably can't happen?!
	 * if we ever get here, argv still points to originally passed
	 * in argument
	 */

	SaveArgv = nil
	SaveArgc = 0
	PsBuffer = nil
	PsBufferSize = 0
	return argv
}
func IsPsTitleAvailable() int {
	if SaveArgv == nil {
		return PS_TITLE_NOT_INITIALIZED
	}
	if PsBuffer == nil {
		return PS_TITLE_BUFFER_NOT_AVAILABLE
	}
	return PS_TITLE_SUCCESS
}
func PsTitleErrno(rc int) *byte {
	switch rc {
	case PS_TITLE_SUCCESS:
		return "Success"
	case PS_TITLE_NOT_AVAILABLE:
		return "Not available on this OS"
	case PS_TITLE_NOT_INITIALIZED:
		return "Not initialized correctly"
	case PS_TITLE_BUFFER_NOT_AVAILABLE:
		return "Buffer not contiguous"
	}
	return "Unknown error code"
}
func SetPsTitle(title *byte) int {
	var rc int = IsPsTitleAvailable()
	if rc != PS_TITLE_SUCCESS {
		return rc
	}
	strncpy(PsBuffer, title, PsBufferSize)
	PsBuffer[PsBufferSize-1] = '0'
	PsBufferCurLen = strlen(PsBuffer)

	/* pad unused memory */

	if PsBufferCurLen < PsBufferSize {
		memset(PsBuffer+PsBufferCurLen, PS_PADDING, PsBufferSize-PsBufferCurLen)
	}
	return PS_TITLE_SUCCESS
}
func GetPsTitle(displen *int, string **byte) int {
	var rc int = IsPsTitleAvailable()
	if rc != PS_TITLE_SUCCESS {
		return rc
	}
	*displen = int(PsBufferCurLen)
	*string = PsBuffer
	return PS_TITLE_SUCCESS
}
func CleanupPsArgs(argv **byte) {
	if SaveArgv != nil {
		SaveArgv = nil
		SaveArgc = 0
		var i int
		for i = 0; FrozenEnviron[i] != nil; i++ {
			zend.Free(FrozenEnviron[i])
		}
		zend.Free(FrozenEnviron)
		zend.Free(NewEnviron)

		/* leave a sane environment behind since some atexit() handlers
		   call getenv(). */

		Environ = EmptyEnviron

		/* leave a sane environment behind since some atexit() handlers
		   call getenv(). */

		var i int
		for i = 0; argv[i] != nil; i++ {
			zend.Free(argv[i])
		}
		zend.Free(argv)
	}
	return
}
