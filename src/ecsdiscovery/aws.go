package main

import (
    "os"
    "log"
    "reflect"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ecs"
    "github.com/aws/aws-sdk-go/aws/awserr"
)

func main () {
    log.Println("test")

    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
    log.Println(sess)
    svc := ecs.New(sess)
    input := &ecs.ListClustersInput{
            MaxResults: aws.Int64(10),
            NextToken: nil,
    }
    clusters, err := svc.ListClusters(input)
    if err != nil {
        if awsErr, ok := err.(awserr.Error); ok {
            log.Panicln(awsErr)
        }
        os.Exit(1)
    }
    log.Println(clusters)
    log.Println(reflect.TypeOf(clusters))
}

/*
if err != nil {
        if awsErr, ok := err.(awserr.Error); ok {
                    // process SDK error
                        }
}
*/
