package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
)

type SplashScreen struct {
	tview.Flex
	app *tview.Application

	view *tview.TextView

	backScreen [][]rune
}

func NewSplashScreen(app *tview.Application) *SplashScreen {
	s := &SplashScreen{
		Flex: *tview.NewFlex().SetDirection(tview.FlexRow),
		app:  app,

		view: tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignCenter),
	}

	s.view.SetScrollable(false).SetBorder(true).SetBorderPadding(0, 0, 0, 0)

	s.prepareScreen()

	go func() {
		stepColors := []struct {
			txColor, shColor, bgColor string
		}{
			{"#000000", "#000000", "#000000"},
			{"#070101", "#070504", "#000000"},
			{"#0e0302", "#0f0a09", "#000000"},
			{"#150403", "#170f0e", "#000000"},
			{"#1d0605", "#1f1413", "#000000"},
			{"#240706", "#271918", "#000000"},
			{"#2b0907", "#2f1e1d", "#000000"},
			{"#330a09", "#372322", "#000000"},
			{"#3a0c0a", "#3f2827", "#000000"},
			{"#410d0b", "#472d2c", "#000000"},
			{"#490f0d", "#4e3231", "#000000"},
			{"#50110e", "#563736", "#000000"},
			{"#57120f", "#5e3c3b", "#000000"},
			{"#5f1411", "#664140", "#000000"},
			{"#661512", "#6e4645", "#000000"},
			{"#6d1713", "#764c49", "#000000"},
			{"#741814", "#7e514e", "#000000"},
			{"#7c1a16", "#865653", "#000000"},
			{"#831b17", "#8e5b58", "#000000"},
			{"#8a1d18", "#96605d", "#000000"},
			{"#921f1a", "#9d6562", "#000000"},
			{"#99201b", "#a56a67", "#000000"},
			{"#a0221c", "#ad6f6c", "#000000"},
			{"#a8231e", "#b57471", "#000000"},
			{"#af251f", "#bd7976", "#000000"},
			{"#b62620", "#c57e7b", "#000000"},
			{"#be2822", "#cd8380", "#000000"},
			{"#c52923", "#d58885", "#000000"},
			{"#cc2b24", "#dd8d8a", "#000000"},
			{"#d42d26", "#e5938f", "#000000"},
			// {"#d42d26", "#e5938f", "#000000"},
		}
		time.Sleep(100 * time.Millisecond)
		for _, colors := range stepColors {
			text := s.drawScreen(colors.txColor, colors.shColor, colors.bgColor)
			s.view.SetText(text)
			time.Sleep(50 * time.Millisecond)
			app.Draw()
		}

	}()

	s.Flex.Clear().SetDirection(tview.FlexRow).
		AddItem(s.view, 10, 1, true)

	return s
}

func (s *SplashScreen) Width() int {
	return len(s.backScreen[0]) + 10
}

func (s *SplashScreen) Height() int {
	return len(s.backScreen)
}

func (s *SplashScreen) drawScreen(txColor, shColor, bgColor string) string {
	b := strings.Builder{}
	for _, line := range s.backScreen {
		const (
			fg = "fg"
			bg = "bg"
			sh = "sh"
		)
		prev := ""
		for _, r := range line {
			if r == '▓' && prev != fg {
				b.WriteString(fmt.Sprintf("[%s:%s]", txColor, bgColor))
				prev = fg
			} else if r == '░' && prev != sh {
				b.WriteString(fmt.Sprintf("[%s:%s]", shColor, bgColor))
				prev = sh
			} else if r == ' ' && prev != bg {
				b.WriteString(fmt.Sprintf("[%s:%s]", bgColor, bgColor))
				prev = bg
			}
			b.WriteRune(r)
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func (s *SplashScreen) prepareScreen() [][]rune {
	text := `
  ▓▓▓▓▓▓▓░    ▓▓▓▓▓▓▓▓░       ▓▓▓░       ▓▓▓▓▓▓▓░    ▓▓▓▓▓▓▓▓▓░  
 ▓▓░    ▓▓░   ▓▓░    ▓▓░     ▓▓░▓▓░     ▓▓░    ▓▓░   ▓▓░         
 ▓▓░          ▓▓░    ▓▓░    ▓▓░  ▓▓░    ▓▓░          ▓▓░         
 ▓▓░   ▓▓▓▓░  ▓▓▓▓▓▓▓▓░    ▓▓░    ▓▓░    ▓▓▓▓▓▓▓░    ▓▓▓▓▓▓▓░    
 ▓▓░    ▓▓░   ▓▓░    ▓▓░   ▓▓▓▓▓▓▓▓▓░          ▓▓░   ▓▓░         
 ▓▓░    ▓▓░   ▓▓░    ▓▓░   ▓▓░    ▓▓░   ▓▓░    ▓▓░   ▓▓░         
  ▓▓▓▓▓▓▓░    ▓▓▓▓▓▓▓▓░    ▓▓░    ▓▓░    ▓▓▓▓▓▓▓░    ▓▓▓▓▓▓▓▓▓░  `
	lines := strings.Split(text, "\n")[1:]
	var backScreen [][]rune
	for _, line := range lines {
		backScreen = append(backScreen, []rune(line))
	}
	s.backScreen = backScreen
	return backScreen
}
