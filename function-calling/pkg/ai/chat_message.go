package ai

import "github.com/sashabaranov/go-openai"

type ChatMessage []openai.ChatCompletionMessage

func (cm *ChatMessage) Clear() {
	*cm = make([]openai.ChatCompletionMessage, 0)
}

// Add 添加角色和对应的 prompt
func (cm *ChatMessage) Add(role, msg string, toolCalls []openai.ToolCall) {
	*cm = append(*cm, openai.ChatCompletionMessage{
		Role:      role,
		Content:   msg,
		ToolCalls: toolCalls,
	})
}

// AddForTool 添加 Tool 角色的 prompt
func (cm *ChatMessage) AddForTool(msg, name, toolCallID string) {
	*cm = append(*cm, openai.ChatCompletionMessage{
		Role:       RoleTool,
		Content:    msg,
		Name:       name,
		ToolCallID: toolCallID,
	})
}

func (cm *ChatMessage) ToMessages() []openai.ChatCompletionMessage {
	msgs := make([]openai.ChatCompletionMessage, len(*cm))
	for i, msg := range *cm {
		msgs[i] = msg
	}
	return msgs
}
