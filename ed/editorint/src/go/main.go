package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() {
		e.cursor = e.cursor.Prev()
		return
	}
	if e.line != e.lines.Front() {
		e.line = e.line.Prev()      
		e.cursor = e.line.Value.End()
	}
}

func (e *Editor) KeyEnter() {
	newline := NewList[rune]()
	e.lines.Insert(e.line.Next(), newline)
	cursor := e.cursor
	for cursor != e.line.Value.Back().Next() {
		next := cursor.Next()
		newline.PushBack(cursor.Value)
		e.line.Value.Erase(cursor)
		cursor = next
	}
	e.line = e.line.Next()
	e.cursor = e.line.Value.Front()
}


func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.Back().Next() {
		e.cursor = e.cursor.Next()
		return
	}
	if e.line != e.lines.Back() {
		e.line = e.line.Next() 
		e.cursor = e.line.Value.Front()
	}
}

func (e *Editor) KeyUp() {
	prevLine := e.line.Prev()
	if prevLine == prevLine.root {
		return
	}
	index := e.line.Value.IndexOf(e.cursor)
	if index == -1 {
		return
	}
	if e.line.Value.Size() > prevLine.Value.Size() {
		e.line = prevLine
		e.cursor = e.line.Value.Back()
		return
	}
	e.line = prevLine
	e.cursor = e.line.Value.Front()
	if e.cursor == nil {
		return 
	}
	for i := 0; i < index; i++ {
		nextCursor := e.cursor.Next()
		if nextCursor == nil {
			break
		}
		e.cursor = nextCursor
	}
}

func (e *Editor) KeyDown() {
	nextLine := e.line.Next()
	if nextLine == nextLine.root {
		return
	}
	index := e.line.Value.IndexOf(e.cursor)
	if index == -1 {
		return
	}
	if e.line.Value.Size() > nextLine.Value.Size() {
		e.line = nextLine
		e.cursor = e.line.Value.Back()
		return
	}
	e.line = nextLine
	e.cursor = e.line.Value.Front()
	if e.cursor == nil {
		return 
	}
	for i := 0; i < index; i++ {
		nextCursor := e.cursor.Next()
		if nextCursor == nil {
			break
		}
		e.cursor = nextCursor
	}
}

func (e *Editor) KeyBackspace() {
	if e.cursor != e.line.Value.Front(){
		e.line.Value.Erase(e.cursor.Prev())
	} else if (e.cursor == e.line.Value.Front()) && (e.line != e.lines.Front()){
		index := e.line.Value.Size()
		prevLine := e.line.Prev()
		caracter := e.line.Value.Front()
		for i := 0; i < index; i++ {
			prevLine.Value.PushBack(caracter.Value)
			caracter = caracter.Next()
		} 
		e.lines.Erase(e.line)
		e.line = prevLine
		e.cursor = e.line.Value.End()

	}
}

func (e *Editor) KeyDelete() {
	if e.cursor != e.line.Value.Back() && e.cursor != e.line.Value.End() {
		newCursor := e.cursor.Next()
		e.line.Value.Erase(e.cursor)
		e.cursor = newCursor
	} else if e.cursor == e.line.Value.Back() && e.line.Next() != e.lines.End() {
		nextLine := e.line.Next()
		index := nextLine.Value.Size()
		caracter := nextLine.Value.Front()
		
		for i := 0; i < index; i++ {
			e.line.Value.PushBack(caracter.Value)
			caracter = caracter.Next()
		}
		e.lines.Erase(nextLine)
	}
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
