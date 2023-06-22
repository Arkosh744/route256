package cart

import (
	"context"
	"math/rand"

	"route256/checkout/internal/models"
)

func (s *Suite) TestAddToCartAndGetCount() {
	ctx := context.Background()

	// Given
	user := int64(rand.Int31())
	sku := uint32(rand.Int31())
	item := &models.ItemData{
		SKU:   sku,
		Count: 5,
	}
	s.Require().NoError(s.repo.AddToCart(ctx, user, item))

	// When
	count, err := s.repo.GetCount(ctx, user, sku)

	// Then
	s.Require().NoError(err)
	s.Require().Equal(item.Count, count)
}

func (s *Suite) TestDeleteFromCartAndGetCount() {
	ctx := context.Background()

	// Given
	user := int64(rand.Int31())
	sku := uint32(rand.Int31())
	item := &models.ItemData{
		SKU:   sku,
		Count: 5,
	}
	s.Require().NoError(s.repo.AddToCart(ctx, user, item))

	// When
	err := s.repo.DeleteFromCart(ctx, user, item)
	s.Require().NoError(err)

	// Then
	count, err := s.repo.GetCount(ctx, user, sku)
	s.Require().NoError(err)
	s.Require().Equal(uint16(0), count)
}

func (s *Suite) TestGetUserCartForNonExistingUser() {
	ctx := context.Background()

	// Given
	user := int64(rand.Int31()) // this user does not exist and does not have a cart

	// When
	items, err := s.repo.GetUserCart(ctx, user)

	// Then
	s.Require().NoError(err)
	s.Require().Empty(items)
}

func (s *Suite) TestDeleteUserCartAndGetCount() {
	ctx := context.Background()

	// Given
	user := int64(rand.Int31())
	sku := uint32(rand.Int31())
	item := &models.ItemData{
		SKU:   sku,
		Count: 5,
	}
	s.Require().NoError(s.repo.AddToCart(ctx, user, item))

	// When
	err := s.repo.DeleteUserCart(ctx, user)

	// Then
	s.Require().NoError(err)
	count, err := s.repo.GetCount(ctx, user, sku)
	s.Require().NoError(err)
	s.Require().Equal(uint16(0), count)
}

func (s *Suite) TestRemoveItemsFromCart() {
	ctx := context.Background()

	// Given
	user := int64(rand.Int31())
	sku := uint32(rand.Int31())
	initialCount := uint16(10)
	removeCount := uint16(5)
	expectedFinalCount := initialCount - removeCount

	item := &models.ItemData{
		SKU:   sku,
		Count: initialCount,
	}
	s.Require().NoError(s.repo.AddToCart(ctx, user, item))

	// When
	itemToRemove := &models.ItemData{
		SKU:   sku,
		Count: removeCount,
	}
	err := s.repo.DeleteFromCart(ctx, user, itemToRemove)

	// Then
	s.Require().NoError(err)
	finalCount, err := s.repo.GetCount(ctx, user, sku)
	s.Require().NoError(err)
	s.Require().Equal(expectedFinalCount, finalCount)
}

func (s *Suite) TestRemoveMoreItemsThanExist() {
	ctx := context.Background()

	// Given
	user := int64(rand.Int31())
	sku := uint32(rand.Int31())
	initialCount := uint16(5)
	removeCount := uint16(10) // greater than the initialCount

	item := &models.ItemData{
		SKU:   sku,
		Count: initialCount,
	}
	s.Require().NoError(s.repo.AddToCart(ctx, user, item))

	// When
	itemToRemove := &models.ItemData{
		SKU:   sku,
		Count: removeCount,
	}
	err := s.repo.DeleteFromCart(ctx, user, itemToRemove)
	// Then
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "stock insufficient")
}
