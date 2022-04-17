package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bishehngliu/posts/controller"
	"github.com/bishehngliu/posts/repository"
	"github.com/bishehngliu/posts/router"
	"github.com/bishehngliu/posts/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	postRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const PORT string = ":8000"

	postRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "up and running")
	})

	postRouter.GET("/posts", postController.GetPosts)
	postRouter.POST("/posts", postController.AddPost)

	log.Println("server is listening on port ", PORT)
	postRouter.SERVE(PORT)
}
