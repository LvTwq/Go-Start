package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/getb", GetDataB)

	r.Run(":8080")
}

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}
