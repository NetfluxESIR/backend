package cmd

import (
	"github.com/NetfluxESIR/backend/internal/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the Netflux API server.",
		Long:  "Start the Netflux API server.",
		RunE: func(cmd *cobra.Command, args []string) error {
			setLogLevel(loglevel)
			cfg := &server.Config{
				Port: port,
				Host: host,
				DSN:  dsn,
				Logger: log.WithFields(log.Fields{
					"service": "netflux",
				}),
				AdminAccount:  adminAccount,
				AdminPassword: adminPassword,
				S3Bucket:      s3Bucket,
				S3Region:      s3Region,
				S3AccessKey:   os.Getenv("AWS_ACCESS_KEY_ID"),
				S3SecretKey:   os.Getenv("AWS_SECRET_ACCESS_KEY"),
			}
			if err := cfg.Validate(); err != nil {
				return err
			}
			srv, err := server.New(cmd.Context(), cfg)
			if err != nil {
				return err
			}
			go func() {
				err := srv.Run(cmd.Context())
				if err != nil {
					panic(err)
				}
			}()
			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt)
			<-sig
			return srv.Stop(cmd.Context())
		},
	}

	// port is the port to listen on.
	port int
	// host is the host to listen on.
	host string
	// dsn is the data source name.
	dsn string
	// loglevel is the log level.
	loglevel string
	// adminAccount is the admin account.
	adminAccount string
	// adminPassword is the admin password.
	adminPassword string
	// s3Bucket is the S3 bucket name.
	s3Bucket string
	// s3Region is the S3 region.
	s3Region string
)

func setLogLevel(loglevel string) {
	if loglevel == "trace" {
		log.SetLevel(log.TraceLevel)
	}
	if loglevel == "debug" {
		log.SetLevel(log.DebugLevel)
	}
	if loglevel == "info" {
		log.SetLevel(log.InfoLevel)
	}
	if loglevel == "warn" {
		log.SetLevel(log.WarnLevel)
	}
	if loglevel == "error" {
		log.SetLevel(log.ErrorLevel)
	}
	if loglevel == "fatal" {
		log.SetLevel(log.FatalLevel)
	}
	if loglevel == "panic" {
		log.SetLevel(log.PanicLevel)
	}
}

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to listen on.")
	serveCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to listen on.")
	serveCmd.Flags().StringVarP(&dsn, "dsn", "d", "netflux:netflux@postgres:5432/netflux?sslmode=disable", "Data source name.")
	serveCmd.Flags().StringVarP(&loglevel, "loglevel", "l", "info", "Log level.")
	serveCmd.Flags().StringVarP(&adminAccount, "admin-account", "a", "admin", "Admin account.")
	serveCmd.Flags().StringVarP(&adminPassword, "admin-password", "P", "admin", "Admin password.")
	serveCmd.Flags().StringVarP(&s3Bucket, "s3-bucket", "b", "", "S3 bucket name.")
	serveCmd.Flags().StringVarP(&s3Region, "s3-region", "r", "", "S3 region.")
}
