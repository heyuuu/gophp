package core

import (
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

// todo envGlobals 不是全局变量而应该是请求作用域变量，每个请求有单独的 envGlobals，请求结束后释放(特殊环境变量 TZ* 也需要考虑到)
var baseEnvirons []string = nil
var envGlobals *EnvGlobals = InitEnviron(baseEnvirons)

func Env__() *EnvGlobals { return envGlobals }

type EnvGlobals struct {
	environs []string
	table    *types.Array
}

func InitEnviron(environs []string) *EnvGlobals {
	env := &EnvGlobals{table: types.NewArray(0)}
	for _, environ := range environs {
		env.PutEnv(environ)
	}
	return env
}

func validEnvName(name string) bool {
	for _, c := range name {
		if c == ' ' || c == '.' || c == '[' {
			return false
		}
	}
	return true
}

func (env *EnvGlobals) PutEnv(environ string) bool {
	// 添加环境字符串
	env.environs = append(env.environs, environ)

	// parse variable
	if pos := strings.IndexByte(environ, '='); pos > 0 && validEnvName(environ[:pos]) {
		name := environ[:pos]
		val := environ[pos+1:]
		if num, ok := types.ParseNumericStr(name); ok {
			env.table.IndexUpdate(num, types.NewZvalString(val))
		} else {
			env.table.KeyUpdateIndirect(name, types.NewZvalString(val))
		}
		return true
	}

	return false
}

func (env *EnvGlobals) UnSetEnv(key string) {
	prefix := key + "="
	for i, environ := range env.environs {
		if strings.HasPrefix(environ, prefix) {
			copy(env.environs[i:], env.environs[i+1:])
			env.environs = env.environs[:len(env.environs)-1]
			break
		}
	}

	env.table.KeyDelete(key)
}

func (env *EnvGlobals) LookupEnv(key string) (val string, exists bool) {
	zv := env.table.KeyFind(key)
	if zv != nil {
		return zv.StringVal(), true
	}
	return "", false
}

func (env *EnvGlobals) DupArray() *types.Array {
	return env.table.Copy()
}
