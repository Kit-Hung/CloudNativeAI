package tools

const (
	DefineType = "function"

	AddToolName = "AddTool"
	SubToolName = "SubTool"
)

type InputArgs struct {
	Numbers []int `json:"numbers"`
}
