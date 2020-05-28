package helpers

import (
	"errors"
)

//http code
const BadRequest = 400
const Unauthorized = 401
const InternalServerError = 500
const NotImplemented = 501

//command string
const Post = "post"
const Put = "put"
const Get = "get"
const Delete = "delete"
const CheckAccount = "check-account"
const AccountExist = "account-exist"
const AccountLogin = "account-login"
const JWTLogin = "jwt-login"
const CheckProject = "check-project"
const QueryError = "query-error"
const VoteExist = "vote-exist"
const CheckVote = "check-vote"
const Oauth = "oauth"

const PostMessage = "Error inserting data to database"
const PutMessage = "Error updating data to database"
const GetMessage = "Error getting data from database"
const DeleteMessage = "Error deleting data from database"
const JWTLoginMessage = "Error creating JWT token"
const QueryErrorMessage = "Error query"
const CheckAccountMessage = "Account is not exist"
const AccountExistMessage = "Account with this username is already exist"
const AccountLoginMessage = "Username or password is incorrect"
const CheckProjectMessage = "Project is not exist"
const VoteExistMessage = "User already vote this project"
const CheckVoteMessage = "Vote is not exist"
const OauthMessage = "Unauthorized"

func ErrorMessage(command string) error {
	switch command {
	case Oauth:
		return errors.New(OauthMessage)
	case CheckAccount:
		return errors.New(CheckAccountMessage)
	case AccountExist:
		return errors.New(AccountExistMessage)
	case AccountLogin:
		return errors.New(AccountLoginMessage)
	case CheckProject:
		return errors.New(CheckProjectMessage)
	case VoteExist:
		return errors.New(VoteExistMessage)
	case CheckVote:
		return errors.New(CheckVoteMessage)
	case Post:
		return errors.New(PostMessage)
	case Put:
		return errors.New(PutMessage)
	case Get:
		return errors.New(GetMessage)
	case Delete:
		return errors.New(DeleteMessage)
	case JWTLogin:
		return errors.New(JWTLoginMessage)
	case QueryError:
		return errors.New(QueryErrorMessage)
	default:
		return errors.New("Bad Request")
	}
}

func ErrorCode(err string) int {
	switch err {
	case OauthMessage:
		return 401
	case CheckAccountMessage:
		return 400
	case AccountExistMessage:
		return 400
	case AccountLoginMessage:
		return 400
	case CheckProjectMessage:
		return 400
	case VoteExistMessage:
		return 400
	case CheckVoteMessage:
		return 400
	case PostMessage:
		return 500
	case PutMessage:
		return 500
	case GetMessage:
		return 500
	case DeleteMessage:
		return 500
	case JWTLoginMessage:
		return 500
	case QueryErrorMessage:
		return 500
	default:
		return 400
	}
}
