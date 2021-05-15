package main

import (
	"net/http"
	"wsjgin"
)

func main() {
	// TEST 1:
	//r := wsjgin.Default()
	//r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	//})
	//r.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	for k, v := range request.Header {
	//		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
	//	}
	//})
	//
	//r.Run(":3456")

	// Test 2
	//r := wsjgin.Default()
	//r.GET("/", func(c *wsjgin.Context) {
	//	c.HTML(http.StatusOK, "<h1>Welcome to wsjgin!</h1>")
	//})
	//r.GET("/hello", func(c *wsjgin.Context) {
	//	// /hello?name=wsj
	//	c.String(http.StatusOK, "Hello %s !, Welcome to wsjgin!", c.Query("name"))
	//})
	//r.GET("/hello/:name", func(c *wsjgin.Context) {
	//	c.String(http.StatusOK, "Hello %s !, Welcome to wsjgin!", c.GetParameter("name"))
	//})
	//r.GET("/assets/*filepath", func(c *wsjgin.Context) {
	//	c.JSON(http.StatusOK, wsjgin.H{
	//		"filepath": c.GetParameter("filepath"),
	//	})
	//})
	//r.POST("/login", func(c *wsjgin.Context) {
	//	c.JSON(http.StatusOK, wsjgin.H{
	//		"username": c.PostForm("username"),
	//		"password": c.PostForm("password"),
	//	})
	//})
	//r.Run(":3432")
	r := wsjgin.Default()
	r.GET("/", func(c *wsjgin.Context) {
		c.HTML(http.StatusOK, "<h1>Welcome TO WSJGIN</h1>")
	})
	group1 := r.Group("/group1")
	{
		group1.GET("/", func(c *wsjgin.Context) {
			c.HTML(http.StatusOK, "<h1>Welcome TO Group1</h1>")
		})

		group1.GET("/hello", func(c *wsjgin.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	group2 := r.Group("/group2")
	{
		group2.GET("/hello/:name", func(c *wsjgin.Context) {
			// /hello/wsj
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.GetParameter("name"), c.Path)
		})
		group2.POST("/login", func(c *wsjgin.Context) {
			c.JSON(http.StatusOK, wsjgin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	r.Run(":2345")
}
