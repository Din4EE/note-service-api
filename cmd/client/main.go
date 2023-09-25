package main

import (
	"context"
	"log"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const addr = "localhost:50051"

func main() {
	connect, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %s", err.Error())
	}
	defer connect.Close()
	client := desc.NewNoteServiceClient(connect)
	res, err := client.CreateNote(
		context.Background(), &desc.CreateNoteRequest{
			Title:  "lol223",
			Text:   "kek444",
			Author: "superduper",
		})
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(res.GetId())
	}
	//res, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
	//	Id: "16a8bf86-9196-4be1-8b4f-af25d589cf91",
	//})
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Println(res)
	//}
	//resGetList, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{
	//	Limit:       100,
	//	Offset:      0,
	//	SearchQuery: "super",
	//})
	//if err != nil {
	//	log.Println(err.Error())
	//} else {
	//	log.Printf("Notes: %v \nTotal Count: %d", resGetList.GetNotes(), resGetList.GetTotalCount())
	//}
	resUpdate, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Id:    res.GetId(),
		Title: wrapperspb.String("lol1"),
	})
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(resUpdate)
	}
}
