package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/renosyah/receptionist-core-api/auth"
	mid "github.com/renosyah/receptionist-core-api/midtrans"
	"github.com/renosyah/receptionist-core-api/router"

	"github.com/gorilla/mux"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dbPool  *sql.DB
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use: "app",
	PreRun: func(cmd *cobra.Command, args []string) {

		router.Init(dbPool, &mid.PaymentGateway{
			ServerKey: viper.GetString("midtrans.server_key"),
			ClientKey: viper.GetString("midtrans.client_key"),
			Env:       viper.GetString("midtrans.enviroment"),
			CallBack:  viper.GetString("midtrans.callback"),
		})
		auth.Init(dbPool, &auth.Option{
			Enable: viper.GetBool("auth.enable"),
		})

	},
	Run: func(cmd *cobra.Command, args []string) {

		r := mux.NewRouter()

		// customer auth
		r.Handle("/customer/auth/login", router.HandlerFunc(router.HandlerLoginCustomer)).Methods(http.MethodPost)
		r.Handle("/customer/auth/register", router.HandlerFunc(router.HandlerRegisterCustomer)).Methods(http.MethodPost)

		// owners auth
		r.Handle("/owner/auth/login", router.HandlerFunc(router.HandlerLoginOwner)).Methods(http.MethodPost)
		r.Handle("/owner/auth/register", router.HandlerFunc(router.HandlerRegisterOwner)).Methods(http.MethodPost)

		apiRouter := r.PathPrefix("/api/v1").Subrouter()
		apiRouter.Use(auth.AuthenticationMiddleware)

		// customer
		apiRouter.Handle("/customers", router.HandlerFunc(router.HandlerAddCustomer)).Methods(http.MethodPost)
		apiRouter.Handle("/customers-list", router.HandlerFunc(router.HandlerAllCustomer)).Methods(http.MethodPost)
		apiRouter.Handle("/customers/{id}", router.HandlerFunc(router.HandlerOneCustomer)).Methods(http.MethodGet)
		apiRouter.Handle("/customers/{id}", router.HandlerFunc(router.HandlerUpdateCustomer)).Methods(http.MethodPut)
		apiRouter.Handle("/customers/{id}", router.HandlerFunc(router.HandlerDeleteCustomer)).Methods(http.MethodDelete)

		// owner
		apiRouter.Handle("/owners", router.HandlerFunc(router.HandlerAddOwner)).Methods(http.MethodPost)
		apiRouter.Handle("/owners-list", router.HandlerFunc(router.HandlerAllOwner)).Methods(http.MethodPost)
		apiRouter.Handle("/owners/{id}", router.HandlerFunc(router.HandlerOneOwner)).Methods(http.MethodGet)
		apiRouter.Handle("/owners/{id}", router.HandlerFunc(router.HandlerUpdateOwner)).Methods(http.MethodPut)
		apiRouter.Handle("/owners/{id}", router.HandlerFunc(router.HandlerDeleteOwner)).Methods(http.MethodDelete)

		// store
		apiRouter.Handle("/stores", router.HandlerFunc(router.HandlerAddStore)).Methods(http.MethodPost)
		apiRouter.Handle("/stores-list", router.HandlerFunc(router.HandlerAllStore)).Methods(http.MethodPost)
		apiRouter.Handle("/stores/{id}", router.HandlerFunc(router.HandlerOneStore)).Methods(http.MethodGet)
		apiRouter.Handle("/stores/{id}", router.HandlerFunc(router.HandlerUpdateStore)).Methods(http.MethodPut)
		apiRouter.Handle("/stores/{id}", router.HandlerFunc(router.HandlerDeleteStore)).Methods(http.MethodDelete)

		// seats
		apiRouter.Handle("/seats", router.HandlerFunc(router.HandlerAddSeats)).Methods(http.MethodPost)
		apiRouter.Handle("/seats-list", router.HandlerFunc(router.HandlerAllSeats)).Methods(http.MethodPost)
		apiRouter.Handle("/seats/{id}", router.HandlerFunc(router.HandlerOneSeats)).Methods(http.MethodGet)
		apiRouter.Handle("/seats/{id}", router.HandlerFunc(router.HandlerUpdateSeats)).Methods(http.MethodPut)
		apiRouter.Handle("/seats/{id}", router.HandlerFunc(router.HandlerDeleteSeats)).Methods(http.MethodDelete)

		// product
		apiRouter.Handle("/products", router.HandlerFunc(router.HandlerAddProduct)).Methods(http.MethodPost)
		apiRouter.Handle("/products-list", router.HandlerFunc(router.HandlerAllProduct)).Methods(http.MethodPost)
		apiRouter.Handle("/products/{id}", router.HandlerFunc(router.HandlerOneProduct)).Methods(http.MethodGet)
		apiRouter.Handle("/products/{id}", router.HandlerFunc(router.HandlerUpdateProduct)).Methods(http.MethodPut)
		apiRouter.Handle("/products/{id}", router.HandlerFunc(router.HandlerDeleteProduct)).Methods(http.MethodDelete)

		// booking
		apiRouter.Handle("/bookings", router.HandlerFunc(router.HandlerAddBooking)).Methods(http.MethodPost)
		apiRouter.Handle("/bookings-list", router.HandlerFunc(router.HandlerAllBooking)).Methods(http.MethodPost)
		apiRouter.Handle("/bookings/{id}", router.HandlerFunc(router.HandlerOneBooking)).Methods(http.MethodGet)
		apiRouter.Handle("/bookings/{id}", router.HandlerFunc(router.HandlerUpdateBooking)).Methods(http.MethodPut)
		apiRouter.Handle("/bookings/{id}", router.HandlerFunc(router.HandlerDeleteBooking)).Methods(http.MethodDelete)

		// booking details
		apiRouter.Handle("/bookings-details", router.HandlerFunc(router.HandlerAddBookingDetail)).Methods(http.MethodPost)
		apiRouter.Handle("/bookings-details-list", router.HandlerFunc(router.HandlerAllBookingDetail)).Methods(http.MethodPost)
		apiRouter.Handle("/bookings-details/{id}", router.HandlerFunc(router.HandlerOneBookingDetail)).Methods(http.MethodGet)
		apiRouter.Handle("/bookings-details/{id}", router.HandlerFunc(router.HandlerUpdateBookingDetail)).Methods(http.MethodPut)
		apiRouter.Handle("/bookings-details/{id}", router.HandlerFunc(router.HandlerDeleteBookingDetail)).Methods(http.MethodDelete)

		// transaction
		apiRouter.Handle("/transactions", router.HandlerFunc(router.HandlerAddTransaction)).Methods(http.MethodPost)
		apiRouter.Handle("/transactions-list", router.HandlerFunc(router.HandlerAllTransaction)).Methods(http.MethodPost)
		apiRouter.Handle("/transactions-sum", router.HandlerFunc(router.HandlerSumTransaction)).Methods(http.MethodPost)
		apiRouter.Handle("/transactions/{id}", router.HandlerFunc(router.HandlerOneTransaction)).Methods(http.MethodGet)
		apiRouter.Handle("/transactions/{id}", router.HandlerFunc(router.HandlerUpdateTransaction)).Methods(http.MethodPut)
		apiRouter.Handle("/transactions/{id}", router.HandlerFunc(router.HandlerDeleteTransaction)).Methods(http.MethodDelete)

		// midtrans
		apiRouter.Handle("/midtrans", router.HandlerFunc(router.HandlerAddMidtransTransaction)).Methods(http.MethodPost)
		r.HandleFunc("/midtrans", router.HandlerMidtransNotification)

		r.PathPrefix("/").Handler(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Cache-Control", "max-age=-1")
				http.StripPrefix("/", http.FileServer(http.Dir("files"))).ServeHTTP(w, r)
			}),
		)

		port := viper.GetInt("app.port")
		p := os.Getenv("PORT")
		if p != "" {
			port, _ = strconv.Atoi(p)
		}

		server := &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      r,
			ReadTimeout:  time.Duration(viper.GetInt("read_timeout")) * time.Second,
			WriteTimeout: time.Duration(viper.GetInt("write_timeout")) * time.Second,
			IdleTimeout:  time.Duration(viper.GetInt("idle_timeout")) * time.Second,
		}

		done := make(chan bool, 1)
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, os.Interrupt)

		go func() {
			<-quit
			log.Println("Server is shutting down...")

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			server.SetKeepAlivesEnabled(false)
			if err := server.Shutdown(ctx); err != nil {
				log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
			}
			close(done)
		}()

		log.Println("Server is ready to handle requests at", fmt.Sprintf(":%d", port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", fmt.Sprintf(":%d", port), err)
		}

		<-done
		log.Println("Server stopped")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is github.com/renosyah/receptionist-core-api/.server.toml)")
	cobra.OnInitialize(initConfig, initDB)
}

func initDB() {

	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.sslmode"))

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error open DB: %v\n", err))
		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error ping DB: %v\n", err))
		return
	}

	dbPool = db
}

func initConfig() {
	viper.SetConfigType("toml")
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
	} else {

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/receptionist-core-api")
		viper.SetConfigName(".server")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
