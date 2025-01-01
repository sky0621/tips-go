package domain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUserByID(t *testing.T) {
	tests := []struct {
		testName string
		inputID  int
		expected *User
		err      error
	}{
		{
			testName: "正常ケース",
			inputID:  5,
			expected: &User{ID: 5, Name: "ユーザー１"},
			err:      nil,
		},
		{
			testName: "準正常ケース：ID不正",
			inputID:  0,
			expected: nil,
			err: ErrInvalidInput(ErrorAttribute{
				Key:   "id",
				Value: 0,
			}),
		},
		{
			testName: "準正常ケース：認証失敗",
			inputID:  10,
			expected: nil,
			err: ErrUnauthenticated(ErrorAttribute{
				Key:   "id",
				Value: 10,
			}),
		},
		{
			testName: "準正常ケース：認可失敗",
			inputID:  20,
			expected: nil,
			err: ErrUnauthorized(ErrorAttribute{
				Key:   "id",
				Value: 20,
			}),
		},
		{
			testName: "準正常ケース：リソース無し",
			inputID:  30,
			expected: nil,
			err: ErrResourceNotFound(ErrorAttribute{
				Key:   "id",
				Value: 30,
			}),
		},
		{
			testName: "準正常ケース：予期せぬエラー",
			inputID:  40,
			expected: nil,
			err: ErrUnexpectedError(ErrorAttribute{
				Key:   "id",
				Value: 40,
			}),
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			sut := &userService{}
			actual, err := sut.GetUserByID(context.Background(), test.inputID)
			if test.err != nil {
				assert.ErrorIs(t, err, test.err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, actual)
		})
	}
}
