package internal

type CommandInfo struct {
	Command   string
	ShortInfo string
}

type SupportedCommandType struct {
	Base CommandInfo
	Pull CommandInfo
}

var SupportedCommands = SupportedCommandType{
	Base: CommandInfo{
		Command:   "nomad",
		ShortInfo: "Nomad is a CLI tool that keeps your AI development environment consistent across every machine you work on.",
	},
	Pull: CommandInfo{
		Command:   "pull",
		ShortInfo: "Pull configuration from a storage provider",
	},
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
