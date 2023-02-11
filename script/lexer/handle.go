package lexer

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type RuleParser struct {
	newLines  []string
	functions []string
	funcIndex int
}

var ruleFuncTpl = `
int %s() {
%s
}
`

func (p *RuleParser) addLine(line string) {
	p.newLines = append(p.newLines, line)
}

func (p *RuleParser) handleFunction(lines []string) {
	if len(lines) <= 3 {
		p.newLines = append(p.newLines, lines...)
		return
	}

	var funcCall = p.newFunction(lines)
	p.addLine(funcCall)
}

func (p *RuleParser) newFunction(lines []string) string {
	p.funcIndex++
	var funcName = "sc_lexer_rule_" + strconv.Itoa(p.funcIndex)
	var funcCode = fmt.Sprintf(ruleFuncTpl, funcName, strings.Join(lines, "\n"))

	p.functions = append(p.functions, funcCode)

	return fmt.Sprintf("    return %s()", funcName)
}

func (p *RuleParser) handle(text string) (string, string) {
	var lines = strings.Split(text, "\n")

	var inRule = false
	var ruleLines []string
	for _, line := range lines {
		if inRule && isRuleEnd(line) {
			p.handleFunction(ruleLines)
			ruleLines = nil
			inRule = false
		}

		if !inRule {
			p.newLines = append(p.newLines, line)
		} else {
			ruleLines = append(ruleLines, line)
		}

		if !inRule && isRuleStart(line) {
			inRule = true
		}
	}
	return strings.Join(p.newLines, "\n"), strings.Join(p.functions, "\n")
}

func handle(text string) (string, string) {
	var parser = &RuleParser{}
	return parser.handle(text)
}

var ruleStartRegexp = regexp.MustCompile("^<[\\w,]+>")

func isRuleStart(line string) bool {
	return ruleStartRegexp.MatchString(line)
}

func isRuleEnd(line string) bool {
	return line == "}"
}
