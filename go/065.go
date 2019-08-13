/*                  |------------ (digit) --------->|<-- (digit) --|
 * |----------- (.) ------------>|                 |--- (.) ---> S4 --(e)-->|                                   
 * S0 -- (+/-) --> S1 -- (.) --> S2 -- (digit) --> S3(digit self) -- (e) --> S5 -- (+/-) --> S6 -- (digit) --> S7(digit self)
 * |                                                |                        |--------------- (digit) -------->|
 * |------------------ (digit) -------------------->|
 */
func isNumber(s string) bool {
    start, end, step, dot := 0, len(s), 0, false
    for start < end && unicode.IsSpace(rune(s[start])) {
        start++
    }
    for start < end && unicode.IsSpace(rune(s[end-1])) {
        end--
    }
    if start == end {
        return false
    }
    for i := start; i < end; i++ {
        if s[i] == '+' || s[i] == '-' {
            if step != 0 && step != 5 {
                return false
            }
            step++
            continue
        }
        if s[i] == '.' {
            if dot {
                return false
            }
            dot = true
            if step == 1 || step == 3 {
                step++
            } else if step == 0 {
                step += 2
            } else {
                return false
            }
            continue
        }
        if s[i] == 'e' {
            if step == 3 {
                step += 2
            } else if step == 4 {
                step++
            } else {
                return false
            }
            continue
        }
        if s[i] >= '0' && s[i] <= '9' {
            switch step {
            case 0, 1, 2, 3, 4:
                step = 3
            case 5, 6, 7:
                step = 7
            }
            continue
        }
        return false
    }
    return step == 3 || step == 4 || step == 7
}
