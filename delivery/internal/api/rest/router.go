package rest

import (
	"delivery/internal/api/rest/handlers/delivery"
	"delivery/internal/api/rest/handlers/health"
	"delivery/internal/api/rest/middleware"
)

func (s *Server) SetupMonitoringRoutes(healthHandler *health.HealthHandler) {
	r := s.engine
	r.GET("/health", healthHandler.CheckHealth)
}

func (s *Server) SetupAPIRoutes(
	internalMiddleware middleware.InternalMiddleware,
	deliveryHandler delivery.Handler,
) {
	r := s.engine
	{
		v1 := r.Group("/v1")
		{
			internal := v1.Group("", internalMiddleware.Handle)

			internal.POST("/deliveries", deliveryHandler.SendCourior)
			internal.POST("/webhooks/courier-status", deliveryHandler.ReceiveCouriorStatus)
		}
	}
}
