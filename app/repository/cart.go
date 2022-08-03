package repository

import (
	"context"
	"errors"

	"github.com/furkanarsl/pf-final/database"
	"github.com/furkanarsl/pf-final/pkg/queries"
)

type cartRepo struct {
	database.DbQueries
}

type CartRepo interface {
	GetCartForUser(userID int64) (queries.Cart, error)
	GetCartItems(cartID int64) ([]queries.ListCartItemsRow, error)
	AddToCart(cartID int64, productID int64, quantity int32) (queries.CartProduct, error)
	RemoveFromCart(cartID, itemID int64) error
}

func NewCartRepo(queries database.DbQueries) *cartRepo {
	return &cartRepo{queries}
}

func (r *cartRepo) GetCartForUser(userID int64) (queries.Cart, error) {
	cart, err := r.Queries.GetCartForUser(context.Background(), userID)

	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *cartRepo) GetCartItems(cartID int64) ([]queries.ListCartItemsRow, error) {
	cartItems, err := r.ListCartItems(context.Background(), cartID)

	if err != nil {
		return cartItems, err
	}

	return cartItems, nil
}

func (r *cartRepo) AddToCart(cartID int64, productID int64, quantity int32) (queries.CartProduct, error) {
	args := queries.AddToCartParams{ProductID: productID, CartID: cartID, Quantity: quantity}
	result, err := r.Queries.AddToCart(context.Background(), args)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *cartRepo) RemoveFromCart(cartID, itemID int64) error {
	_, err := r.GetCartItem(context.Background(), itemID)
	if err != nil {
		return err
	}
	args := queries.DeleteCartItemParams{CartID: cartID, ID: itemID}
	err = r.DeleteCartItem(context.Background(), args)
	if err != nil {
		return errors.New("failed to delete product from cart")
	}
	return nil
}
