package likes

//func GetHandlerLikesCreate(usecase usecase.LikeUsecase) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		err := usecase.Set(ctx)
//		if err != nil {
//			return err.HTTP()
//		}
//		user, ok := conn.FindUserByCookie(session)
//		if !ok {
//			return echo.NewHTTPError(http.StatusUnauthorized, "no user session in db")
//		}
//		userID := user.ID
//		rawBody, err := ioutil.ReadAll(ctx.Request().Body)
//		if err != nil {
//			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//		}
//		defer ctx.Request().Body.Close()
//		newLike := like.Like{}
//		err = newLike.UnmarshalJSON(rawBody)
//		if err != nil {
//			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//		}
//		_, err = govalidator.ValidateStruct(newLike)
//		if err != nil {
//			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//		}
//		newLike.UserID = userID
//		e := conn.Upsert(newLike)
//		if e != nil {
//			return echo.NewHTTPError(e.Status, e.Error)
//		}
//		return nil
//	}
//}
