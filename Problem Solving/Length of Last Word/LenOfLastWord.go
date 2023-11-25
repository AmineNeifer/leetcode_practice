func lengthOfLastWord(s string) int {
    splitted := strings.Split(s, " ")
    last_word := splitted[len(splitted) - 1]

    for i := len(splitted) - 1; i >= 0; i-- {
        if strings.ReplaceAll(splitted[i], " ", "")  != "" {
            last_word = splitted[i]
            break
        }
    }
    return len(last_word)
}
