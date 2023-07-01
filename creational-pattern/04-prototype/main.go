package main

import (
	"fmt"

	buttonmaker "github.com/winartodev/go-design-pattern/creational-pattern/04-prototype/button-maker"
	musicplaylist "github.com/winartodev/go-design-pattern/creational-pattern/04-prototype/music-playlist"
	shapedrawer "github.com/winartodev/go-design-pattern/creational-pattern/04-prototype/shape-drawer"
)

func main() {
	registry := musicplaylist.NewPlaylistRegistry()

	regularPlaylist := musicplaylist.RegularPlaylist{
		Name:  "Regular Playlist",
		Track: []string{"Track 1", "Track 2", "Track 3"},
	}

	smartPlaylist := musicplaylist.SmartPlaylist{
		Name:  "Smart Playlist",
		Genre: "Rock",
	}

	registry.RegisterPrototype("Regular Playlist", &regularPlaylist)
	registry.RegisterPrototype("Smart Playlist", &smartPlaylist)

	cloneRegularPlaylist := registry.Clone("Regular Playlist").(*musicplaylist.RegularPlaylist)
	cloneSmartPlaylist := registry.Clone("Smart Playlist").(*musicplaylist.SmartPlaylist)

	fmt.Printf("%+v\n", cloneRegularPlaylist)
	fmt.Printf("%+v\n", cloneSmartPlaylist)

	registry1 := shapedrawer.NewShapeDrawer()

	drawCircle := shapedrawer.Circle{
		Radius: 4.15,
	}

	drawRectangle := shapedrawer.Rectangle{
		Height: 4.10,
		Width:  5.15,
	}

	drawTriangle := shapedrawer.Triangle{
		Base:   5.15,
		Height: 4.10,
	}

	registry1.RegisterPrototype("Circle", &drawCircle)
	registry1.RegisterPrototype("Rectangle", &drawRectangle)
	registry1.RegisterPrototype("Triangle", &drawTriangle)

	registry1.Draw("Circle")
	registry1.Draw("Rectangle")
	registry1.Draw("Triangle")

	buttonRegistry := buttonmaker.NewButtonRegistry()

	buttonLoading := buttonmaker.Button{
		X:     10,
		Y:     40,
		Color: "Red",
	}

	buttonEscape := buttonmaker.Button{
		X:     10,
		Y:     40,
		Color: "Green",
	}

	buttonRegistry.AddItem("Button Loading", &buttonLoading)
	buttonRegistry.AddItem("Button Escape", &buttonEscape)

	buttonLoadingColor := buttonRegistry.GetByColor("Red")
	buttonEscapeColor := buttonRegistry.GetByColor("Green")

	fmt.Printf("button loading color is: %v\n", buttonLoadingColor.GetColor())
	fmt.Printf("button escape color is: %v\n", buttonEscapeColor.GetColor())
}
