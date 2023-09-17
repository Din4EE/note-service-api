package main

import (
	"context"
	"log"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const addr = "localhost:50051"

func main() {
	connect, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %s", err.Error())
	}
	defer connect.Close()
	client := desc.NewNoteServiceClient(connect)
	//res, err := client.CreateNote(
	//	context.Background(), &desc.CreateNoteRequest{
	//		Title:  "lol",
	//		Text:   "kek",
	//		Author: "superpuper",
	//	})
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(res.GetId())
	//}
	//res, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
	//	Id: "50298619-bc7a-4434-8f70-e0cb72102bd8",
	//})
	//resGetList, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{
	//	Limit:       100,
	//	Offset:      0,
	//	SearchQuery: "super",
	//})
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(resGetList.GetIds())
	//}
	resUpdate, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Id:     "d5753100-0e21-48de-873f-dd4aefa3fb55",
		Title:  "lol",
		Text:   "kek",
		Author: "cheburek",
	})
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(resUpdate.GetStatus())
	}
}
