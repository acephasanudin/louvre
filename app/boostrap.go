package app

import (
	"example/service/app/adapter"
	"example/service/app/services"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func TerminateSignal() chan os.Signal {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	return term
}

func RunRestService(configPath string) (chan os.Signal, services.RestService) {
	cfg, r := adapter.InitAdapters(configPath)

	rs := services.NewRestService(&services.RestOptions{
		Config:   &cfg.Get().Transport.Rest,
		Keycloak: &cfg.Get().Datasource.Keycloak,
		Service:  r.Transports.Rest,
	})

	go rs.RunRest()

	fmt.Println("REST Service is running...")

	return TerminateSignal(), rs
}

func RunGRPCService(configPath string) (chan os.Signal, services.GRPCService) {
	cfg, r := adapter.InitAdapters(configPath)

	gs := services.NewGRPCService(&services.GRPCOptions{
		Config:  &cfg.Get().Transport.GRPC,
		Service: r.Transports.GRPC,
	})

	go gs.RunGRPC()

	fmt.Println("GRPC Service is running...")

	return TerminateSignal(), gs
}

func RunCronService(configPath string) (chan os.Signal, services.CronService) {
	cfg, r := adapter.InitAdapters(configPath)

	cs := services.NewCronService(&services.CronOptions{
		Config:  &cfg.Get().Transport.Cron,
		Service: r.Transports.Cron,
	})

	go cs.RunCron()

	fmt.Println("Cron Service is running...")

	return TerminateSignal(), cs
}

func RunkafkaListnerService(configPath string) (chan os.Signal, services.KafkaListenerService) {
	cfg, r := adapter.InitAdapters(configPath)

	kl := services.NewKafkaListener(&services.KafkaListenerOptions{
		Config: &cfg.Get().Transport.Kafka,
		Tracer: r.Tracer,
	})

	go kl.RunListener()

	fmt.Println("Kafka Service is running...")

	return TerminateSignal(), kl
}

func RunAllServices(configPath string) (chan os.Signal, services.RestService, services.GRPCService, services.CronService, services.KafkaListenerService) {
	cfg, r := adapter.InitAdapters(configPath)
	tsCfg := cfg.Get().Transport
	keycloakCfg := cfg.Get().Datasource.Keycloak

	fmt.Println("start process NewRestService")
	rs := services.NewRestService(&services.RestOptions{
		Keycloak: &keycloakCfg,
		Config:   &tsCfg.Rest,
		Service:  r.Transports.Rest,
	})
	fmt.Println("end process NewRestService")

	fmt.Println("start process NewGRPCService")
	gs := services.NewGRPCService(&services.GRPCOptions{
		Config:  &tsCfg.GRPC,
		Service: r.Transports.GRPC,
	})
	fmt.Println("end process NewGRPCService")

	fmt.Println("start process NewCronService")
	cs := services.NewCronService(&services.CronOptions{
		Config:  &tsCfg.Cron,
		Service: r.Transports.Cron,
	})
	fmt.Println("end process NewCronService")

	fmt.Println("start process NewKafkaListener")
	kl := services.NewKafkaListener(&services.KafkaListenerOptions{
		Config:         &cfg.Get().Transport.Kafka,
		Tracer:         r.Tracer,
		KafkaTransport: r.Transports.Kafka,
	})
	fmt.Println("end process NewKafkaListener")

	go rs.RunRest()
	go gs.RunGRPC()
	go cs.RunCron()
	go kl.RunListener()

	return TerminateSignal(), rs, gs, cs, kl
}