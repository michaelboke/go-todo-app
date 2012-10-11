package main

import (
	"flag"
	"reflect"
	"github.com/robfig/revel"
	
  "github.com/yanatan16/go-todo-app/app/models"
  
  "github.com/yanatan16/go-todo-app/app/controllers"
  
)

var (
	addr *string = flag.String("addr", "", "Address to listen on")
	port *int = flag.Int("port", 0, "Port")
	importPath *string = flag.String("importPath", "", "Path to the app.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	rev.INFO.Println("Running revel server")
	flag.Parse()
	rev.Init(*importPath, "dev")
	
	rev.RegisterController((*controllers.Application)(nil),
		[]*rev.MethodType{
			&rev.MethodType{
				Name: "Index",
				Args: []*rev.MethodArg{ 
			  },
				RenderArgNames: map[int][]string{ 
					23: []string{ 
					},
				},
			},
			&rev.MethodType{
				Name: "Register",
				Args: []*rev.MethodArg{ 
			  },
				RenderArgNames: map[int][]string{ 
					27: []string{ 
					},
				},
			},
			&rev.MethodType{
				Name: "SaveUser",
				Args: []*rev.MethodArg{ 
					&rev.MethodArg{Name: "user", Type: reflect.TypeOf((*models.User)(nil)) },
					&rev.MethodArg{Name: "verifyPassword", Type: reflect.TypeOf((*string)(nil)) },
			  },
				RenderArgNames: map[int][]string{ 
				},
			},
			&rev.MethodType{
				Name: "Login",
				Args: []*rev.MethodArg{ 
					&rev.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&rev.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
			  },
				RenderArgNames: map[int][]string{ 
				},
			},
			&rev.MethodType{
				Name: "Logout",
				Args: []*rev.MethodArg{ 
			  },
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	rev.RegisterController((*controllers.Todo)(nil),
		[]*rev.MethodType{
			&rev.MethodType{
				Name: "Index",
				Args: []*rev.MethodArg{ 
			  },
				RenderArgNames: map[int][]string{ 
					61: []string{ 
					},
					67: []string{ 
					},
					71: []string{ 
					},
				},
			},
			&rev.MethodType{
				Name: "JsonReadList",
				Args: []*rev.MethodArg{ 
			  },
				RenderArgNames: map[int][]string{ 
				},
			},
			&rev.MethodType{
				Name: "JsonReadItem",
				Args: []*rev.MethodArg{ 
					&rev.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
			  },
				RenderArgNames: map[int][]string{ 
				},
			},
			&rev.MethodType{
				Name: "JsonUpdateItem",
				Args: []*rev.MethodArg{ 
					&rev.MethodArg{Name: "item", Type: reflect.TypeOf((**models.Item)(nil)) },
			  },
				RenderArgNames: map[int][]string{ 
				},
			},
			&rev.MethodType{
				Name: "JsonCreateItem",
				Args: []*rev.MethodArg{ 
					&rev.MethodArg{Name: "desc", Type: reflect.TypeOf((*string)(nil)) },
					&rev.MethodArg{Name: "done", Type: reflect.TypeOf((*bool)(nil)) },
			  },
				RenderArgNames: map[int][]string{ 
				},
			},
			&rev.MethodType{
				Name: "JsonDeleteItem",
				Args: []*rev.MethodArg{ 
					&rev.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
			  },
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	rev.Run(*addr, *port)
}
