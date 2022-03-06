package constants

const (
	CREATE_BOOK_TABLE = `CREATE TABLE IF NOT EXISTS book ( 
		name text not null,
		id text not null,
		category text,
		author text,
		translator text,
		owner text not null,
		add_time timestamp,
		PRIMARY KEY (name)
	);	
		`
	CREATE_BOOK_TABLE_CATEGORY = `CREATE TABLE IF NOT EXISTS category ( 
		category text not null ,
		name text not null,
		author text not null,
		translator text not null,
		PRIMARY KEY (category) 
		
		);
		`
	GET_BOOK_ON_CATEGORY = `SELECT * FROM category WEHRE category=$1;`

	ADD_BOOK = `insert into book (name, id, category, author, translator, owner, add_time) values ($1, $2, $3, $4, $5, $6, $7);`

	ADD_BOOK_CATEGORY = `INSERT INTO category (category, name, author, translator) VALUES($1,$2,$3,$4);`

	DELETE_BOOK = `DELETE FROM book WHERE name=$1;`

	GET_ONE_BOOK = `select * from book where name=$1;`

	GET_ALL_BOOKS = `SELECT name FROM book;`
)
