package constants

type ErrorCode int

const OK ErrorCode = 0
const NOT_FOUND ErrorCode = 1
const UNKOWN_ERROR ErrorCode = 2
const DB_ERROR ErrorCode = 3
const WRONG_REQ_FORMAT ErrorCode = 4
const JWT_NOT_GENERATED ErrorCode = 5
const JWT_TOKEN_ERROR ErrorCode = 6
const JWT_TOKEN_EXPERIED ErrorCode = 7
const REQUIRED_LOGIN ErrorCode = 8
