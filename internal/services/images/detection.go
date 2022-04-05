package images

import (
	"context"
	"crypto/x509"
	"fmt"
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
	x := 1
	imageDescriptions := map[string][]string{}
	// Batch requests by 15 because google doesn't let your request more than 17 images at once and 17 is a dumb number
	// so I went with 15
	// EDIT: There's no way to tie a result back to the original image, which is also dumb, so I have to do the batches as
	// one at a time
	for x*1 < len(urls){
		request := &vision.BatchAnnotateImagesRequest{}
		var urlToUse string
		for _, url := range urls[(x-1)*1:x*1] {
			urlToUse = url
			if url == "" {
				continue
			}
			urlToUse = url
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

		resp, err := i.client.BatchAnnotateImages(ctx, request)
		if err != nil {
			return nil, err
		}

		for _, detectedContents := range resp.GetResponses() {

			detectedObjects := []string{}
			for _, contents := range detectedContents.GetWebDetection().GetWebEntities() {
				if strings.ToLower(contents.GetDescription()) == "Dungeons & Dragons"{
					continue
				}
				detectedObjects = append(detectedObjects,  strings.Split(strings.ToLower(contents.GetDescription()), " ")...)
			}
			for _, contents := range detectedContents.GetLabelAnnotations() {
				if strings.ToLower(contents.GetDescription()) == "Dungeons & Dragons"{
					continue
				}
				detectedObjects = append(detectedObjects,  strings.Split(strings.ToLower(contents.GetDescription()), " ")...)
			}
			detectedObjects = append(detectedObjects, strings.Split(detectedContents.GetFullTextAnnotation().GetText(), " ")...)
			imageDescriptions[urlToUse] = detectedObjects
		}
		fmt.Println(imageDescriptions)
		x++
	}


	// Swap the [url][]descriptions to be [description]url
	descriptionsToImages := map[string][]string{}
	for url, descriptions := range imageDescriptions{
		for _, description := range descriptions {
			if descriptionsToImages[description] == nil {
				descriptionsToImages[description] = []string{}
			}
			descriptionsToImages[description] = append(descriptionsToImages[description], url)
		}
	}
	return descriptionsToImages, nil
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
