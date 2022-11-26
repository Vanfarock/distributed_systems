package main

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ParametersDFS struct {
	Cities     []City
	StartIndex int
	EndIndex   int
}

type City struct {
	Id        int
	Neighbors []City
}

func dfs(parameters ParametersDFS) ([]City, int) {
	startCity := parameters.Cities[parameters.StartIndex]
	endCity := parameters.Cities[parameters.EndIndex]
	path := make([]City, 0)
	path = append(path, startCity)

	for _, city := range parameters.Cities {
		for i, neighbor := range city.Neighbors {
			city.Neighbors[i] = parameters.Cities[neighbor.Id-1]
		}
	}

	return dfsInternal(startCity, endCity, path, 0)
}

func dfsInternal(startCity, endCity City, path []City, distance int) ([]City, int) {
	if startCity.Id == endCity.Id {
		return path, distance
	}

	answer := math.MaxInt
	bestPath := path

	for _, neighbor := range startCity.Neighbors {
		if contains(path, neighbor) {
			continue
		}

		path = append(path, neighbor)
		shortestPath, shortestDistance := dfsInternal(neighbor, endCity, path, distance+1)
		if shortestDistance < answer {
			answer = shortestDistance
			bestPath = shortestPath
		}
		path = path[:len(path)-1]
	}
	return bestPath, answer
}

func contains(cities []City, city City) bool {
	for _, check := range cities {
		if check.Id == city.Id {
			return true
		}
	}
	return false
}

func main() {
	r := gin.Default()

	r.POST("/dfs", func(c *gin.Context) {
		var parameters ParametersDFS
		err := c.BindJSON(&parameters)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if parameters.StartIndex < 0 || parameters.StartIndex > len(parameters.Cities) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if parameters.StartIndex > parameters.EndIndex {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if parameters.EndIndex < 0 || parameters.EndIndex > len(parameters.Cities) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		path, _ := dfs(parameters)
		idsPath := make([]int, 0)
		for _, city := range path {
			idsPath = append(idsPath, city.Id)
		}
		c.JSON(http.StatusOK, idsPath)
	})

	r.Run(":3333")
}
