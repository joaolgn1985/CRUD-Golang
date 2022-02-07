# CRUD-Golang

### First Hello-Word
The first step is to create the hello folder, which will contain our first Go script. Create the /hello folder inside the src folder of your Go Workspace, which we created earlier in the Preparing the Environment exercise.

``` 
folder-user/
└── go
    ├── bin
    ├── pkg
    └── src
        └── hello
```

All projects that you will develop in the Go language will be in their own folders in the src directory.

Now let's create the first file, hello.go inside the /hello folder:

```
folder-user/
└── go
    ├── bin
    ├── pkg
    └── src
        └── hello
           └── hello.go
```

And in this file we are going to create our first program.

Like every project in Go, we need to define what the initial package will be with the instruction:

```
//hello.go
package main

import "fmt"

func main(){
    fmt.Println("Hello World com Go!")
}
```

To run our code, just use the command `go run hello.go` in the terminal inside the folder that contains our file with the program source code and the executable will be automatically created and executed:

```
// Terminal
go run hello.go
Hello World com Go!
```

# Function fmt.Scan()

In this chapter we will create the script's intro message and the menu that will be displayed to the user, in addition to capturing the choice he made.

```
//hello.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Joao"
	fmt.Println("variable name is:", reflect.TypeOf(name))
	version := 1.1

	fmt.Println("Hello, Mr.", name)
	fmt.Println("This is program is in version", version)

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	var comand int
	fmt.Scanf("%d", &comand)
	fmt.Println("Value of the command variable is:", comand)
}

```

We were able to create our intro message and user options menu, as well as display which option they chose. We are already advancing in our program!

# Control flux using if or swith


```
package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Joao"
	fmt.Println("variable name is:", reflect.TypeOf(name))
	version := 1.1

	fmt.Println("Hello, Mr.", name)
	fmt.Println("This is program is in version", version)

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	var comand int
	fmt.Scanf("%d", &comand)
	fmt.Println("Value of the command variable is:", comand)

	switch comand {
	case 1:
		fmt.Println("Starting monitoring...")
	case 2:
		fmt.Println("Opening logs...")
	case 0:
		fmt.Println("Exiting of the program...")
	default:
		fmt.Println("Please choose one the options above...")
	}
}
```

#### Simple CRUD made in Golang to exercise language learning.

For those unfamiliar with this term, CRUD is an English acronym for Create, Read, Update and Delete. In this 'program type', we can easily identify CRUD operations. Are they:
C = Create or add entries.
R = Read, retrieve or view existing entries.
U = Update/Edit existing entries.
D= Remove existing entries.

Before cloning the project, make sure you have installed and configured GO on your machine. To run the project, just type at your prompt: go run main.go Remember to include your database information.
