package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	configure "ordering_service/internal/config"
	"time"
)

type DynamoDB struct {
	DB                 *dynamodb.Client
	OrderTableName     string
	OrderItemTableName string
}

func InitDynamoDB() (*DynamoDB, error) {
	cfg, err := configure.LoadConfig()

	// Create custom resolver for local DynamoDB
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           cfg.DynamoEndpoint,
			SigningRegion: cfg.DynamoRegion,
		}, nil
	})

	// Load AWS configuration with local settings
	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.DynamoRegion),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.DynamoAccessKey,
			cfg.DynamoSecretKey,
			"local-session",
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load local DynamoDB config: %w", err)
	}
	// Create DynamoDB client
	newDbInstance := &DynamoDB{
		DB:                 dynamodb.NewFromConfig(awsConfig),
		OrderTableName:     "orders",
		OrderItemTableName: "menu_items",
	}
	if err := newDbInstance.createOrderTable(); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	if err := newDbInstance.CreateMenuItemTable(); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}
	return newDbInstance, nil
}

func (d *DynamoDB) createOrderTable() error {
	// Check if the table already exists
	exists, err := d.doesTableExists(d.OrderTableName)
	if err != nil {
		return fmt.Errorf("failed to check if order table exists: %w", err)
	}

	if exists {
		// Table already exists, no need to create
		return nil
	}

	// Define the table schema
	input := &dynamodb.CreateTableInput{
		TableName: &d.OrderTableName,
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	// Create the table
	_, err = d.DB.CreateTable(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	// Wait for the table to become active
	waiter := dynamodb.NewTableExistsWaiter(d.DB)
	err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(d.OrderTableName)}, 5*time.Minute)
	if err != nil {
		return fmt.Errorf("failed to wait for table to become active: %w", err)
	}

	return nil
}

func (d *DynamoDB) CreateMenuItemTable() error {
	// Check if the table already exists
	exists, err := d.doesTableExists(d.OrderItemTableName)
	if err != nil {
		return fmt.Errorf("failed to check if menu item table exists: %w", err)
	}

	if exists {
		// Table already exists, no need to create
		return nil
	}

	// Define the table schema
	input := &dynamodb.CreateTableInput{
		TableName: &d.OrderItemTableName,
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	// Create the table
	_, err = d.DB.CreateTable(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	// Wait for the table to become active
	waiter := dynamodb.NewTableExistsWaiter(d.DB)
	err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(d.OrderItemTableName)}, 5*time.Minute)
	if err != nil {
		return fmt.Errorf("failed to wait for table to become active: %w", err)
	}

	return nil
}

func (d *DynamoDB) doesTableExists(tableName string) (bool, error) {
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}

	_, err := d.DB.DescribeTable(context.TODO(), input)
	if err != nil {
		var notFoundErr *types.ResourceNotFoundException
		if errors.As(err, &notFoundErr) {
			return false, nil // Table does not exist
		}
		return false, fmt.Errorf("failed to describe table: %w", err)
	}

	return true, nil // Table exists
}
