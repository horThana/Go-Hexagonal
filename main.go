package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/horThana/Backend/adapters/http"
	"github.com/horThana/Backend/adapters/repository"
	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    app := fiber.New()

    //Core middleware
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET, POST, PUT, DELETE",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    // Connect to product database
    productDB, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
    if err != nil {
        panic("ไม่สามารถเชื่อมต่อฐานข้อมูลผลิตภัณฑ์ได้")
    }

    // Connect to user database
    userDB, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
    if err != nil {
        panic("ไม่สามารถเชื่อมต่อฐานข้อมูลผู้ใช้ได้")
    }

    // Connect to payment database
    paymentDB, err := gorm.Open(sqlite.Open("payment.db"), &gorm.Config{})
    if err != nil {
        panic("ไม่สามารถเชื่อมต่อฐานข้อมูลการชำระเงินได้")
    }

    // Migrate the schema for both databases
    productDB.AutoMigrate(&domain.Product{})
    userDB.AutoMigrate(&domain.User{})
    paymentDB.AutoMigrate(&domain.PayMentSubscription{})

    // Set up repositories, services, and handlers for products
    productRepo := repository.NewGormProductRepository(productDB)
    productService := services.NewProductService(productRepo)
    productHandler := http.NewHttpProductAdapter(productService)

    // Set up repositories, services, and handlers for users
    userRepo := repository.NewGormUserRepository(userDB)
    userService := services.NewUserService(userRepo)
    userHandler := http.NewHttpUserAdapter(userService)

    // Set up repositories, services, and handlers for payments
    paymentRepo := repository.NewGormPaymentSubscriptionRepository(paymentDB)
    paymentService := services.NewPaymentService(paymentRepo)
    paymentHandler := http.NewHttpPaymentAdapter(paymentService)


    // Define routes for products
    app.Post("/product", productHandler.CreateProduct)
    app.Get("/product/:id", productHandler.FindProductByID)
    app.Get("/product", productHandler.FindAllProducts)
    app.Delete("/product/:id", productHandler.DeleteProduct)

    // Define routes for users
    app.Post("/user", userHandler.CreateUser)
    app.Get("/user/:id", userHandler.FindUserByID)
    app.Get("/user", userHandler.FindAllUsers)
    app.Delete("/user/:id", userHandler.DeleteUser)

    // Define routes for payments
    app.Post("/payment", paymentHandler.CreatePaymentSubscription)
    app.Get("/payment/:id", paymentHandler.FindPaymentSubscriptionByID)
    app.Get("/payment", paymentHandler.FindAllPaymentSubscriptions)
    app.Delete("/payment/:id", paymentHandler.DeletePaymentSubscription)

    // Start the server
    app.Listen(":8000")
}