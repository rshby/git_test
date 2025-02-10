package factory

import (
	"fmt"
	"time"
)

type Content interface {
	Play()
}

type CloudContent struct {
}

func (c *CloudContent) Play() {
	fmt.Println("this is cloud content about kubernetes")
}

type CodingContent struct {
}

func (c *CodingContent) Play() {
	fmt.Println("konten coding")
}

type GrebekRumah struct {
}

func (g *GrebekRumah) Play() {
	fmt.Println("ini adalah konten yang sangat penting")
}

type ContentCreator interface {
	Produce(t time.Time) Content
}

type Imre struct {
}

func (i *Imre) Produce(t time.Time) Content {
	if t.Weekday() == time.Tuesday {
		return &CloudContent{}
	} else if time.Weekday(t.Day()) == time.Thursday {
		return &CodingContent{}
	} else {
		return nil
	}
}

type Atta struct {
}

func (a *Atta) Produce() Content {
	return &GrebekRumah{}
}
