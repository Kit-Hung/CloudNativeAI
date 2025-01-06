package ai

import "github.com/sashabaranov/go-openai"

type ChatMessage struct {
	Msg openai.ChatCompletionMessage
}

type ChatMessages []*ChatMessage

// Clear 重新初始化并设定人设
func (cm *ChatMessages) Clear(systemRole string) {
	*cm = make(ChatMessages, 0)
	cm.AddForSystem(systemRole)
}

func (cm *ChatMessages) Add(msg, role string) {
	*cm = append(*cm, &ChatMessage{
		Msg: openai.ChatCompletionMessage{
			Role:    role,
			Content: msg,
		},
	})
}

func (cm *ChatMessages) AddForToolCall(resp openai.ChatCompletionMessage, role string) {
	*cm = append(*cm, &ChatMessage{
		Msg: openai.ChatCompletionMessage{
			Role:         role,
			Content:      resp.Content,
			FunctionCall: resp.FunctionCall,
			ToolCalls:    resp.ToolCalls,
		},
	})
}

func (cm *ChatMessages) AddForAssistant(resp openai.ChatCompletionMessage) {
	cm.AddForToolCall(resp, RoleAssistant)
}

func (cm *ChatMessages) AddForUser(msg string) {
	cm.Add(msg, RoleUser)
}

func (cm *ChatMessages) AddForSystem(msg string) {
	cm.Add(msg, RoleSystem)
}

func (cm *ChatMessages) AddForTool(msg, name, toolCallID string) {
	*cm = append(*cm, &ChatMessage{
		Msg: openai.ChatCompletionMessage{
			Role:       RoleTool,
			Content:    msg,
			Name:       name,
			ToolCallID: toolCallID,
		},
	})
}

func (cm *ChatMessages) ToMessage() []openai.ChatCompletionMessage {
	result := make([]openai.ChatCompletionMessage, len(*cm))
	for i, msg := range *cm {
		result[i] = msg.Msg
	}
	return result
}

func (cm *ChatMessages) GetLast() string {
	if len(*cm) == 0 {
		return "什么都没找到"
	}
	return (*cm)[len(*cm)-1].Msg.Content
}
