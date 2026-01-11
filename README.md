# cagent RAG Test

Sample project to test RAG configurations for the blog post https://www.ajeetraina.com/how-to-rag-with-docker-cagent/

## Setup

```bash
# Set your OpenAI API key
export OPENAI_API_KEY=your-key-here

# Run cagent
cd cagent-rag-test
cagent run cagent-config.yaml
```



## Test Queries

Try these queries to verify RAG is working:

| Query | Should Find | File |
|-------|-------------|------|
| "how does authentication work?" | `TokenValidator`, `CheckCredentials` | src/auth.go |
| "retry logic with backoff" | `Client.Do()` with exponential backoff | pkg/httpclient.go |
| "HandleRequest" | `func HandleRequest(...)` | src/handlers.go |
| "validate user token" | `validateToken()` | src/auth.go |


<img width="615" height="479" alt="image" src="https://github.com/user-attachments/assets/26408a87-3189-4499-a93e-dc424c1ea1cf" />

<img width="795" height="624" alt="image" src="https://github.com/user-attachments/assets/dbdedf30-7e59-4dfa-977e-12da4bec58ea" />

## ReRanking

```
cagent run cagent-reranking-config.yaml
```

## Chunking

```
cagent run cagent-chunking-config.yaml
```



## What to Check

1. **Semantic search works**: "authentication" finds code using "token", "credentials"
2. **Keyword search works**: "HandleRequest" finds exact function name
3. **Hybrid fusion works**: Results include both exact matches and related code

## Files

```
cagent-rag-demo/
├── cagent-config.yaml   # RAG config (hybrid: embeddings + bm25)
├── src/
│   ├── auth.go          # Token validation, credentials check
│   └── handlers.go      # HTTP handlers
└── pkg/
    └── httpclient.go    # HTTP client with retry/backoff
```
