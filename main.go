package main

import (
	"net/http"
	"wsjgin"
)

// For Test4
//func Logger() wsjgin.HandleFunc {
//	return func(c *wsjgin.Context) {
//		// 开始计时
//		t := time.Now()
//		c.Next()
//		// 计算时间
//		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
//	}
//}

// for test5
//type student struct {
//	Name string
//	Age  int8
//}
//func FormatAsDate(t time.Time) string {
//	year, month, day := t.Date()
//	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
//}

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

	// Test 3
	//r := wsjgin.Default()
	//r.GET("/", func(c *wsjgin.Context) {
	//	c.HTML(http.StatusOK, "<h1>Welcome TO WSJGIN</h1>")
	//})
	//group1 := r.Group("/group1")
	//{
	//	group1.GET("/", func(c *wsjgin.Context) {
	//		c.HTML(http.StatusOK, "<h1>Welcome TO Group1</h1>")
	//	})
	//
	//	group1.GET("/hello", func(c *wsjgin.Context) {
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//	})
	//}
	//group2 := r.Group("/group2")
	//{
	//	group2.GET("/hello/:name", func(c *wsjgin.Context) {
	//		// /hello/wsj
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.GetParameter("name"), c.Path)
	//	})
	//	group2.POST("/login", func(c *wsjgin.Context) {
	//		c.JSON(http.StatusOK, wsjgin.H{
	//			"username": c.PostForm("username"),
	//			"password": c.PostForm("password"),
	//		})
	//	})
	//}
	//r.Run(":2345")

	// Test 4
	//r := wsjgin.Default()
	//r.Use(Logger())
	//r.GET("/", func(c *wsjgin.Context) {
	//	c.HTML(http.StatusOK, "<h1>Welcome To WSJGIN</h1>")
	//})
	//r.Run(":9999")

	// Test 5
	//r := wsjgin.Default()
	//r.Use(wsjgin.Logger())
	//r.SetFuncMap(template.FuncMap{
	//	"FormatAsDate": FormatAsDate,
	//})
	//r.LoadHTMLGlob("templates/*")
	//r.Static("/assets", "./static")
	//
	//stu1 := &student{Name: "Shijian", Age: 20}
	//stu2 := &student{Name: "Sword", Age: 19}
	//r.GET("/", func(c *wsjgin.Context) {
	//	c.HTML(http.StatusOK, "testcss1.tmpl", nil)
	//})
	//r.GET("/students_info", func(c *wsjgin.Context) {
	//	c.HTML(http.StatusOK, "test1.tmpl", wsjgin.H{
	//		"title":  "wsjgin",
	//		"stuArr": [2]*student{stu1, stu2},
	//	})
	//})
	//r.GET("/date", func(c *wsjgin.Context) {
	//	c.HTML(http.StatusOK, "testfunc1.tmpl", wsjgin.H{
	//		"title": "wsjgin",
	//		"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
	//	})
	//})
	//r.Run(":3478")

	// test 6
	r := wsjgin.NewDefault()
	r.GET("/", func(c *wsjgin.Context) {
		c.String(http.StatusOK, "Hello WSJGIN\n")
	})
	// 索引超出界限错误,以测试recovery
	r.GET("/panic", func(c *wsjgin.Context) {
		names := []string{"ShijianWang"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":2348")

}
