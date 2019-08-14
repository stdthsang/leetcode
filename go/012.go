func intToRoman(num int) string {
    v1 := []string{"", "M", "MM", "MMM"}
    v2 := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
    v3 := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
    v4 := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
    return v1[num/1000] + v2[(num%1000)/100] + v3[(num%100)/10] + v4[num%10]
}
