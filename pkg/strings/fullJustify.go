package strings

import "strings"

func fillWords(words []string, bg, ed, maxWidth int, lastLine bool) string {
	sb := &strings.Builder{}
	wordCount := ed - bg + 1
	spaceCount := maxWidth + 1 - wordCount
	for i := bg; i <= ed; i++ {
		spaceCount -= len(words[i])
	}
	t := 0
	if wordCount != 1 {
		t = spaceCount % (wordCount - 1)
		spaceCount /= wordCount - 1
	}
	for i := bg; i < ed; i++ {
		sb.WriteString(words[i])
		if lastLine {
			sb.WriteByte(' ')
			continue
		}
		spaceSize := spaceCount + 1
		if i - bg < t {
			spaceSize ++
		}
		for j := 0; j < spaceSize; j++ {
			sb.WriteByte(' ')
		}
	}
	sb.WriteString(words[ed])
	lens := sb.Len()
	if lens < maxWidth {
		for i := 0; i < maxWidth - lens; i++ {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}

func FullJustify(words []string, maxWidth int) []string {
	var (
		result []string
		count int
		bg int
	)
	for i := 0; i < len(words); i++ {
		count += len(words[i]) + 1
		if i + 1 == len(words) || count + len(words[i + 1]) > maxWidth {
			result = append(result, fillWords(words, bg, i, maxWidth, i + 1 == len(words)))
			bg = i + 1
			count = 0
		}
	}
	return result
}
