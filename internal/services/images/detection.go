package images

import (
	"context"
	"crypto/x509"
	"google.golang.org/genproto/googleapis/cloud/vision/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"strings"
)

type ImageDescriber interface {
	GetDescriptionWords(ctx context.Context, urls []string) (map[string][]string, error)
}

type imageDescriber struct {
	baseURL          string
	authFileLocation string
	client           vision.ImageAnnotatorClient
}

func (i imageDescriber) GetDescriptionWords(ctx context.Context, urls []string) (map[string][]string, error) {
	request := &vision.BatchAnnotateImagesRequest{}
	for _, url := range urls {
		request.Requests = append(request.Requests, &vision.AnnotateImageRequest{
			Image: &vision.Image{
				Source: &vision.ImageSource{
					ImageUri: url,
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
			},
			ImageContext: nil,
		})
	}
	resp, err := i.client.BatchAnnotateImages(ctx, request
	if err != nil {
		return nil, err
	}
	imageDescriptions := map[string][]string{}
	for _, detectedContents := range resp.GetResponses() {

		detectedObjects := []string{}
		for _, contents := range detectedContents.GetWebDetection().GetWebEntities() {
			detectedObjects = append(detectedObjects,  strings.Split(strings.ToLower(contents.GetDescription()), " ")...)
		}
		for _, contents := range detectedContents.GetLabelAnnotations() {
			detectedObjects = append(detectedObjects,  strings.Split(strings.ToLower(contents.GetDescription()), " ")...)
		}
		detectedObjects = append(detectedObjects, strings.Split(detectedContents.GetFullTextAnnotation().GetText(), " ")...)
		for _, contents := range detectedContents.GetLabelAnnotations() {
			detectedObjects = append(detectedObjects, strings.Split(strings.ToLower(contents.GetDescription()), " ")...)
		}
		imageDescriptions[detectedContents.GetContext().GetUri()] = detectedObjects
	}
	return imageDescriptions, nil
}

func NewImageDescriber(url, authFile string) ImageDescriber {
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
	return imageDescriber{
		baseURL:          url,
		authFileLocation: authFile,
		client:           client,
	}
}
