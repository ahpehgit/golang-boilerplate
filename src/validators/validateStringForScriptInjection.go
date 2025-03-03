package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateStringForScriptInjection(field validator.FieldLevel) bool {
	value := field.Field().String()

	if value == "" {
		return true
	}

	// Comprehensive validation rules for common script injection patterns:
	if isScriptPatternMatch(value, "<script>") ||
		isScriptPatternMatch(value, "</script>") ||
		isScriptPatternMatch(value, "<iframe>") ||
		isScriptPatternMatch(value, "</iframe>") ||
		isScriptPatternMatch(value, "<alert>(.*?)</alert>") ||
		isScriptPatternMatch(value, "javascript:") ||
		isScriptPatternMatch(value, "vbscript:") ||
		isScriptPatternMatch(value, "data:") ||
		isScriptPatternMatch(value, "on[a-z]+=") { // Catch events like onclick, onmouseover, etc.
		//|| strings.Contains(value, "&lt;") ||
		//strings.Contains(value, "&gt;") {
		return false
	}

	return true
}

// Helper function to match potential script patterns more accurately:
func isScriptPatternMatch(value, pattern string) bool {
	r, _ := regexp.Compile(pattern)
	return r.MatchString(value)
}
