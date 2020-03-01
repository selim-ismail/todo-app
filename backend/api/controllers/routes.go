package controllers

func (server *Server) initializeRoutes() {
	v1 := server.Router.Group("/api/v1")
	{
		v1.GET("/todo-lists", server.GetAllTodoLists)
		v1.GET("/todo/:id", server.GetTodoById)
	}
}
