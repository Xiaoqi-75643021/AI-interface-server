package openai

// Request sended to OpenAI
type ChatRequest struct {
	Model		string		`json:"model"`		// The language model to use(e.g., GPT-4, GPT-3.5...)
	Messages	[]Message	`json:"messages"`	// An array of Message structs(Representing the conversation history)
}

// Response from OpenAI
type ChatResponse struct {
    ID                string   `json:"id"`					// A unique identifier for the chat session.
    Object            string   `json:"object"`				
    Created           int64    `json:"created"`				// A timestamp indicating when the response was created.
    Model             string   `json:"model"`				// The language model that was used for the response.
    SystemFingerprint string   `json:"system_fingerprint"`	// An identifier for the system state after the completion.
    Choices           []Choice `json:"choices"`				// An array of Choice structs representing possible response
    Usage             Usage    `json:"usage"`	
}

// A single message in a conversation(User or AI)
type Message struct {
	Role	string `json:"role"`	// Indicating whether the message is from the 'user' or 'assistant'
	Content string `json:"content"`	// The actual text of the message
}

// Choices returned by OpenAI
type Choice struct {
	Index			int		`json:"index"`			// The position of the choice in the array of responses
	Message			Message	`json:"message"`
	FinishReason	string	`json:"finish_reason"`	// The reason why the model stopped generating text(e.g., 'length' or 'stop)
}

// Token usage of the completion(Providing metrics on how many tokens were used.)
type Usage struct {
	PromptTokens		int `json:"prompt_tokens"`		// The count of tokens used in the prompt
	CompletionTokens	int `json:"completion_tokens"`	// The count of tokens generated as the completion
	TotalTokens			int `json:"total_tokens"`		// The total count of tokens used(Including prompts and completions)
}