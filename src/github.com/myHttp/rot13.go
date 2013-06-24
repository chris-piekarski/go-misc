package myHttp

// rot13(rot13(x)) = x
// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
// NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm

// Package has two rot13 versions. One is a look up table and the other is a real-time mod operator version.

var table = map[byte]byte{'A':'N','B':'O','C':'P','D':'Q','E':'R','F':'S','G':'T','H':'U','I':'V','J':'W','K':'X','L':'Y','M':'Z','N':'A','O':'B','P':'C','Q':'D','R':'E','S':'F','T':'G','U':'H','V':'I','W':'J','X':'K','Y':'L','Z':'M','a':'n','b':'o','c':'p','d':'q','e':'r','f':'s','g':'t','h':'u','i':'v','j':'w','k':'x','l':'y','m':'z','n':'a','o':'b','p':'c','q':'d','r':'e','s':'f','t':'g','u':'h','v':'i','w':'j','x':'k','y':'l','z':'m'} 

func Rot13Table(input []byte) {
	for i, c := range input {
		input[i] = table[c]
	}
}

func Rot13Mod(input []byte) {
        for i, c := range input {
                base := byte(0)
                if c >= 'a' && c <= 'z' {
                        base = 'a'
                }else if c >= 'A' && c <= 'Z' {
                        base = 'A'
                }
                if base != 0 {
                        input[i] = byte((c - base + 13) % 26 + base)
                }
        }
}
