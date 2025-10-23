package services

import (
	"example/service/app"
	"example/service/pkg/logger"
	"context"
	"fmt"
	"time"

	pkgEnv "example/service/pkg/env"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:              "start",
		Short:            "Start all services",
		Long:             "Start all services",
		PersistentPreRun: rootPreRun,
		RunE:             runAllServices,
		//RunE: runRestServices,
	}
	serviceName = "Example/Service"
)

func rootPreRun(cmd *cobra.Command, args []string) {
	// initiate logger
	logLvl := zerolog.InfoLevel
	if pkgEnv.NewEnv().GetEnvironmentName() != "production" {
		logLvl = zerolog.DebugLevel
	}

	logger.InitGlobalLogger(&logger.Config{
		ServiceName: serviceName,
		Level:       logLvl,
	})
}

func StartCmd() *cobra.Command {
	return startCmd
}

func runRestServices(cmd *cobra.Command, args []string) error {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}

	osSignal, rest := app.RunRestService(configPath)

	fmt.Println("ALL Services is running...")

	select {
	case err := <-rest.ListenError():
		return fmt.Errorf("error starting rest server: %w", err)
	case <-osSignal:
		// Graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := rest.Shutdown(ctx); err != nil {
			fmt.Printf("Failed to shutdown REST server gracefully: %v\n", err)
		}

		fmt.Println("Exiting gracefully...")
	}

	return nil
}

func runAllServices(cmd *cobra.Command, args []string) error {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}

	osSignal, rest, gRpc, cron, kafka := app.RunAllServices(configPath)

	fmt.Println("ALL Services is running...")

	select {
	case err := <-rest.ListenError():
		return fmt.Errorf("error starting rest server: %w", err)
	case err := <-gRpc.ListenError():
		return fmt.Errorf("error starting gRpc server: %w", err)
	case err := <-cron.ListenError():
		return fmt.Errorf("error starting cron server: %w", err)
	case err := <-kafka.ListenError():
		return fmt.Errorf("error starting kafka server: %w", err)
	case <-osSignal:
		// Graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := rest.Shutdown(ctx); err != nil {
			fmt.Printf("Failed to shutdown REST server gracefully: %v\n", err)
		}
		if err := gRpc.Shutdown(ctx); err != nil {
			fmt.Printf("Failed to shutdown gRPC server gracefully: %v\n", err)
		}
		if err := cron.Shutdown(ctx); err != nil {
			fmt.Printf("Failed to shutdown gRPC server gracefully: %v\n", err)
		}
		if err := kafka.Shutdown(ctx); err != nil {
			fmt.Printf("Failed to shutdown Kafka server gracefully: %v\n", err)
		}
		fmt.Println("Exiting gracefully...")
	}

	return nil
}