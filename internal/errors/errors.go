package errors

const (
	// User-related errors
	CodeInvalidUserType            = 1001 // Неправильный тип пользователя
	CodeUserNotFound               = 1002 // Пользователь не найден
	CodeInvalidCredentials         = 1003 // Неправильные учетные данные
	CodeUserAlreadyExists          = 1004 // Пользователь уже существует
	CodeEmptyAuthorizationHeader   = 1005 // Пустой заголовок авторизации
	CodeInvalidAuthorizationHeader = 1006 // Неправильный заголовок авторизации

	// Service-related errors
	CodeServiceError  = 2001 // Общая ошибка сервиса
	CodeDatabaseError = 2002 // Ошибка базы данных
	CodeUnauthorized  = 2003 // Неавторизованный доступ
	CodeForbidden     = 2004 // Доступ запрещен
	CodeInvalidInput  = 2005 // Неправильные входные данные

	// Flat-related errors
	CodeFlatNotFound    = 3001 // Квартира не найдена
	CodeFlatStatusError = 3003 // Ошибка статуса квартиры

	// House-related errors
	CodeHouseNotFound      = 4001 // Дом не найден
	CodeHouseAlreadyExists = 4002 // Дом уже существует

	// Subscription-related errors
	CodeSubscriptionError = 5001 // Ошибка подписки

	// General errors
	CodeInternalServerError = 6001 // Внутренняя ошибка сервера
	CodeBadRequest          = 6002 // Неправильный запрос
	CodeConflict            = 6003 // Конфликт данных
	CodeRateLimitExceeded   = 6004 // Превышено ограничение на количество запросов
)
