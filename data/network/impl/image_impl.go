package impl

import (
	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"context"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
	"log"
	"os"
	"vita-message-service/data/entity"
	"vita-message-service/data/entity/image"
	"vita-message-service/data/network"
)

type imageService struct{}

func NewImageService() network.ImageService {
	return &imageService{}
}

func (m *imageService) Scan(message entity.Message) []image.Possibility {
	ctx := context.Background()
	client, err := network.GetGoogleVision()
	if err != nil {
		log.Fatalf("error when init google vision : %v", err)
		return nil
	}
	defer client.Close()

	// Sets the name of the image file to annotate.
	filename := "upload/images/" + message.Message

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	localImage, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	res, err := client.AnnotateImage(ctx, &pb.AnnotateImageRequest{
		Image: localImage,
		Features: []*visionpb.Feature{
			{Type: visionpb.Feature_LANDMARK_DETECTION, MaxResults: 25},
			{Type: visionpb.Feature_DOCUMENT_TEXT_DETECTION, MaxResults: 25},
			{Type: visionpb.Feature_TEXT_DETECTION, MaxResults: 25},
			{Type: visionpb.Feature_LOGO_DETECTION, MaxResults: 25},
			{Type: visionpb.Feature_OBJECT_LOCALIZATION, MaxResults: 25},
		},
	})

	if err != nil {
		log.Fatalf("failed to detect object : %v", err)
	}

	var possibilities []image.Possibility

	if len(res.LandmarkAnnotations) > 0 {
		possibilities = append(possibilities, image.Possibility{
			Type:        visionpb.Feature_Type_name[int32(visionpb.Feature_LANDMARK_DETECTION)],
			Description: res.LandmarkAnnotations[0].Description,
		})
	}

	if fullTextNotation := res.FullTextAnnotation; fullTextNotation != nil {
		possibilities = append(possibilities, image.Possibility{
			Type:        visionpb.Feature_Type_name[int32(visionpb.Feature_TEXT_DETECTION)],
			Description: fullTextNotation.Text,
		})
	}

	if len(res.LogoAnnotations) > 0 {
		possibilities = append(possibilities, image.Possibility{
			Type:        visionpb.Feature_Type_name[int32(visionpb.Feature_LOGO_DETECTION)],
			Description: res.LogoAnnotations[0].Description,
		})
	}

	if len(possibilities) == 0 && len(res.LocalizedObjectAnnotations) > 0 {
		for _, object := range res.LocalizedObjectAnnotations {
			possibilities = append(possibilities, image.Possibility{
				Type:        visionpb.Feature_Type_name[int32(visionpb.Feature_OBJECT_LOCALIZATION)],
				Description: object.Name,
			})
		}
	}

	return possibilities
}
