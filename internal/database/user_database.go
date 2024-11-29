package database

import (
	"ecommerce-app/config"
	"ecommerce-app/internal/model"
)

func GetAllUsers() ([]model.User, error) {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.FullName,
			&user.PhoneNumber,
			&user.Address,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func InsertUser(user *model.User) error {
	sqlStatement := `INSERT INTO users  (
                                            username,
                                            email,
                                            password,
                                            full_name,
                                            phone_number,
                                            address,
                                            status,
                                            created_at,
                                            updated_at
                                        )
                    VALUES  (
                                $1, $2, $3, $4, $5, $6, $7, $8, $9
                            ) 
                    RETURNING user_id`
	return config.DB.QueryRow(
		sqlStatement,
		user.Username,
		user.Email,
		user.Password,
		user.FullName,
		user.PhoneNumber,
		user.Address,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.UserID)
}

func UpdateUser(id string, user *model.User) error {
	sqlStatement := `UPDATE users SET

                    username=$1,
                    email=$2,
                    password=$3,
                    full_name=$4,
                    phone_number=$5,
                    address=$6,
                    status=$7,
                    created_at=$8,
                    updated_at=$9
                    
                    WHERE user_id=$10 RETURNING user_id`
	return config.DB.QueryRow(
		sqlStatement,
		user.Username,
		user.Email,
		user.Password,
		user.FullName,
		user.PhoneNumber,
		user.Address,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
		id,
	).Scan(&user.UserID)
}

func DeleteUser(id string) error {
	sqlStatement := `DELETE FROM users WHERE user_id =$1`
	_, err := config.DB.Exec(sqlStatement, id)
	return err
}
