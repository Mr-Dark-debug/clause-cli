package utils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Truncate truncates a string to maxLen characters, adding "..." if truncated.
func Truncate(s string, maxLen int) string {
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}

	if maxLen <= 3 {
		return "..."
	}

	runes := []rune(s)
	truncated := runes[:maxLen-3]
	return string(truncated) + "..."
}

// TruncateWithEllipsis truncates a string with custom ellipsis.
func TruncateWithEllipsis(s string, maxLen int, ellipsis string) string {
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}

	ellipsisLen := utf8.RuneCountInString(ellipsis)
	if maxLen <= ellipsisLen {
		return ellipsis[:maxLen]
	}

	runes := []rune(s)
	truncated := runes[:maxLen-ellipsisLen]
	return string(truncated) + ellipsis
}

// PadLeft pads a string on the left to reach the specified length.
func PadLeft(s string, pad string, length int) string {
	if utf8.RuneCountInString(s) >= length {
		return s
	}

	padRunes := []rune(pad)
	if len(padRunes) == 0 {
		padRunes = []rune(" ")
	}

	targetRunes := []rune(s)
	padLen := length - len(targetRunes)

	result := make([]rune, length)
	for i := 0; i < padLen; i++ {
		result[i] = padRunes[i%len(padRunes)]
	}
	copy(result[padLen:], targetRunes)

	return string(result)
}

// PadRight pads a string on the right to reach the specified length.
func PadRight(s string, pad string, length int) string {
	if utf8.RuneCountInString(s) >= length {
		return s
	}

	padRunes := []rune(pad)
	if len(padRunes) == 0 {
		padRunes = []rune(" ")
	}

	targetRunes := []rune(s)
	_ = length - len(targetRunes) // padLen - used for logic below

	result := make([]rune, length)
	copy(result, targetRunes)
	for i := len(targetRunes); i < length; i++ {
		result[i] = padRunes[(i-len(targetRunes))%len(padRunes)]
	}

	return string(result)
}

// Center centers a string within a field of the specified length.
func Center(s string, length int) string {
	strLen := utf8.RuneCountInString(s)
	if strLen >= length {
		return s
	}

	leftPad := (length - strLen) / 2
	rightPad := length - strLen - leftPad

	return strings.Repeat(" ", leftPad) + s + strings.Repeat(" ", rightPad)
}

// TitleCase converts a string to Title Case.
func TitleCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

// CamelCase converts a string to camelCase.
func CamelCase(s string) string {
	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}

	result := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		result += TitleCase(words[i])
	}
	return result
}

// PascalCase converts a string to PascalCase.
func PascalCase(s string) string {
	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}

	result := ""
	for _, word := range words {
		result += TitleCase(word)
	}
	return result
}

// SnakeCase converts a string to snake_case.
func SnakeCase(s string) string {
	words := splitWords(s)
	return strings.ToLower(strings.Join(words, "_"))
}

// ScreamingSnakeCase converts a string to SCREAMING_SNAKE_CASE.
func ScreamingSnakeCase(s string) string {
	words := splitWords(s)
	return strings.ToUpper(strings.Join(words, "_"))
}

// KebabCase converts a string to kebab-case.
func KebabCase(s string) string {
	words := splitWords(s)
	return strings.ToLower(strings.Join(words, "-"))
}

// TrainCase converts a string to Train-Case.
func TrainCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		words[i] = TitleCase(word)
	}
	return strings.Join(words, "-")
}

// ConstantCase converts a string to CONSTANT_CASE.
func ConstantCase(s string) string {
	return ScreamingSnakeCase(s)
}

// DotCase converts a string to dot.case.
func DotCase(s string) string {
	words := splitWords(s)
	return strings.ToLower(strings.Join(words, "."))
}

// PathCase converts a string to path/case.
func PathCase(s string) string {
	words := splitWords(s)
	return strings.ToLower(strings.Join(words, "/"))
}

// splitWords splits a string into words based on various delimiters and case changes.
func splitWords(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}

	// Replace common separators with spaces
	replacer := strings.NewReplacer(
		"_", " ",
		"-", " ",
		".", " ",
		"/", " ",
		"\\", " ",
	)
	s = replacer.Replace(s)

	// Handle camelCase/PascalCase by inserting spaces before uppercase letters
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) && (unicode.IsLower(rune(s[i-1])) || (i < len(s)-1 && unicode.IsLower(rune(s[i+1])))) {
			result.WriteRune(' ')
		}
		result.WriteRune(r)
	}

	// Split on whitespace and filter empty strings
	fields := strings.Fields(result.String())
	words := make([]string, 0, len(fields))
	for _, field := range fields {
		if field != "" {
			words = append(words, field)
		}
	}

	return words
}

