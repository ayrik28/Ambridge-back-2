package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// responseBodyWriter is a custom response writer that captures the response body
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write captures the response and writes it to the original writer
func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// LoggerMiddleware logs all requests with their paths, methods, request bodies, status codes, and response bodies
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Read the request body
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// Restore the request body for further processing
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// Create a custom response writer to capture the response
		responseBody := &bytes.Buffer{}
		writer := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           responseBody,
		}
		c.Writer = writer

		// Process request
		c.Next()

		// End timer
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Format request body for logging (if JSON)
		var requestJSON interface{}
		if len(requestBody) > 0 {
			if err := json.Unmarshal(requestBody, &requestJSON); err != nil {
				// If not valid JSON, use raw string
				requestJSON = string(requestBody)
			}
		} else {
			requestJSON = "empty body"
		}

		// Format response body for logging (if JSON)
		var responseJSON interface{}
		if responseBody.Len() > 0 {
			if err := json.Unmarshal(responseBody.Bytes(), &responseJSON); err != nil {
				// If not valid JSON, use raw string
				responseJSON = responseBody.String()
			}
		} else {
			responseJSON = "empty body"
		}

		// Log the request details
		log.Printf("[REQUEST] %s %s | Status: %d | Latency: %v\n", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), latency)
		log.Printf("[REQUEST BODY] %s\n", formatJSON(requestJSON))

		// Log the response details
		if c.Writer.Status() >= 400 {
			log.Printf("[ERROR RESPONSE] %s\n", formatJSON(responseJSON))
		} else {
			log.Printf("[RESPONSE] %s\n", formatJSON(responseJSON))
		}
	}
}

// formatJSON formats the JSON for pretty printing
func formatJSON(v interface{}) string {
	if v == nil {
		return "null"
	}

	switch value := v.(type) {
	case string:
		return value
	default:
		jsonBytes, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return fmt.Sprintf("%v", v)
		}
		return string(jsonBytes)
	}
}
