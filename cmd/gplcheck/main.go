package main

import (
	"context"

	"go.uber.org/fx"

	"gplcheck/pkg/app"
	"gplcheck/pkg/common"
	"gplcheck/pkg/controllers"
	"gplcheck/pkg/tui"
)

func main() {
	// app := app.NewApp()
	// app.Run()
	di := fx.New(
		fx.Provide(app.NewApp,
			tui.NewTui,
			tui.InitApp,
			tui.NewMainFrame,
			tui.NewFileView,
			func() string { return "." },
			tui.NewResultView,
			tui.NewStatusView,
			controllers.NewFileViewController,
			controllers.NewResultViewController,
			controllers.NewStatusViewController,
			common.NewNotifier,
		),
		fx.Invoke(Run),
	)
	di.Start(context.Background())

	di.Stop(context.Background())
}

func Run(lifecycle fx.Lifecycle, app *app.App) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			app.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
