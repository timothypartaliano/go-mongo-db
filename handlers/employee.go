package handlers

import (
	"context"
	"net/http"
	"strconv"
	"ugc-2/config"
	"ugc-2/model"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateEmployee(c echo.Context) error {
    employee := new(model.Employee)
    if err := c.Bind(employee); err != nil {
        return err
    }

    result, err := config.Collection.InsertOne(context.Background(), employee)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusCreated, result.InsertedID)
}

func GetEmployee(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid employee ID")
    }

    filter := bson.M{"_id": id}
    var employee model.Employee
    err = config.Collection.FindOne(context.Background(), filter).Decode(&employee)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.JSON(http.StatusNotFound, "Employee not found")
        }
        return err
    }

    return c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid employee ID")
    }

    employee := new(model.Employee)
    if err := c.Bind(employee); err != nil {
        return err
    }

    filter := bson.M{"_id": id}
    update := bson.M{
        "$set": bson.M{
            "name":     employee.Name,
            "position": employee.Position,
            "salary":   employee.Salary,
        },
    }

    _, err = config.Collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, "Employee updated")
}

func DeleteEmployee(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid employee ID")
    }

    filter := bson.M{"_id": id}
    _, err = config.Collection.DeleteOne(context.Background(), filter)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, "Employee deleted")
}