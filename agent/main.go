package main

import (
	"Kit-Hung/cloud-native-AI/agent/pkg/ai"
	prompt_tpl "Kit-Hung/cloud-native-AI/agent/pkg/prompt-tpl"
	"Kit-Hung/cloud-native-AI/agent/pkg/tools"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	query := "1+2-3+4-5+6=? Just give me a number result"

	toolFormat := "%s:%s\nparam\n%s"
	addTool := fmt.Sprintf(toolFormat, tools.AddToolName, tools.AddToolDescription, tools.AddToolParam)
	subTool := fmt.Sprintf(toolFormat, tools.SubToolName, tools.SubToolDescription, tools.SubToolParam)

	toolList := make([]string, 2)
	toolList = append(toolList, addTool, subTool)

	toolNames := make([]string, 2)
	toolNames = append(toolNames, tools.AddToolName, tools.SubToolName)

	prompt := fmt.Sprintf(prompt_tpl.Template, toolList, toolNames, query)
	fmt.Println("prompt: ", prompt)

	// 注入用户 prompt
	ai.MessageStore.AddForUser(prompt)
	i := 1

	for {
		response := ai.NormalChat(ai.MessageStore.ToMessage())
		fmt.Printf("============= 第 %d 轮回答 =============\n", i)
		fmt.Println(response)

		regexPattern := regexp.MustCompile(`Final Answer:\s*(.*)`)
		finalAnswer := regexPattern.FindStringSubmatch(response.Content)
		if len(finalAnswer) > 1 {
			fmt.Println("============= 最终回复 =============")
			fmt.Println(response.Content)
			break
		}

		ai.MessageStore.AddForAssistant(response)

		regexAction := regexp.MustCompile(`Action:\s*(.*?)[.\n]`)
		regexActionInput := regexp.MustCompile(`Action Input:\s*(.*?)[.\n]`)
		action := regexAction.FindStringSubmatch(response.Content)
		actionInput := regexActionInput.FindStringSubmatch(response.Content)

		if len(action) > 1 && len(actionInput) > 1 {
			i++
			result := 0

			toolName := action[1]
			fmt.Printf("calling %s\n", toolName)

			switch toolName {
			case tools.AddToolName:
				result = tools.AddTool(actionInput[1])
			case tools.SubToolName:
				result = tools.SubTool(actionInput[1])
			}
			fmt.Println("============= 函数返回结果 =============")
			fmt.Println("result: ", result)

			observation := "Observation: " + strconv.Itoa(result)
			prompt = response.Content + observation
			fmt.Printf("============= 第 %d 轮的 prompt =============", i)
			fmt.Println(prompt)
			ai.MessageStore.AddForUser(prompt)
		}
	}
}
