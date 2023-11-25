func romanToInt(s string) int {
    result := 0
    RtoI := map[byte]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
    }
    i:= 0
    for i < len(s) - 1 {
        if RtoI[s[i+1]] > RtoI[s[i]] {
            result -= RtoI[s[i]]
        } else {
            result += RtoI[s[i]]
        }
        i++
    }
    return result + RtoI[s[len(s) - 1]]
}
