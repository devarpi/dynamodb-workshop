package model

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
