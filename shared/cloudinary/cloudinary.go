package cloudinary

import (
	"context"
	"errors"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func getCloudinaryClient() (*cloudinary.Cloudinary, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		return nil, errors.New("credenciales de cloudinary no configuradas")
	}

	return cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
}

// uploadFile es el manejador central para evitar código repetido
func uploadFile(fileHeader *multipart.FileHeader, resourceType string, folder string) (string, error) {
	cld, err := getCloudinaryClient()
	if err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	ctx := context.Background()
	// La subida determina el ResourceType (imagen o video-para-audio)
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:       folder,
		ResourceType: resourceType,
	})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}

// UploadImage maneja subida de imágenes
func UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	return uploadFile(fileHeader, "image", "biblioteca/imagenes")
}

// UploadAudio maneja la subida de audios (Cloudinary los trata como "video")
func UploadAudio(fileHeader *multipart.FileHeader) (string, error) {
	return uploadFile(fileHeader, "video", "biblioteca/audios")
}

// DeleteFile (Opcional) remueve el recurso de la nube
func DeleteFile(publicID string, resourceType string) error {
	cld, err := getCloudinaryClient()
	if err != nil {
		return err
	}
	ctx := context.Background()
	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     publicID,
		ResourceType: resourceType,
	})
	return err
}
