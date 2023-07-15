package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"route256/notifications/internal/models"
	"testing"
	"time"
)

func TestListUserHistoryDay(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockCache := NewMockCache(ctrl)
	service := NewService(mockRepo, mockCache)

	testUserID := int64(1)
	ctx := context.Background()
	testTime := time.Now()

	tests := []struct {
		name           string
		cacheReturn    []models.OrderMessage
		cacheError     error
		repoReturn     []models.OrderMessage
		repoError      error
		expectedResult []models.OrderMessage
		expectedError  string
	}{
		{
			name:           "Empty cache and repository",
			cacheReturn:    nil,
			cacheError:     nil,
			repoReturn:     nil,
			repoError:      nil,
			expectedResult: []models.OrderMessage{},
			expectedError:  "",
		},
		{
			name:           "Cache error",
			cacheReturn:    nil,
			cacheError:     errors.New("cache error"),
			repoReturn:     nil,
			repoError:      nil,
			expectedResult: nil,
			expectedError:  "cache error",
		},
		{
			name:           "Repository error",
			cacheReturn:    nil,
			cacheError:     nil,
			repoReturn:     nil,
			repoError:      errors.New("repo error"),
			expectedResult: nil,
			expectedError:  "repo error",
		},
		{
			name: "Successful retrieval from cache and repository",
			cacheReturn: []models.OrderMessage{
				{
					UserID:    1,
					OrderID:   1,
					Status:    "new",
					CreatedAt: testTime,
				},
			},
			cacheError: nil,
			repoReturn: []models.OrderMessage{
				{
					UserID:    1,
					OrderID:   2,
					Status:    "new",
					CreatedAt: testTime,
				},
			},
			repoError: nil,
			expectedResult: []models.OrderMessage{
				{
					UserID:    1,
					OrderID:   1,
					Status:    "new",
					CreatedAt: testTime,
				},
				{
					UserID:    1,
					OrderID:   2,
					Status:    "new",
					CreatedAt: testTime,
				},
			},
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCache.EXPECT().GetUserHistoryDay(ctx, testUserID).Return(tt.cacheReturn, tt.cacheError)

			if tt.cacheError == nil && len(tt.cacheReturn) > 0 {
				mockCache.EXPECT().GetLatestMessageTime(ctx, testUserID).Return(testTime, nil)
				mockRepo.EXPECT().ListUserHistoryDay(ctx, testUserID, gomock.Any()).Return(tt.repoReturn, tt.repoError)

				for _, msg := range tt.repoReturn {
					mockCache.EXPECT().AddToUserHistoryDay(ctx, msg).Return(nil)
				}
			} else if tt.cacheError == nil && len(tt.cacheReturn) == 0 {
				mockRepo.EXPECT().ListUserHistoryDay(ctx, testUserID, nil).Return(tt.repoReturn, tt.repoError)

				for _, msg := range tt.repoReturn {
					mockCache.EXPECT().AddToUserHistoryDay(ctx, msg).Return(nil)
				}
			}

			result, err := service.ListUserHistoryDay(ctx, testUserID)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.expectedResult), len(result))
			}
		})
	}
}
