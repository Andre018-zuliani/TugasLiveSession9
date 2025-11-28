<<<<<<< HEAD
package main

import (
	"fmt"
	"session-9/handler"
	"session-9/repository"
	"session-9/service"
)

func main() {
	repo := repository.NewStudentRepository("data/student.json")
	svc := service.NewStudentService(repo)
	h := handler.NewStudentHandler(svc)

	for {
		fmt.Println("=== Student CLI ===")
		fmt.Println("1. List students")
		fmt.Println("2. Create student")
		fmt.Println("0. Exit")
		fmt.Print("Choose menu: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			out, err := h.ListStudents()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(out)

		case "2":
			var name string
			var age int

			fmt.Print("Name: ")
			fmt.Scanln(&name)

			fmt.Print("Age: ")
			fmt.Scanln(&age)

			out, err := h.CreateStudent(name, age)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(out)

		case "0":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
	}
}
=======
package main

import (
	"fmt"
	"session-9/handler"
	"session-9/repository"
	"session-9/service"
)

func main() {
	repo := repository.NewStudentRepository("data/student.json")
	svc := service.NewStudentService(repo)
	h := handler.NewStudentHandler(svc)

	for {
		fmt.Println("=== Student CLI ===")
		fmt.Println("1. List students")
		fmt.Println("2. Create student")
		fmt.Println("0. Exit")
		fmt.Print("Choose menu: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			out, err := h.ListStudents()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(out)

		case "2":
			var name string
			var age int

			fmt.Print("Name: ")
			fmt.Scanln(&name)

			fmt.Print("Age: ")
			fmt.Scanln(&age)

			out, err := h.CreateStudent(name, age)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(out)

		case "0":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
	}
}
>>>>>>> 0ce222e71a1ed1664ffb65f41ab219bc70574b1f
