package repository

import "go-project/internal/domain/entity"

// User はテスト
type User interface {
	Find() *entity.User
}
