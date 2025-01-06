package main

import (
	ai2 "Kit-Hung/cloud-native-AI/function-calling/pkg/ai"
	tools2 "Kit-Hung/cloud-native-AI/function-calling/pkg/tools"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"strconv"
)

func main() {
	toolList := make([]openai.Tool, 0)
	toolList = append(toolList, tools2.AddToolDefine, tools2.SubToolDefine)

	prompt := "1+2-3+4-5+6=? Just give me a number result"
	ai2.MessageStore.Add(ai2.RoleUser, prompt, nil)

	response := ai2.ToolChat(ai2.MessageStore.ToMessages(), toolList)
	toolCalls := response.ToolCalls

	for {
		if toolCalls == nil {
			fmt.Println("大模型的最终回复： ", response.Content)
			break
		}

		fmt.Println("大模型的回复是： ", response.Content)
		fmt.Println("大模型选择的工具是： ", toolCalls)

		selectToolCall := toolCalls[0]
		selectFunction := selectToolCall.Function
		functionName := selectFunction.Name

		var result int
		var args tools2.InputArgs
		err := json.Unmarshal([]byte(selectFunction.Arguments), &args)
		if err != nil {
			log.Fatalln("json unmarshal err: ", err.Error())
		}

		switch functionName {
		case tools2.AddToolDefine.Function.Name:
			result = tools2.AddTool(args.Numbers)
		case tools2.SubToolDefine.Function.Name:
			result = tools2.SubTool(args.Numbers)
		}

		fmt.Println("函数计算结果： ", result)
		ai2.MessageStore.Add(ai2.RoleAssistant, response.Content, toolCalls)
		ai2.MessageStore.AddForTool(strconv.Itoa(result), functionName, selectToolCall.ID)

		response = ai2.ToolChat(ai2.MessageStore.ToMessages(), toolList)
		toolCalls = response.ToolCalls
	}
}
