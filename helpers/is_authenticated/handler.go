package isauthenticated

import (
	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// token, err := request.ParseFromRequest(c.Request(), request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		// 	return []byte("secret"), nil
		// })

		// if err != nil {
		// 	return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
		// }

		// if !token.Valid {
		// 	return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
		// }

		// claims, ok := token.Claims.(jwt.MapClaims)
		// if !ok {
		// 	return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
		// }

		// userID, ok := claims["user_id"].(float64)
		// if !ok {
		// 	return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
		// }

		// c.Set("user_id", int(userID))

		return next(c)
	}
}
