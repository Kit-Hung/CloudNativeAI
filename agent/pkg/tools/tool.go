package tools

import (
	"github.com/sashabaranov/go-openai"
	"strconv"
	"strings"
)

func AddToolTpl() openai.Tool {
	funDefine := openai.FunctionDefinition{
		Name:        AddToolName,
		Description: AddToolDescription,
		Parameters:  AddToolParam,
	}

	tool := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &funDefine,
	}
	return tool
}

func SubToolTpl() openai.Tool {
	funDefine := openai.FunctionDefinition{
		Name:        SubToolName,
		Description: SubToolDescription,
		Parameters:  SubToolParam,
	}

	tool := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &funDefine,
	}
	return tool
}

func ToolTpl(name, description, param string, tools []openai.Tool) []openai.Tool {
	funDefine := openai.FunctionDefinition{
		Name:        name,
		Description: description,
		Parameters:  param,
	}

	tool := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &funDefine,
	}

	tools = append(tools, tool)
	return tools
}

func AddTool(numbers string) int {
	nums := strings.Split(numbers, ",")
	num0, _ := strconv.Atoi(nums[0])
	num1, _ := strconv.Atoi(nums[1])
	return num0 + num1
}

func SubTool(numbers string) int {
	nums := strings.Split(numbers, ",")
	num0, _ := strconv.Atoi(nums[0])
	num1, _ := strconv.Atoi(nums[1])
	return num0 - num1
}
