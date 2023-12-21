// baidu/models.go
package baidu

type Message struct {
    Role        string `json:"role"`
    Content     string `json:"content"`
    // Name        string `json:"name,omitempty"`
    // FunctionCall *FunctionCall `json:"function_call,omitempty"`
}

// type FunctionCall struct {
//     Name      string `json:"name"`
//     Arguments string `json:"arguments"`
//     Thoughts  string `json:"thoughts,omitempty"`
// }

type ChatRequest struct {
    Messages []Message `json:"messages"`
}

type ChatResponse struct {
    ID                 string `json:"id"`
    Object             string `json:"object"`
    Created            int    `json:"created"`
    Result             string `json:"result"`
    IsTruncated        bool   `json:"is_truncated"`
    NeedClearHistory   bool   `json:"need_clear_history"`
    Usage              Usage  `json:"usage"`
}

type Usage struct {
    PromptTokens      int `json:"prompt_tokens"`
    CompletionTokens  int `json:"completion_tokens"`
    TotalTokens       int `json:"total_tokens"`
}