// IsEmptyStr checks if a string is empty or contains only whitespace.
func IsEmptyStr(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsBlank checks if a string contains only whitespace.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotEmptyStr checks if a string is not empty and not just whitespace.
func IsNotEmptyStr(s string) bool {
	return strings.TrimSpace(s) != ""
}

// DefaultIfEmpty returns the default value if the string is empty.
func DefaultIfEmpty(s, defaultValue string) string {
	if IsEmptyStr(s) {
		return defaultValue
	}
	return s
}

// TrimPrefix removes a prefix from a string if present.
func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// TrimSuffix removes a suffix from a string if present.
func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// RemoveStr removes all occurrences of substr from s.
func RemoveStr(s, substr string) string {
	return strings.ReplaceAll(s, substr, "")
}

// ReverseStr reverses a string.
func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Repeat repeats a string n times.
func Repeat(s string, n int) string {
	return strings.Repeat(s, n)
}

// ContainsStr checks if a string contains a substring (case-sensitive).
func ContainsStr(s, substr string) bool {
	return strings.Contains(s, substr)
}

// ContainsIgnoreCase checks if a string contains a substring (case-insensitive).
func ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// HasPrefix checks if a string starts with a prefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix checks if a string ends with a suffix.
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Capitalize capitalizes the first character of a string.
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Uncapitalize lowercases the first character of a string.
func Uncapitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// Upper converts a string to uppercase.
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Lower converts a string to lowercase.
func Lower(s string) string {
	return strings.ToLower(s)
}

// SwapCase swaps the case of each character in a string.
func SwapCase(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}
	return string(runes)
}

// CountStr counts the number of occurrences of substr in s.
func CountStr(s, substr string) int {
	return strings.Count(s, substr)
}

// IsNumeric checks if a string contains only numeric characters.
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsAlpha checks if a string contains only alphabetic characters.
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsAlphaNumeric checks if a string contains only alphanumeric characters.
func IsAlphaNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsASCII checks if a string contains only ASCII characters.
func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// Abbreviate abbreviates a string to maxLen, keeping the end.
func Abbreviate(s string, maxLen int) string {
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}

	if maxLen <= 3 {
		return "..."
	}

	runes := []rune(s)
	return string(runes[:maxLen-3]) + "..."
}

// AbbreviateMiddle abbreviates the middle of a string.
func AbbreviateMiddle(s string, maxLen int, middle string) string {
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}

	middleLen := utf8.RuneCountInString(middle)
	if maxLen <= middleLen {
		return middle
	}

	runes := []rune(s)
	available := maxLen - middleLen
	leftLen := available / 2
	rightLen := available - leftLen

	return string(runes[:leftLen]) + middle + string(runes[len(runes)-rightLen:])
}

// Wrap wraps text to a maximum line length.
func Wrap(s string, maxLen int) string {
	if maxLen <= 0 {
		return s
	}

	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}

	var result strings.Builder
	lineLen := 0

	for _, word := range words {
		wordLen := utf8.RuneCountInString(word)

		if lineLen > 0 && lineLen+1+wordLen > maxLen {
			result.WriteString("\n")
			lineLen = 0
		}

		if lineLen > 0 {
			result.WriteString(" ")
			lineLen++
		}

		result.WriteString(word)
		lineLen += wordLen
	}

	return result.String()
}

// Indent indents each line of a string with the given prefix.
func Indent(s, prefix string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = prefix + line
		}
	}
	return strings.Join(lines, "\n")
}

// Dedent removes common leading whitespace from each line.
func Dedent(s string) string {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return s
	}

	// Find minimum common indentation
	minIndent := -1
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		indent := 0
		for _, r := range line {
			if r == ' ' || r == '\t' {
				indent++
			} else {
				break
			}
		}
		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}

	if minIndent <= 0 {
		return s
	}

	// Remove common indentation
	for i, line := range lines {
		if strings.TrimSpace(line) != "" && len(line) >= minIndent {
			lines[i] = line[minIndent:]
		}
	}

	return strings.Join(lines, "\n")
}

// StripPrefix removes a prefix if present, with optional case-insensitive matching.
func StripPrefix(s, prefix string, ignoreCase bool) string {
	if ignoreCase {
		if strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix)) {
			return s[len(prefix):]
		}
		return s
	}
	return strings.TrimPrefix(s, prefix)
}

// StripSuffix removes a suffix if present, with optional case-insensitive matching.
func StripSuffix(s, suffix string, ignoreCase bool) string {
	if ignoreCase {
		if strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix)) {
			return s[:len(s)-len(suffix)]
		}
		return s
	}
	return strings.TrimSuffix(s, suffix)
}

// Ellipsis adds ellipsis to a string if it exceeds maxLen.
func Ellipsis(s string, maxLen int) string {
	return Truncate(s, maxLen)
}

// Quote quotes a string with double quotes.
func Quote(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `\"`) + `"`
}

// SingleQuote quotes a string with single quotes.
func SingleQuote(s string) string {
	return `'` + strings.ReplaceAll(s, `'`, `\'`) + `'`
}

// Unquote removes surrounding quotes from a string.
func Unquote(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// Match reports whether a string matches a regex pattern.
func Match(s, pattern string) (bool, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.MatchString(s), nil
}

// MustMatch is like Match but panics on error.
func MustMatch(s, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

// Extract extracts the first match of a regex pattern from a string.
func Extract(s, pattern string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return re.FindString(s), nil
}

// ExtractAll extracts all matches of a regex pattern from a string.
func ExtractAll(s, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.FindAllString(s, -1), nil
}

// TruncateText is an alias for Truncate for compatibility.
func TruncateText(s string, maxLen int) string {
	return Truncate(s, maxLen)
}

// FormatNumber formats a number with thousand separators.
func FormatNumber(n int64) string {
	if n < 0 {
		return "-" + FormatNumber(-n)
	}

	s := strconv.FormatInt(n, 10)
	result := ""
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += ","
		}
		result += string(c)
	}
	return result
}

// MatchesRegex checks if a string matches a regex pattern.
// Returns true if the string matches, false otherwise.
func MatchesRegex(s, pattern string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(s)
}
