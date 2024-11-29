package database

import (
	"ecommerce-app/config"
	"ecommerce-app/internal/model"
)

func GetAllUsersRoleMap() ([]model.UserRoleMap, error) {
	rows, err := config.DB.Query("SELECT * FROM user_role_map")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersRoleMap []model.UserRoleMap
	for rows.Next() {
		var userRoleMap model.UserRoleMap
		if err := rows.Scan(
			&userRoleMap.UserID,
			&userRoleMap.RoleID,
		); err != nil {
			return nil, err
		}
		usersRoleMap = append(usersRoleMap, userRoleMap)
	}
	return usersRoleMap, nil
}

func InsertUserRoleMap(userRoleMap *model.UserRoleMap) error {
	sqlStatement := `INSERT INTO user_role_map 
					(
                    	user_id,
						role_id
               		)
					VALUES  
					(
						$1,
						$2
					) 
					RETURNING user_id`
	return config.DB.QueryRow(
		sqlStatement,
		userRoleMap.UserID,
		userRoleMap.RoleID,
	).Scan(&userRoleMap.UserID)
}

func UpdateUserRoleMap(id string, userRoleMap *model.UserRoleMap) error {
	sqlStatement := `UPDATE user_role_map SET

                     role_id=$1             
    
					 WHERE 
	
					 user_id=$2 
	
					 RETURNING user_id`
	return config.DB.QueryRow(
		sqlStatement,
		userRoleMap.RoleID,
		id,
	).Scan(&userRoleMap.RoleID)
}

func DeleteUserRoleMap(id string) error {
	sqlStatement := `DELETE 
						FROM
						
						user_role_map
						
						WHERE 
						user_id =$1 `
	_, err := config.DB.Exec(sqlStatement, id)
	return err
}
