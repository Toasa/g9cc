package token

import (
    . "g9cc/common"
    . "g9cc/util"
    "fmt"
    "strings"
)

// Tokenizer
func add_token(v *Vector, ty int, input string) *Token {
    t := new(Token)
    t.Ty = ty
    t.Input = input
    Vec_push(v, t)
    return t
}

var keywords *Map

func Scan(s string) *Vector {
    var v *Vector = New_vec()

    // index of input
    i_input := 0

    for s[i_input] != '\000' {

        // white space
        if isspace(s[i_input]) {
            i_input++
            continue
        }

        // single-letter token
        if strings.Contains("+-*/;", string(s[i_input])) {
            add_token(v, int(s[i_input]), string(s[i_input]))
            i_input++
            continue
        }

        // keyword
        if isalpha(s[i_input]) || s[i_input] == '_' {

            len := 1
            for i := len + i_input; isalpha(s[i]) || isdigit(s[i]) || s[i] == '_'; {
                len++
                i = len + i_input
            }
            var name string = s[i_input:len + i_input]

            ty, _ := Map_get(keywords, name).(int)
            if ty == 0 {
                Error(fmt.Sprintf("unknown identifier: %s", name))
            }

            add_token(v, ty, name)
            i_input += len
            continue
        }

        // number
        if isdigit(s[i_input]) {
            var num int = int(s[i_input] - '0')
            i_input++
            for ; isdigit(s[i_input]); i_input++ {
                num = num * 10 + int(s[i_input] - '0')
            }

            var t *Token = add_token(v, TK_NUM, string(num))

            t.Val = num
            continue
        }

        fmt.Println("what's up guys")
        Error(fmt.Sprintf("cannot tokenize: %s", s));
    }

    add_token(v, TK_EOF, s);
    return v
}

func Tokenize(s string) *Vector {
    keywords = New_map()
    Map_put(keywords, "return", TK_RETURN)

    return Scan(s)
}

func isdigit(c uint8) bool {
    if '0' <= c && c <= '9' {
        return true
    } else {
        return false
    }
}

func isspace(c uint8) bool {
    if c == ' ' {
        return true
    } else {
        return false
    }
}

func isalpha(c uint8) bool {
    if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
        return true
    } else {
        return false
    }
}
