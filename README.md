# GDSC Fall 2024 Workshop: Real-Time Chat App Workshop with Go, Fiber, and HTMX

Welcome! In this workshop, we will be using the following technologies to build a real-time chat app:

## Go

Go, often called Golang, is a beginner-friendly programming language created by Google. Imagine it as a modern, streamlined version of C or Java. Go is great for newcomers because:

- It has a simple, easy-to-read syntax
- It compiles quickly, so you can see results fast
- It includes helpful tools like a formatter and package manager
- It's excellent for building web servers and command-line tools

Go is particularly good at handling multiple tasks at once (concurrency), which is super useful for web applications.

## Fiber

Fiber is a web framework for Go. If you've ever used Express.js with Node.js, Fiber will feel familiar. It's designed to make building web applications in Go easier and faster. With Fiber, you can:

- Create routes for your web app (like "/home" or "/about")
- Handle web requests and responses
- Serve static files (like images or CSS)
- Use middleware to add extra functionality

Fiber is a great choice for beginners because it's straightforward to use and has clear documentation.

## HTMX

HTMX is a cool tool that lets you add interactive features to your web pages without writing much (or any) JavaScript. It works by adding special attributes to your HTML elements. For example:

- You can make a button load new content without refreshing the page
- You can update parts of your page automatically
- You can even handle things like file uploads

HTMX is perfect for beginners because it lets you create dynamic web pages using mostly HTML, which you're probably already familiar with. 

This stack (Go + Fiber + HTMX) works great for developers looking to build high-performance, scalable web applications with a lean front-end approach, leveraging server-side rendering and minimal client-side JavaScript. Basically, Go and Fiber handle the server-side logic, while HTMX makes your front-end interactive with minimal effort.

**For a more thorough breakdown of these technologies, watch the following brief videos:**
- [Go in 100 seconds](https://www.youtube.com/watch?v=446E-r0rXHI)
- [HTMX in 100 seconds](https://www.youtube.com/watch?v=r-GSGH2RxJs)

By the end of this workshop, you'll have created a reactive frontend without JavaScript and learned how to leverage Fiber's versatility with WebSockets.

## Prerequisites

- Solid understanding of Go and HTTP servers (optional)
- Go installed (version 1.22 or later)

## Workshop Outline

1. Project Setup
2. Installing Dependencies
3. Creating the Main Application
4. Setting Up Static Files
5. Implementing Handlers
6. Creating the Message Structure
7. Implementing WebSocket Functionality
8. Integrating WebSocket with Routes and HTMX

## 1. Project Setup

Create a new folder named `go-chat` and initialize your Go module:

```bash
mkdir go-chat
cd go-chat
go mod init github.com/yourusername/go-chat
```

## 2. Installing Dependencies

Install the required libraries:

```bash
go get -u github.com/gofiber/fiber/v2
go get -u github.com/gofiber/websocket/v2
go get -u github.com/gofiber/template/html/v2
```

## 3. Creating the Main Application

Create a `main.go` file in the root directory with a basic Fiber server:

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()
    
    app.Get("/ping", func(ctx *fiber.Ctx) error {
        return ctx.SendString("Welcome to fiber")
    })
    
    app.Listen(":3000")
}
```

## 4. Setting Up Static Files

Create `static` and `views` folders. In the `views` folder, create `index.html` and `messages.html`. In the `static` folder, create `style.css`[1].

Update `main.go` to handle static files:

```go
viewsEngine := html.New("./views", ".html")
app := fiber.New(fiber.Config{
    Views: viewsEngine,
})
app.Static("/static/", "./static")
```

## 5. Implementing Handlers

Create a `handlers.go` file:

```go
package handlers

import "github.com/gofiber/fiber/v2"

type AppHandler struct{}

func NewAppHandler() *AppHandler {
    return &AppHandler{}
}

func (a *AppHandler) HandleGetIndex(ctx *fiber.Ctx) error {
    return ctx.Render("index", fiber.Map{})
}
```

Update `main.go` to use the new handler:

```go
appHandler := handlers.NewAppHandler()
app.Get("/", appHandler.HandleGetIndex)
```

## 6. Creating the Message Structure

Create a `message.go` file:

```go
package main

type Message struct {
    Text string `json:"text"`
}
```

## 7. Implementing WebSocket Functionality

Create a `websocket.go` file to handle WebSocket connections and message broadcasting[1].

## 8. Integrating WebSocket with Routes and HTMX

Update `main.go` to include WebSocket routes:

```go
server := NewWebSocket()
app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
    server.HandleWebSocket(ctx)
}))
go server.HandleMessages()
```

Modify `index.html` to include HTMX WebSocket attributes:

```html
<div class="chat-window" hx-ext="ws" ws-connect="/ws">
    <div class="messages" id="messages" hx-swap="beforeend" hx-swap-oob="beforeend">
        <!-- Messages will be appended here -->
    </div>
    <form id="form" ws-send>
        <!-- Form content -->
    </form>
</div>
```

## Running the Application

Start the server:

```bash
go run main.go
```

Visit `http://localhost:3000` in your browser to test the chat application[1].

## Conclusion

Congratulations! You've built a real-time chat application using Go, Fiber, and HTMX. This project demonstrates how to create a reactive frontend without JavaScript and leverage WebSockets for real-time communication[1].

Feel free to extend this project by adding features like user authentication, multiple chat rooms, or message persistence.

## Acknowledgments

Thanks to FreeCodeCamp for providing the template to this workshop [here](https://www.freecodecamp.org/news/real-time-chat-with-go-fiber-htmx/)!
