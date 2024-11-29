package database

import (
	"ecommerce-app/config"
	"ecommerce-app/internal/model"
)

func GetAllUsersRole() ([]model.UserRole, error) {
	rows, err := config.DB.Query("SELECT * FROM user_roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersRole []model.UserRole
	for rows.Next() {
		var userRole model.UserRole
		if err := rows.Scan(
			&userRole.RoleID,
			&userRole.RoleName,
		); err != nil {
			return nil, err
		}
		usersRole = append(usersRole, userRole)
	}
	return usersRole, nil
}

func InsertUserRole(userRole *model.UserRole) error {
	sqlStatement := `INSERT INTO user_roles (
                                            role_name
                                       		)
									VALUES  (
												$1
											) 
									RETURNING role_id`
	return config.DB.QueryRow(
		sqlStatement,
		userRole.RoleName,
	).Scan(&userRole.RoleID)
}

func UpdateUserRole(id string, userRole *model.UserRole) error {
	sqlStatement := `UPDATE user_roles SET
                    role_name=$1             
                    WHERE 
					role_id=$2 
					RETURNING role_id`
	return config.DB.QueryRow(
		sqlStatement,
		userRole.RoleName,
		id,
	).Scan(&userRole.RoleID)
}

func DeleteUserRole(id string) error {
	sqlStatement :=    `DELETE 
						FROM
						
						user_roles
						
						WHERE 
						role_id =$1`
	_, err := config.DB.Exec(sqlStatement, id)
	return err
}
