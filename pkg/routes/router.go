package routes

import (
	"time"
	"github.com/go-chi/chi/v5"
	"go-chi-app.com/go-chi-app/pkg/handlers"
	"go-chi-app.com/go-chi-app/pkg/middleware"
	
	"github.com/go-chi/httprate"
)


func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	// Apply rate limiting for every IP address
	router.Use(httprate.LimitByIP(3, 1*time.Minute))

	// Public routes
	router.Post("/login", handlers.Login)      // Login endpoint
	router.Post("/register", handlers.Register)// Endpoint to get a JWT
	router.Get("/oauth2/login", handlers.LoginHandler)       // Google Login
	router.Get("/oauth2/callback", handlers.CallbackHandler) // Callback
	router.Get("/verify", handlers.VerifyEmail)
	// Protected routes
	router.Route("/posts", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware) // Protect with JWT

		// Apply role-based access control
		r.Get("/", handlers.GetUsers)    // Only admins can GET /users
		r.Post("/", handlers.CreateUser) // Only admins can POST /users
	})
	router.Route("/users", func(r chi.Router) {
		r.Use(middleware.JWTMiddleware) // Protect with JWT

		// Apply role-based access control
		r.With(middleware.RoleMiddleware("admin")).Get("/", handlers.GetUsers)    // Only admins can GET /users
		r.With(middleware.RoleMiddleware("admin")).Post("/", handlers.CreateUser) // Only admins can POST /users
	})

	return router
}

