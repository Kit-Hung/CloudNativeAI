package tools

const (
	AddToolName        = "AddTool"
	AddToolDescription = `
	Use this tool for addition calculations.
	example:
		1+2 =?
	then Action Input is: 1,2
	`
	AddToolParam = `
	{
		"type": "object",
		"properties": {
			"numbers": {
				"type": "array",
				"items": {"type": "integer"}
			},
		},
	}
	`
)

const (
	SubToolName        = "SubTool"
	SubToolDescription = `
	Use this tool for subtraction calculations.
	example:
		1-2 =?
	then Action Input is: 1,2
	`
	SubToolParam = `
	{
		"type": "object",
		"properties": {
			"numbers": {
				"type": "array",
				"items": {"type": "integer"}
			},
		},
	}
	`
)
