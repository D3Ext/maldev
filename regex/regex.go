package regex

import (
  "regexp"
  "strings"
)

func FindSubstringBetween(input string, first_delim string, second_delim string) ([]string) {
  var final_out []string

  re := regexp.MustCompile(first_delim + `(.*?)` + second_delim) // Create regular expression
  matches := re.FindAllString(input, -1) // Get all matched
  for _, entry := range matches { // Iterate through them to remove given delimiters from final output
    mod_matches := strings.Split(entry, first_delim)
    mod2_matches := strings.Join(mod_matches, "")
    mod3_matches := strings.Split(mod2_matches, second_delim)
    final_out = append(final_out, mod3_matches[0])
  }

  return final_out // Return a slice so no substrings misses
}






