// Code generated by "stringer -type=Stage -output=enum_string.go"; DO NOT EDIT.

package def

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StageVariableName-1]
	_ = x[StageFuncName-2]
	_ = x[StageMethodName-3]
	_ = x[StageClassName-4]
}

const _Stage_name = "StageVariableNameStageFuncNameStageMethodNameStageClassName"

var _Stage_index = [...]uint8{0, 17, 30, 45, 59}

func (i Stage) String() string {
	i -= 1
	if i < 0 || i >= Stage(len(_Stage_index)-1) {
		return "Stage(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Stage_name[_Stage_index[i]:_Stage_index[i+1]]
}