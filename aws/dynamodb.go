package dynamodb

import (
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type ExpenseDBItem {
	usr_id int
	exp_date int
	exp_name string
	exp_total float32
}

func AddExpenseItem(svc dynamodbiface.DynamoDBAPI, usrId int, date int, name string, total float32) (error) {
	item := ExpenseDBItem{
		usr_id:    usrId,
		exp_date:    date,
		exp_name: name,
		exp_total:  total,
	}

	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: &tablename,
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
