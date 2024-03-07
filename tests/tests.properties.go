package tests

// properties for TestCase
func (tc *TestCase) FileName() string {
	return tc.fileName
}
func (tc *TestCase) FilePath() string {
	return tc.filePath
}
func (tc *TestCase) Sections() map[string]string {
	return tc.sections
}
