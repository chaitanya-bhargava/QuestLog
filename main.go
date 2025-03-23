package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/chaitanya-bhargava/QuestLog/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	key := os.Getenv("SESSION_SECRET")
	if key == "" {
		log.Fatal("SESSION_SECRET is not found in the environment")
	}

	maxAge := 86400 * 30
	isProd := true     // Set to true in production (when using HTTPS)

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd // Set to true in production

	gothic.Store = store

	publicURL := os.Getenv("PUBLIC_URL")
    if publicURL == "" {
        log.Fatal("PUBLIC_URL is not found in the environment")
    }

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), publicURL+"/v1/auth/google/callback","email","profile"),
	)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Get("/genres", middlewareRawg(handlerGetGenres))
	v1Router.Get("/search", middlewareRawg(handlerSearchGames))
	v1Router.Post("/gameLog", apiCfg.middlewareAuthRawg(apiCfg.handlerCreateGameLog))
	v1Router.Get("/gameLog",apiCfg.middlewareAuth(apiCfg.handlerGetGameLogByGameID))
	v1Router.Delete("/gameLog",apiCfg.middlewareAuth(apiCfg.handlerDeleteGameLogByGameID))
	v1Router.Get("/games", middlewareRawg(handlerGetGamesByGenre))
	v1Router.Get("/games/gameLog",apiCfg.middlewareAuth(apiCfg.handlerGetGamesByGameLog))
	v1Router.Post("/game", apiCfg.middlewareAuthRawg(apiCfg.handlerCreateGame))
	v1Router.Get("/auth/{provider}", handlerBeginAuth)
	v1Router.Get("/auth/{provider}/callback", apiCfg.handlerCompleteAuth)
	v1Router.Get("/auth/logout",handlerLogout)
	v1Router.Get("/auth/me", apiCfg.handlerGetCurrentUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
