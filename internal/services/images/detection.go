package images

import (
	"context"
	"crypto/x509"
	"google.golang.org/genproto/googleapis/cloud/vision/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"strings"
	"sync"
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

	imageDescriptions := map[string][]string{}
	wg := sync.WaitGroup{}
	// Batch requests by 15 because google doesn't let your request more than 17 images at once and 17 is a dumb number
	// so I went with 15
	// EDIT: There's no way to tie a result back to the original image, which is also dumb, so I have to do the batches as
	// one at a time
	for _, url := range urls {
		if url == "" {
			continue
		}
		wg.Add(1)
		go i.getImageDetails(ctx, url, imageDescriptions, &wg)
	}
	wg.Wait()
	// Swap the [url][]descriptions to be [description]url
	descriptionsToImages := map[string][]string{}
	for url, descriptions := range imageDescriptions {
		for _, description := range descriptions {
			if descriptionsToImages[description] == nil {
				descriptionsToImages[description] = []string{}
			}
			descriptionsToImages[description] = append(descriptionsToImages[description], url)
		}
	}
	return descriptionsToImages, nil
}

func (i imageDescriber) getImageDetails(ctx context.Context, url string, imageDescriptions map[string][]string, wg *sync.WaitGroup) {
	request := &vision.BatchAnnotateImagesRequest{}
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
	resp, _ := i.client.BatchAnnotateImages(ctx, request)

	for _, detectedContents := range resp.GetResponses() {

		detectedObjects := []string{}
		for _, contents := range detectedContents.GetWebDetection().GetWebEntities() {
			if strings.ToLower(contents.GetDescription()) == "Dungeons & Dragons" {
				continue
			}
			detectedObjects = append(detectedObjects, strings.Split(strings.ToLower(contents.GetDescription()), " ")...)
		}
		for _, contents := range detectedContents.GetLabelAnnotations() {
			if strings.ToLower(contents.GetDescription()) == "Dungeons & Dragons" {
				continue
			}
			detectedObjects = append(detectedObjects, strings.Split(strings.ToLower(contents.GetDescription()), " ")...)
		}
		detectedObjects = append(detectedObjects, strings.Split(detectedContents.GetFullTextAnnotation().GetText(), " ")...)

		imageDescriptions[url] = detectedObjects

	}

	wg.Done()
}

func NewImageDescriber(url, authFile string) ImageDescriber {
	pool, _ := x509.SystemCertPool()
	// error handling omitted
	creds := credentials.NewClientTLSFromCert(pool, "")
	perRPC, _ := oauth.NewServiceAccountFromFile(authFile, "https://www.googleapis.com/auth/cloud-platform")
	conn, _ := grpc.Dial(
		url,
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
