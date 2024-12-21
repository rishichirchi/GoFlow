## GoFlow: AI-Powered GitHub Workflow Automation and Chrome Extension

## Overview
GoFlow is a project that leverages the power of generative AI to streamline and enhance tasks such as code review generation, social media post creation, email drafting, and pull request analysis. The project integrates the Gemini generative AI model along with GitHub and Twitter APIs to automate workflows effectively.

---

## Folder Structure

```
Gotta-Go/
├── backend/            
│   ├── handlers/       # Contains the main handler functions for various features
│   │   ├── code-reviewer.go 
│   │   ├── email-outreach.go 
│   │   ├── gemini-chatbot.go 
│   │   ├── generate-posts.go
│   │   └── make-post.go     
│   ├── instructions/   # Markdown files for prompts and templates
│   │   ├── code-reviewer.md    
│   │   ├── gofr-introduction.md 
│   │   └── sm-post.md          
│   ├── models/         # Contains data models used across the application
│   │   ├── chatbot-response.go 
│   │   ├── code-review.go      
│   │   └── github-pr.go       
│   ├── .gitignore      
│   ├── go.mod          
│   ├── go.sum          
│   └── main.go         
├── frontend/           # Frontend for the application
│   ├── public/         
│   └── src/            # Source code for the frontend
│       ├── .gitignore  
│       ├── README.md   
│       ├── background.js      
│       ├── manifest.json      
│       ├── package.json       
│       ├── package-lock.json  
│       ├── tsconfig.json     
│       └── webpack.config.js  
├── middleware/         # Middleware layer
│   ├── .gitignore      
│   ├── index.js        # Middleware entry point
│   ├── package.json
│   └── package-lock.json 
├── README.md           # Project documentation
```

---

## Features

### 1. Code Review Generation
- **Handler:** `GenerateCodeReview`
- **Description:** Reads a file and combines its content with a predefined set of instructions to generate a detailed code review using the Gemini generative AI model.
- **Input:** Filepath of the code to be reviewed.
- **Output:** JSON response containing the review.

### 2. Cold Email Generation
- **Handler:** `GenerateColdEmailForGolang`
- **Description:** Creates tailored cold outreach emails targeted at Golang developers.
- **Input:** Purpose of the email.
- **Output:** Email subject and body.

### 3. Social Media Post Creation
- **Handler:** `GenerateSocialMediaPost`
- **Description:** Combines GitHub pull request data with predefined templates to create engaging social media posts.
- **Input:** Pull request data.
- **Output:** A complete social media post.

### 4. Chatbot Integration
- **Handler:** `GeminiChatBot`
- **Description:** Provides an interactive chatbot interface for answering questions or generating content based on predefined instructions.
- **Input:** User prompt.
- **Output:** Generated content.

### 5. GitHub Pull Request Fetching
- **Handler:** `FetchPullRequests`
- **Description:** Fetches pull requests from a specified GitHub repository using the GitHub API.
- **Input:** Repository and owner details.
- **Output:** Pull request data.

### 6. Twitter Posting
- **Handler:** `MakeTweet`
- **Description:** Posts tweets to Twitter using a backend API.
- **Input:** Tweet content.
- **Output:** Response from the Twitter API.

---

## Environment Variables
The following environment variables are required:

| Variable           | Description                                  |
|--------------------|----------------------------------------------|
| `GEMINI_API_KEY`   | API key for Gemini generative AI.            |
| `GITHUB_TOKEN`     | Personal access token for GitHub API.        |

---

## Dependencies
- **Go:** Ensure Go is installed and properly configured.
- **Modules:** Run `go mod tidy` to install dependencies.

---

## How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/rishichirchi/gotta-go.git
   cd gotta-go
   ```

2. Set up the `.env` file with your API keys.

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   # Development environment
   go run main.go

   # Production environment
   go build -o gotta-go main.go
   ./gotta-go
   ```

---

## API Endpoints

### 1. Generate Code Review
- **Endpoint:** `/generate-code-review`
- **Method:** `POST`
- **Input:** `{ "name": "filepath" }`
- **Output:** Code review content.

### 2. Generate Cold Email
- **Endpoint:** `/generate-cold-email`
- **Method:** `GET`
- **Query Param:** `purpose=<email purpose>`
- **Output:** Email subject and body.

### 3. Generate Social Media Post

- **Handler:** `GenerateSocialMediaPost`
- **Description:** Combines GitHub pull request data with predefined templates to create engaging social media posts.
- **Input:** Pull request data.
- **Output:** A complete social media post.

#### Sample Pull Request Data Structure
```json
{
  "title": "Fix memory leak in data parser",
  "author": "johndoe",
  "description": "This pull request addresses the memory leak issue occurring in the data parser module. Unit tests have been added to cover edge cases.",
  "repository": "example-repo",
  "url": "https://github.com/example-org/example-repo/pull/123",
  "labels": ["bugfix", "performance"]
}
```

This data is processed to generate a concise and engaging post for social media platforms.
- **Endpoint:** `/generate-sm-post`
- **Method:** `GET`
- **Output:** Social media post content.

### 4. Chatbot Interaction
- **Endpoint:** `/chatbot`
- **Method:** `GET`
- **Query Param:** `prompt=<user prompt>`
- **Output:** Generated chatbot response.

### 5. Fetch Pull Requests
- **Endpoint:** `/fetch-pull-requests`
- **Method:** `GET`
- **Output:** Pull request data.

### 6. Make Tweet
- **Endpoint:** `/make-tweet`
- **Method:** `POST`
- **Input:** `{ "content": "tweet content" }`
- **Output:** Twitter API response.

---

## Contributing

1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Create a pull request.

---

## License
This project is licensed under the MIT License, a permissive free software license. This means contributors retain copyright over their work, but grant permission to use, modify, and distribute the project, provided proper attribution is given. See the LICENSE file for more details.

