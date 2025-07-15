package dao

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBAdapter struct {
	client *dynamodb.Client
	table  string
}

func NewDynamoDBAdapter(tableName string) (*DynamoDBAdapter, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBAdapter{
		client: client,
		table:  tableName,
	}, nil
}

func (a *DynamoDBAdapter) CreateItem(ctx context.Context, item map[string]types.AttributeValue) error {
	_, err := a.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(a.table),
		Item:      item,
	})
	return err
}

func (a *DynamoDBAdapter) GetItem(ctx context.Context, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	result, err := a.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(a.table),
		Key:       key,
	})
	if err != nil {
		return nil, err
	}
	return result.Item, nil
}

func (a *DynamoDBAdapter) UpdateItem(ctx context.Context, key map[string]types.AttributeValue, updateExpression string, expressionValues map[string]types.AttributeValue) error {
	_, err := a.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(a.table),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionValues,
	})
	return err
}

func (a *DynamoDBAdapter) DeleteItem(ctx context.Context, key map[string]types.AttributeValue) error {
	_, err := a.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(a.table),
		Key:       key,
	})
	return err
}

func (a *DynamoDBAdapter) QueryItems(ctx context.Context, keyCondition string, expressionValues map[string]types.AttributeValue) ([]map[string]types.AttributeValue, error) {
	input := &dynamodb.QueryInput{
		TableName:                 aws.String(a.table),
		KeyConditionExpression:    aws.String(keyCondition),
		ExpressionAttributeValues: expressionValues,
	}

	result, err := a.client.Query(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}
