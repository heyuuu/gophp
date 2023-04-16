package date

import (
	"fmt"
	"strings"
	"time"
)

func Format(format string, t time.Time, localtime bool) string {
	if format == "" {
		return ""
	}
	//if localtime {
	//	zone, offset := t.Zone()
	//
	//}

	var buf strings.Builder
	for _, c := range []byte(format) {
		switch c {
		// day
		case 'd':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Day())
		case 'D':
			_, _ = fmt.Fprint(&buf, dayShortNames[t.Weekday()])
		case 'j':
			_, _ = fmt.Fprintf(&buf, "%d", t.Day())
		case 'l':
			_, _ = fmt.Fprint(&buf, dayFullNames[t.Weekday()])
		case 'S':
			_, _ = fmt.Fprintf(&buf, englishSuffix(t.Day()))
		case 'w':
			_, _ = fmt.Fprintf(&buf, "%d", t.Weekday())
		case 'N':
			_, _ = fmt.Fprintf(&buf, "%d", (t.Weekday()+6)%7+1)
		case 'z':
			_, _ = fmt.Fprintf(&buf, "%d", t.YearDay()-1)
		// week
		case 'W':
			_, week := t.ISOWeek()
			_, _ = fmt.Fprintf(&buf, "%02d", week)
		// mouth
		case 'F':
			_, _ = fmt.Fprintf(&buf, monFullNames[t.Month()-1])
		case 'm':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Month())
		case 'M':
			_, _ = fmt.Fprintf(&buf, monShortNames[t.Month()-1])
		case 'n':
			_, _ = fmt.Fprintf(&buf, "%d", t.Month())
		case 't':
			_, _ = fmt.Fprintf(&buf, "%02d", daysInMouth(t.Year(), int(t.Month())))
		// year
		case 'L':
			_, _ = fmt.Fprintf(&buf, "%d", cond(isLeapYear(t.Year()), 1, 0))
		case 'o':
			year, _ := t.ISOWeek()
			_, _ = fmt.Fprintf(&buf, "%d", year)
		case 'Y':
			_, _ = fmt.Fprintf(&buf, "%04d", t.Year())
		case 'y':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Year()%100)
		// time
		case 'a':
			_, _ = fmt.Fprintf(&buf, cond(t.Hour() >= 12, "pm", "am"))
		case 'A':
			_, _ = fmt.Fprintf(&buf, cond(t.Hour() >= 12, "PM", "AM"))
		case 'B':
			// Swatch 互联网时间 (Swatch Internal time)
			// - 以Swatch总部(UTC+1)为基准时间，将一天划分为 1000 等分，称为拍(peat)
			unix := t.Unix()
			peat := (unix%86400 + 3600) * 10
			if peat < 0 {
				peat += 864000
			}
			peat = peat / 864 % 1000
			_, _ = fmt.Fprintf(&buf, "%03d", peat)
		case 'g':
			_, _ = fmt.Fprintf(&buf, "%d", (t.Hour()+11)%12+1)
		case 'G':
			_, _ = fmt.Fprintf(&buf, "%d", t.Hour())
		case 'h':
			_, _ = fmt.Fprintf(&buf, "%02d", (t.Hour()+11)%12+1)
		case 'H':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Hour())
		case 'i':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Minute())
		case 's':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Second())
		case 'u':
			_, _ = fmt.Fprintf(&buf, "%06d", t.Nanosecond()/1000)
		case 'v':
			_, _ = fmt.Fprintf(&buf, "%03d", t.Nanosecond()/10000000)
		// timezone
		case 'e', 'I', 'O', 'P', 'p', 'T', 'Z':
			panic("todo") // todo
		//
		case 'c':
			_, _ = fmt.Fprint(&buf, t.Format("2006-01-02T15:04:05-07:00")) // ISO8601 / RFC3339
		case 'r':
			_, _ = fmt.Fprintf(&buf, t.Format("Mon, 02 Jan 2006 15:04:05 -0700")) // RFC2822 / RFC5322
		case 'U':
			_, _ = fmt.Fprintf(&buf, "%d", t.Unix())
		default:
			buf.WriteByte(c)
		}

	}

	return buf.String()
}
