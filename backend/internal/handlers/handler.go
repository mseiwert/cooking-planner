package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Name string `json:"name"`
}

type RecipeHandler struct {
	h *RecipeHandler
}

func NewRecipeHandler() *RecipeHandler {
	return &RecipeHandler{}
}

func (h *RecipeHandler) RequestNewRecipes(g *gin.Context) {
	g.JSON(http.StatusOK, Recipe{Name: "Pasta"})
}
