package tests

import (
	"fmt"
	"regexp"
	"strings"
)

func generateDiff(wanted, wantedReg, output string) string {
	return new(diffCase).diff(wanted, wantedReg, output)
}

type diffCase struct {
	isReg          bool
	ar1, ar2, w    []string
	count1, count2 int
}

func (c *diffCase) diff(wanted, wantedReg, output string) string {
	c.init(wanted, wantedReg, output)
	return c.generateArrayDiff()
}

func (c *diffCase) init(wanted string, wantedReg string, output string) {
	c.isReg = wantedReg != ""
	c.w = strings.Split(wanted, "\n")
	if c.isReg {
		c.ar1 = strings.Split(wantedReg, "\n")
	} else {
		c.ar1 = c.w
	}
	c.ar2 = strings.Split(output, "\n")

	c.count1, c.count2 = len(c.ar1), len(c.ar2)
}

func (c *diffCase) generateArrayDiff() string {
	var lines1, lines2 []int
	l1, l2 := 0, 0
	for l1 < c.count1 && l2 < c.count2 {
		if c.compareLine(l1, l2) {
			l1++
			l2++
		} else {
			c1 := c.countArrayDiff(l1+1, l2, 10)
			c2 := c.countArrayDiff(l1, l2+1, 10)

			var inc1, inc2 bool
			if c1 > c2 {
				inc1 = true
			} else if c2 > 0 {
				inc2 = true
			} else {
				inc1, inc2 = true, true
			}
			if inc1 {
				lines1 = append(lines1, l1)
				l1++
			}
			if inc2 {
				lines2 = append(lines2, l2)
				l2++
			}
		}
	}
	for ; l1 < c.count1; l1++ {
		lines1 = append(lines1, l1)
	}
	for ; l2 < c.count2; l2++ {
		lines2 = append(lines2, l2)
	}

	return c.buildDiffString(lines1, lines2)
}

func (c *diffCase) countArrayDiff(l1 int, l2 int, steps int) int {
	equal := 0
	for l1 < c.count1 && l2 < c.count2 && c.compareLine(l1, l2) {
		l1++
		l2++
		equal++
		steps--
	}
	steps--
	if steps > 0 {
		eq1 := 0
		st := steps/2 + 1
		for i := 1; i <= st && l1+i < c.count1; i++ {
			eq := c.countArrayDiff(l1+i, l2, st-i)
			if eq > eq1 {
				eq1 = eq
			}
		}

		eq2 := 0
		st = steps
		for i := 1; i <= st && l2+i < c.count2; i++ {
			eq := c.countArrayDiff(l1, l2+i, st-i)
			if eq > eq2 {
				eq2 = eq
			}
		}

		if eq1 > eq2 {
			equal += eq1
		} else {
			equal += eq2
		}
	}
	return equal
}

func (c *diffCase) compareLine(l1 int, l2 int) bool {
	line1, line2 := c.ar1[l1], c.ar2[l2]
	if !c.isReg {
		return line1 == line2
	}

	reg, err := regexp.Compile(line1)
	if err != nil {
		return false
	}
	return reg.MatchString(line2)
}

func (c *diffCase) buildDiffString(lines1, lines2 []int) string {
	size := 0
	for _, l1 := range lines1 {
		size += 10 + len(c.w[l1])
	}
	for _, l2 := range lines2 {
		size += 10 + len(c.ar2[l2])
	}

	var buf strings.Builder
	buf.Grow(size)

	lastL1, lastL2 := -2, -2
	for len(lines1) > 0 && len(lines2) > 0 {
		l1, l2 := lines1[0], lines2[0]
		if l1 == lastL1+1 || (l1 < l2 && l2 != lastL2+1) {
			c.writeLine1(&buf, l1)
			lastL1 = l1
			lines1 = lines1[1:]
		} else {
			c.writeLine2(&buf, l2)
			lastL2 = l2
			lines2 = lines2[1:]
		}
	}
	for _, l1 := range lines1 {
		c.writeLine1(&buf, l1)
	}
	for _, l2 := range lines2 {
		c.writeLine2(&buf, l2)
	}
	return buf.String()
}

func (c *diffCase) writeLine1(buf *strings.Builder, idx1 int) {
	_, _ = fmt.Fprintf(buf, "%03d- %s\n", idx1+1, c.w[idx1])
}
func (c *diffCase) writeLine2(buf *strings.Builder, idx2 int) {
	_, _ = fmt.Fprintf(buf, "%03d+ %s\n", idx2+1, c.ar2[idx2])
}
