package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/compat_oai/openai"
	"github.com/openai/openai-go/option"
)

func main() {
	
	startTime := time.Now()
	ctx := context.Background()
	g := genkit.Init(ctx, genkit.WithPlugins(&openai.OpenAI{
		APIKey: "tada",
		Opts: []option.RequestOption{
			option.WithBaseURL(os.Getenv("MODEL_RUNNER_BASE_URL")),
		},
	}))

	modelId := os.Getenv("MODEL_ID")
	quantizationFormat := os.Getenv("QUANTIZATION_FORMAT")
	//_ = quantizationFormat

	temperature, err := strconv.ParseFloat(os.Getenv("MODEL_TEMPERATURE"), 32)
	if err != nil {
		temperature = 0.7 // default value
	}
	topP, err := strconv.ParseFloat(os.Getenv("MODEL_TOP_P"), 32)
	if err != nil {
		topP = 0.8 // default value
	}
	score := 0
	fmt.Println("🌺🍕 Hawaiian Pizza Test for", modelId)
	// Loop from 1 to 5 to run multiple prompts
	for i := 1; i <= 5; i++ {
		promptEnvVar := fmt.Sprintf("USER_MESSAGE_%d", i)
		wordsToSearchFor := fmt.Sprintf("WORDS_TO_SEARCH_FOR_%d", i)

		fmt.Println("\n----- Running prompt:", os.Getenv(promptEnvVar), "-----")
		resp, err := genkit.Generate(ctx, g,
			ai.WithModelName("openai/"+modelId),
			ai.WithSystem(os.Getenv("SYSTEM_INSTRUCTION")),
			ai.WithPrompt(os.Getenv(promptEnvVar)),

			ai.WithConfig(map[string]any{
				"temperature": float32(temperature),
				"top_p":       float32(topP),
			}),
			// Print the final response to stdout
			ai.WithStreaming(func(ctx context.Context, chunk *ai.ModelResponseChunk) error {
				// Do something with the chunk...
				fmt.Print(chunk.Text())
				return nil
			}),
		)
		if err != nil {
			log.Fatal(err)
		}
		// search for all words in wordsToSearchFor
		fmt.Println("\n----- Cheking answer:")
		wordsToSearchForList := strings.Split(os.Getenv(wordsToSearchFor), ",")
		fmt.Println("📝 Words to search for:", wordsToSearchForList)
		numberOfWordsToFind := len(wordsToSearchForList)
		numberOfFoundWords := 0
		for _, word := range wordsToSearchForList {
			if !strings.Contains(strings.ToLower(resp.Text()), strings.ToLower(word)) {
				fmt.Printf("🔴 Error: word '%s' not found in response\n", word)
			} else {
				fmt.Printf("🟢 Success: word '%s' found in response\n", word)
				numberOfFoundWords++
			}
		}
		// calculate score
		if numberOfFoundWords == numberOfWordsToFind {
			score += 1
		}
		fmt.Println("----- End of prompt -----")
	}
	fmt.Printf("\n🌺🍕 Final Score: %d/5\n", score)

	// Calculate and display total processing time
	totalTime := time.Since(startTime)
	fmt.Printf("\n⏱️  Total processing time: %v\n", totalTime)

	// create a json file with the score
	jsonRecord := struct {
		Test string `json:"test"`
		Score int `json:"score"`
		ModelId string `json:"model_id"`
		QuantizationFormat string `json:"quantization_format"`
		Temperature float64 `json:"temperature"`
		TopP float64 `json:"top_p"`
		TotalTime string `json:"total_time"`
		Architecture string `json:"architecture"`
		OS string `json:"os"`
		NumCPU int `json:"num_cpu"`
	}{
		Test: "Hawaiian Pizza Test",
		Score: score,
		ModelId: modelId,
		QuantizationFormat: quantizationFormat,
		Temperature: temperature,
		TopP: topP,
		TotalTime: totalTime.String(),
		Architecture: runtime.GOARCH,
		OS: runtime.GOOS,
		NumCPU: runtime.NumCPU(),
	}
	

	// save jsonRecord to a json file
	jsonData, err := json.MarshalIndent(jsonRecord, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(os.Getenv("SCORE_FILE"), jsonData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("💾 Score saved to ", os.Getenv("SCORE_FILE"))

}
