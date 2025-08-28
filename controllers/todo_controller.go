package controllers

import (
    "net/http"
    "strconv"
    "todo-api/config"
    "todo-api/models"

    "github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
    rows, err := config.DB.Query("SELECT id, title, completed FROM todos")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var todos []models.Todo
    for rows.Next() {
        var t models.Todo
        if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        todos = append(todos, t)
    }

    c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    res, err := config.DB.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", todo.Title, todo.Completed)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    id, _ := res.LastInsertId()
    todo.ID = int(id)
    c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := config.DB.Exec("UPDATE todos SET title=?, completed=? WHERE id=?", todo.Title, todo.Completed, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    todo.ID, _ = strconv.Atoi(id)
    c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    _, err := config.DB.Exec("DELETE FROM todos WHERE id=?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}