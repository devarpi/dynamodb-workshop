package model

import (
	"time"
)

// BaseItem represents the basic structure of items stored in DynamoDB
type BaseItem struct {
	ID        string `json:"id" dynamodbav:"id"`
	CreatedAt int64  `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt int64  `json:"updated_at" dynamodbav:"updated_at"`
}

// TableMetadata contains information about DynamoDB table
type TableMetadata struct {
	TableName string
	PKName    string
	SKName    string
}

// Tenant represents a tenant entity in the system
type Tenant struct {
	PK             string    `json:"pk" dynamodbav:"pk"`
	SK             string    `json:"sk" dynamodbav:"sk"`
	TenantID       string    `json:"tenantId" dynamodbav:"tenantId"`
	TenantName     string    `json:"tenantName" dynamodbav:"tenantName"`
	LineOfBusiness string    `json:"lineOfBusiness" dynamodbav:"lineOfBusiness"`
	CostGroupID    string    `json:"costGroupId" dynamodbav:"costGroupId"`
	CreatedAt      time.Time `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" dynamodbav:"updatedAt"`
}
