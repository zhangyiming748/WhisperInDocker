package constant

var (
	root     string = "/mnt/c/Users/zen/Github/FastYt-dlp/joi"
	language string = "en"
	pattern  string = "mp4"
	model    string = "large"
)

func GetRoot() string {
	return root
}

func GetLanguage() string {
	return language
}
func GetPattern() string {
	return pattern
}
func GetModel() string {
	return model
}
func SetRoot(s string) {
	root = s
}
func SetLanguage(s string) {
	language = s
}
func SetPattern(s string) {
	pattern = s
}
func SetModel(s string) {
	model = s
}
