package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"dynamodb-301-workshop/src/model"
)

type TenantDAO struct {
	db *DynamoDBAdapter
}

func NewTenantDAO(tableName string) (*TenantDAO, error) {
	adapter, err := NewDynamoDBAdapter(tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to create tenant DAO: %w", err)
	}
	return &TenantDAO{db: adapter}, nil
}

func (dao *TenantDAO) CreateTenant(ctx context.Context, tenant *model.Tenant) error {
	if tenant.TenantID == "" {
		return ErrInvalidInput
	}

	tenant.PK = fmt.Sprintf("TENANT#%s", tenant.TenantID)
	tenant.SK = "METADATA#TENANT"
	tenant.CreatedAt = time.Now()
	tenant.UpdatedAt = time.Now()

	av, err := attributevalue.MarshalMap(tenant)
	if err != nil {
		return fmt.Errorf("failed to marshal tenant: %w", err)
	}

	return dao.db.CreateItem(ctx, av)
}

func (dao *TenantDAO) GetTenant(ctx context.Context, tenantID string) (*model.Tenant, error) {
	key := map[string]types.AttributeValue{
		"pk": &types.AttributeValueMemberS{Value: fmt.Sprintf("TENANT#%s", tenantID)},
		"sk": &types.AttributeValueMemberS{Value: "METADATA#TENANT"},
	}

	item, err := dao.db.GetItem(ctx, key)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, ErrItemNotFound
	}

	tenant := &model.Tenant{}
	err = attributevalue.UnmarshalMap(item, tenant)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal tenant: %w", err)
	}

	return tenant, nil
}

func (dao *TenantDAO) UpdateTenant(ctx context.Context, tenant *model.Tenant) error {
	if tenant.TenantID == "" {
		return ErrInvalidInput
	}

	key := map[string]types.AttributeValue{
		"pk": &types.AttributeValueMemberS{Value: fmt.Sprintf("TENANT#%s", tenant.TenantID)},
		"sk": &types.AttributeValueMemberS{Value: "METADATA#TENANT"},
	}

	updateExpr := "set tenantName = :name, lineOfBusiness = :lob, costGroupId = :cgid, updatedAt = :updated"
	exprValues := map[string]types.AttributeValue{
		":name":    &types.AttributeValueMemberS{Value: tenant.TenantName},
		":lob":     &types.AttributeValueMemberS{Value: tenant.LineOfBusiness},
		":cgid":    &types.AttributeValueMemberS{Value: tenant.CostGroupID},
		":updated": &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
	}

	return dao.db.UpdateItem(ctx, key, updateExpr, exprValues)
}

func (dao *TenantDAO) QueryTenantsByLOB(ctx context.Context, lineOfBusiness string) ([]*model.Tenant, error) {
	keyCondition := "begins_with(pk, :prefix) AND lineOfBusiness = :lob"
	expressionValues := map[string]types.AttributeValue{
		":prefix": &types.AttributeValueMemberS{Value: "TENANT#"},
		":lob":    &types.AttributeValueMemberS{Value: lineOfBusiness},
	}

	items, err := dao.db.QueryItems(ctx, keyCondition, expressionValues)
	if err != nil {
		return nil, err
	}

	tenants := make([]*model.Tenant, 0, len(items))
	for _, item := range items {
		tenant := &model.Tenant{}
		err = attributevalue.UnmarshalMap(item, tenant)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal tenant: %w", err)
		}
		tenants = append(tenants, tenant)
	}

	return tenants, nil
}
