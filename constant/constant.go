package constant

var (
	root     string = "E:\\pikpak\\MILF"
	language string = "en"
	pattern  string = "mp4"
	//model    string = "large"
	model string = "base"
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
