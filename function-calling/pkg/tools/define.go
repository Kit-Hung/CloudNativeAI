package tools

import "github.com/sashabaranov/go-openai"

var AddToolDefine = openai.Tool{
	Type: DefineType,
	Function: &openai.FunctionDefinition{
		Name: AddToolName,
		Description: `
		Use this tool for addition calculations.
			example:
				1+2=?
			then Action Input is: 1, 2
		`,
		Parameters: `
		{
			"type": "object",
			"properties": {
				"numbers": {
					"type": "array",
					"items": {"type": "integer"}
				},
			},
		}
		`,
	},
}

var SubToolDefine = openai.Tool{
	Type: DefineType,
	Function: &openai.FunctionDefinition{
		Name: SubToolName,
		Description: `
		Use this tool for subtraction calculations.
			example:
				1-2=?
			then Action Input is: 1, 2
		`,
		Parameters: `
		{
			"type": "object",
			"properties": {
				"numbers": {
					"type": "array",
					"items": {"type": "integer"}
				},
			},
		}
		`,
	},
}
