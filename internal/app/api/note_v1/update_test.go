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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestNoteUpdate(t *testing.T) {
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
		input       *desc.UpdateRequest
		expectedRes *emptypb.Empty
		expectError bool
		errorText   string
	}{
		{
			desc: "Valid request",
			setupMock: func() {
				noteMock.EXPECT().Update(ctx, gomock.Any(), &model.NoteInfoUpdate{
					Title:  &title,
					Text:   &text,
					Author: &author,
					Email:  &email,
				}).Return(nil)
			},
			input: &desc.UpdateRequest{
				Id: fakeId,
				Info: &desc.UpdateNoteInfo{
					Title:  wrapperspb.String(title),
					Text:   wrapperspb.String(text),
					Author: wrapperspb.String(author),
					Email:  wrapperspb.String(email),
				},
			},
			expectedRes: &emptypb.Empty{},
		},
		{
			desc: "Valid request with empty fields",
			setupMock: func() {
				noteMock.EXPECT().Update(ctx, gomock.Any(), &model.NoteInfoUpdate{
					Title:  new(string),
					Text:   &text,
					Author: &author,
					Email:  &email,
				}).Return(nil)
			},
			input: &desc.UpdateRequest{
				Id: fakeId,
				Info: &desc.UpdateNoteInfo{
					Title:  wrapperspb.String(""),
					Text:   wrapperspb.String(text),
					Author: wrapperspb.String(author),
					Email:  wrapperspb.String(email),
				},
			},
			expectedRes: &emptypb.Empty{},
		},
		{
			desc: "Repo error",
			setupMock: func() {
				noteMock.EXPECT().Update(ctx, gomock.Any(), &model.NoteInfoUpdate{
					Title: &title,
					Text:  &text,
				}).Return(errors.New(errorText))
			},
			input: &desc.UpdateRequest{
				Id: fakeId,
				Info: &desc.UpdateNoteInfo{
					Title: wrapperspb.String(title),
					Text:  wrapperspb.String(text),
				},
			},
			expectError: true,
			errorText:   errorText,
		},
		{
			desc: "Id not found",
			setupMock: func() {
				noteMock.EXPECT().Update(ctx, fakeId, &model.NoteInfoUpdate{
					Title: &title,
					Text:  &text,
				}).Return(errors.New(errorText))
			},
			input: &desc.UpdateRequest{
				Id: fakeId,
				Info: &desc.UpdateNoteInfo{
					Title: wrapperspb.String(title),
					Text:  wrapperspb.String(text),
				},
			},
			expectError: true,
			errorText:   errorText,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.setupMock()

			res, err := api.Update(ctx, tc.input)

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
