package main

/*
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8181") // listen and serve on 0.0.0.0:8080
}
*/
import (
	db "go-gin-blog/database"
 router "go-gin-blog/routers"
)

func main() {
   //���ݿ�
   defer db.SqlDB.Close()

   //·�ɲ���
   router:=router.InitRouter()

   //��̬��Դ
   router.Static("/static", "./static")

   //���еĶ˿�
   router.Run(":8000")

}
