CREATE TABLE IF NOT EXISTS spaceships (
    id INT AUTO_INCREMENT,
    name VARCHAR(255),
    class VARCHAR(255),
    crew INT,
    image VARCHAR(255),
    value FLOAT,
    status VARCHAR(255),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS armaments (
    id INT AUTO_INCREMENT,
    spaceship_id INT,
    title VARCHAR(255),
    qty VARCHAR(255),
    PRIMARY KEY (id),
    FOREIGN KEY (spaceship_id) REFERENCES spaceships(id)
);

-- How to run the script
/**
*  #connect to the database
*  - mysql -u username -p database_name
*  suggestion for database name -> galatic
* 
*  #execute sql script
*  - source [/path/to/create_tables.sql] -> source ./sql/create_tables.sql
*/