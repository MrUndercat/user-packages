package hidespam

func HideSpam(line string) string {
	result := []rune(line)
	for i := 0; i < len(result)-6; i++ {
		if string(result[i:i+7]) == "http://" {
			for j := i + 7; j < len(result) && result[j] != ' '; j++ {
				result[j] = '*'
			}
		}
	}
	return string(result)
}
