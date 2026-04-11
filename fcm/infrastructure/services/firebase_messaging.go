package services

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type FirebaseMessagingService interface {
	SendMessage(token string, data map[string]string) error
}

type firebaseMessagingImpl struct {
	client *messaging.Client
}

func NewFirebaseMessagingService() FirebaseMessagingService {
	projectID := os.Getenv("FIREBASE_PROJECT_ID")
	privateKey := os.Getenv("FIREBASE_PRIVATE_KEY")
	clientEmail := os.Getenv("FIREBASE_CLIENT_EMAIL")

	if projectID == "" || privateKey == "" || clientEmail == "" {
		log.Println("Advertencia: Credenciales de Firebase no configuradas. Las notificaciones FCM estarán desactivadas.")
		return &firebaseMessagingImpl{}
	}

	privateKey = strings.ReplaceAll(privateKey, "\\n", "\n")

	creds := map[string]string{
		"type":         "service_account",
		"project_id":   projectID,
		"private_key":  privateKey,
		"client_email": clientEmail,
	}
	credsJSON, _ := json.Marshal(creds)

	opt := option.WithCredentialsJSON(credsJSON)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("Error inicializando Firebase App: %v", err)
		return &firebaseMessagingImpl{}
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Printf("Error inicializando Firebase Messaging: %v", err)
		return &firebaseMessagingImpl{}
	}

	log.Println("Firebase Messaging inicializado correctamente")
	return &firebaseMessagingImpl{client: client}
}

func (s *firebaseMessagingImpl) SendMessage(token string, data map[string]string) error {
	if s.client == nil {
		log.Println("Advertencia: Firebase no inicializado, omitiendo notificación a token:", token)
		return nil
	}

	msg := &messaging.Message{
		Token: token,
		Data:  data,
	}

	_, err := s.client.Send(context.Background(), msg)
	if err != nil {
		log.Printf("Error enviando mensaje FCM a token %s: %v", token, err)
		return err
	}

	log.Printf("Mensaje FCM enviado exitosamente al token: %s", token)
	return nil
}
