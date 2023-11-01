package note_v1

import (
	"context"
	"errors"
	"testing"

	noteMocks "github.com/Din4EE/note-service-api/internal/repo/mocks"
	"github.com/Din4EE/note-service-api/internal/service/note"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestNoteDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	api := NewNote(note.NewService(noteMock))
	ctx := context.Background()

	var (
		fakeId    = gofakeit.Uint64()
		errorText = gofakeit.Sentence(1)
	)

	testCases := []struct {
		desc        string
		setupMock   func()
		input       *desc.DeleteRequest
		expectedRes *emptypb.Empty
		expectError bool
		errorText   string
	}{
		{
			desc: "Valid request",
			setupMock: func() {
				noteMock.EXPECT().Delete(ctx, gomock.Any()).Return(nil)
			},
			input: &desc.DeleteRequest{
				Id: 1,
			},
			expectedRes: &emptypb.Empty{},
		},
		{
			desc: "Error from repo",
			setupMock: func() {
				noteMock.EXPECT().Delete(ctx, gomock.Any()).Return(errors.New(errorText))
			},
			input: &desc.DeleteRequest{
				Id: 1,
			},
			expectError: true,
			errorText:   errorText,
		},
		{
			desc: "Id not found",
			setupMock: func() {
				noteMock.EXPECT().Delete(ctx, fakeId).Return(errors.New(errorText))
			},
			input: &desc.DeleteRequest{
				Id: fakeId,
			},
			expectError: true,
			errorText:   errorText,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.setupMock()

			res, err := api.Delete(ctx, tc.input)

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
