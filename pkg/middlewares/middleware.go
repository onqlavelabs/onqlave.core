package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"onqlave.io/onqlave.core/pkg/common"
	"onqlave.io/onqlave.core/pkg/data"
	"onqlave.io/onqlave.core/pkg/utils/auth"
)

func WithServerHeader() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, common.ServerType)
			return hf(c)
		}
	}
}

func WithDatabaseContext(dbConnection data.DbConnection) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(data.DBConnectionKey, dbConnection)
			return next(c)
		}
	}
}

func WithContentType() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			contentType := c.Request().Header.Get("Content-Type")
			if contentType != "application/json" {
				fmt.Printf("content type is %s", contentType)
				c.Request().Header.Set("Content-Type", "application/json")
			}
			return hf(c)
		}
	}
}

func WithErrorHandler() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := hf(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, common.NewBaseErrorResponse(int32(http.StatusBadRequest), err.Error()))
			}
			return nil
		}
	}
}

func WithRequestValidation(disable bool) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			newly, context := common.GetOnqlaveContext(c)
			message := ""
			routeId := c.Request().Header.Get(string(common.OnqlaveRoute))
			if routeId == "" {
				message = fmt.Sprintf("Missing %s in request header", common.OnqlaveRoute)
			}
			context.SetRouteId(common.OnqlaveRouteHeader(routeId))
			onqlaveId := c.Request().Header.Get(string(common.OnqlaveId))
			if onqlaveId == "" {
				message = fmt.Sprintf("Missing %s in request header", common.OnqlaveId)
			}
			context.SetOnqlaveId(common.OnqlaveIdHeader(onqlaveId))
			version := c.Request().Header.Get(string(common.OnqlaveVersion))
			if version == "" {
				message = fmt.Sprintf("Missing %s in request header", common.OnqlaveVersion)
			}
			context.SetVersion(common.OnqlaveVersionHeader(version))
			if newly {
				c.Set(common.OnqlaveContext, context)
			}
			if message != "" && !disable {
				return c.JSON(http.StatusForbidden, common.NewBaseErrorResponse(int32(http.StatusForbidden), message))
			}
			return hf(c)
		}
	}
}

func WithLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().Host, "localhost1")
		},
	})
}

func WithCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowOriginFunc: func(origin string) (bool, error) {
			return true, nil
		},
	})
}

func WithRateLimiting() echo.MiddlewareFunc {
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 5, ExpiresIn: 1 * time.Second},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	return middleware.RateLimiterWithConfig(config)
}

func WithTimeout() echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 120 * time.Second,
	})
}

func WithBodyLimit() echo.MiddlewareFunc {
	return middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{
		Limit: "10M",
	})
}

func WithTracker(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request to tokenizer service routes")
		return next(c)
	}
}

func WithKeyAuthentication(authentication auth.KeyAuthenticationService) echo.MiddlewareFunc {
	config := middleware.KeyAuthConfig{
		KeyLookup: "header:" + string(common.OnqlaveApiKey),
		Validator: func(key string, c echo.Context) (bool, error) {
			valid, err := authentication.VerifyKey(c.Request().Context(), key)
			if valid {
				newly, context := common.GetOnqlaveContext(c)
				context.SetApiKey(common.OnqlaveApiKeyHeader(key))
				if newly {
					c.Set(common.OnqlaveContext, context)
				}
			}
			return valid, err
		},
		ErrorHandler: func(e error, c echo.Context) error {
			return e
		},
	}
	return middleware.KeyAuthWithConfig(config)
}

func WithJWTAuthentication(authentication auth.JwtAuthenticationService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get(echo.HeaderAuthorization)
			token, err := authentication.VerifyToken(c.Request().Context(), header)
			if err != nil {
				return err
			}
			_, context := common.GetOnqlaveContext(c)
			context.SetUser(token)
			return next(c)
		}
	}
}

func WithAdminJWT() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.String() == "/get-token"
		},
		SigningKey: []byte("secret"),
	})
}
