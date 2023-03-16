// <<generate>>

package core

import "math"

func ZendIsnan(f float64) bool  { return math.IsNaN(f) }
func ZendIsInf(f float64) bool  { return math.IsInf(f, 1) || math.IsInf(f, -1) }
func ZendFinite(f float64) bool { return !ZendIsInf(f) }
