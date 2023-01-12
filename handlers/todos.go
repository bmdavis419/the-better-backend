package handlers

import (
	"github.com/bmdavis419/the-better-backend/database"
	"github.com/bmdavis419/the-better-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Summary Get all todos.
// @Description fetch every todo available.
// @Tags todos
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Todo
// @Router /todos [get]
func HandleAllTodos(c *fiber.Ctx) error {
	// fetch all todos
	coll := database.GetCollection("todos")

	// return all todos
	filter := bson.M{}
	opts := options.Find().SetSkip(0).SetLimit(100)

	// find all todos
	cursor, err := coll.Find(c.Context(), filter, opts)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// convert cursor to slice
	todos := make([]models.Todo, 0)
	if err = cursor.All(c.Context(), &todos); err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return todos
	return c.Status(200).JSON(todos)
}

type CreateTodoDTO struct {
	Title       string `json:"title" bson:"title"`
	Completed   bool   `json:"completed" bson:"completed"`
	Description string `json:"description" bson:"description"`
	Date        string `json:"date" bson:"date"`
}

type CreateTodoResDTO struct {
	InsertedId primitive.ObjectID `json:"inserted_id" bson:"_id"`
}

// @Summary Create a todo.
// @Description create a single todo.
// @Tags todos
// @Accept json
// @Param todo body CreateTodoDTO true "Todo to create"
// @Produce json
// @Success 200 {object} CreateTodoResDTO
// @Router /todos [post]
func HandleCreateTodo(c *fiber.Ctx) error {
	// get the todo from the request body
	nTodo := new(CreateTodoDTO)

	// validate the request body
	if err := c.BodyParser(nTodo); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	// insert the todo into the database
	coll := database.GetCollection("todos")
	res, err := coll.InsertOne(c.Context(), nTodo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return the inserted todo
	return c.Status(200).JSON(fiber.Map{"inserted_id": res.InsertedID})
}

type UpdateTodoDTO struct {
	Title       string `json:"title" bson:"title"`
	Completed   bool   `json:"completed" bson:"completed"`
	Description string `json:"description" bson:"description"`
	Date        string `json:"date" bson:"date"`
}

type UpdateTodoResDTO struct {
	UpdatedCount int64 `json:"updated_count" bson:"updated_count"`
}

// @Summary Update a todo.
// @Description update a single todo.
// @Tags todos
// @Accept json
// @Param todo body UpdateTodoDTO true "Todo update data"
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} UpdateTodoResDTO
// @Router /todos/:id [put]
func HandleUpdateTodo(c *fiber.Ctx) error {
	// get the id from the request params
	id := c.Params("id")
	dbId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}

	// get the todo from the request body
	uTodo := new(UpdateTodoDTO)

	// validate the request body
	if err := c.BodyParser(uTodo); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	// update the todo in the database
	coll := database.GetCollection("todos")
	filter := bson.M{"_id": dbId}
	update := bson.M{"$set": uTodo}
	res, err := coll.UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return the updated todo
	return c.Status(200).JSON(fiber.Map{"updated_count": res.ModifiedCount})
}

// @Summary Get a single todo.
// @Description fetch a single todo.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} models.Todo
// @Router /todos/:id [get]
func HandleGetOneTodo(c *fiber.Ctx) error {
	// get the id from the request params
	id := c.Params("id")
	dbId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}

	// fetch the todo from the database
	coll := database.GetCollection("todos")
	filter := bson.M{"_id": dbId}
	var todo models.Todo
	err = coll.FindOne(c.Context(), filter).Decode(&todo)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	// return the todo
	return c.Status(200).JSON(todo)
}

type DeleteTodoResDTO struct {
	DeletedCount int64 `json:"deleted_count" bson:"deleted_count"`
}

// @Summary Delete a single todo.
// @Description delete a single todo by id.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} DeleteTodoResDTO
// @Router /todos/:id [delete]
func HandleDeleteTodo(c *fiber.Ctx) error {
	// get the id from the request params
	id := c.Params("id")
	dbId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}

	// delete the todo from the database
	coll := database.GetCollection("todos")
	filter := bson.M{"_id": dbId}
	res, err := coll.DeleteOne(c.Context(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return the deleted todo
	return c.Status(200).JSON(fiber.Map{"deleted_count": res.DeletedCount})
}
