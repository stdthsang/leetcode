/*                  |------------ (digit) --------->|<-- (digit) --|
 * |----------- (.) ------------>|                 |--- (.) ---> S4 --(e)-->|                                   
 * S0 -- (+/-) --> S1 -- (.) --> S2 -- (digit) --> S3(digit self) -- (e) --> S5 -- (+/-) --> S6 -- (digit) --> S7(digit self)
 * |                                                |                        |--------------- (digit) -------->|
 * |------------------ (digit) -------------------->|
 */
class Solution {
public:
    bool isNumber(string s) {
        int start = 0, end = s.size(), step = 0;
        while (start < end && isspace(s[start]))
            ++start;
        while (start < end && isspace(s[end-1]))
            --end;
        if (start == end)
            return false;
        bool dot = false;
        for (int i = start; i < end; ++i) {
            if (s[i] == '+' || s[i] == '-') {
                if (step == 0 || step == 5)
                    ++step;
                else return false;
                continue;
            }
            if (s[i] == '.') {
                if (dot)
                    return false;
                dot = true;
                if (step == 1 || step == 3)
                    ++step;
                else if (step == 0)
                    step += 2;
                else return false;
                continue;
            }
            if (s[i] == 'e') {
                if (step == 3)
                    step += 2;
                else if (step == 4)
                    ++step;
                else return false;
                continue;
            }
            if (isdigit(s[i])) {
                switch (step) {
                    case 0:case 1:case 2:case 3:case 4:
                        step = 3;
                        break;
                    case 5:case 6:case 7:
                        step = 7;
                        break;
                }
                continue;
            }
            return false;
        }
        return step == 3 || step == 4 || step == 7;
    }
};
