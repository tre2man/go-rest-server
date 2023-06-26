package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album 구조체 타입 정의
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album 객체 (메모리 db)
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// albums 핸들러 함수
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// albums 데이터 추가
func postAlbums(c *gin.Context) {
	// 새 앨범 객체 생성
	var newAlbum album

	// 데이터 파싱하기
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// 실제 데이터 추가
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// id로 특정 앨범 가져오기
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// albums에서 id로 앨범 찾기
	// albums 순회
	for _, a := range albums {
		// id 구분하기
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// get hello
func getHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}

func main() {
	router := gin.Default()
	router.GET("/", getHello)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run(":8080")
}
