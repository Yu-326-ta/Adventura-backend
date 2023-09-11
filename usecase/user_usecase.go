package usecase

import (
	"echo-todo-api/model"
	"echo-todo-api/repository"
	"echo-todo-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	// 関数に対して情報を与えたいだけなので値渡し、関数内での変更を呼び出し元にも反映したい場合は参照渡し
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	// usecaseはripositoryのinterfaceにのみ依存なのでフィールドはurのみ（validatorも依存）
	ur repository.IUserRepository
	uv validator.IUserValidator
	
}

// 関数の最初にNewがつく関数をコンストラクタ関数といい、インスタンスの作成や初期化で使う
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	// 返り値のIUserUsecaseを満たすためにSignUpとLoginを実装する必要がある0
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	// 10は暗号化の複雑さを表す
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	// CreateUserが成功すると参照渡しのためnewUserの値が新しくなっているのでUserResponseを新しい情報に変更して返す
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// データベースに保存しているハッシュ化したパスワードと送られてくる平文のパスワード比較
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		// jwtの有効期限の設定（ここでは12時間）
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	// tokenに対してSignedStringメソッドを実行でjwttokenを生成、その時にシークレットキーを渡す
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
