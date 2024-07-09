package routes

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sikarwar2010/doc-fiber-app/database"
	"github.com/sikarwar2010/doc-fiber-app/models"
)

type Order struct {
	Id       uint      `json:"id"`
	User     User      `json:"user"`
	Product  Product   `json:"product"`
	CreateAt time.Time `json:"order_date"`
}

func CreateOrderResponce(order models.Order, user User, product Product) Order {
	return Order{
		Id:       order.ID,
		User:     user,
		Product:  product,
		CreateAt: order.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)
	responceUser := CreateResponceUser(user)
	reponceProduct := CreateResponceProduct(product)

	return c.Status(200).JSON(CreateOrderResponce(order, responceUser, reponceProduct))

}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}
	for _, order := range orders {
		var user models.User
		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		var product models.Product
		database.Database.Db.Find(&product, "id = ?", order.ProductRefer)
		responseOrders = append(responseOrders, CreateOrderResponce(order, CreateResponceUser(user), CreateResponceProduct(product)))
	}

	return c.Status(200).JSON(responseOrders)

}

func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetOrderById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	database.Database.Db.First(&user, order.UserRefer)

	var product models.Product
	database.Database.Db.First(&product, order.ProductRefer)

	responseOrder := CreateOrderResponce(order, CreateResponceUser(user), CreateResponceProduct(product))

	return c.Status(200).JSON(responseOrder)

}
