package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents the API client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	AuthToken  string
}

// NewClient creates a new API client
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SetAuthToken sets the authentication token
func (c *Client) SetAuthToken(token string) {
	c.AuthToken = token
}

// RegisterRequest represents the registration request
type RegisterRequest struct {
	GPUType           string `json:"gpuType"`
	GPUVram           int    `json:"gpuVram"`
	GPUCount          int    `json:"gpuCount"`
	RosWalletAddress  string `json:"rosWalletAddress"`
	ContributorName   string `json:"contributorName,omitempty"`
}

// RegisterResponse represents the registration response
type RegisterResponse struct {
	Success        bool   `json:"success"`
	NodeID         int    `json:"node_id"`
	NodeAuthToken  string `json:"node_auth_token"`
	Message        string `json:"message"`
}

// Register registers a new worker node
func (c *Client) Register(req *RegisterRequest) (*RegisterResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.HTTPClient.Post(
		c.BaseURL+"/api/worker/register",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("registration failed (status %d): %s", resp.StatusCode, string(body))
	}

	var result RegisterResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// HeartbeatRequest represents the heartbeat request
type HeartbeatRequest struct {
	Status string `json:"status"`
}

// HeartbeatResponse represents the heartbeat response
type HeartbeatResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Heartbeat sends a heartbeat to the server
func (c *Client) Heartbeat(status string) error {
	req := &HeartbeatRequest{Status: status}
	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest(
		"POST",
		c.BaseURL+"/api/worker/heartbeat",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.AuthToken)

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("heartbeat failed (status %d): %s", resp.StatusCode, string(body))
	}

	return nil
}

// JobPayload represents the job payload
type JobPayload struct {
	DockerImage   string                 `json:"docker_image"`
	InputS3URL    string                 `json:"input_s3_url"`
	OutputS3Path  string                 `json:"output_s3_path"`
	Prompt        string                 `json:"prompt,omitempty"`
	InitVideoURL  string                 `json:"init_video_url,omitempty"`
	WorkflowJSON  interface{}            `json:"workflow_json,omitempty"`
	Extra         map[string]interface{} `json:"-"`
}

// Job represents a job
type Job struct {
	JobID    string      `json:"job_id"`
	TaskType string      `json:"task_type"`
	Payload  *JobPayload `json:"payload"`
}

// GetJobResponse represents the get job response
type GetJobResponse struct {
	Success bool   `json:"success"`
	Job     *Job   `json:"job"`
	Message string `json:"message"`
}

// GetJob gets a new job from the server
func (c *Client) GetJob() (*Job, error) {
	httpReq, err := http.NewRequest(
		"GET",
		c.BaseURL+"/api/worker/get-job",
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.AuthToken)

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get job failed (status %d): %s", resp.StatusCode, string(body))
	}

	var result GetJobResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result.Job, nil
}

// SubmitResultRequest represents the submit result request
type SubmitResultRequest struct {
	JobID        string `json:"job_id"`
	Status       string `json:"status"`
	OutputS3URL  string `json:"output_s3_url,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// SubmitResultResponse represents the submit result response
type SubmitResultResponse struct {
	Success    bool    `json:"success"`
	Message    string  `json:"message"`
	RewardPaid float64 `json:"reward_paid"`
}

// SubmitResult submits the result of a job
func (c *Client) SubmitResult(req *SubmitResultRequest) (*SubmitResultResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest(
		"POST",
		c.BaseURL+"/api/worker/submit-result",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.AuthToken)

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("submit result failed (status %d): %s", resp.StatusCode, string(body))
	}

	var result SubmitResultResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

