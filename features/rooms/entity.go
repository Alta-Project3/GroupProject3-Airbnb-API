package rooms

import "groupproject3-airbnb-api/features/user"

type RoomEntity struct {
	Id          uint
	UserId      uint `validate:"required"`
	User        user.Core
	RoomName    string  `validate:"required"`
	Price       int     `validate:"required"`
	Description string  `validate:"required"`
	Latitude    float64 `validate:"required"`
	Longitude   float64 `validate:"required"`
}

type RoomServiceInterface interface {
	GetAll() ([]RoomEntity, error)
	GetById(id uint) (RoomEntity, error)
	GetByUserId(id uint) ([]RoomEntity, error)
	Create(roomEntity RoomEntity) (RoomEntity, error)
	Update(roomEntity RoomEntity, id uint) (RoomEntity, error)
	Delete(id uint) error
}

type RoomDataInterface interface {
	SelectAll() ([]RoomEntity, error)
	SelectById(id uint) (RoomEntity, error)
	SelectByUserId(user_id uint) ([]RoomEntity, error)
	Store(roomEntity RoomEntity) (uint, error)
	Edit(roomEntity RoomEntity, id uint) error
	Destroy(id uint) error
}