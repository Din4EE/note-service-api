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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNoteGet(t *testing.T) {
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
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()
	)

	testCases := []struct {
		desc        string
		setupMock   func()
		input       *desc.GetRequest
		expectedRes *desc.GetResponse
		expectError bool
		errorText   string
	}{
		{
			desc: "Valid request",
			setupMock: func() {
				noteMock.EXPECT().Get(ctx, gomock.Any()).Return(&model.Note{
					ID: fakeId,
					NoteInfo: &model.NoteInfo{
						Title:  title,
						Text:   text,
						Author: author,
						Email:  email,
					},
					CreatedAt: createdAt,
					UpdatedAt: &updatedAt,
				}, nil)
			},
			input: &desc.GetRequest{
				Id: fakeId,
			},
			expectedRes: &desc.GetResponse{
				Note: &desc.Note{
					Id:        fakeId,
					Info:      &desc.NoteInfo{Title: title, Text: text, Author: author, Email: email},
					CreatedAt: timestamppb.New(createdAt),
					UpdatedAt: timestamppb.New(updatedAt),
				},
			},
		},
		{
			desc: "Error from repo",
			setupMock: func() {
				noteMock.EXPECT().Get(ctx, gomock.Any()).Return(nil, errors.New(errorText))
			},
			input: &desc.GetRequest{
				Id: fakeId,
			},
			expectError: true,
			errorText:   errorText,
		},
		{
			desc: "Id not found",
			setupMock: func() {
				noteMock.EXPECT().Get(ctx, fakeId).Return(nil, errors.New(errorText))
			},
			input: &desc.GetRequest{
				Id: fakeId,
			},
			expectError: true,
			errorText:   errorText,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.setupMock()

			res, err := api.Get(ctx, tc.input)

			if tc.expectError {
				require.Error(t, err)
				require.Equal(t, tc.errorText, err.Error())
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expectedRes, res)
		})
	}

}
