package string

import "unicode/utf8"

func matchAuctor(s string) (end int) {
	end = -1
	var r rune
	var rlen int
	i := 0
	_, _, _ = r, rlen, i
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 97:
		goto s2
	}
	return
s2:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 117:
		goto s3
	}
	return
s3:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 99:
		goto s4
	}
	return
s4:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 116:
		goto s5
	}
	return
s5:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 111:
		goto s6
	}
	return
s6:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 114:
		end = i
	}
	return
}
