package controller

import (
	"docto/constants"
	"docto/models"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func uploadFile(ctx *fiber.Ctx, s3Client *s3.Client, wg *sync.WaitGroup, mutex *sync.Mutex, fileReader *multipart.FileHeader, results *[]models.File) {
	defer wg.Done()
	key := uuid.New().String()
	contentType := fileReader.Header.Get("Content-Type")

	file, _ := fileReader.Open()
	defer file.Close()

	_, err := s3Client.PutObject(ctx.Context(), &s3.PutObjectInput{
		Bucket:      aws.String(constants.S3_BUCKET),
		Key:         aws.String(key),
		Body:        io.Reader(file),
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	*results = append(*results, models.File{
		FileName: fileReader.Filename,
		Key:      key,
		Url:      fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", constants.S3_BUCKET, constants.S3_REGION, key),
	})
}

func UploadFiles(ctx *fiber.Ctx, s3 *s3.Client, fileReaders []*multipart.FileHeader) (*[]models.File, error) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	results := make([]models.File, 0)

	for _, fileReader := range fileReaders {
		wg.Add(1)

		go uploadFile(ctx, s3, &wg, &mutex, fileReader, &results)
	}

	wg.Wait()

	if len(results) != len(fileReaders) {
		return nil, errors.New("files not uploaded successfully")
	}

	return &results, nil

}
