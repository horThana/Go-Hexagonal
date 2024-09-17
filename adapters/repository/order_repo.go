package repository

// //ใช้เพื่อส่งข้อมูลของ Order ไปยัง Database = เก็บไว้ที่ gorm_adapter.go
// import (
// 	"fmt"

// 	core "github.com/horThana/Backend/core/ports"
// 	"gorm.io/gorm"
// )

// type GormOrderRepository struct {
// 	db *gorm.DB
// }

// func NewGormOrderRepository(db *gorm.DB) core.Order_Repository {
// 	return &GormOrderRepository{db: db}

// }

// func (r *GormOrderRepository) Save(order core.Order_Repository) error {
// 	if result := r.db.Create(&order); result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func(r *GormOrderRepository) FindByID(id string) (core.Order_Repository, error) {
// 	var order core.Order
// 	result := r.db.First(&order, "id = ?", id)
// 		fmt.Println(result)
// 	if result := r.db.First(&order, id); result.Error != nil {
// 		return core.Order{}, result.Error
// 	}
// 	return order, nil
// }

// func(r *GormOrderRepository) FindAll()([]core.Order, error){
// 	var orders []core.Order
// 	result := r.db.Find(&orders)
// 	   if result.Error != nil {
// 		return []core.Order{}, result.Error
// 		 }
// 	return orders, nil
// }