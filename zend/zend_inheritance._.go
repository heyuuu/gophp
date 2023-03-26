package zend

/* Instanceof that's safe to use on unlinked classes. */

/* Unresolved means that class declarations that are currently not available are needed to
 * determine whether the inheritance is valid or not. At runtime UNRESOLVED should be treated
 * as an ERROR. */

type InheritanceStatus = int

const (
	INHERITANCE_UNRESOLVED InheritanceStatus = -1
	INHERITANCE_ERROR                        = 0
	INHERITANCE_SUCCESS                      = 1
)
const MAX_ABSTRACT_INFO_CNT = 3
const MAX_ABSTRACT_INFO_FMT = "%s%s%s%s"

type VarianceObligationType = int

const (
	OBLIGATION_DEPENDENCY = iota
	OBLIGATION_COMPATIBILITY
	OBLIGATION_PROPERTY_COMPATIBILITY
)
