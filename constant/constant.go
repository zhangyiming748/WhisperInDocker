package constant

var (
	root     string = "/home/zen/docker/yt-dlp"                // 视频所在位置
	location string = "/home/zen/github/WhisperInDocker/model" // 模型所在位置
	language string = "en"                                     // 视频语言
	pattern  string = "webm"                                   // 视频扩展名
	model    string = "medium"                                 // 模型等级
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
func GetLocation() string {
	return location
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
func SetLocation(s string) {
	location = s
}
