package rooms

import "groupproject3-airbnb-api/features/user"

type RoomEntity struct {
	Id          uint
	UserId      uint
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
	GetByUserId(userId, userIdLogin uint) ([]RoomEntity, error)
	Create(roomEntity RoomEntity, userId uint) (RoomEntity, error)
	Update(roomEntity RoomEntity, id, userId uint) (RoomEntity, error)
	Delete(id, userId uint) error
}

type RoomDataInterface interface {
	SelectAll() ([]RoomEntity, error)
	SelectById(id uint) (RoomEntity, error)
	SelectByUserId(user_id uint) ([]RoomEntity, error)
	Store(roomEntity RoomEntity, userId uint) (uint, error)
	Edit(roomEntity RoomEntity, id uint) error
	Destroy(id uint) error
}
