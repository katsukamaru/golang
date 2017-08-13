package main

func basedir(s string) string {
	for i := len(s) - 1 ; i >= 0 ;i-- {
		// if s[i] == "/" と　if s[i] == '/' では異なる
		if s[i] == '/' {
			s = s[ i+1:]
			break
		}
	}
	for i := len(s) - 1 ; i >= 0 ;i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
