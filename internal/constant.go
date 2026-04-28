package internal

var SupportedCommands = struct {
	Base string
	Pull string
}{
	Base: "nomad",
	Pull: "nomad pull",
}

var SupportedTools = []string{"claude-code"}
var SupportedStorages = []string{"github"}

var ToolFlag = struct {
	Short string
	Long  string
	Usage string
}{
	Short: "t",
	Long:  "tool",
	Usage: "Name of the tool (e.g., claude-code)",
}

var StorageFlag = struct {
	Short string
	Long  string
	Usage string
}{
	Short: "s",
	Long:  "storage",
	Usage: "Name of the storage provider (e.g., github)",
}
