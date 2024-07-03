package render

import (
	"github.com/heyuuu/gophp/kits/mapkit"
	"strconv"
	"strings"
)

type pkgName string

func (n pkgName) DefaultAlias() string {
	name := string(n)
	if idx := strings.LastIndexByte(name, '/'); idx >= 0 {
		return name[idx+1:]
	}
	return name
}

type imports struct {
	m    map[pkgName]string
	used map[string]bool
}

func newImports() *imports {
	return &imports{
		m:    make(map[pkgName]string),
		used: make(map[string]bool),
	}
}

func (i *imports) getOrAdd(pkg pkgName) string {
	// get
	if alias, ok := i.m[pkg]; ok {
		return alias
	}

	// add
	defaultAlias := pkg.DefaultAlias()
	alias := defaultAlias
	for num := 2; i.used[alias]; num++ {
		alias = defaultAlias + strconv.Itoa(num)
	}

	i.m[pkg] = alias
	i.used[alias] = true

	return alias
}

func (i *imports) Len() int {
	return len(i.m)
}

func (i *imports) SortedEach(f func(pkg pkgName, alias string)) {
	for _, pkg := range mapkit.SortedKeys(i.m) {
		f(pkg, i.m[pkg])
	}
}
