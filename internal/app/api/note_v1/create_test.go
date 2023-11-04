package note_v1

import (
	"context"
	"errors"
	"testing"

	noteMocks "github.com/Din4EE/note-service-api/internal/repo/mocks"
	"github.com/Din4EE/note-service-api/internal/service/model"
	"github.com/Din4EE/note-service-api/internal/service/note"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestNoteCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	api := NewNote(note.NewService(noteMock))
	ctx := context.Background()

	var (
		fakeId    = gofakeit.Uint64()
		errorText = gofakeit.Sentence(1)
		title     = gofakeit.BookTitle()
		text      = gofakeit.Phrase()
		author    = gofakeit.Name()
		email     = gofakeit.Email()
	)

	testCases := []struct {
		desc        string
		setupMock   func()
		input       *desc.CreateRequest
		expectedRes *desc.CreateResponse
		expectError error
	}{
		{
			desc: "Valid request",
			setupMock: func() {
				noteMock.EXPECT().Create(ctx, &model.Note{
					NoteInfo: &model.NoteInfo{
						Title:  title,
						Text:   text,
						Author: author,
						Email:  email,
					},
				}).Return(fakeId, nil)
			},
			input: &desc.CreateRequest{
				Info: &desc.NoteInfo{
					Title:  title,
					Text:   text,
					Author: author,
					Email:  email,
				},
			},
			expectedRes: &desc.CreateResponse{Id: fakeId},
		},
		{
			desc: "Error from repo",
			setupMock: func() {
				noteMock.EXPECT().Create(ctx, &model.Note{
					NoteInfo: &model.NoteInfo{
						Title:  title,
						Text:   text,
						Author: author,
						Email:  email,
					},
				}).Return(uint64(0), errors.New(errorText))
			},
			input: &desc.CreateRequest{
				Info: &desc.NoteInfo{
					Title:  title,
					Text:   text,
					Author: author,
					Email:  email,
				},
			},
			expectError: errors.New(errorText),
		},
		{
			desc: "Empty request",
			setupMock: func() {
				noteMock.EXPECT().Create(ctx, &model.Note{
					NoteInfo: &model.NoteInfo{},
				}).Return(fakeId, nil)
			},
			input: &desc.CreateRequest{
				Info: &desc.NoteInfo{},
			},
			expectedRes: &desc.CreateResponse{Id: fakeId},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.setupMock()

			res, err := api.Create(ctx, tc.input)

			require.Equalf(t, tc.expectError, err, "want: %v, got: %v", tc.expectError, err)
			require.Equalf(t, tc.expectedRes, res, "want: %v, got: %v", tc.expectedRes, res)
		})
	}
}
