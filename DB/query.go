package DB

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/KristianXi3/crud/entity1"
)

func (s *Dbstruct) GetUsers(ctx context.Context) ([]entity1.User, error) {
	var result []entity1.User

	err := s.SqlDb.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := s.SqlDb.QueryContext(ctx, "select id, username, email, password, age, createdat, updatedat from users")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row entity1.User
		err := rows.Scan(
			&row.Id,
			&row.Username,
			&row.Email,
			&row.Password,
			&row.Age,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (s *Dbstruct) GetUserByID(ctx context.Context, userid int) (*entity1.User, error) {
	result := &entity1.User{}

	err := s.SqlDb.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := s.SqlDb.QueryContext(ctx, "select id, username, email, password, age, createdat, updatedat from users where id = @ID", sql.Named("ID", userid))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&result.Id,
			&result.Username,
			&result.Email,
			&result.Password,
			&result.Age,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return result, nil
}

func (s *Dbstruct) CreateUser(ctx context.Context, user entity1.User) (string, error) {
	var result string

	err := s.SqlDb.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	_, err = s.SqlDb.ExecContext(ctx, "insert into users (id, username, email, password, age, createdat, updatedat) values (@id, @username, @email, @password, @age, @createdat, @updatedat)",
		sql.Named("id", user.Id),
		sql.Named("username", user.Username),
		sql.Named("email", user.Email),
		sql.Named("password", user.Password),
		sql.Named("age", user.Age),
		sql.Named("createdat", time.Now()),
		sql.Named("updatedat", time.Now()))
	if err != nil {
		return "", err
	}

	result = "Inserted"

	return result, nil
}

func (s *Dbstruct) UpdateUser(ctx context.Context, userId int, user entity1.User) (string, error) {
	var result string

	err := s.SqlDb.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	_, err = s.SqlDb.ExecContext(ctx, "update users set username = @username,email = @email, password = @password, age = @age, updatedat = @updatedat where id = @id",
		sql.Named("id", userId),
		sql.Named("username", user.Username),
		sql.Named("email", user.Email),
		sql.Named("password", user.Password),
		sql.Named("age", user.Age),
		sql.Named("updatedat", time.Now()))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	result = "Updated"

	return result, nil
}

func (s *Dbstruct) DeleteUser(ctx context.Context, userId int) (string, error) {
	var result string

	err := s.SqlDb.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	_, err = s.SqlDb.ExecContext(ctx, "delete from users where id=@id", sql.Named("id", userId))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	result = "Deleted"

	return result, nil
}
