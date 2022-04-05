package main

import (
	"context"
	"crypto/x509"
	"fmt"
	"google.golang.org/genproto/googleapis/cloud/vision/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func main(){
	pool, _ := x509.SystemCertPool()
	// error handling omitted
	creds := credentials.NewClientTLSFromCert(pool, "")
	perRPC, _ := oauth.NewServiceAccountFromFile("dndai-346223-ef10bf69a9d1.json", "https://www.googleapis.com/auth/cloud-platform")
	conn, _ := grpc.Dial(
		"vision.googleapis.com:443",
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(perRPC),
	)
	client := vision.NewImageAnnotatorClient(conn)

	resp, err := client.BatchAnnotateImages(context.Background(), &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{
			{
				Image: &vision.Image{
					Source: &vision.ImageSource{
						ImageUri:    "https://www.dndbeyond.com/attachments/thumbnails/6/305/850/546/ud5xx-00-01.png",
					},

				},
				Features: []*vision.Feature{
					{
						Type:       vision.Feature_LABEL_DETECTION,
						MaxResults: 10,
						Model:      "builtin/latest",
					},
					{
						Type:       vision.Feature_WEB_DETECTION,
						MaxResults: 10,
						Model:      "builtin/latest",
					},
					{
						Type:       vision.Feature_IMAGE_PROPERTIES,
						MaxResults: 10,
						Model:      "builtin/latest",
					},
				},
			},
		},
	})
	fmt.Println(resp)
	if err != nil{
		panic(err)
	}

	for _, detectedContents := range resp.GetResponses() {
		fmt.Println(detectedContents)
		fmt.Println(detectedContents.GetContext())
		detectedObjects := []string{}
		for _, contents := range detectedContents.GetWebDetection().GetWebEntities(){
			detectedObjects = append(detectedObjects, contents.GetDescription())
		}
		for _, contents := range detectedContents.GetLabelAnnotations(){
			detectedObjects = append(detectedObjects, contents.GetDescription())
		}
	}

}