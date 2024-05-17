package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

    e := echo.New()
    e.Use(middleware.Logger())
    
    player := newPlayer(1, 1)
    e.Renderer = newTemplate()

    e.Static("/js", "js")
    e.Static("/css", "css")
    e.Static("/images", "images")
    e.Static("/frames", "frames")

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "player", player)
    })
    
    e.GET("/video/:id/buffer-init", func(c echo.Context) error {
        idStr := c.Param("id")
        e.Logger.Print(idStr)
        first := true
        for i := 0; i < 150; i++ {
            nextFrame := player.BufferData.EndBuffer + 1
            path := fmt.Sprintf("frames/%s/%d.jpg", idStr, nextFrame)
            if _, err := os.Stat(path); err != nil {
                player.BufferData.EndBuffer = 1
            } else {
                player.BufferData.EndBuffer += 1
            }
            if first {
                c.Render(200, "bufferFrame-first-oob", player.BufferData)
                first = false
            } else {
                c.Render(200, "bufferFrame-oob", player.BufferData)
            }
        }
        c.Render(200, "frameSwap-oob", player.BufferData)
        return c.Render(200, "buffer-fetch", player.BufferData)
    })

    e.GET("/video/:id/buffer/:last-frame", func(c echo.Context) error {
        idStr := c.Param("id")
        lastFrameStr := c.Param("last-frame")
        lastFrame, err := strconv.Atoi(lastFrameStr)
        if err != nil {
            player.BufferData.EndBuffer = player.VideoData.Time
        } else {
            player.BufferData.EndBuffer = int32(lastFrame)
        }
        for i := 0; i < 60; i++ {
            nextFrame := player.BufferData.EndBuffer + 1
            path := fmt.Sprintf("frames/%s/%d.jpg", idStr, nextFrame)
            if _, err := os.Stat(path); err != nil {
                player.BufferData.EndBuffer = 1
            } else {
                player.BufferData.EndBuffer += 1
            }
            c.Render(200, "bufferFrame-oob", player.BufferData)
        }
        return c.Render(200, "buffer-fetch", player.BufferData)
    })
    
    e.GET("/video/:id/next-frame", func(c echo.Context) error {
        nextFrame := player.VideoData.Time + 1
        path := fmt.Sprintf("frames/%s/%d.jpg", c.Param("id"), nextFrame)
        if _, err := os.Stat(path); err != nil {
            player.VideoData.Time = 1
        } else {
            player.VideoData.Time += 1
        }
        return c.Render(200, "frame", player.VideoData)
    })

    e.Logger.Fatal(e.Start(":3902"))

}

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
    return &Templates{
        templates: template.Must(template.ParseGlob("./views/*.html")),
    }
}
