package main
/*
	单元测试。
	命令：go test
 */
import (
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {
// 	testcases := []struct {
// 		in, want string
// 	} {
// 		{"hello, world", "dlrow ,olleh"},
// 		{" ", " "},
// 		{"!12345", "54321!"},
// 	}
// 	for _, tc := range testcases {
// 		rev := Reverse(tc.in)
// 	if rev != tc.want {
// 		t.Errorf("Reverse: %q, want %q", rev, tc.want)
// 	}
// 	}
// }

func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus。fuzzing基于语料库生成随机输入
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
        doubleRev, err2 := Reverse(rev)
		// t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
        if err2 != nil {
			return
		}
		if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}