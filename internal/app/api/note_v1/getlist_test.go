package note_v1

import (
	"context"
	"errors"
	"reflect"
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

func TestNoteGetList(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	api := NewNote(note.NewService(noteMock))
	ctx := context.Background()

	var (
		fakeId      = gofakeit.Uint64()
		errorText   = gofakeit.Sentence(1)
		searchQuery = gofakeit.Sentence(1)
		limit       = gofakeit.Uint64()
		offset      = gofakeit.Uint64()
		title       = gofakeit.BookTitle()
		text        = gofakeit.Phrase()
		author      = gofakeit.Name()
		email       = gofakeit.Email()
		createdAt   = gofakeit.Date()
		updatedAt   = gofakeit.Date()
	)

	testCases := []struct {
		desc        string
		setupMock   func()
		input       *desc.GetListRequest
		expectedRes *desc.GetListResponse
		expectError bool
		errorText   string
	}{
		{
			desc: "Valid request",
			setupMock: func() {
				noteMock.EXPECT().GetList(ctx, searchQuery, limit, offset).Return([]*model.Note{
					{
						ID: fakeId,
						NoteInfo: &model.NoteInfo{
							Title:  title,
							Text:   text,
							Author: author,
							Email:  email,
						},
						CreatedAt: createdAt,
						UpdatedAt: &updatedAt,
					},
					{
						ID: fakeId + 1,
						NoteInfo: &model.NoteInfo{
							Title:  title + "1",
							Text:   text + "1",
							Author: author + "1",
							Email:  email + "1",
						},
						CreatedAt: createdAt,
						UpdatedAt: &updatedAt,
					},
				}, nil)
			},
			input: &desc.GetListRequest{
				SearchQuery: searchQuery,
				Limit:       limit,
				Offset:      offset,
			},
			expectedRes: &desc.GetListResponse{
				Notes: []*desc.Note{
					{
						Id: fakeId,
						Info: &desc.NoteInfo{
							Title:  title,
							Text:   text,
							Author: author,
							Email:  email,
						},
						CreatedAt: timestamppb.New(createdAt),
						UpdatedAt: timestamppb.New(updatedAt),
					},
					{
						Id: fakeId + 1,
						Info: &desc.NoteInfo{
							Title:  title + "1",
							Text:   text + "1",
							Author: author + "1",
							Email:  email + "1",
						},
						CreatedAt: timestamppb.New(createdAt),
						UpdatedAt: timestamppb.New(updatedAt),
					},
				},
			},
		},
		{
			desc: "Error from repo",
			setupMock: func() {
				noteMock.EXPECT().GetList(ctx, searchQuery, limit, offset).Return(nil, errors.New(errorText))
			},
			input: &desc.GetListRequest{
				SearchQuery: searchQuery,
				Limit:       limit,
				Offset:      offset,
			},
			expectError: true,
			errorText:   errorText,
		},
		{
			desc: "Empty result",
			setupMock: func() {
				noteMock.EXPECT().GetList(ctx, searchQuery, limit, offset).Return([]*model.Note{}, nil)
			},
			input: &desc.GetListRequest{
				SearchQuery: searchQuery,
				Limit:       limit,
				Offset:      offset,
			},
			expectedRes: &desc.GetListResponse{
				Notes: []*desc.Note{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.setupMock()

			res, err := api.GetList(ctx, tc.input)

			if tc.expectError {
				require.Error(t, err)
				require.Equal(t, tc.errorText, err.Error())
				return
			}

			require.NoError(t, err)
			require.True(t, reflect.DeepEqual(tc.expectedRes, res))
		})
	}
}
