/*
            2  010              010
            3  011              011
-------------------------------------
   a=(2&3)<<1= 100;    b=(2^3)= 001
-------------------------------------
   a | b = 101
*/
func getSum(a int, b int) int {
    for b != 0 {
        var carry int32 = int32((a & b & 0x7fffffff) << 1)
        a = a ^ b
        b = int(carry)
    }
    return a
}
