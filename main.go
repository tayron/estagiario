package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
)

func GetResponse(client gpt3.Client, ctx context.Context, quesiton string) {
	var responseText = ""
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			quesiton,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
		Stream:      false,
	}, func(resp *gpt3.CompletionResponse) {
		responseText += resp.Choices[0].Text
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(13)
	}

	// fmt.Printf(responseText)
	ExecuteShellScript(responseText)
	fmt.Printf("\n")
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func main() {
	log.SetOutput(new(NullWriter))
	apiKey := os.Getenv("API_KEY_CHAT_GPT")
	//apiKey := "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"
	if apiKey == "" {
		panic("Missing API KEY for CHAT GPT")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with ChatGPT in console.",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("\nEstagiário: ")

				if !scanner.Scan() {
					break
				}

				question := scanner.Text()
				questionParam := validateQuestion(question)
				switch questionParam {
				case "quit":
					quit = true
				case "":
					continue

				default:
					questionParam = "Crie um arquivo shell script que: " + questionParam
					GetResponse(client, ctx, questionParam)
				}
			}
		},
	}

	log.Fatal(rootCmd.Execute())
}

func validateQuestion(question string) string {
	quest := strings.Trim(question, " ")
	keywords := []string{"", "loop", "break", "continue", "cls", "exit", "block"}
	for _, x := range keywords {
		if quest == x {
			return ""
		}
	}
	return quest
}

func ExecuteShellScript(cmdCommand string) {
	cmd := exec.Command("/bin/sh", "-c", cmdCommand)
	out, err := cmd.Output()

	fmt.Println("\n-----------------------------------------------------------------------")

	if err != nil {
		fmt.Println("- Não foi possível executar a tarefa:")
		fmt.Println(err)
	} else {

		outString := string(out)

		if outString != "" {
			fmt.Printf("- %s \n", outString)
		} else {
			fmt.Println("- Tarefa executada com sucesso")
		}
	}

	fmt.Println("-----------------------------------------------------------------------")
}
